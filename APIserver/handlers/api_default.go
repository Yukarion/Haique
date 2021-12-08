package handlers

import (
	"context"
	"net/http"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

var ctxBG = context.Background()

// DeleteHaiku -
//需要がそんなにないので実装は後回し！
func (c *Container) DeleteHaiku(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetHaiku - get_haiku
func (c *Container) GetHaiku(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// GetUser - user_info
func (c *Container) GetUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
