package models

type User struct {

	UserId int64 `json:"user_id"`

	Subscription []int64 `json:"subscription"`

	// タイムラインをサクッとつくるためにUserがポストのリストを持つ
	Posts []int64 `json:"posts"`
}
