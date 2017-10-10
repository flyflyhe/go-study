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

var dir string

func init() {
	dir, _ = filepath.Abs("./config")
}

func GetMysql() Mysql {
	var db Mysql
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
		parseDb(dir+"/l3db_dev.json", &db)
	} else {
		parseDb(dir+"/l3db_prod.json", &db)
	}
	return db
}
