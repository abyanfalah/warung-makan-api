package server

import (
	"strconv"
	"warung-makan/config"
	"warung-makan/controller"
	"warung-makan/manager"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type appServer struct {
	ucMan        manager.UsecaseManager
	engine       *gin.Engine
	config       config.Config
	tokenService authenticator.AccessToken
}

func NewAppServer() *appServer {
	config := config.NewConfig()
	infraMan := manager.NewInfraManager(config)
	repoMan := manager.NewRepoManager(infraMan)

	return &appServer{
		ucMan:        manager.NewUsecaseManager(repoMan),
		engine:       gin.Default(),
		config:       config,
		tokenService: authenticator.NewAccessToken(config.TokenConfig),
	}
}

func (a *appServer) initHandlers() {
	controller.NewController(a.ucMan, a.engine)

	controller.NewUserController(a.ucMan, a.engine)
	controller.NewMenuController(a.ucMan, a.engine)
	controller.NewTransactionController(a.ucMan, a.engine)
	controller.NewLoginController(a.ucMan, a.engine)
}

func (a *appServer) Run() {
	a.initHandlers()
	apiPort := a.config.ApiConfig.Port
	if apiPort == "" {
		apiPort = "8000"
	}

	for {
		err := a.engine.Run(":" + apiPort)
		if err != nil {
			apiPortInt, _ := strconv.Atoi(apiPort)
			apiPort = strconv.Itoa(apiPortInt + 1)
		} else {
			break
		}
	}

}
