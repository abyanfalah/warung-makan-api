package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifetime time.Duration
}

type Config struct {
	DbConfig
	ApiConfig
	TokenConfig
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

	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("APP_NAME"),
		JwtSignatureKey:     "asdf",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifetime: time.Minute * 30,
	}
}

func NewConfig() Config {
	conf := new(Config)
	conf.readConfig()
	return *conf
}
