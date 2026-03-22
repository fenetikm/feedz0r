# TODO.md

## Now
- [ ] Don't put the ID in parens, confusing when post title has parens
- [ ] Output number of posts for a feed via `list`
- [ ] Fetch, refresh all feeds
- [ ] Update README.md

## Done
- [X] Read config
- [X] `Help` command (but not output anything)
- [X] Shift help to it's own file etc.
- [X] `Add` command (don't actually add)
- [X] Add a feed:
    - [X] Take in the arg, print
    - [X] SQLite using config, connect
    - [X] Config file for db config
    - [X] Define schema for a feed
    - [X] Add to db
    - [X] Test, manual
- [X] Fetch a feed, RSS
- [X] Store feed posts
- [X] Basic daemon going
- [X] Don't fetch if it has been a while
- [X] List of feeds
- [X] Get a list of entries for a feed
- [X] Get post by ID
- [X] Output handling
- [X] Try using FZF, raw output

## Soon
- `help <cmd>` to run `Help` func
- Config file specifies location of db e.g. `~/.fz/fz.db`
- Have a default config file which can be overridden via `~/.config/fz/config.yml`
- Atom feeds and RSS, detect can look at `mmcdole/gofeed`
- Tests, create more fixture examples
- see `misc/fixes_01.md`, found by Claude

## Later
- Tests
- Create DB, run migrations etc.
- Dockerised versioned
- Vim keys, at least up / down
- Have a lookg at

## Refs
- Markdown output via `JohannesKaufmann/html-to-markdown → glamour`
