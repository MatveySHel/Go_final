FROM golang:1.23 AS builder

WORKDIR /hotels

COPY go.mod go.sum db/postgres.go ./
RUN go mod download

COPY . .
COPY cmd/hotels/main.go ./main.go

RUN go build -o app ./main.go

CMD ["./app"]