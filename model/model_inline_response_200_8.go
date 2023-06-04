package model

type InlineResponse2008 struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data InlineResponse2008Data `json:"data"`
}
