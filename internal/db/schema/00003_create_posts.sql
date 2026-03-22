-- +goose Up
CREATE TABLE posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at INTEGER NOT NULL DEFAULT (unixepoch()),
  updated_at INTEGER NOT NULL DEFAULT (unixepoch()),
  title TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  description TEXT,
  published_at INTEGER NOT NULL,
  feed_id INTEGER NOT NULL
);

-- +goose Down
DROP TABLE posts;
