#!/usr/bin/env bash
#
# Usage
# cd ~/src/afyadigital/maagizo/backend
# bash ./scripts/dev.sh
set -e

# echo ""
# export DATABASE_URL="${DATABASE_URL-postgres://postgres:supersecret@0.0.0.0/maagizo?sslmode=disable}"

# echo "- Executing database migrations..."
# migrate -path ./pkg/postgres/migrations/ -database $DATABASE_URL up
# echo ""

# echo "- Seeding dev database"
# docker exec -i maagizo_db psql -U postgres -d maagizo < ./scripts/temp_seed.sql
# echo ""

echo "Starting api..."

go run cmd/server/main.go