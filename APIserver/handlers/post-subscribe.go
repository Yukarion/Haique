package handlers

import (
	"net/http"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// PostSubscribe -
func (c *Container) PostSubscribe(ctx echo.Context) error {
	var payload models.SessionId

	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}
	session_id := payload.SessionId

	subscriber_id_str, err := c.RedisClient.Get(ctxBG, session_id+":linked_user_id").Result()
	if err != nil {
		return ctx.HTML(http.StatusBadRequest, "invalid session id")
	}
	receiver_id_str := ctx.Param("user_id")

	_, err = c.RedisClient.SAdd(ctxBG, "user_id:"+subscriber_id_str+"subscription", receiver_id_str).Result()
	if err != nil {
		return err
	}
	_, err = c.RedisClient.SAdd(ctxBG, "user_id:"+receiver_id_str+"subscribed_by", subscriber_id_str).Result()
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
