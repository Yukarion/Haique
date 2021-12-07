package handlers

import (
	"net/http"
	"strconv"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// GetTop - top
func (c *Container) GetTop(ctx echo.Context) error {
	var haiku_list []models.Haiku

	haiku_id_list, err := c.RedisClient.LRange(ctxBG, "global:top_haiku_id_list", 0, -1).Result()
	if err != nil {
		return err
	}
	for _, haiku_id_str := range haiku_id_list {
		if haiku_id_str == "-1" {
			break
		}
		var tmp_haiku models.Haiku

		id, _ := strconv.Atoi(haiku_id_str)
		tmp_haiku.HaikuId = int64(id)

		author_id_str, _ := c.RedisClient.Get(ctxBG, "haiku_id:"+haiku_id_str+":author_id").Result()
		author_id, _ := strconv.Atoi(author_id_str)
		tmp_haiku.AuthorId = int64(author_id)

		content, _ := c.RedisClient.LRange(ctxBG, "haiku_id:"+haiku_id_str+":content", 0, -1).Result()
		tmp_haiku.Content.First = content[0]
		tmp_haiku.Content.Second = content[1]
		tmp_haiku.Content.Third = content[2]
		tmp_haiku.Content.AuthorName = content[3]

		created_at_str, _ := c.RedisClient.Get(ctxBG, "haiku_id:"+haiku_id_str+":created_at").Result()
		created_at, _ := strconv.Atoi(created_at_str)
		tmp_haiku.CreatedAt = int64(created_at)

		haiku_list = append(haiku_list, tmp_haiku)
	}
	return ctx.JSON(http.StatusOK, haiku_list)
}
