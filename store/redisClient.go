package store

import (
	"Golang-assignment-srikomm/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(ctx context.Context, environment string) (*RedisClient, error) {
	redisConfig := config.GetRedisConfig(environment)
	newClient := redis.NewClient(
		&redis.Options{
			Addr:     redisConfig.ClientAddress,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})

	_, err := newClient.Ping(ctx).Result()

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &RedisClient{
		client: newClient,
		ctx:    ctx,
	}, nil
}

func (rdb *RedisClient) GetValue(key string) (string, error) {
	value, err := rdb.client.Get(rdb.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rdb *RedisClient) SetValue(key string, val interface{}, expiry time.Duration) error {
	_, err := rdb.client.Set(rdb.ctx, key, val, expiry).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		return err
	}
	return nil
}
