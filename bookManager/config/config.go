package config

type Config struct {
	Db struct {
		Host     string `env:"database_host" env-default:"localhost"`
		Port     int    `env:"database_port" env-default:"5432"`
		Name     string `env:"database_name" env-default:"yourDbName"`
		UserName string `env:"database_user" env-default:"yourUserName"`
		Password string `env:"database_password" env-default:"yourPassword"`
	}
}
