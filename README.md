# Feedz0r

RSS feed wrangler TUI.

## Installation
```sh
git clone ...
go build -o fz ./cmd/fz
# move fz somewhere on your $PATH
```

## Get going
```sh
# add a feed
fz add hn https://news.ycombinator.com/rss

# fetch posts for it
fz fetch

# list the latest posts
fz list --type posts

# view one post
fz post 1
```

Watch feeds:
```sh
fz watch
```

## FZF
Make a feed reader:

