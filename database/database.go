package database

import (
	"context"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ljcbaby/select/config"
	"github.com/redis/go-redis/v9"
)

var (
	MySQL *sql.DB
	Redis *redis.Client
)

func ConnectMySQL() error {
	conf := config.Conf.MySQL
	var err error
	MySQL, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	))
	if err != nil {
		return err
	}
	err = MySQL.Ping()
	return err
}

func ConnectRedis() error {
	conf := config.Conf.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		DB:   conf.Database,
	})
	_, err := Redis.Ping(context.Background()).Result()
	return err
}
