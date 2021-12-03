package models

import (
	"github.com/google/uuid"
)

type SessionId struct {

	// 言語固有のUUID型を使う
	Id uuid.UUID `json:"id"`
}
