package model

type InlineResponse2004Data struct {
	Name      string                          `json:"name"`
	Identify  string                          `json:"identify"`
	Selection InlineResponse2004DataSelection `json:"selection"`
}
