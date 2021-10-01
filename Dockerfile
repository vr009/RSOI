FROM golang:1.16 AS builder

ENV GO111MODULE=on

ENV DATABASE_URL "postgres://postgres:postgres@127.0.0.1:5432/postgres?pool_max_conns=10"

WORKDIR /opt/app

COPY . .

RUN go build ./service/cmd/main.go
