package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

// GetHaiku - get_haiku
func (c *Container) GetHaiku(ctx echo.Context) error {

	var tmp_res models.InlineResponse2001

	haiku_id_str := ctx.Param("haiku_id")
	haiku_id, _ := strconv.Atoi(haiku_id_str)
	tmp_res.Haiku.HaikuId = int64(haiku_id)

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

	tmp_res.Author.Name, _ = c.RedisClient.Get(ctxBG, "user_id:"+author_id_str+":name").Result()

	subscription_id_str_list, _ := c.RedisClient.SMembers(ctxBG, "user_id:"+author_id_str+":subscription").Result()
	log.Println(subscription_id_str_list)
	for _, subscription_id_str := range subscription_id_str_list {
		subscription_id, _ := strconv.Atoi(subscription_id_str)
		tmp_res.Author.Subscription = append(tmp_res.Author.Subscription, int64(subscription_id))
	}

	subscribed_by_id_str_list, _ := c.RedisClient.SMembers(ctxBG, "user_id:"+author_id_str+":subscribed_by").Result()
	for _, subscribed_by_id_str := range subscribed_by_id_str_list {
		subscribed_by_id, _ := strconv.Atoi(subscribed_by_id_str)
		tmp_res.Author.SubscribedBy = append(tmp_res.Author.SubscribedBy, int64(subscribed_by_id))
	}

	author_haiku_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+author_id_str+":author_haiku_id_list", 0, -1).Result()
	for _, author_haiku_id_str := range author_haiku_id_str_list {
		author_haiku_id, _ := strconv.Atoi(author_haiku_id_str)
		tmp_res.Author.AuthorHaikuIdList = append(tmp_res.Author.AuthorHaikuIdList, int64(author_haiku_id))
	}

	timeline_haiku_id_str_list, _ := c.RedisClient.LRange(ctxBG, "user_id:"+author_id_str+":timeline_haiku_id_list", 0, -1).Result()
	for _, timeline_haiku_id_str := range timeline_haiku_id_str_list {
		timeline_haiku_id, _ := strconv.Atoi(timeline_haiku_id_str)
		tmp_res.Author.TimelineHaikuIdList = append(tmp_res.Author.TimelineHaikuIdList, int64(timeline_haiku_id))
	}
	log.Println(tmp_res)
	return ctx.JSON(http.StatusOK, tmp_res)
}
