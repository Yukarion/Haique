package models

type InlineResponse2001 struct {

	Author User `json:"author"`

	Haiku Haiku `json:"haiku"`
}
