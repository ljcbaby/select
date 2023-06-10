package service

import (
	"github.com/ljcbaby/select/database"
	"github.com/ljcbaby/select/model"
)

type GroupService struct{}

func (s *GroupService) GetGroups(poolId int64, groups *[]model.GroupRole) error {
	rows, err := database.MySQL.Query("SELECT ID, name FROM groups WHERE PoolID = ?", poolId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var group model.GroupRole
		err := rows.Scan(&group.Id, &group.Name)
		if err != nil {
			return err
		}
		*groups = append(*groups, group)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if len(*groups) == 0 {
		*groups = []model.GroupRole{}
	}

	return nil
}

func (s *GroupService) CreateGroup(poolId int64, c string) error {
	_, err := database.MySQL.Exec("INSERT INTO groups (PoolID, Name) VALUES (?, ?)", poolId, c)
	if err != nil {
		return err
	}
	return nil
}

func (s *GroupService) UpdateGroup(id int64, c string) error {
	_, err := database.MySQL.Exec("UPDATE groups SET Name = ? WHERE ID = ?", c, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *GroupService) DeleteGroup(id int64) error {
	_, err := database.MySQL.Exec("DELETE FROM groups WHERE ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *GroupService) VerifyGroup(poolId int64, id int64) (bool, error) {
	var count int
	err := database.MySQL.QueryRow("SELECT COUNT(*) FROM groups WHERE PoolID = ? AND ID = ?", poolId, id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
