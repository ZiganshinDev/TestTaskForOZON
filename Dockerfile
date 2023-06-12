FROM golang:1.14-buster

ENV GOPATH=/
ENV DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

COPY ./ ./

RUN apt-get update && apt-get -y install postgresql-client && \
    chmod +x wait-for-postgres.sh && \
    go mod download && \
    go build -o app ./cmd/app/main.go

CMD ["./app", "-db"]