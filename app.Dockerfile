FROM golang:1.24.2-alpine3.21 AS builder
WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY config.yaml .
COPY . .

RUN go build -o application .

FROM alpine:3.21
WORKDIR /app

COPY --from=builder /usr/src/app/config.yaml .
COPY --from=builder /usr/src/app/application ./app

EXPOSE 8000

CMD ["./app"]
