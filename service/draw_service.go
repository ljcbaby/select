package service

import (
	"github.com/ljcbaby/select/database"
	"github.com/ljcbaby/select/model"
)

type DrawService struct{}

func (s *DrawService) Draw(draws *[]model.Result) error {
	return nil
}

func (s *DrawService) GetDraws(poolID int64, draws *[]model.Result) error {
	rows, err := database.MySQL.Query("SELECT ID, Name, identify, roleName,uid From RS WHERE PoolID = ?", poolID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var draw model.Result
		err := rows.Scan(&draw.Id, &draw.Name, &draw.Identify, &draw.RoleName, &draw.UID)
		if err != nil {
			return err
		}
		*draws = append(*draws, draw)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if len(*draws) == 0 {
		*draws = []model.Result{}
	}

	return nil
}
