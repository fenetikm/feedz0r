# TODO.md

## Now
- [X] List of feeds
- [X] Get a list of entries for a feed
- [ ] Get post by ID
- [ ] Try using FZF, raw output
- [ ] Output markdown
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

## Notes
Fetching:
- we have the command `fetch`
- want to have something that can be called by daemon and also the command

Using FZF we could do something via this command for simplicity.

## Soon
- `help <cmd>` to run `Help` func
- Config file specifies location of db e.g. `~/.fz/fz.db`
- Have a default config file which can be overridden via `~/.config/fz/config.yml`
- Atom feeds and RSS, detect can look at `mmcdole/gofeed`
- Test against a local file for fetching
- Markdown output via `JohannesKaufmann/html-to-markdown → glamour.`

## Later
- Tests
- Create DB, run migrations etc.
- Dockerised versioned
- Vim keys, at least up / down
- Have a lookg at
