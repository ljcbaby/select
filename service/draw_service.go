package service

import (
	"database/sql"
	"errors"
	"math/rand"

	"github.com/ljcbaby/select/database"
	"github.com/ljcbaby/select/model"
)

type DrawService struct{}

func (s *DrawService) Draw(req model.Draw, result *model.Result) error {
	var c int64 = 0
	err := database.MySQL.QueryRow("SELECT COUNT(*) FROM result WHERE poolID = ? AND userIdentify = ?", req.PoolId, req.Identify).Scan(&c)
	if err != nil {
		return err
	}
	if c != 0 {
		return errors.New("hasDrown")
	}

	var flag bool = true
	var t int
	err = database.MySQL.QueryRow("SELECT type FROM pools WHERE ID = ? ", req.PoolId).Scan(&t)
	if err != nil {
		return err
	}
	for flag {
		var rows *sql.Rows
		var err error
		tx, err := database.MySQL.Begin()
		if err != nil {
			return err
		}
		if t != 3 {
			rows, err = tx.Query("SELECT ID, number-used L from selections where poolID = ? and number != used ", req.PoolId)
		} else {
			rows, err = tx.Query("SELECT ID, number-used L from selections where poolID = ? and roleID = ? and number != used ", req.PoolId, req.RoleId)
		}
		if err != nil {
			return err
		}
		defer rows.Close()
		var ss []struct {
			ID int64
			L  int
		}
		var sum int = 0
		for rows.Next() {
			var s struct {
				ID int64
				L  int
			}
			err := rows.Scan(&s.ID, &s.L)
			if err != nil {
				return err
			}
			ss = append(ss, s)
			sum += s.L
		}
		if err := rows.Err(); err != nil {
			return err
		}
		if sum == 0 {
			if t == 3 {
				var count int
				err := tx.QueryRow("SELECT count(*) FROM selections WHERE poolID = ? and number != used", req.PoolId).Scan(&count)
				if err != nil {
					return err
				}
				if count != 0 {
					return errors.New("roleFinished")
				}
			}
			_, err = tx.Exec("UPDATE pools SET status = 3 WHERE ID = ?", req.PoolId)
			if err != nil {
				return err
			}
			err = tx.Commit()
			if err != nil {
				return err
			}
			return errors.New("finished")
		}
		var r int = 0
		r = rand.Intn(sum)
		for _, s := range ss {
			r -= s.L
			if r < 0 {
				c = s.ID
				break
			}
		}
		_, err = tx.Exec("UPDATE selections SET used = used + 1 WHERE ID = ?", c)
		if err != nil {
			tx.Rollback()
			return err
			// continue
		}
		_, err = tx.Exec("INSERT INTO result (poolID, selectID, name, userIdentify) VALUES(?, ?, ?, ?)", req.PoolId, c, req.Name, req.Identify)
		if err != nil {
			tx.Rollback()
			return err
			// continue
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
		flag = false
	}

	result.Name = req.Name
	result.Identify = req.Identify
	result.Selection.Id = c
	if t != 3 {
		err = database.MySQL.QueryRow("SELECT name FROM selections WHERE ID = ?", c).Scan(&result.Selection.Name)
		if err != nil {
			return err
		}
	} else {
		err = database.MySQL.QueryRow("SELECT name FROM `groups` WHERE ID = (SELECT groupID FROM selections WHERE ID = ?)", c).Scan(&result.Selection.GroupName)
		if err != nil {
			return err
		}
		err = database.MySQL.QueryRow("SELECT name FROM roles WHERE ID = (SELECT roleID FROM selections WHERE ID = ?)", c).Scan(&result.Selection.RoleName)
		if err != nil {
			return err
		}
	}

	err = database.MySQL.QueryRow("SELECT count(*) FROM selections WHERE poolID = ? and number != used", req.PoolId).Scan(&c)
	if err != nil {
		return err
	}
	if c == 0 {
		_, err = database.MySQL.Exec("UPDATE pools SET status = 3 WHERE ID = ?", req.PoolId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DrawService) GetDraws(poolID int64, draws *[]model.Result) error {
	var t int
	err := database.MySQL.QueryRow("SELECT type FROM pools WHERE ID = ? ", poolID).Scan(&t)
	if err != nil {
		return err
	}
	var rows *sql.Rows
	if t != 3 {
		rows, err = database.MySQL.Query("SELECT ID, Name, userIdentify, selectID From result WHERE PoolID = ?", poolID)
	} else {
		rows, err = database.MySQL.Query("SELECT ID, Name, identify, groupID, roleName From RS WHERE PoolID = ?", poolID)
	}
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var draw model.Result
		var err error
		if t != 3 {
			err = rows.Scan(&draw.Id, &draw.Name, &draw.Identify, &draw.UID)
		} else {
			err = rows.Scan(&draw.Id, &draw.Name, &draw.Identify, &draw.UID, &draw.RoleName)
		}
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
