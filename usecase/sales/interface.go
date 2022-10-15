package sales

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type Reader interface {
	Search(ctx context.Context, start, end time.Time) ([]entity.Sale, error)
}

type Writer interface {
	Create(ctx context.Context, s entity.Sale) error
}

type Repository interface {
	Reader
	Writer
}
