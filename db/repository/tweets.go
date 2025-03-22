package db

import (
	"GoTwitter/dto"
	"GoTwitter/errors"
	"GoTwitter/models"
	"context"
	"database/sql"
	"net/http"
)

type TweetsRepository interface {
	Create(context context.Context, tweet *dto.CreateTweetDTO) (*models.Tweet, *errors.AppError)
	GetByID(context context.Context, id int64) (*models.Tweet, *errors.AppError)
	GetAll(context context.Context) ([]*models.Tweet, *errors.AppError)
	Update(context context.Context, id int64) *errors.AppError
	Delete(context context.Context, id int64) (bool, *errors.AppError)
}

type TweetsStore struct {
	db *sql.DB
}

func NewTweetsStore(db *sql.DB) *TweetsStore {
	return &TweetsStore{db}
}

func (s *TweetsStore) Create(ctx context.Context, tweet *dto.CreateTweetDTO) (*models.Tweet, *errors.AppError) {
	// Step 1: Insert the tweet
	query := `INSERT INTO tweets (tweet, user_id) VALUES (?, ?)`
	result, err := s.db.ExecContext(ctx, query, tweet.Tweet, tweet.UserId)
	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error inserting tweet", err)
	}

	// Step 2: Get the inserted tweet ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error fetching last insert ID", err)
	}
	Id := int64(id)

	// Step 3 (optional): Fetch the full row (to get created_at, updated_at)
	newtweet := &models.Tweet{}
	row := s.db.QueryRowContext(ctx, "SELECT id, tweet, user_id, created_at, updated_at FROM tweets WHERE id = ?", Id)
	if err := row.Scan(&newtweet.Id, &newtweet.Tweet, &newtweet.UserId, &newtweet.CreatedAt, &newtweet.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewAppError(http.StatusNotFound, "Tweet not found after insertion", err)
		}
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error fetching inserted tweet", err)
	}

	return newtweet, nil
}

func (s *TweetsStore) GetByID(ctx context.Context, id int64) (*models.Tweet, *errors.AppError) {

	query := `SELECT id, tweet, user_id, created_at, updated_at FROM tweets WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, id)

	tweet := &models.Tweet{}
	if err := row.Scan(&tweet.Id, &tweet.Tweet, &tweet.UserId, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewAppError(http.StatusNotFound, "Tweet not found", err)
		}
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error fetching tweet", err)
	}

	return tweet, nil
}

func (s *TweetsStore) GetAll(ctx context.Context) ([]*models.Tweet, *errors.AppError) {

	query := `SELECT id, tweet, user_id, created_at, updated_at FROM tweets`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error fetching tweets", err)
	}
	defer rows.Close()

	var tweets []*models.Tweet
	for rows.Next() {
		tweet := &models.Tweet{}
		if err := rows.Scan(&tweet.Id, &tweet.Tweet, &tweet.UserId, &tweet.CreatedAt, &tweet.UpdatedAt); err != nil {
			return nil, errors.NewAppError(http.StatusInternalServerError, "Error scanning tweets", err)
		}
		tweets = append(tweets, tweet)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error scanning tweets", err)
	}
	return tweets, nil
}

func (s *TweetsStore) Update(ctx context.Context, id int64) error {
	return nil
}

func (s *TweetsStore) Delete(ctx context.Context, id int64) (bool, *errors.AppError) {

	query := `DELETE FROM tweets WHERE id = ?`
	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, errors.NewAppError(http.StatusInternalServerError, "Error deleting tweet", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, errors.NewAppError(http.StatusInternalServerError, "Error checking rows affected", err)
	}

	if rowsAffected == 0 {
		return false, errors.NewAppError(http.StatusNotFound, "Tweet not found", nil)
	}

	return true, nil
}
