package model

type PoolBase struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int64  `json:"type"`
	Status      int64  `json:"status"`
}

type Selection struct {
	Id      int64  `json:"id"`
	Number  int64  `json:"number"`
	Name    string `json:"name,omitempty"`
	GroupID int64  `json:"groupID,omitempty"`
	RoleID  int64  `json:"roleID,omitempty"`
}

type GroupRole struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Pool struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        int64       `json:"type"`
	Status      int64       `json:"status"`
	Selections  []Selection `json:"selections"`
	Groups      []GroupRole `json:"groups,omitempty"`
	Roles       []GroupRole `json:"roles,omitempty"`
}

type Result struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Identify  string    `json:"identify"`
	Selection Selection `json:"selection"`
	RoleName  string    `json:"roleName"`
	UID       int64
}

type Results struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	GroupName string   `json:"groupName"`
	Result    []Result `json:"result"`
}
