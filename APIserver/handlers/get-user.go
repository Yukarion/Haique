package handlers

import (
	"net/http"
	"strconv"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// GetUser - user_info
func (c *Container) GetUser(ctx echo.Context) error {

	var tmp_res models.InlineResponse200

	user_id_str := ctx.Param("user_id")
	user_id, _ := strconv.Atoi(user_id_str)
	tmp_res.User.UserId = int64(user_id)

	tmp_res.User.Name, _ = c.RedisClient.Get(ctxBG, "user_id:"+user_id_str+":name").Result()

	subscription_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+user_id_str+":subscription", 0, -1).Result()
	for _, subscription_id_str := range subscription_id_str_list {
		subscription_id, _ := strconv.Atoi(subscription_id_str)
		tmp_res.User.Subscription = append(tmp_res.User.Subscription, int64(subscription_id))
	}

	subscribed_by_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+user_id_str+":subscribed_by", 0, -1).Result()
	for _, subscribed_by_id_str := range subscribed_by_id_str_list {
		subscribed_by_id, _ := strconv.Atoi(subscribed_by_id_str)
		tmp_res.User.SubscribedBy = append(tmp_res.User.SubscribedBy, int64(subscribed_by_id))
	}

	author_haiku_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+user_id_str+":haiku_id_list", 0, -1).Result()
	for _, author_haiku_id_str := range author_haiku_id_str_list {
		author_haiku_id, _ := strconv.Atoi(author_haiku_id_str)
		tmp_res.User.AuthorHaikuIdList = append(tmp_res.User.AuthorHaikuIdList, int64(author_haiku_id))
	}

	timeline_haiku_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+user_id_str+":haiku_id_list", 0, -1).Result()
	for _, timeline_haiku_id_str := range timeline_haiku_id_str_list {
		timeline_haiku_id, _ := strconv.Atoi(timeline_haiku_id_str)
		tmp_res.User.TimelineHaikuIdList = append(tmp_res.User.TimelineHaikuIdList, int64(timeline_haiku_id))
	}

	haiku_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+user_id_str+"haiku_id_list", 0, -1).Result()
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

		tmp_res.Haikus = append(tmp_res.Haikus, tmp_haiku)
	}

	return ctx.JSON(http.StatusOK, tmp_res)
}
