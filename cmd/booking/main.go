package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ArtemShamro/Go_Booking/booking/pkg/api"
)

func main() {
	bookingServer := api.NewBookingServer()

	bookingServer.Init(context.Background())

	r := http.NewServeMux()

	h := api.HandlerFromMux(bookingServer, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8082",
	}

	if err := s.ListenAndServe(); err != nil {
		slog.Debug(err.Error())
		print(err.Error())
	}
}