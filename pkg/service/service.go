package service

import (
	"github.com/namalkin/go_wserve"
	"github.com/namalkin/go_wserve/pkg/repository"
)

type Authorisation interface {
	CreateUser(user go_wserve.User) (int, error)
	GenerateToken(isername, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
	}
}
