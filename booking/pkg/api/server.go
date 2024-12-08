package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MatveyShel/Go_final/booking/app"
	"github.com/MatveyShel/Go_final/booking/domain"
	"github.com/MatveyShel/Go_final/booking/pkg/msger"
	"github.com/MatveyShel/Go_final/booking/pkg/repository"
	"github.com/MatveyShel/Go_final/db"
	"github.com/MatveyShel/Go_final/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingServer struct {
	service *app.Service
}

func NewBookingServer() *BookingServer {
	return &BookingServer{
	}
}

type ConfirmResponse struct {
	Booking 	*domain.Booking		`json:"booking,omitempty"`
	Price 		int					`json:"price,omitempty"`
	Days		int					`json:"days,omitempty"`
	TotalPrice	int					`json:"totalprice,omitempty"`
}

func (a *BookingServer) Init(ctx context.Context) error {  // ADD CONFIG
	//инициализация grpc, http, роутинг, адаптеров, репозиториев, кафка, коннекторов к другим микросервисам,
	
	//инициализация postgres
	psotgresConn, err := db.NewPostgres("postgres", "12345", "localhost", "13501", "postgres")
    if err != nil {
        log.Fatalf("Error : %v", err)
    }
	repo := repository.NewRepository(psotgresConn)

	//инициализация grpc  client
	grpcConn, err := grpc.NewClient("0.0.0.0:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewMessengerServerClient(grpcConn)
	msger := msger.NewRpcMessager(c)

	a.service = app.NewService(repo, msger)
	return nil
}

func (server *BookingServer) CreateNewBooking (w http.ResponseWriter, r *http.Request) {
	var m CreateNewBookingJSONRequestBody
	
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
			http.Error(w, "Bad argument", http.StatusBadRequest)

		}
	
	h := domain.Booking{
		Client : *m.Client,
		Hotel : *m.Hotel,
		CheckIn: m.CheckIn.Time,
		CheckOut: m.CheckOut.Time,
	}
	
	// GET PRICE PER DAY FROM HOTEL SERVICE
	price, _ := server.service.GetPrice(context.Background(), *m.Hotel)
	days := int(h.CheckOut.Sub(h.CheckIn).Hours() / 24)
	totalPrice := price * days
	
	// CREATE NEW BOOKING CORE
	res, err := server.service.CreateNewBooking(context.Background(), h)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	
	err = json.NewEncoder(w).Encode(ConfirmResponse{Booking: res, Price: price, 
		Days: days, TotalPrice: totalPrice})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (server *BookingServer) GetClientBookingList (w http.ResponseWriter, r *http.Request) {
	var m GetClientBookingListJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Bad argument", http.StatusBadRequest)
	}

	bookings, err := server.service.GetClientBookingList(context.Background(), *m.Client)
	if err != nil {
		fmt.Printf("ERRRO")
	}
	fmt.Println(m)
	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (server *BookingServer) GetHotelBookingList (w http.ResponseWriter, r *http.Request) {
	var m GetHotelBookingListJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Bad argument", http.StatusBadRequest)
	}

	bookings, err := server.service.GetHotelBookingList(context.Background(), *m.Hotel)
	if err != nil {
		fmt.Printf("ERRRO")
	}
	fmt.Println(m)
	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}