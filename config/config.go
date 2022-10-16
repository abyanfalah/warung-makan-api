package config

import "os"

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	DbName   string
	DbDriver string
}

type ApiConfig struct {
	Host string
	Port string
}

type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) readConfig() {
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		DbName:   os.Getenv("DB_NAME"),
		DbDriver: os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		Host: os.Getenv("API_HOST"),
		Port: os.Getenv("API_PORT"),
	}
}

func NewConfig() Config {
	conf := new(Config)
	conf.readConfig()
	return *conf
}
