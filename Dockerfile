FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download -x

RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=docs --build="go build ./cmd/server/main.go" --command=./main