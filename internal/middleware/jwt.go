package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yossdev/mypoints-rest-api/configs"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	"net/http"
	"strings"
)

// Package for handling Auth Middleware using JWT

// JwtVerifyAuth0TokenRSA func for verify jwt token with signed string RS265 from Auth0 App
//func JwtVerifyAuth0TokenRSA(c *fiber.Ctx) error {
//	JwtToken := strings.Replace(c.Get("Authorization"), fmt.Sprintf("%s ", "Bearer"), "", 1)
//
//	if JwtToken == "" {
//		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, errors.New(web.InvalidJwt), web.BadCredential)
//	}
//
//	req := new(http.Request)
//	req.Header = http.Header{}
//	req.Header.Set("Authorization", JwtToken)
//
//	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		} else {
//			return jwt.ParseRSAPublicKeyFromPEM([]byte(configs.Get().PublicKey))
//		}
//	})
//
//	if err != nil || !token.Valid {
//		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, errors.New(web.InvalidJwt), err)
//	}
//
//	return c.Next()
//}

// JwtVerifyToken func for verify jwt token with signed string HS265
func JwtVerifyToken(c *fiber.Ctx) error {
	jwtToken := strings.Replace(c.Get("Authorization"), fmt.Sprintf("%s ", "Bearer"), "", 1)

	if jwtToken == "" {
		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, errors.New(web.InvalidJwt), web.BadCredential)
	}

	req := new(http.Request)
	req.Header = http.Header{}
	req.Header.Set("Authorization", jwtToken)

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]

		if tokenType != "access_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(configs.Get().JwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return web.JsonErrorResponse(c, fiber.StatusUnauthorized, errors.New(web.InvalidJwt), err)
	}

	return c.Next()
}
