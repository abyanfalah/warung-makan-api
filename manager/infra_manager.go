package manager

import (
	"fmt"
	"warung-makan/config"

	"github.com/jmoiron/sqlx"
)

type infraManager struct {
	*sqlx.DB
	config.Config
}

type InfraManager interface {
	GetSqlDb() *sqlx.DB
}

func (i *infraManager) GetSqlDb() *sqlx.DB {
	return i.DB
}

func (i *infraManager) initDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", i.DbConfig.Host, i.DbConfig.User, i.DbConfig.Pass, i.DbConfig.DbName, i.DbConfig.Port)

	connection, err := sqlx.Connect(i.DbConfig.DbDriver, dsn)
	if err != nil {
		panic(err)
	}
	i.DB = connection
	fmt.Println("Connected to database ->", i.DbConfig.DbName)
}

func NewInfraManager(config config.Config) InfraManager {
	infraMan := new(infraManager)
	infraMan.Config = config
	infraMan.initDb()
	return infraMan
}
