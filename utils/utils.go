package utils

import (
	"encoding/json"
	"errors"
	"merchant-service/config"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	Secret string `json:"secret"`
	jwt.StandardClaims
}

func ResponseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application-json")
	json.NewEncoder(w).Encode(data)
}

func ValidateToken(signedToken string) (err error) {
	serviceConfig := config.GetServiceConfig()

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(serviceConfig.JwtSecret), nil
		},
	)
	if err != nil {
		return errors.New("tokan can not validate")
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.Secret != string(serviceConfig.JwtSecret) {
		return errors.New("wrong secret")
	}
	return
}
