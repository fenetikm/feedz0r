# TODO.md

## Now
- [X] Basic daemon going
- [X] Don't fetch if it has been a while
- [ ] List of feeds
- [ ] Get a list of entries for a feed
- [ ] Try using FZF, raw output
- [ ] Output markdown

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

## Notes
Fetching:
- we have the command `fetch`
- want to have something that can be called by daemon and also the command

Using FZF we could do something via this command for simplicity.

## Soon
- Fetch feed
- Wrap db handling
- Commands print their help
- Config file specifies location of db e.g. `~/.fz/fz.db`
- Atom feeds and RSS, detect can look at `mmcdole/gofeed`
- Test against a local file for fetching
- Markdown output via `JohannesKaufmann/html-to-markdown → glamour.`

## Later
- Tests
- Create DB, run migrations etc.
- Dockerised versioned
- Vim keys, at least up / down
- Have a lookg at
