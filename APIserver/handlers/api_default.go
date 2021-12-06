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
