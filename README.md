## [WIP] Quiz APP API

This API is a work in progress being developed as a side project to exemplify the creation of a
RESTful API using [Go](https://golang.org), [Mux](https://github.com/gorilla/mux) as the routing framework and [Postgres](https://www.postgresql.org) as the database.

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The API requires **Go 1.16 or above**.

```shell
# clone the repo
git clone https://github.com/jotajay/go-quiz-api.git

cd go-quiz-api
go run /cmd/server/main.go
```

Now you have the API at `http://127.0.0.1:3000` providing the following endpoints:

As of now the API offers the following endpoints:

## HealthCheck

- `GET /`: Get API health status

## User

- `GET /users`: Get users
- `GET /users/{id}`: Get user by id
- `POST /users`: Create a new user

## TODO

- [] Add env vars
- [x] Change postgres pk from autoincrement to UUID
- [] Test all endpoints
- [] Add Dockerfile for local developement
- [] Add log system
- [] Add CI/CD
- [] Create endpoints for Quiz, CompletedQuiz & Score
