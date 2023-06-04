package model

type InlineResponse2001 struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data InlineResponse2001Data `json:"data"`
}
