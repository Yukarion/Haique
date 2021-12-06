package models

type User struct {

	UserId int64 `json:"user_id"`

	Name string `json:"name"`

	Subscription []int64 `json:"subscription"`

	SubscribedBy []int64 `json:"subscribed_by"`

	HaikuIdList []int64 `json:"haiku_id_list"`
}
