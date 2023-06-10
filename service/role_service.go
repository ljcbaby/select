package service

import (
	"github.com/ljcbaby/select/database"
	"github.com/ljcbaby/select/model"
)

type RoleService struct{}

func (s *RoleService) GetRoles(poolId int64, roles *[]model.GroupRole) error {
	rows, err := database.MySQL.Query("SELECT ID, name FROM roles WHERE PoolID = ?", poolId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var role model.GroupRole
		err := rows.Scan(&role.Id, &role.Name)
		if err != nil {
			return err
		}
		*roles = append(*roles, role)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if len(*roles) == 0 {
		*roles = []model.GroupRole{}
	}

	return nil
}

func (s *RoleService) CreateRole(poolId int64, c string) error {
	_, err := database.MySQL.Exec("INSERT INTO roles (PoolID, Name) VALUES (?, ?)", poolId, c)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) UpdateRole(id int64, c string) error {
	_, err := database.MySQL.Exec("UPDATE roles SET Name = ? WHERE ID = ?", c, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) DeleteRole(id int64) error {
	_, err := database.MySQL.Exec("DELETE FROM roles WHERE ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) VerifyRole(poolId int64, id int64) (bool, error) {
	var count int
	err := database.MySQL.QueryRow("SELECT COUNT(*) FROM roles WHERE PoolID = ? AND ID = ?", poolId, id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
