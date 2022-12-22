package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/logger"
	"github.com/urfave/negroni"
)

type Options func(s *http.Server)

const TIMEOUT = 30 * time.Second

func Start(
	port int, svc, version string, h http.Handler, logger logger.Logger, option ...Options,
) error {
	n := negroni.New()

	n.UseHandler(h)

	srv := &http.Server{
		ReadHeaderTimeout: TIMEOUT,
		ReadTimeout:       TIMEOUT,
		WriteTimeout:      TIMEOUT,
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           n,
	}

	for _, opt := range option {
		opt(srv)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()

	logger.Infof("Starting server on port %d", port)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func WithReadTimeout(timeout time.Duration) Options {
	return func(s *http.Server) {
		s.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Options {
	return func(s *http.Server) {
		s.WriteTimeout = timeout
	}
}

func WithMiddlewares(middlwares ...negroni.HandlerFunc) Options {
	return func(s *http.Server) {
		n, ok := s.Handler.(*negroni.Negroni)
		if !ok {
			n = negroni.New()
		}

		for _, m := range middlwares {
			n.Use(m)
		}
	}
}
