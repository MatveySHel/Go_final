FROM golang:1.23 AS builder

WORKDIR /tg_bot

COPY go.mod go.sum ./
RUN go mod download


COPY . .
COPY cmd/telegram_bot/main.go ./main.go

RUN go build -o app ./main.go

CMD ["./app"]