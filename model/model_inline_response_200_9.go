package model

type InlineResponse2009 struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data InlineResponse2009Data `json:"data"`
}
