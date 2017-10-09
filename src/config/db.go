package config

import (
	"fmt"
	"path/filepath"
)

type Mysql struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
	Dsn      string
}

func GetMysql() Mysql {
	var db Mysql
	dir, _ := filepath.Abs("./config")
	if TESTING {
		err := parseDb(dir+"/db_dev.json", &db)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		parseDb(dir+"/db_prod.json", &db)
	}
	return db
}

func GetL3Mysql() Mysql {
	var db Mysql
	if TESTING {
		db.Driver = "mysql"
		db.Username = "root"
		db.Password = ""
		db.Port = "3306"
		db.Host = "127.0.0.1"
		db.Database = "go"
		db.Dsn = "root:@tcp(127.0.0.1:3306)/go?charset=utf8"
	} else {
	}
	return db
}
