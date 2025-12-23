FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache bash

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app/data

COPY --from=builder /app/app .

CMD ["./app"]