FROM golang:1.23 AS builder

WORKDIR /booking

COPY go.mod go.sum ./
RUN go mod download


COPY . .
COPY cmd/booking/main.go ./main.go

RUN go build -o app ./main.go

CMD ["./app"]