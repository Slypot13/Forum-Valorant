package services

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"strings"
	"unicode"

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