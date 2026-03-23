# Feedz0r

RSS feed wrangler CLI.

## Installation
Note!
This doesn't quite work yet - requires setting up the db file.
```sh
git clone ...
go build -o fz ./cmd/fz
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

