## [WIP] Quiz APP API

This API is a work in progress being developed as a side project to exemplify the creation of a
RESTful API using [Go](https://golang.org), [Mux](https://github.com/gorilla/mux) as the routing framework and [Postgres](https://www.postgresql.org) as the database.

## API Status

As of now the API offers the following endpoints:

## HealthCheck

- `GET /`: Get API health status

## User

- `GET /users`: Get users
- `GET /users/{id}`: Get user by id
- `POST /users`: Create a new user

## TODO

- [] Change postgres pk from autoincrement to UUID
- [] Test all endpoints
- [] Add Dockerfile for local developement
- [] Add log system
- [] Add CI/CD
- [] Create endpoints for Quiz, CompletedQuiz & Score
