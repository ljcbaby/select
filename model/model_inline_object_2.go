package model

type InlineObject2 struct {
	Name     string `json:"name"`
	RoleID   int32  `json:"roleID,omitempty"`
	Identify string `json:"identify"`
}
