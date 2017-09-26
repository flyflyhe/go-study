package mysql

import (
	"config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var l3db *sql.DB

func GetMysql() (*sql.DB, error) {
	var err error

	if db != nil {
		return db, err
	}
	config := config.GetMysql()
	db, err = sql.Open(config.Driver, config.Dsn)

	return db, err
}

func GetL3Mysql() (*sql.DB, error) {
	var err error
	if l3db != nil {
		return l3db, err
	}
	config := config.GetL3Mysql()
	l3db, err = sql.Open(config.Driver, config.Dsn)

	return l3db, err
}
