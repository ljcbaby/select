package model

type InlineResponse2001Data struct {
	Id          int32                              `json:"id"`
	Name        string                             `json:"name"`
	Description string                             `json:"description"`
	Type        int32                              `json:"type"`
	Status      int32                              `json:"status"`
	Selections  []InlineResponse2001DataSelections `json:"selections"`
	Groups      []InlineResponse2001DataGroups     `json:"groups,omitempty"`
	Roles       []InlineResponse2001DataGroups     `json:"roles,omitempty"`
}
