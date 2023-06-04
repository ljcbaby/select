package model

type InlineResponse2005Data struct {
	Id        int32                           `json:"id"`
	Name      string                          `json:"name"`
	Identify  string                          `json:"identify"`
	Selection InlineResponse2004DataSelection `json:"selection"`
}
