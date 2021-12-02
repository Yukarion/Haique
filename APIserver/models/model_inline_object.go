package models

type InlineObject struct {

	SessionId string `json:"session_id"`

	Content ApiPostHaikuContent `json:"content"`
}
