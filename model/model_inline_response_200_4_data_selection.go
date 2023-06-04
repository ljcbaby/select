package model

type InlineResponse2004DataSelection struct {
	Id      int32  `json:"id"`
	Name    string `json:"name,omitempty"`
	GroupID int32  `json:"groupID,omitempty"`
	RoleID  int32  `json:"roleID,omitempty"`
}
