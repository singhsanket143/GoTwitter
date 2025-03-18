package db

import (
	"context"
	"database/sql"
)

type UsersRepository interface {
	Create(context context.Context) error
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context) error {
	return nil
}
