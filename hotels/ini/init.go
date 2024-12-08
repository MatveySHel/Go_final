package ini

// инифциализация приложения (методы которые вызываются в main

import (
	"context"
	"log"

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
	//инициализация grpc, http, роутинг, адаптеров, репозиториев, кафка, коннекторов к другим микросервисам,
	conn, err := db.NewPostgres("postgres", "123", "localhost", "5432", "db")
    if err != nil {
        log.Fatalf("Error : %v", err)
    }
	repo := repository.NewRepository(conn)
	a.service = app.NewService(repo)

	return a.service, nil
}