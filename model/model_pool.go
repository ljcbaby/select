package model

type Pool struct {
	Id int32 `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	// 1抽奖2分组3角色分组4抽签
	Type int32 `json:"type"`

	// 1未发布2抽签中3已完成
	Status int32 `json:"status"`

	Selections []InlineResponse2001DataSelections `json:"selections"`

	Groups []InlineResponse2001DataGroups `json:"groups,omitempty"`

	Roles []InlineResponse2001DataGroups `json:"roles,omitempty"`
}
