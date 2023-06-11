package service

import (
	"database/sql"

	"github.com/ljcbaby/select/database"
	"github.com/ljcbaby/select/model"
)

type SelectionService struct{}

func (s *SelectionService) GetSelectionsByOrder(poolId int64, selections *[]model.Results) error {
	rows, err := database.MySQL.Query("SELECT ID, Name FROM selections WHERE PoolID = ?", poolId)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		var selection model.Results
		err := rows.Scan(&selection.Id, &selection.Name)
		if err != nil {
			return err
		}
		*selections = append(*selections, selection)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if len(*selections) == 0 {
		*selections = []model.Results{}
	}

	return nil
}

func (s *SelectionService) GetSelections(poolId int64, selections *[]model.Selection) error {
	var t int
	err := database.MySQL.QueryRow("SELECT type FROM pools WHERE ID = ?", poolId).Scan(&t)
	if err != nil {
		return err
	}
	var rows *sql.Rows
	if t != 3 {
		rows, err = database.MySQL.Query("SELECT ID, Number, Name FROM selections WHERE poolID = ?", poolId)
	} else {
		rows, err = database.MySQL.Query("SELECT ID, Number, GroupID, RoleID FROM selections WHERE poolID = ?", poolId)
	}
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var selection model.Selection
		var err error
		if t != 3 {
			err = rows.Scan(&selection.Id, &selection.Number, &selection.Name)
		} else {
			err = rows.Scan(&selection.Id, &selection.Number, &selection.GroupID, &selection.RoleID)
		}
		if err != nil {
			return err
		}
		*selections = append(*selections, selection)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if len(*selections) == 0 {
		*selections = []model.Selection{}
	}

	return nil
}

func (s *SelectionService) CreateSelection(poolId int64, c model.Selection) error {
	var err error
	if c.GroupID == 0 && c.RoleID == 0 {
		_, err = database.MySQL.Exec("INSERT INTO selections (PoolID, Number, Name) VALUES (?, ?, ?)", poolId, c.Number, c.Name)
	} else {
		_, err = database.MySQL.Exec("INSERT INTO selections (PoolID, Number, Name, GroupID, RoleID) VALUES (?, ?, ?, ?, ?)", poolId, c.Number, c.Name, c.GroupID, c.RoleID)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *SelectionService) UpdateSelection(id int64, c model.Selection) error {
	var err error
	if c.GroupID == 0 && c.RoleID == 0 {
		_, err = database.MySQL.Exec("UPDATE selections SET Number = ?, Name = ? WHERE ID = ?", c.Number, c.Name, id)
	} else {
		_, err = database.MySQL.Exec("UPDATE selections SET Number = ?, Name = ?, GroupID = ?, RoleID = ? WHERE ID = ?", c.Number, c.Name, c.GroupID, c.RoleID, id)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *SelectionService) DeleteSelection(id int64) error {
	_, err := database.MySQL.Exec("DELETE FROM selections WHERE ID = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SelectionService) VerifySelection(poolID int64, id int64) (bool, error) {
	var count int
	err := database.MySQL.QueryRow("SELECT COUNT(*) FROM selections WHERE PoolID = ? AND ID = ?", poolID, id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
