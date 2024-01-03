#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
set -e

# Run DB migration from main.go instead of here
# echo "run db migration"
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
# Pass all arguments and execute the command
exec "$@"
