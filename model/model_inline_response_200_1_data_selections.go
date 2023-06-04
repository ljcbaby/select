package model

type InlineResponse2001DataSelections struct {
	Id      int32  `json:"id"`
	Number  int32  `json:"number"`
	Name    string `json:"name,omitempty"`
	GroupID int32  `json:"groupID,omitempty"`
	RoleID  int32  `json:"roleID,omitempty"`
}
