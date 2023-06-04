package model

type InlineResponse2005 struct {
	Code int32                    `json:"code"`
	Msg  string                   `json:"msg"`
	Data []InlineResponse2005Data `json:"data"`
}
