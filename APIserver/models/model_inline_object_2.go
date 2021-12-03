package models

type InlineObject2 struct {

	SessionId SessionId `json:"session_id"`

	Content ApiPostHaikuContent `json:"content"`
}
