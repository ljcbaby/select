package model

type InlineResponse2004 struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data InlineResponse2004Data `json:"data"`
}
