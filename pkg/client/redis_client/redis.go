package redis_client

import (
	"aero-internship/pkg/config"
	"github.com/go-redis/redis/v7"
)

func NewRedisClient(cfg *config.Config, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{Addr: cfg.GetRedisAddr(), Password: cfg.GetRedisPass(), DB: db})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
