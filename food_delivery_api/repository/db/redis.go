package db

import (
	"context"
	"fmt"
	"food_delivery/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	redisConnStr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr: redisConnStr,
	})

	// Test the connection
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong)

	return client, nil
}
