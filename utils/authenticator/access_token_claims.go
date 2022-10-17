package authenticator

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	jwt.StandardClaims
	Id       string `json:"user_id"`
	Username string `json:"username"`
}
