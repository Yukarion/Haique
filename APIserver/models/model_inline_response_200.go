package models

type InlineResponse200 struct {

	User User `json:"user"`

	Haikus []Haiku `json:"haikus"`
}
