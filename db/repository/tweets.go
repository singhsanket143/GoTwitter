package db

import (
	"GoTwitter/models"
	"context"
	"database/sql"
)

type TweetsRepository interface {
	Create(context context.Context, tweet *models.Tweet) error
	GetByID(context context.Context, id int64) (*models.Tweet, error)
	GetAll(context context.Context) ([]*models.Tweet, error)
	Update(context context.Context, id int64) error
	Delete(context context.Context, id int64) error
}

type TweetsStore struct {
	db *sql.DB
}

func NewTweetsStore(db *sql.DB) *TweetsStore {
	return &TweetsStore{db}
}

func (s *TweetsStore) Create(ctx context.Context, tweet *models.Tweet) error {
	// Step 1: Insert the tweet
	query := `INSERT INTO tweets (tweet, user_id) VALUES (?, ?)`
	result, err := s.db.ExecContext(ctx, query, tweet.Tweet, tweet.UserId)
	if err != nil {
		return err
	}

	// Step 2: Get the inserted tweet ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	tweet.Id = int64(id)

	// Step 3 (optional): Fetch the full row (to get created_at, updated_at)
	row := s.db.QueryRowContext(ctx, "SELECT created_at, updated_at FROM tweets WHERE id = ?", tweet.Id)
	if err := row.Scan(&tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (s *TweetsStore) GetByID(ctx context.Context, id int64) (*models.Tweet, error) {
	return nil, nil
}

func (s *TweetsStore) GetAll(ctx context.Context) ([]*models.Tweet, error) {
	return nil, nil
}

func (s *TweetsStore) Update(ctx context.Context, id int64) error {
	return nil
}

func (s *TweetsStore) Delete(ctx context.Context, id int64) error {
	return nil
}
