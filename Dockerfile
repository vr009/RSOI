FROM golang:1.16 AS builder

ENV GO111MODULE=on

ENV DATABASE_URL "user=postgres password=postgres host=postgres port=5432 dbname=postgres"

WORKDIR /opt/app

COPY . .

EXPOSE 8080

RUN go build ./cmd/main.go

CMD ./main