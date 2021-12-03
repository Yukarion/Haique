package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// PostHaiku -
func (c *Container) PostHaiku(ctx echo.Context) error {
	var payload models.InlineObject2
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}
	session_id := payload.SessionId
	content := payload.Content
	if !isValidSessionId(session_id) {
		ctx.NoContent(http.StatusBadRequest)
	}
	if content.First == "" || content.Second == "" || content.Third == "" {
		return ctx.HTML(http.StatusBadRequest, "Empty haiku is not allowed")
	}
	if err := c.RedisClient.Set(context.Background(), "key", "yo!", 0).Err(); err != nil {
		log.Println("omg!")
		return err
	}
	log.Println("created")
	return ctx.HTML(http.StatusCreated, ""+content.Second)

	//return ctx.NoContent(http.StatusOK)
}
