-- +goose Up
-- +goose StatementBegin
ALTER TABLE Tag ADD CONSTRAINT unique_tag_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Tag DROP CONSTRAINT unique_tag_name;
-- +goose StatementEnd
