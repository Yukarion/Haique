package models

type Haiku struct {

	HaikuId int32 `json:"haiku_id"`

	Author HaikuAuthor `json:"author"`

	Content ApiPostHaikuContent `json:"content"`

	Likes int32 `json:"likes"`

	CreatedAt string `json:"created_at"`
}
