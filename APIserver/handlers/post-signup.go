package handlers

import (
	"net/http"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// PostApiSignup -
func (c *Container) PostApiSignup(ctx echo.Context) error {
	var payload models.InlineObject1
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}
	name := payload.Name
	pw := payload.Pw
	session_id, err := genUUID()
	if err != nil {
		return err
	}
	/*
		{session_id}:linked_user_id -> {user_id}
		{id}:pw -> {pw}
		{id}:user_id -> {user_id}
	*/
	c.RedisClient.Set(ctxBG, name+":password", hashPW(pw), 0)
	userId := c.RedisClient.Incr(ctxBG, "global:nextUserId")
	c.RedisClient.Set(ctxBG, name+":user_id", userId, 0)
	c.RedisClient.Set(ctxBG, session_id.String()+":linked_user_id", userId, 0)
	return ctx.JSON(http.StatusCreated, models.SessionId{Id: session_id})
}
