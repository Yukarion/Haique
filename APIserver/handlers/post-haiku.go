package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

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
	content = eraseSpaceInContent(content)
	if content.First == "" || content.Second == "" || content.Third == "" {
		return ctx.HTML(http.StatusBadRequest, "containing empty clause")
	}
	author_id, err := c.RedisClient.Get(ctxBG, session_id+":linked_user_id").Result()
	if err != nil {
		return ctx.HTML(http.StatusBadRequest, "invalid session id")
	}
	haiku_id, err := c.RedisClient.Incr(ctxBG, "global:next_haiku_id").Result()
	if err != nil {
		return err
	}
	haiku_id_str := strconv.FormatInt(haiku_id, 10)
	current_unix_time := time.Now().Unix()
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":content", strings.Join([]string{content.First, content.Second, content.Third}, " "), 0)
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":author_id", author_id, 0)
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":likes", 0, 0)
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":createdAt", current_unix_time, 0)

	return ctx.NoContent(http.StatusCreated)
}
