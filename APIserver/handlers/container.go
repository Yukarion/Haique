package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const TOP_HAIKUS_NUM = 30
const DUMMY_HAIKU_ID = -1

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
	_, err := rClient.SetNX(context.Background(), "global:next_user_id", 0, 0).Result()
	if err != nil {
		return err
	}
	_, err = rClient.SetNX(context.Background(), "global:next_haiku_id", 0, 0).Result()
	if err != nil {
		return err
	}

	//top_haiku_idsの大きさがTOP_HAIKUS_NUM未満のときに、調整で末尾に-1を追加（基本的にTOP_HAIKUS_NUM or 0のはずだが、TOP_HAIKUS_NUMが変更されたときは便利）
	length, err := rClient.LLen(context.Background(), "global:top_haiku_id_list").Result()

	for i := 0; int64(i) < TOP_HAIKUS_NUM-length; i++ {
		_, err = rClient.RPush(context.Background(), "global:top_haiku_id_list", DUMMY_HAIKU_ID).Result()
		if err != nil {
			return err
		}
	}

	return nil

}
