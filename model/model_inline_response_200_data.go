package model

type InlineResponse200Data struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int32  `json:"type"`
	Status      int32  `json:"status"`
}
