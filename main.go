package main

import (
	"fmt"
	"os"
	"strings"
	"warung-makan/config"
	"warung-makan/server"
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

	fmt.Println("Setting finished")
	fmt.Println(strings.Repeat("=", 50))

}

func viewConfigs() {

	config := config.NewConfig()

	fmt.Println("configs: ")
	fmt.Println("db config:", config.DbConfig)
	fmt.Println("api config:", config.ApiConfig)

	fmt.Println(strings.Repeat("=", 50))

	fmt.Println()

}
