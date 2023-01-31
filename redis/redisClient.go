package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

var ctx = context.Background()

func (rdb *RedisClient) ConfigureRedisClient() {
	rdb.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.client.Ping(rdb.client.Context()).Result()

	if err != nil {
		// TODO Create custom error and wrap this
		panic(err)
	}
}

func (rdb *RedisClient) GetValue(key string) string {
	rdb.ConfigureRedisClient()
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			// TODO Create custom error and wrap this
			panic(err)
		}
	}(rdb.client)
	value, _ := rdb.client.Get(ctx, key).Result()
	return value
}

func (rdb *RedisClient) SetValue(key string, val interface{}, expiry time.Duration) string {
	rdb.ConfigureRedisClient()
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			// TODO Create custom error and wrap this
			panic(err)
		}
	}(rdb.client)
	value, err := rdb.client.Set(ctx, key, val, expiry).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		panic(err)
	}
	return value
}
