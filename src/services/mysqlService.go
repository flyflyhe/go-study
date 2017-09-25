package services

import (
	"config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysql() (*sql.DB, error) {
	config := config.GetMysql()
	db, err := sql.Open(config.Driver, config.Dsn)

	return db, err
}
