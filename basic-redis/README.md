# CLI Todo App with Redis

A command-line todo application built with Go using Redis as the database backend.

## Prerequisites

- Go 1.21+
- Docker Compose (or local Redis on `localhost:6379`)

## Quick Start

```bash
# Install dependencies
go mod download

# Start Redis
docker compose up -d

# Run commands
go run . add "Buy groceries"
go run . list
go run . get <todo-id>
go run . update <todo-id> "Updated text"
go run . delete <todo-id>
```

## Commands

- `add <text>` - Create a new todo
- `list` - List all todos
- `get <id>` - Get a todo by ID
- `update <id> <text>` - Update a todo
- `delete <id>` - Delete a todo
- `--redis=ADDR` - Custom Redis address (default: `localhost:6379`)

## Build

```bash
go build -o todo
./todo list
```

## How It Works

- Todos stored as Redis hashes: `todo:{id}` with fields `id`, `text`, `created_at`
- Index set `todos:all` maintains all todo IDs for listing
- Uses `github.com/redis/go-redis/v9` and `github.com/google/uuid`

## Docker

```bash
docker compose up -d    # Start Redis
docker compose down     # Stop Redis
```
