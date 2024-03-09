#!/bin/sh

set -e

echo "start the app"
exec "$@"

# echo "run db migration"
# /app/migrate -path /app/migrate -database "$DB_SOURCE" -verbose up