package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MatveyShel/Go_final/booking/app"
	"github.com/MatveyShel/Go_final/booking/domain"
	"github.com/MatveyShel/Go_final/booking/pkg/msger"
	"github.com/MatveyShel/Go_final/booking/pkg/repository"
	"github.com/MatveyShel/Go_final/booking/pkg/kafka"
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
	psotgresConn, err := db.NewPostgres("postgre", "123", "localhost", "6432", "booking")
    if err != nil {
        log.Fatalf("Error : %v", err)
    }
	repo := repository.NewRepository(psotgresConn)


	kafkaBroker := "0.0.0.0:9092"
	kafkaTopic := "notifications"
	kafkaProducer := kafka.NewProducer([]string{kafkaBroker}, kafkaTopic)

	//инициализация grpc  client
	grpcConn, err := grpc.NewClient("0.0.0.0:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewMessengerServerClient(grpcConn)
	msger := msger.NewRpcMessager(c)
	a.service = app.NewService(repo, msger, kafkaProducer)
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
	
	price, _ := server.service.GetPrice(context.Background(), *m.Hotel)
	days := int(h.CheckOut.Sub(h.CheckIn).Hours() / 24)
	totalPrice := price * days
	
	res, err := server.service.CreateNewBooking(context.Background(), h)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	message := fmt.Sprintf(
		"New booking created: {Client: %s, Hotel: %s, CheckIn: %s, CheckOut: %s, TotalPrice: %d}",
		*m.Client, *m.Hotel, m.CheckIn.Time.Format(time.RFC3339), m.CheckOut.Time.Format(time.RFC3339), totalPrice,
	)
	err = server.service.KafkaProducer.SendMessage(context.Background(), []byte(message))
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	} else {
		log.Println("Message successfully sent to Kafka.")
	}

	err = json.NewEncoder(w).Encode(ConfirmResponse{Booking: res, Price: price, 
		Days: days, TotalPrice: totalPrice})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	log.Printf("The booking was created successfully!")
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
		fmt.Printf("ERROR")
	}
	fmt.Println(m)
	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
