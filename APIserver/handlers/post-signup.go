package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

const STRETCH_NUM = 5

// PostSignup -
func (c *Container) PostSignup(ctx echo.Context) error {
	var payload models.InlineObject1
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}
	name := payload.Name
	rawPw := payload.Pw

	hashedPw, err := hashPassword(rawPw, STRETCH_NUM)

	if err != nil {
		return err
	}
	/*
		{session_id}:linked_user_id -> {user_id}
		{name}:pw -> {pw}
		{name}:user_id -> {user_id}
	*/
	isNameUnique, err := c.RedisClient.SetNX(ctxBG, name+":pw", hashedPw, 0).Result()
	if err != nil {
		return err
	}
	if !isNameUnique {
		// 登録済みのuser名は不可
		return ctx.NoContent(http.StatusConflict)
	}
	session_id, err := c.UUIDgenerator()
	if err != nil {
		return err
	}
	//pwは上でsetされていることに注意
	user_id, err := c.RedisClient.Incr(ctxBG, "global:next_user_id").Result()
	if err != nil {
		return err
	}
	user_id_str := strconv.Itoa(int(user_id))
	c.RedisClient.Set(ctxBG, name+":user_id", user_id, 0)
	c.RedisClient.Set(ctxBG, session_id+":linked_user_id", user_id, time.Hour*1)
	c.RedisClient.Set(ctxBG, "user_id:"+user_id_str+":name", name, 0)
	return ctx.JSON(http.StatusCreated, models.InlineObject3{SessionId: session_id})
}
