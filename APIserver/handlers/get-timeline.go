package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// GetTimeline - timeline
func (c *Container) GetTimeline(ctx echo.Context) error {
	var payload models.InlineObject5
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}

	session_id := payload.SessionId
	start := payload.Start
	stop := payload.Stop
	if stop == 0 {
		stop = 30
	}

	user_id_str, err := c.RedisClient.Get(ctxBG, session_id+":linked_user_id").Result()
	if err != nil {
		return ctx.HTML(http.StatusBadRequest, "invalid session id")
	}
	var haiku_list []models.Haiku
	haiku_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+user_id_str+":timeline_haiku_id_list", start, stop).Result()
	log.Println(haiku_id_str_list)
	for _, haiku_id_str := range haiku_id_str_list {

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

/*
//あわれな旧実装くんを戒めとして残しておく。
func (c *Container) GetTimeline(ctx echo.Context) error {
	var payload models.SessionId

	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}
	session_id := payload.SessionId

	user_id_str, err := c.RedisClient.Get(ctxBG, session_id+":linked_user_id").Result()
	if err != nil {
		return ctx.HTML(http.StatusBadRequest, "invalid session id")
	}
	var haiku_list []models.Haiku
	subscription_id_str_list, err := c.RedisClient.SMembers(ctxBG, "user_id:"+user_id_str+"subscription").Result()
	for _, subscription_id_str := range subscription_id_str_list { //subscribeしているユーザのidから

		haiku_id_str_list, _ := c.RedisClient.SMembers(ctxBG, "user_id:"+subscription_id_str+":author_haiku_id_list").Result() //そのユーザのhaikuのidを読み込んで
		for _, haiku_id_str := range haiku_id_str_list {                                                                       //subscribeしているユーザが投稿した全てのhaikuのデータを取得

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
	}
	return ctx.JSON(http.StatusOK, haiku_list)
}
*/
