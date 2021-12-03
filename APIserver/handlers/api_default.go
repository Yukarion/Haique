package handlers

import (
	"context"
	"net/http"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

var ctxBG = context.Background()

// DeleteApiHaikuId -
func (c *Container) DeleteApiHaikuId(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// DeleteApiSubscribeUserId -
func (c *Container) DeleteApiSubscribeUserId(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetApiHaikuId - get_haiku
func (c *Container) GetApiHaikuId(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetApiTimeline - timeline
func (c *Container) GetApiTimeline(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetApiUser - user_info
func (c *Container) GetApiUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetTop - top
func (c *Container) GetTop(ctx echo.Context) error {
	val, err := c.RedisClient.Get(ctxBG, "key").Result()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: val,
	})
}

// PostSubscribe -
func (c *Container) PostSubscribe(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
