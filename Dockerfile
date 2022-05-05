# syntax=docker/dockerfile:1

FROM golang:1.18.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY tests ./tests

# RUN go build -o /go-blog

CMD [ "go", "test", "./tests"]