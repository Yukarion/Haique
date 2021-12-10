package handlers

import (
	"net/http"
	"time"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// PostSignin -
func (c *Container) PostSignin(ctx echo.Context) error {
	/*
		今の実装では、ログインでもセッションIDが新しく生成されて、どんどん溜まっていくので注意
	*/
	var payload models.InlineObject1
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}
	name := payload.Name
	rawPw := payload.Pw

	hashedPw, err := c.RedisClient.Get(ctxBG, name+":pw").Result()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "not registered (maybe typo)")
	}

	if isValidPassword(hashedPw, rawPw) {
		user_id, err := c.RedisClient.Get(ctxBG, name+":user_id").Result()
		if err != nil {
			return err
		}
		session_id, err := c.UUIDgenerator()
		if err != nil {
			return err
		}
		c.RedisClient.Set(ctxBG, session_id+":linked_user_id", user_id, time.Hour*1)
		return ctx.JSON(http.StatusOK, models.InlineObject3{SessionId: session_id})
	}
	return ctx.NoContent(http.StatusBadRequest)
}
