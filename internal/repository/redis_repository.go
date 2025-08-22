package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string) (Repository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}
	return &redisRepository{
		client: client,
	}, nil
}

func (r *redisRepository) SaveURL(ctx context.Context, longURL string) (uint64, error) {
	id, err := r.client.Incr(ctx, "url_counter").Uint64()
	if err != nil {
		return 0, err
	}
	err = r.client.Set(ctx, r.getKey(id), longURL, 0).Err()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *redisRepository) GetURL(ctx context.Context, id uint64) (string, bool) {
	val, err := r.client.Get(ctx, r.getKey(id)).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (r *redisRepository) getKey(id uint64) string {
	return "url:" + strconv.FormatUint(id, 10)
}
