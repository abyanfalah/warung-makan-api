package manager

import (
	"warung-makan/controller"

	"github.com/gin-gonic/gin"
)

type controllerManager struct {
	usecaseManager
	engine *gin.Engine
}

type ControllerManager interface {
	UserController() controller.UserController
}

func (cm *controllerManager) UserController() controller.UserController {
	return *controller.NewUserController(cm.usecaseManager, cm.engine)
}

func NewControllerManager(ucman usecaseManager, engine *gin.Engine) ControllerManager {
	return &controllerManager{
		usecaseManager: ucman,
		engine:         engine,
	}
}
