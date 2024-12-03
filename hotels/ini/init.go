package ini

// инифциализация приложения (методы которые вызываются в main

import (
	"context"
	"log"

	"github.com/ArtemShamro/Go_Booking/db"
	"github.com/ArtemShamro/Go_Booking/hotels/app"
	"github.com/ArtemShamro/Go_Booking/hotels/pkg/repository"
)

type HotelService struct {
	service *app.Service
}

func NewHotelService() *HotelService {
	return &HotelService{}
}

func (a *HotelService) Init(ctx context.Context) (*app.Service, error) {  // ADD CONFIG
	//инициализация grpc, http, роутинг, адаптеров, репозиториев, кафка, коннекторов к другим микросервисам,
	conn, err := db.NewPostgres("postgres", "12345", "localhost", "13500", "postgres")
    if err != nil {
        log.Fatalf("Error : %v", err)
    }
	repo := repository.NewRepository(conn)
	a.service = app.NewService(repo)

	return a.service, nil
}