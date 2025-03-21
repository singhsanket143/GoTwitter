-- +goose Up
-- +goose StatementBegin
CREATE TABLE Tweet_Tags (
    tweet_id BIGINT UNSIGNED NOT NULL,
    tag_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (tweet_id, tag_id),
    FOREIGN KEY (tweet_id) REFERENCES Tweets(id),
    FOREIGN KEY (tag_id) REFERENCES Tag(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE TWEET_TAGS;
-- +goose StatementEnd
