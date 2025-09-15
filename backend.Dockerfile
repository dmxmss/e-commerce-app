FROM golang:1.24.4-alpine3.22 AS builder
WORKDIR /usr/src/app

COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod download

COPY ./backend/config.yaml .
COPY main.go .
COPY ./backend/ .

RUN go build -o application .

FROM alpine:3.21
WORKDIR /app

COPY --from=builder /usr/src/app/config.yaml .
COPY --from=builder /usr/src/app/application ./app

EXPOSE 8080

CMD ["./app"]
