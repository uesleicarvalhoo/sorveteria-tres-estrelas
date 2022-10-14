package auth

import "context"

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
