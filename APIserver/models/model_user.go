package models

type User struct {

	UserId int64 `json:"user_id"`

	Name string `json:"name"`

	Subscription []int64 `json:"subscription"`

	SubscribedBy []int64 `json:"subscribed_by"`

	AuthorHaikuIdList []int64 `json:"author_haiku_id_list"`

	TimelineHaikuIdList []int64 `json:"timeline_haiku_id_list"`
}
