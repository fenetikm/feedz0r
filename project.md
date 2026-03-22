# Feedz0r

My personal RSS CLI feed reader and watcher.
Bring it up in a tmux popup!
Keep it running to update dem feeds.

## Stack
- Go
- Charm stuff
- SQLite

## Features
Me to implement.
- CRUD feeds
- Fetch a feed, see changes
- Poll feeds for updates (configurable re how often)
- TUI for reading feeds:
    * Show which ones have new things
    * Select one, show to the right, formatted

## AI implement ideas
Things for AI to do:
- Add in "help" messages for each command, get "help" command working
- Send email digest

## Specs
- get one random post
- crud stuff

Examples:
```sh
# add
feedzor add --url https://its.mw/atom.xml
```

## Later
- Watch a web page for changes instead of a feed.
- AI stuff:
    * Summary of what is in an article
    * Convert HTML to markdown
    * Configuration on which model to use, support local

## Feeds to start with
- https://its.mw/atom.xml

## Architecture

Note: out of date now.
```text
## Option 1: Flat cmd/ with subcommands (simplest)
    myapp/
    ├── cmd/
    │   └── myapp/
    │       └── main.go          # flag root, subcommands wired here
    ├── internal/
    │   ├── tui/                 # bubbletea/tview components
    │   ├── daemon/              # daemon loop, signal handling, pidfile
    │   └── config/
    └── go.mod

  • Use subcommands: myapp tui, myapp run, myapp daemon
  • Good for small-to-medium projects
  • Everything under internal/ keeps logic out of main.go
```
