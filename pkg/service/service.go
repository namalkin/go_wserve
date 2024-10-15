package service

import "github.com/namalkin/go_wserve/pkg/repository"

type Authorisation interface {
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
	return &Service{}
}
