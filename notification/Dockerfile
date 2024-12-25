FROM golang:1.23-alpine AS builder

WORKDIR /notes

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY cmd/notification/main.go ./main.go

RUN go build -o app ./main.go

CMD ["./app"]