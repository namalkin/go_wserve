package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/namalkin/go_wserve"
	"github.com/namalkin/go_wserve/pkg/repository"
)

const (
	salt       = "erijj4or-3j4or34r"
	signingKey = "fjijerwqdqw-dewfewf"
	TokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorisation
}

type tokenClasis struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user go_wserve.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClasis{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New() // хеширование
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))) // хэш+соль пароля
}
