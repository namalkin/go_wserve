package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/namalkin/go_wserve"
)

type Authorisation interface {
	CreateUser(user go_wserve.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
