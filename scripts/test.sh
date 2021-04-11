#!/usr/bin/env bash
#
# Usage
# cd ~/src/github.com/jotajay/go-quiz-app
# bash ./scripts/test.sh
set -e

TEST_DATABASE=quiztestdb
TEST_DATABASE_URL="${TEST_DATABASE_URL-postgres://postgres:supersecret@0.0.0.0/$TEST_DATABASE?sslmode=disable}"

export TEST_DATABASE_URL=$TEST_DATABASE_URL

docker exec pg-quiz-app psql -U postgres -c "DROP DATABASE IF EXISTS $TEST_DATABASE"
docker exec pg-quiz-app psql -U postgres -c "CREATE DATABASE $TEST_DATABASE"

go vet ./...

go test -race -v ./...
