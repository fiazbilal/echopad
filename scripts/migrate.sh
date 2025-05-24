#!/bin/bash

# Set default values
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}
DB_NAME=${DB_NAME:-company_db}
DB_SSLMODE=${DB_SSLMODE:-disable}

# Load environment variables from .env file if it exists
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Construct the PostgreSQL URL
POSTGRESQL_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"

# Check if command argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 [up|down|create <name>]"
    exit 1
fi

# Execute migrations based on command
case "$1" in
    up)
        echo "Running migrations up..."
        migrate -database "${POSTGRESQL_URL}" -path ./migrations up
        ;;
    down)
        echo "Running migrations down..."
        migrate -database "${POSTGRESQL_URL}" -path ./migrations down
        ;;
    create)
        if [ -z "$2" ]; then
            echo "Migration name is required for create command"
            echo "Usage: $0 create <name>"
            exit 1
        fi
        echo "Creating new migration: $2"
        migrate create -ext sql -dir ./migrations -seq "$2"
        ;;
    *)
        echo "Invalid command. Use 'up', 'down', or 'create <name>'"
        exit 1
        ;;
esac
