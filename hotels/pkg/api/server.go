package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MatveyShel/Go_final/hotels/app"
	"github.com/MatveyShel/Go_final/hotels/domain"
	"github.com/MatveyShel/Go_final/pkg/pb"
)

type HotelServer struct {
	pb.UnimplementedMessengerServerServer

	service *app.Service
}

func NewHotelServe(service *app.Service) *HotelServer {
	return &HotelServer{
		service: service,
	}
}

func (server *HotelServer) GetHotelsList (w http.ResponseWriter, _ *http.Request) {
	m, err := server.service.GetHotelsList(context.Background())
	if err != nil {
		fmt.Printf("ERRRO")
	}
	fmt.Println(m)
	err = json.NewEncoder(w).Encode(m)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (server *HotelServer) CreateNewHotel (w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE")
	var m CreateNewHotelJSONRequestBody
	
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Bad argument", http.StatusBadRequest)
	}
	
	h := domain.Hotel{
		Name: *m.Name,
		Price: *m.Price,
		City: *m.City,
	}
	
	res, err := server.service.CreateNewHotel(context.Background(), &h)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func (server *HotelServer) EditHotel (w http.ResponseWriter, r *http.Request) {
	fmt.Println("EDIT")

	var m EditHotelJSONRequestBody
	
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Bad argument", http.StatusBadRequest)
	}
	
	h := domain.Hotel{
		Name: *m.Name,
		Price: *m.Price,
		City: *m.City,
	}
	
	res, err := server.service.EditHotel(context.Background(), &h)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (server *HotelServer) AskPrice(_ context.Context, in *pb.AskRequest) (*pb.AskResponse, error) {
	price, err := server.service.GetPrice(context.Background(), in.Hotel)
	if err != nil {
        log.Printf("Cant get price")
        return nil, err
    }
	return &pb.AskResponse{Price: int32(price)}, nil
}