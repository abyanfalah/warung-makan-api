package main

import (
	"fmt"
	"os"
	"strings"
	"warung-makan/config"
	"warung-makan/server"
	"warung-makan/utils/authenticator"
)

const (
	DB_HOST   = "localhost"
	DB_PORT   = "5432"
	DB_USER   = "postgres"
	DB_PASS   = "12345"
	DB_NAME   = "warung_makan"
	DB_DRIVER = "postgres"

	API_HOST = "localhost"
	API_PORT = "8000"

	APP_NAME = "warung_makan_enigma"
)

func main() {
	setEnv()
	viewConfigs()

	server.NewAppServer().Run()

}

func setEnv() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Auto setting environment variables")
	fmt.Println("You can disable this feature on main.go")

	os.Setenv("DB_HOST", DB_HOST)
	os.Setenv("DB_PORT", DB_PORT)
	os.Setenv("DB_USER", DB_USER)
	os.Setenv("DB_PASS", DB_PASS)
	os.Setenv("DB_NAME", DB_NAME)
	os.Setenv("DB_DRIVER", DB_DRIVER)

	os.Setenv("API_HOST", API_HOST)
	os.Setenv("API_PORT", API_PORT)

	os.Setenv("APP_NAME", APP_NAME)

	fmt.Println("Setting finished")
	fmt.Println(strings.Repeat("=", 50))

}

func viewConfigs() {

	config := config.NewConfig()

	fmt.Println("configs: ")
	fmt.Println("db config:", config.DbConfig)
	fmt.Println("api config (port maybe auto set):", config.ApiConfig)

	fmt.Println(strings.Repeat("=", 50))

	fmt.Println()

}

func testVerifyToken() {
	config := config.NewConfig()

	accessToken := authenticator.NewAccessToken(config.TokenConfig)
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjYwMjg5NDMsImlhdCI6MTY2NjAyNTM0MywiaXNzIjoid2FydW5nX21ha2FuX2VuaWdtYSIsInVzZXJfaWQiOiIxOTRlZWM5NC1mNjkxLTRkOTgtYWU5Yi1jNGU2NTllNzcyNTciLCJ1c2VybmFtZSI6ImFub24ifQ.fO5gML3Mi2dTmgnAPwEL7w_gd6uBxiXZITQm1ZdWDJw"
	t, err := accessToken.VerifyToken(tokenString)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
}
