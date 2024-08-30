package db

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func InitRedis() *redis.Client {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))
	if err != nil {
		log.Fatalf("failed to parse Redis connection URL: %v", err)
	}

	rdb := redis.NewClient(opt)

	_, err = rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	return rdb
}
