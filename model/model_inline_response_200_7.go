package model

type InlineResponse2007 struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data InlineResponse2007Data `json:"data"`
}
