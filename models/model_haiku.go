package models

import (
	"time"
)

type Haiku struct {

	HaikuId int64 `json:"haiku_id"`

	Content ApiPostHaikuContent `json:"content"`

	// いいね数（実装しないかも？）
	Likes int32 `json:"likes"`

	CreatedAt time.Time `json:"created_at"`
}
