## [WIP] Quiz APP API

This API is a work in progress being developed as a side project to exemplify the creation of a
RESTful API using [Go](https://golang.org), [Mux](https://github.com/gorilla/mux) as the routing framework and [Postgres](https://www.postgresql.org) as the database.

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The API requires **Go 1.16 or above**.

```shell
# clone the repo
git clone https://github.com/jotajay/go-quiz-api.git

# you need a postgres container running
docker run -dp "5432:5432" --name <name> -e POSTGRES_PASSWORD=<your-password> postgres
# To initialize the API
cd go-quiz-api
make dev

# To test the API
make test
```

Now you have the API at `http://127.0.0.1:3000` providing the following endpoints:

## HealthCheck

- `GET /`: Get API health status

## User

- `GET /users`: Get users
- `GET /users/{id}`: Get user by id
- `POST /users`: Create a new user

## Quiz

- `POST /quiz` : Create a new quiz

## TODO

- [] Add env vars
- [] Test all endpoints
- [] Add Dockerfile for local developement
- [] Add log system
- [] Add CI/CD
- [] Create endpoints for Quiz, CompletedQuiz & Score
- [] Create config env
