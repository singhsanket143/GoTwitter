package db

import (
	"context"
	"database/sql"
)

type Tweet struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TweetsRepository interface {
	Create(context context.Context, tweet *Tweet) error
	GetByID(context context.Context, id int64) (*Tweet, error)
	GetAll(context context.Context) ([]*Tweet, error)
	Update(context context.Context, id int64) error
	Delete(context context.Context, id int64) error
}

type TweetsStore struct {
	db *sql.DB
}

func (s *TweetsStore) Create(ctx context.Context, tweet *Tweet) error {
	query := `INSERT INTO tweets (content, user_id) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err := s.db.QueryRowContext(
		ctx,
		query,
		tweet.Content,
		tweet.UserId,
	).Scan(&tweet.ID, &tweet.CreatedAt, &tweet.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *TweetsStore) GetByID(ctx context.Context, id int64) (*Tweet, error) {
	return nil, nil
}

func (s *TweetsStore) GetAll(ctx context.Context) ([]*Tweet, error) {
	return nil, nil
}

func (s *TweetsStore) Update(ctx context.Context, id int64) error {
	return nil
}

func (s *TweetsStore) Delete(ctx context.Context, id int64) error {
	return nil
}
