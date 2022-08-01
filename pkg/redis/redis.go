package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
)

func NewRedisClient() *redis.Client {
	var err error = nil
	if os.Getenv("APP_ENV") != "PRODUCTION" {
		err = godotenv.Load(".env")
	}

	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redis := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return redis
}
