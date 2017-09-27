package config

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
	if TESTING {
		db.Driver = "mysql"
		db.Username = "root"
		db.Password = ""
		db.Port = "3306"
		db.Host = "127.0.0.1"
		db.Database = "go"
		db.Dsn = "root:@tcp(127.0.0.1:3306)/go?charset=utf8"
	} else {
		db.Driver = "mysql"
		db.Username = "root"
		db.Password = ""
		db.Host = "127.0.0.1"
		db.Port = "3306"
		db.Database = "go"
		db.Dsn = ""
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
		db.Driver = "mysql"
		db.Username = ""
		db.Password = ""
		db.Host = ""
		db.Port = "3306"
		db.Database = ""
		db.Dsn = ""
	}
	return db
}
