package server

import (
	"warung-makan/config"
	"warung-makan/manager"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type appServer struct {
	ucMan  manager.UsecaseManager
	engine gin.Engine
	port   string
}

func NewAppServer() *appServer {
	config := config.NewConfig()
	infraMan := manager.NewInfraManager(config)
	repoMan := manager.NewRepoManager(infraMan)

	return &appServer{
		ucMan:  manager.NewUsecaseManager(repoMan),
		engine: *gin.Default(),
		port:   config.ApiConfig.Port,
	}
}
