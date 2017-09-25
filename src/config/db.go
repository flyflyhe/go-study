package config

type Mysql struct {
	Driver   string
	Username string
	Password string
	Host     string
	Database string
}

func GetMysql() Mysql {
	var db Mysql
	if TESTING {
		db.Driver = "mysql"
		db.Username = "root"
		db.Password = ""
		db.Host = "127.0.0.1"
		db.Database = "go"
	} else {
		db.Driver = "mysql"
		db.Username = "root"
		db.Password = "^YHNMJU&8ikm"
		db.Host = "127.0.0.1"
		db.Database = "go"
	}
	return db
}
