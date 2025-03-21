package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type TweetTagsRepository interface {
	BulkCreate(context context.Context, tagIds []int64, tweetId int64) (bool, error)
}

type TweetTagsStore struct {
	db *sql.DB
}

func NewTweetTagsStore(db *sql.DB) *TweetTagsStore {
	return &TweetTagsStore{db}
}

func (s *TweetTagsStore) BulkCreate(ctx context.Context, tagIds []int64, tweetId int64) (bool, error) {

	if len(tagIds) == 0 {
		return false, sql.ErrNoRows
	}

	var placeholders []string
	var args []interface{}

	for _, tagId := range tagIds {
		placeholders = append(placeholders, "(?, ?)")
		args = append(args, tagId, tweetId)
	}

	query := fmt.Sprintf(`
		INSERT INTO tweet_tags (tag_id, tweet_id)
		VALUES %s
	`, strings.Join(placeholders, ", "))

	_, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		return false, err
	}

	return true, nil
}
