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
		db.Username = "l3dbmasteruser"
		db.Password = "qn4u8b5HW1Xa9L3f"
		db.Host = "rm-wz9xl80aji73zkk7e.mysql.rds.aliyuncs.com"
		db.Port = "3306"
		db.Database = "lonlife"
		db.Dsn = "l3dbmasteruser:qn4u8b5HW1Xa9L3f@tcp(rm-wz9xl80aji73zkk7e.mysql.rds.aliyuncs.com:3306)/lonlife?charset=utf8"
	}
	return db
}
