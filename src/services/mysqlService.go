package services

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetClient() *sql.DB {
	db,err = sql.Open("mysql", "user:@/test")
}