package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/namalkin/go_wserve"
	"github.com/namalkin/go_wserve/pkg/repository"
)

const salt = "erijj4or-3j4or34r"

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user go_wserve.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New() // хеширование
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))) // хэш+соль пароля
}
