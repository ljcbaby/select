package model

type InlineResponse2003 struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data InlineResponse2003Data `json:"data"`
}
