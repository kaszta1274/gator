# gator

An RSS feed aggregator.

### Features

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

### Prerequisites

- [Postgres v15 or later](https://www.postgresql.org/download/)
- [Go 1.26+](https://go.dev/dl/)

### Installation

```bash
go install github.com/kaszta1274/gator
```

### Setup

1. Create a new PostgreSQL database.

2. Create a config file in your home directory, ```~/.gatorconfig.json```, with the database connection string:
    ```
    {
        "db_url": "protocol://username:password@host:port/database?sslmode=disable"
    }
    ```

### Usage

- ```gator login <name>``` - sets the current user
- ```gator register <name>``` - adds a new user to the database
- ```gator reset``` - deletes all users
- ```gator users``` - lists all the users in the database
- ```gator agg <time_between_reqs>``` - colects RSS feeds
- ```gator addfeed <name> <url>``` - adds an RSS feed to the database
- ```gator feeds``` - lists all the feeds in the database
- ```gator follow <url>``` - follows the RSS feed for the current user
- ```gator unfollow <url>``` - follows the RSS feed for the current user
- ```gator following``` - lists all the followed feeds for the current user
- ```gator browse [limit]``` - displays summaries of aggregated posts
