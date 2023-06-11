FROM golang:1.19.3-alpine

RUN apk add --no-cache git ca-certificates postgresql-client bash

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/cmd/app
RUN go build -o my-service .

FROM alpine:latest

COPY --from=0 /app/cmd/app/my-service /usr/local/bin/my-service

COPY wait-for-postgres.sh .

RUN chmod +x wait-for-postgres.sh

CMD ["my-service"]