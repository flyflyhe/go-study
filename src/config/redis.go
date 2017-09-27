package config

type Redis struct {
	Host     string
	Port     string
	Password string
	Select   int
}

func GetRedis() Redis {
	var config Redis
	if TESTING {
		config.Host = "127.0.0.1"
		config.Port = "6379"
		config.Password = "123456"
		config.Select = 0
	} else {
		config.Host = "127.0.0.1"
		config.Port = "6379"
		config.Password = ""
		config.Select = 0
	}
	return config
}
