FROM golang:1.19.3-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/cmd/app

RUN go build -o my-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=build /app/cmd/app/my-service .

CMD ["./my-service"]