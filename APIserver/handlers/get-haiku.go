package handlers

import (
	"net/http"
	"strconv"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// GetHaiku - get_haiku
func (c *Container) GetHaiku(ctx echo.Context) error {

	var tmp_res models.InlineResponce2001

	haiku_id_str := ctx.Param("haiku_id")
	haiku_id, _ := strconv.Atoi(haiku_id_str)
	tmp_res.Haiku.HaikuId = int64(id)

	author_id_str, err := c.RedisClient.Get(ctxBG, "haiku_id:"+haiku_id_str+":author_id").Result()
	if err != nil {
		return ctx.HTML(http.StatusBadRequest, "invalid haiku id")
	}
	author_id, _ := strconv.Atoi(author_id_str)
	tmp_res.Haiku.AuthorId = int64(author_id)
	tmp_res.Author.UserId = int64(author_id)

	content, _ := c.RedisClient.LRange(ctxBG, "haiku_id:"+haiku_id_str+":content", 0, -1).Result()
	tmp_res.Haiku.Content.First = content[0]
	tmp_res.Haiku.Content.Second = content[1]
	tmp_res.Haiku.Content.Third = content[2]
	tmp_res.Haiku.Content.AuthorName = content[3]

	created_at_str, _ := c.RedisClient.Get(ctxBG, "haiku_id:"+haiku_id_str+":created_at").Result()
	created_at, _ := strconv.Atoi(created_at_str)
	tmp_res.Haiku.CreatedAt = int64(created_at)

	tmp_res.Author.Name = c.RedisClient.Get(ctxBG, "user_id:"+author_id_str+":name").Result()

	subscription_id_str_list := c.RedisClient.Get(ctxBG, "user_id:"+author_id_str+":subscription").Result()
	var tmp_sub_id_list []int
	for _, subscription_id_str := range subscription_id_str_list {
		subscription_id, _ := strconv.Atoi(subscription_id_str)
		tmp_sub_id_list = append(temp_sub_id_list, int64(subscription_id))
	}
	tmp_res.Author.Subscription = temp_sub_id_list

	subscribed_by_id_str_list := c.RedisClient.Get(ctxBG, "user_id:"+author_id_str+":subscribed_by").Result()
	var tmp_subscribed_by_id_list []int
	for _, subscribed_by_id_str := range subscribed_by_id_str_list {
		subscribed_by_id, _ := strconv.Atoi(subscribed_by_id_str)
		tmp_subscribed_by_id_list = append(temp_subscribed_by_id_list, int64(subscribed_by_id))
	}
	tmp_res.Author.SubscribedBy = tmp_subscribed_by_id_list

	author_haiku_id_str_list := c.RedisClient.Get(ctxBG, "user_id:"+author_id_str+":haiku_id_list").Result()
	var tmp_haiku_id_list []int
	for _, author_haiku_id_str := range author_haiku_id_str_list {
		author_haiku_id, _ := strconv.Atoi(author_haiku_id_str)
		tmp_haiku_id_list = append(temp_haiku_id_list, int64(author_haiku_id))
	}
	tmp_res.Author.AuthorHaikuIdList = tmp_haiku_id_list

	timeline_haiku_id_str_list := c.RedisClient.Get(ctxBG, "user_id:"+author_id_str+":timeline_haiku_id_list").Result()
	var tmp_timeline_id_list []int
	for _, timeline_haiku_id_str := range timeline_haiku_id_str_list {
		timeline_haiku_id, _ := strconv.Atoi(timeline_haiku_id_str)
		tmp_timeline_id_list = append(temp_timeline_id_list, int64(timeline_haiku_id))
	}
	tmp_res.Author.TimelineHaikuIdList = temp_timeline_id_list

	return ctx.JSON(http.StatusOK, tmp_res)
}
