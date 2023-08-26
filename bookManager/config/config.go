package config

type Config struct {
	Db struct {
		Host     string `env:"POSTGRES_HOST"`
		Port     int    `env:"POSTGRES_PORT"`
		Name     string `env:"POSTGRES_DB"`
		UserName string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
	}
}
