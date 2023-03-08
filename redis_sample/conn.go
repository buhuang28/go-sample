package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

type RedisConfig struct {
	Ip       string
	Port     string
	Password string
	DB       int
}

func initClient(c RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Ip, c.Port),
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
