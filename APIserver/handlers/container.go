package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var rClient *redis.Client

// Container will hold all dependencies for your application.
type Container struct {
	RedisClient *redis.Client
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer() (Container, error) {
	err := containerInit()
	c := Container{RedisClient: rClient}
	return c, err
}

func containerInit() error {
	rClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := rClient.SetNX(context.Background(), "global:nextUserId", 0, 0).Result()
	if err != nil {
		return err
	}
	_, err = rClient.SetNX(context.Background(), "global:nextHaikuId", 0, 0).Result()
	if err != nil {
		return err
	}
	return nil

}
