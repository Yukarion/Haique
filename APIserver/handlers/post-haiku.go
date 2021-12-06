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

	author_id_str, err := c.RedisClient.Get(ctxBG, session_id+":linked_user_id").Result()
	if err != nil {
		return ctx.HTML(http.StatusBadRequest, "invalid session id")
	}
	author_id, _ := strconv.Atoi(author_id_str)

	haiku_id, err := c.RedisClient.Incr(ctxBG, "global:next_haiku_id").Result()
	if err != nil {
		return err
	}
	haiku_id_str := strconv.FormatInt(haiku_id, 10)

	author_name, err := c.RedisClient.Get(ctxBG, "user_id:"+author_id_str+":name").Result()
	if err != nil {
		return err
	}

	current_unix_time := time.Now().Unix()
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":content", strings.Join([]string{content.First, content.Second, content.Third, author_name}, " "), 0)
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":author_id", author_id, 0)
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":likes", 0, 0)
	c.RedisClient.Set(ctxBG, "haiku_id:"+haiku_id_str+":created_at", current_unix_time, 0)

	c.RedisClient.LPush(ctxBG, "global:top_haiku_id_list", haiku_id)
	c.RedisClient.RPop(ctxBG, "global:top_haiku_id_list")

	return ctx.NoContent(http.StatusCreated)
}
