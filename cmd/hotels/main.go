package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"

	"github.com/ArtemShamro/Go_Booking/hotels/ini"
	"github.com/ArtemShamro/Go_Booking/hotels/pkg/api"
	"github.com/ArtemShamro/Go_Booking/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// SERVICE
	hotelService := ini.NewHotelService()

	coreService, _ := hotelService.Init(context.Background())

	hotelServer := api.NewHotelServe(coreService)
	

	// GRCP Server
	// addr := os.Getenv("SERVER_ADDR")
	addr := "0.0.0.0:8081"
	
	// Creating a TCP socket.
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	// If you want to connect to the server via grpcurl, you have to register the reflection service.
	reflection.Register(grpcServer)
	
	// Creating and registering implementation of the storage service.
	pb.RegisterMessengerServerServer(grpcServer, hotelServer)
	
	go func(){
		// Starting the server.
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("server failed")
		}
	}()


	
	// HTTP Server
	r := http.NewServeMux()

	h := api.HandlerFromMux(hotelServer, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	if err := s.ListenAndServe(); err != nil {
		slog.Debug(err.Error())
		print(err.Error())
	}

}

