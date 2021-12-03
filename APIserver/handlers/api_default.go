package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/models"
	"github.com/labstack/echo/v4"
)

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
	val, err := c.RedisClient.Get(context.Background(), "key").Result()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: val,
	})
}

// PostApiSignup -
func (c *Container) PostApiSignup(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PostHaiku -
func (c *Container) PostHaiku(ctx echo.Context) error {
	var payload models.ApiPostHaikuContent
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}

	//if !isValidSessionId(session_id){ reject! }
	if payload.First == "" || payload.Second == "" || payload.Third == "" {
		return ctx.HTML(http.StatusBadRequest, "Empty haiku is not allowed")
	}
	if err := c.RedisClient.Set(context.Background(), "key", "yo!", 0).Err(); err != nil {
		log.Println("omg!")
		return err
	}
	log.Println("created")
	return ctx.HTML(http.StatusCreated, ""+payload.Second)

	//return ctx.NoContent(http.StatusOK)
}

// PostSignin -
func (c *Container) PostSignin(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PostSubscribe -
func (c *Container) PostSubscribe(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
