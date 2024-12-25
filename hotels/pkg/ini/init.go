package ini

// инифциализация приложения (методы которые вызываются в main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MatveyShel/Go_final/db"
	"github.com/MatveyShel/Go_final/hotels/app"
	"github.com/MatveyShel/Go_final/hotels/pkg/repository"
)


type HotelService struct {
	service *app.Service
}

func NewHotelService() *HotelService {
	return &HotelService{}
}

func (a *HotelService) Init(ctx context.Context) (*app.Service, error) {  // ADD CONFIG
	DB_USER := os.Getenv("HOTEL_DB_USER")
	DB_PASSWORD := os.Getenv("HOTEL_DB_PASSWORD")
	DB_HOST := os.Getenv("HOTEL_DB_HOST")
	DB_PORT := os.Getenv("HOTEL_DB_PORT")
	DB_NAME := os.Getenv("HOTEL_DB_NAME")
	time.Sleep(3 * time.Second) //Ждем пока развернется постгря
	//инициализация grpc, http, роутинг, адаптеров, репозиториев, кафка, коннекторов к другим микросервисам,
	conn, err := db.NewPostgres(DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
    if err != nil {
        log.Fatalf("Error : %v", err)
    }
	repo := repository.NewRepository(conn)
	a.service = app.NewService(repo)

	return a.service, nil
}