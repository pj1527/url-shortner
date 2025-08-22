package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisRepository struct {
	client        *redis.Client
	saveURLScript *redis.Script
}

const saveURLScriptBody = `
    local id = redis.call('INCR', KEYS[1])
    redis.call('SET', 'url:' .. id, ARGV[1])
    return id
`

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
	script := redis.NewScript(saveURLScriptBody)
	return &redisRepository{
		client:        client,
		saveURLScript: script,
	}, nil
}

func (r *redisRepository) SaveURL(ctx context.Context, longURL string) (uint64, error) {
	keys := []string{"url_counter"}
	args := []any{longURL}
	result, err := r.saveURLScript.Run(ctx, r.client, keys, args...).Result()
	if err != nil {
		log.Println("Failed to execute redis script", err)
		return 0, fmt.Errorf("failed to execute redis script")
	}
	id, ok := result.(int64)
	if !ok {
		return 0, fmt.Errorf("unexpected result type from redis script: %T", result)
	}
	return uint64(id), nil
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
