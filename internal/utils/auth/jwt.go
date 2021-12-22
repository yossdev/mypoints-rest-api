package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/configs"
	"time"
)

type Token struct {
	ID           uuid.UUID `json:"id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

func Sign(id uuid.UUID, claims jwt.MapClaims) Token {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(configs.Get().JwtTokenExpired).Unix()

	if claims["id"] == nil {
		return Token{}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// setup userdata
	var _, checkExp = claims["exp"]
	var _, checkIat = claims["exp"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}

	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}
	claims["token_type"] = "access_token"

	token.Claims = claims

	authToken := new(Token)
	tokenString, err := token.SignedString([]byte(configs.Get().JwtSecretKey))
	if err != nil {
		return Token{}
	}

	authToken.ID = id
	authToken.AccessToken = tokenString

	//create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenExpired := timeNow.Add(configs.Get().JwtRefreshExpired).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims

	refreshTokenString, err := refreshToken.SignedString([]byte(configs.Get().JwtSecretKey))

	if err != nil {
		return Token{}
	}
	authToken.RefreshToken = refreshTokenString

	return Token{
		ID:           authToken.ID,
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
	}
}
