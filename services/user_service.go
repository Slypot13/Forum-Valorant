package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"strings"
	"unicode"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"forum-valorant/config"

	"forum-valorant/models"
	"forum-valorant/repositories"
	)

type UserService struct {
	userRepository *repositories.UserRepository
}

func InitUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Register(user models.User) error {

	if len(user.Password) < 12 {
		return errors.New("le mot de passe doit contenir au moins 12 caractères")
	}

	hasUpper := false
	hasSpecial := false

	for _, char := range user.Password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}

		if strings.ContainsAny(string(char), "!@#$%^&*()-_=+[]{};:,.<>?/") {
			hasSpecial = true
		}
	}

	if hasUpper == false {
		return errors.New("le mot de passe doit contenir au moins une majuscule")
	}

	if hasSpecial == false {
		return errors.New("le mot de passe doit contenir au moins un caractère spécial")
	}

	hash := sha512.Sum512([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])

	user.Role = "user"
	user.IsBanned = false

	return s.userRepository.CreateUser(user)
}

func (s *UserService) Login(identifier string, password string) (string, error) {
	user, err := s.userRepository.FindByIdentifier(identifier)

	if err != nil {
		return "", errors.New("identifiant ou mot de passe incorrect")
	}

	if user.IsBanned {
		return "", errors.New("ce compte est banni")
	}

	hash := sha512.Sum512([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	if hashedPassword != user.Password {
		return "", errors.New("identifiant ou mot de passe incorrect")
	}

	secret := config.GetRequiredEnv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("erreur lors de la création du token")
	}

	return tokenString, nil
}