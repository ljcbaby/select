package model

type InlineResponse200 struct {
	Code int32                   `json:"code"`
	Msg  string                  `json:"msg"`
	Data []InlineResponse200Data `json:"data"`
}
