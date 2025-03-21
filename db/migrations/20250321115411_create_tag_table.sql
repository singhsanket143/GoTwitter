-- +goose Up
-- +goose StatementBegin
CREATE TABLE Tag (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Tag;
-- +goose StatementEnd
