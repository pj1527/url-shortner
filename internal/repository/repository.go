package repository

import "context"

type Repository interface {
	SaveURL(ctx context.Context, longURL string) (uint64, error)
	GetURL(ctx context.Context, id uint64) (string, bool)
}
