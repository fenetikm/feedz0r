-- +goose Up
CREATE TABLE feeds (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at INTEGER NOT NULL DEFAULT (unixepoch()),
  updated_at INTEGER NOT NULL DEFAULT (unixepoch()),
  name TEXT,
  url TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE feeds;
