# RSS Feed Aggregator

A command-line tool that aggregates posts from RSS feeds you follow. Stay updated with content from multiple sources in one place.

## Prerequisites

Before using RSS Feed Aggregator, ensure you have:

- **Go** (version 1.16 or newer): [Install Go](https://golang.org/doc/install)
- **PostgreSQL** (version 12 or newer): [Install PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the RSS Feed Aggregator CLI using Go:

```bash
# Install directly from GitHub
go install github.com/shivtriv12/RSSFeedAggregator@latest

# Or clone and install locally
git clone https://github.com/shivtriv12/RSSFeedAggregator.git
cd RSSFeedAggregator
go install .
```

## Database Setup

1. Create a PostgreSQL database for the application:

```bash
# Login to PostgreSQL
sudo -u postgres psql

# Create the database
CREATE DATABASE gator;

# Exit PostgreSQL
\q
```

2. Run the database migrations (included in the repository):

```bash
cd RSSFeedAggregator
go get github.com/pressly/goose/v3/cmd/goose
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

## Configuration

On first run, the application will create a config file at `~/.gatorconfig.json`. You can manually edit this file to change your database connection:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Usage

Here are some commands to get you started:

### User Management

```bash
# Register a new user
gator register username

# Log in as an existing user
gator login username
```

### Feed Management

```bash
# Add a new RSS feed
gator addfeed "Tech News" "https://example.com/feed.xml"

# List feeds you're following
gator following
```

### Reading Content

```bash
# Browse the latest posts (default: shows 2 posts)
gator browse

# Browse more posts
gator browse 10
```

### Automatic Feed Collection

```bash
# Start collecting feeds every minute
gator agg 1m
```

## Command Reference

| Command                | Description                                                              |
| ---------------------- | ------------------------------------------------------------------------ |
| `register <username>`  | Creates a new user account and logs you in                               |
| `login <username>`     | Logs in as an existing user                                              |
| `users`                | Lists all registered users, marking the current user                     |
| `addfeed <name> <url>` | Adds a new RSS feed and automatically follows it                         |
| `feeds`                | Lists all feeds in the system (name, URL and owner)                      |
| `follow <url>`         | Follows an existing feed by its URL                                      |
| `following`            | Lists all feeds the current user is following                            |
| `unfollow <url>`       | Stops following a feed                                                   |
| `browse [limit]`       | Shows latest posts from feeds you follow (default: 2 posts)              |
| `agg <interval>`       | Starts continuous feed collection at specified intervals (e.g., 60s, 5m) |
| `reset`                | Resets the database by clearing all users (use with caution)             |
