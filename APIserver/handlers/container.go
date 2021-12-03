package handlers

import (
	"github.com/go-redis/redis/v8"
)

var rClient *redis.Client

// Container will hold all dependencies for your application.
type Container struct {
	RedisClient *redis.Client
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer() (Container, error) {
	containerInit()
	c := Container{RedisClient: rClient}
	return c, nil
}

func containerInit() {
	rClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
