package models

type Haiku struct {
	HaikuId int64 `json:"haiku_id"`

	Author HaikuAuthor `json:"author"`

	Content ApiPostHaikuContent `json:"content"`

	Likes int32 `json:"likes"`

	// unixtime
	CreatedAt Int64 `json:"created_at"`
}
