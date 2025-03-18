-- +goose Up
-- +goose StatementBegin
ALTER TABLE Tweets ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Tweets DROP CONSTRAINT fk_user_id;
-- +goose StatementEnd
