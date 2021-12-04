package handlers

import (
	"log"
	"net/http"

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
		return err
	}

	if isValidPassword(hashedPw, rawPw) {
		userId, err := c.RedisClient.Get(ctxBG, name+":user_id").Result()
		if err != nil {
			log.Println(name, err)
			return err
		}
		session_id, err := genUUID()
		if err != nil {
			return err
		}
		c.RedisClient.Set(ctxBG, session_id+":linked_user_id", userId, 0)
		return ctx.JSON(http.StatusOK, models.SessionId{SessionId: session_id})
	}
	return ctx.NoContent(http.StatusBadRequest)
}
