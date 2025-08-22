package repository

import (
	"context"
	"log"
)

type Repository interface {
	SaveURL(ctx context.Context, longURL string) (uint64, error)
	GetURL(ctx context.Context, id uint64) (string, bool)
}

type Config struct {
	UseRedis  bool
	RedisAddr string
}

func NewRepository(cfg Config) (Repository, error) {
	if cfg.UseRedis {
		log.Println("Using Redis")
		return NewRedisRepository(cfg.RedisAddr)
	}
	log.Println("Using InMemory")
	return NewInMemoryRepository()
}
