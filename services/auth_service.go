package services

import (
	"errors"
	"net/http"

	"forum-valorant/config"

	"github.com/golang-jwt/jwt/v5"
)

func GetUserIdFromRequest(r *http.Request) (int, error) {
	cookie, err := r.Cookie("token")

	if err != nil {
		return 0, errors.New("utilisateur non connecté")
	}

	secret := config.GetRequiredEnv("JWT_SECRET")

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || token.Valid == false {
		return 0, errors.New("token invalide")
	}

	claims := token.Claims.(jwt.MapClaims)

	userIdFloat, ok := claims["user_id"].(float64)

	if !ok {
		return 0, errors.New("id utilisateur invalide")
	}

	userId := int(userIdFloat)

	return userId, nil
}

func GetUserRoleFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")

	if err != nil {
		return "", errors.New("utilisateur non connecté")
	}

	secret := config.GetRequiredEnv("JWT_SECRET")

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || token.Valid == false {
		return "", errors.New("token invalide")
	}

	claims := token.Claims.(jwt.MapClaims)

	role, ok := claims["role"].(string)

	if !ok {
		return "", errors.New("role invalide")
	}

	return role, nil
}