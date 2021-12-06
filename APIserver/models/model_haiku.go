package models

type Haiku struct {

	HaikuId int64 `json:"haiku_id"`

	AuthorId int64 `json:"author_id"`

	Content HaikuContent `json:"content"`

	Likes int32 `json:"likes"`

	// unixtime
	CreatedAt int64 `json:"created_at"`
}
