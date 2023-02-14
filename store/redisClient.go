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

func NewRedisClient(ctx context.Context) (*RedisClient, error) {
	newClient := redis.NewClient(
		&redis.Options{
			Addr:     config.GetRedisClientAddress(),
			Password: config.GetRedisClientPassword(),
			DB:       config.GetRedisDB(),
		})

	_, err := newClient.Ping(newClient.Context()).Result()

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
	defer rdb.client.Close()
	value, err := rdb.client.Get(rdb.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rdb *RedisClient) SetValue(key string, val interface{}, expiry time.Duration) error {
	defer rdb.client.Close()
	_, err := rdb.client.Set(rdb.ctx, key, val, expiry).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		return err
	}
	return nil
}
