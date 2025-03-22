package db

import (
	"GoTwitter/errors"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type TweetTagsRepository interface {
	BulkCreate(context context.Context, tagIds []int64, tweetId int64) (bool, *errors.AppError)
}

type TweetTagsStore struct {
	db *sql.DB
}

func NewTweetTagsStore(db *sql.DB) *TweetTagsStore {
	return &TweetTagsStore{db}
}

func (s *TweetTagsStore) BulkCreate(ctx context.Context, tagIds []int64, tweetId int64) (bool, *errors.AppError) {

	if len(tagIds) == 0 {
		return false, errors.NewAppError(http.StatusBadRequest, "No tags provided", nil)
	}

	var placeholders []string
	var args []any

	for _, tagId := range tagIds {
		if tagId <= 0 {
			return false, errors.NewAppError(http.StatusBadRequest, "Invalid tag ID provided", nil)
		}
		placeholders = append(placeholders, "(?, ?)")
		args = append(args, tagId, tweetId)
	}

	query := fmt.Sprintf(`
		INSERT INTO tweet_tags (tag_id, tweet_id)
		VALUES %s
	`, strings.Join(placeholders, ", "))

	// Prepare the SQL statement for execution. This helps in optimizing the query execution
	// and prevents SQL injection by safely substituting the arguments.
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		// If there is an error preparing the statement, return an internal server error.
		return false, errors.NewAppError(http.StatusInternalServerError, "Error preparing query for tweet_tags", err)
	}
	// Ensure the prepared statement is closed after execution to free up resources.
	defer stmt.Close()

	// Execute the prepared statement with the provided arguments (tag IDs and tweet ID).
	_, execErr := stmt.ExecContext(ctx, args...)
	if execErr != nil {
		// Check if the error is related to foreign key constraints, indicating invalid IDs.
		if strings.Contains(execErr.Error(), "foreign key constraint") {
			return false, errors.NewAppError(http.StatusBadRequest, "Invalid tag ID or tweet ID", execErr)
		}
		// For other execution errors, return an internal server error.
		return false, errors.NewAppError(http.StatusInternalServerError, "Error executing query for tweet_tags", execErr)
	}

	return true, nil
}
