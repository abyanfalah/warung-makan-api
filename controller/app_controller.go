package controller

import (
	"net/http"
	"strings"
	"warung-makan/config"
	"warung-makan/manager"
	"warung-makan/middleware"
	"warung-makan/model"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ucMan  manager.UsecaseManager
	router *gin.Engine
}

func NewController(usecaseManager manager.UsecaseManager, router *gin.Engine) *Controller {
	controller := Controller{
		ucMan:  usecaseManager,
		router: router,
	}
	accessToken := authenticator.NewAccessToken(config.NewConfig().TokenConfig)
	tokenMdw := middleware.NewAuthTokenMiddleware(accessToken)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	test := router.Group("/test")

	// ======= GENERATE TOKEN
	test.POST("/generate_token", func(ctx *gin.Context) {
		var user model.User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			utils.JsonErrorBadRequest(ctx, err, "cannot bind struct")
		}

		token, err := accessToken.GenerateAccessToken(&user)
		if err != nil {
			utils.JsonErrorBadGateway(ctx, err, "cannot generate token")
			return
		}

		utils.JsonNamedDataMessageResponse(ctx, "token", token, "token generated")
	})

	// ======= REQUIRE TOKEN
	test.GET("/require_token", func(ctx *gin.Context) {
		var h middleware.FakeAuthHeader
		err := ctx.ShouldBindHeader(&h)
		if err != nil {
			utils.JsonErrorBadRequest(ctx, err, "cant bind header")
			return
		}

		utils.JsonNamedDataMessageResponse(ctx, "token", h.Authorization, "token received")
	})

	// ======= VERIFY TOKEN
	test.GET("/verify_token", func(ctx *gin.Context) {
		var h middleware.FakeAuthHeader
		err := ctx.ShouldBindHeader(&h)
		if err != nil {
			utils.JsonErrorBadRequest(ctx, err, "cant bind header")
			return
		}

		tokenString := strings.Replace(h.Authorization, "Bearer ", "", -1)
		if tokenString == "" {
			utils.JsonErrorBadGateway(ctx, nil, "token string became empty")
			ctx.Abort()
			return
		}

		mapClaim, err := accessToken.VerifyToken(tokenString)
		if err != nil {
			utils.JsonErrorBadGateway(ctx, err, "cant verify token")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":      "token verified",
			"token_string": tokenString,
			"map_claim":    mapClaim,
		})
	})

	protectedRoute := router.Group("/test/protected", tokenMdw.RequireToken())
	protectedRoute.GET("/secret_place", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "welcome to the secret place. your token is verified! You can now access all protected endpoints!")
	})

	return &controller
}
