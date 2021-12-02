package models

type User struct {

	UserId int32 `json:"user_id"`

	Name string `json:"name"`

	Subscription []int32 `json:"subscription"`

	SubscribedBy []int32 `json:"subscribed_by"`

	Posts []int32 `json:"posts"`
}
