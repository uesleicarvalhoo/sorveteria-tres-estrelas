package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/logger"
	"golang.org/x/sync/errgroup"
)

const TIMEOUT = time.Second * 30

func Subscribe(start, stop func(ctx context.Context) error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGALRM)

		<-quit
		cancel()
	}()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return start(ctx)
	})

	g.Go(func() error {
		<-gCtx.Done()

		return stop(ctx)
	})

	if err := g.Wait(); err != nil {
		logger.Errorf("exiting with error: %s", err)
	}
}
