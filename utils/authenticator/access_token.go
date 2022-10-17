package authenticator

import (
	"time"
	"warung-makan/config"
	"warung-makan/model"

	"github.com/golang-jwt/jwt/v4"
)

type accessToken struct {
	config config.TokenConfig
}

type AccessToken interface {
	GenerateAccessToken(user *model.User) (string, error)
}

func (at *accessToken) GenerateAccessToken(user *model.User) (string, error) {
	now := time.Now().UTC()
	end := now.Add(at.config.AccessTokenLifetime)

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    at.config.ApplicationName,
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
		Id:       user.Id,
		Username: user.Username,
	}

	token := jwt.NewWithClaims(
		at.config.JwtSigningMethod,
		claims,
	)

	return token.SignedString([]byte(at.config.JwtSignatureKey))
}

func NewAccessToken(config config.TokenConfig) AccessToken {
	return &accessToken{
		config: config,
	}
}
