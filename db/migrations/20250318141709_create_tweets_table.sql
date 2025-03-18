-- +goose Up
-- +goose StatementBegin
CREATE TABLE Tweets (
 id SERIAL PRIMARY KEY,
 user_id BIGINT UNSIGNED NOT NULL,
 tweet VARCHAR(255) NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- FOREIGN KEY (user_id) REFERENCES Users(id) not adding now will show in next migration
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Tweets;
-- +goose StatementEnd
