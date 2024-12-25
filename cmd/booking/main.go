package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/MatveyShel/Go_final/booking/pkg/api"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	bookingServer := api.NewBookingServer()

	bookingServer.Init(context.Background())

	r := http.NewServeMux()

	h := api.HandlerFromMux(bookingServer, r)

	s := &http.Server{
		Handler: h,
		Addr:    os.Getenv("BOOKING_SERVICE_ADDR"),
	}

	if err := s.ListenAndServe(); err != nil {
		slog.Debug(err.Error())
		print(err.Error())
	}
}