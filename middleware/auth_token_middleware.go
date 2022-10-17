package middleware

import (
	"strings"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type FakeAuthHeader struct {
	Authorization string `header:"Authorization" binding:"required"`
}

type authHeader struct {
	Authorization string `header:"Authorization" binding:"required"`
}

type authTokenMiddleware struct {
	accessToken authenticator.AccessToken
}

type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
}

func (atm *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var headerAuth authHeader
		err := ctx.ShouldBindHeader(&headerAuth)
		if err != nil {
			utils.JsonErrorUnauthorized(ctx, err, "cant bind header")
			ctx.Abort()
			return
		}

		tokenString := strings.Replace(headerAuth.Authorization, "Bearer ", "", -1)
		if tokenString == "" {
			utils.JsonErrorBadGateway(ctx, nil, "token string became empty")
			ctx.Abort()
			return
		}

		_, err = atm.accessToken.VerifyToken(tokenString)
		if err != nil {
			utils.JsonErrorUnauthorized(ctx, err, "cannot verify token")
			ctx.Abort()
			return
		}
	}
}

func NewAuthTokenMiddleware(accessToken authenticator.AccessToken) AuthTokenMiddleware {
	return &authTokenMiddleware{
		accessToken: accessToken,
	}
}
