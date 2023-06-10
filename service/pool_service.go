package service

import (
	"errors"

	"github.com/ljcbaby/select/database"
	"github.com/ljcbaby/select/model"
)

type PoolService struct{}

func (s *PoolService) GetPools(pools *[]model.PoolBase) error {
	rows, err := database.MySQL.Query("SELECT ID, Name, Description, type, status FROM pools")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var pool model.PoolBase
		err := rows.Scan(&pool.Id, &pool.Name, &pool.Description, &pool.Type, &pool.Status)
		if err != nil {
			return err
		}
		*pools = append(*pools, pool)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if len(*pools) == 0 {
		*pools = []model.PoolBase{}
	}

	return nil
}

func (s *PoolService) CreatePool(c model.PoolBase) (int64, error) {
	tx, err := database.MySQL.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO pools (Name, Description, Type) VALUES (?, ?, ?)", c.Name, c.Description, c.Type)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return lastInsertID, nil
}

// GetPool 获取pool
func (s *PoolService) GetPool(id int64, pool *model.Pool) error {
	return nil
}

// UpdatePool 更新pool
func (s *PoolService) UpdatePool(id int64, c model.PoolBase) error {
	return nil
}

// DeletePool 删除pool
func (s *PoolService) DeletePool(id int64) error {
	var count int = 0
	err := database.MySQL.QueryRow("SELECT COUNT(*) FROM pools WHERE ID = ?", id).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("noPool")
	}
	c, err := database.MySQL.Exec("DELETE FROM pools WHERE ID = ?", id)
	if err != nil {
		return err
	}
	_, err = c.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
