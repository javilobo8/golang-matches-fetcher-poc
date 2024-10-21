FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/ ./main.go

FROM alpine:3.20

WORKDIR /opt/app
COPY --from=builder /opt/app/bin/main .
COPY ./config/* ./config/