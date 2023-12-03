FROM golang:1.21-alpine3.18 AS builder

COPY . /github.com/k6mil6/pocket-project-bot
WORKDIR /github.com/k6mil6/pocket-project-bot

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/k6mil6/pocket-project-bot/bin/bot .
COPY --from=0 /github.com/k6mil6/pocket-project-bot/configs configs/

EXPOSE 80

CMD ["./bot"]