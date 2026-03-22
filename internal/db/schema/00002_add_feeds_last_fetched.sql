-- +goose Up
ALTER TABLE feeds
ADD COLUMN last_fetched_at INTEGER;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_fetched_at;

