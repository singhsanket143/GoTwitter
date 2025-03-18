package db

import "database/sql"

type Storage struct {
	TweetsRepository TweetsRepository
	UsersRepository  UsersRepository
}

func NewMySQLStorage(db *sql.DB) Storage {
	return Storage{
		TweetsRepository: &TweetsStore{db},
		UsersRepository:  &UsersStore{db},
	}
}
