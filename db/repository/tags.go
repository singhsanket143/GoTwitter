package db

import (
	"GoTwitter/errors"
	"GoTwitter/models"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type TagsRepository interface {
	BulkCreate(context context.Context, tag []string) ([]*models.Tag, *errors.AppError)
}

type TagsStore struct {
	db *sql.DB
}

func NewTagsStore(db *sql.DB) *TagsStore {
	return &TagsStore{db}
}

func (s *TagsStore) BulkCreate(ctx context.Context, tags []string) ([]*models.Tag, *errors.AppError) {

	if len(tags) == 0 {
		return nil, errors.NewAppError(http.StatusBadRequest, "No tags provided", nil)
	}

	var placeholders []string
	var args []any

	for _, tag := range tags {
		placeholders = append(placeholders, "(?)")
		args = append(args, tag)
	}

	query := fmt.Sprintf(`
		INSERT IGNORE INTO tag (name)
		VALUES %s
	`, strings.Join(placeholders, ", "))

	_, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error creating tags", err)
	}

	selectQuery := fmt.Sprintf(`
		SELECT id, name
		FROM tag
		WHERE name IN (%s)
	`, strings.Join(placeholders, ", "))

	rows, err := s.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, "Error fetching newly created tags", err)
	}
	defer rows.Close()

	var result []*models.Tag
	for rows.Next() {
		tag := &models.Tag{}
		if err := rows.Scan(&tag.Id, &tag.Name); err != nil {
			return nil, errors.NewAppError(http.StatusInternalServerError, "Error scanning tags", err)
		}
		result = append(result, tag)
	}

	return result, nil

}
