package models

type InlineResponse2001 struct {

	User User `json:"user"`

	Haikus []Haiku `json:"haikus"`
}
