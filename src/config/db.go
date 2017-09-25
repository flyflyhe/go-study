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
		db.Dsn = "root:@127.0.0.1:3306/go"
	} else {
		db.Driver = "mysql"
		db.Username = "root"
		db.Password = "^YHNMJU&8ikm"
		db.Host = "127.0.0.1"
		db.Port = "3306"
		db.Database = "go"
		db.Dsn = "root:^YHNMJU&8ikm@127.0.0.1:3306/go"
	}
	return db
}
