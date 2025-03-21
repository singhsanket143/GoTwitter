package db

import (
	"GoTwitter/models"
	"context"
	"database/sql"
)

type UsersRepository interface {
	Create(context context.Context, user *models.User) error
}

type UsersStore struct {
	db *sql.DB
}

func NewUsersStore(db *sql.DB) *UsersStore {
	return &UsersStore{db}
}

func (s *UsersStore) Create(ctx context.Context, user *models.User) error {
	return nil
}
