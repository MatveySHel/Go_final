package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/MatveyShel/Go_final/booking/domain"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	pgx *pgx.Conn // сторонняя библиотека для общения с бд
}

func NewRepository(pgx *pgx.Conn) *Repository {
	return&Repository{
		pgx: pgx,
	}
}

func (r *Repository) CreateNewBooking(ctx context.Context, booking domain.Booking) (*domain.Booking, error) {
	query := `
        INSERT INTO bookings (client, hotel, checkIn, checkOut) VALUES (@client, @hotel, @checkIn, @checkOut) RETURNING id;
    `
    // Define the named arguments for the query.
    args := pgx.NamedArgs{
        "client"	:   booking.Client,
        "hotel"     :   booking.Hotel,
        "checkIn"   :   booking.CheckIn,
        "checkOut"  :   booking.CheckOut,
    }
    // Execute the query with named arguments to insert the book details into the database.
    err := r.pgx.QueryRow(context.Background(), query, args).Scan(&booking.ID)
    if err != nil {
        log.Println("Error Adding Hotel")
        return nil, err
    }
	return &booking, nil
}

func (r *Repository) GetClientBookingList(ctx context.Context, client string) (*[]domain.Booking, error) {
    
    query := `
        SELECT * FROM bookings WHERE client = $1
    `

    rows, err := r.pgx.Query(context.Background(), query, client)
    if err != nil {
        log.Printf("Error Querying the Table")
        fmt.Println("ERROR : ", err)
        return nil, err
    }
    defer rows.Close()

    
    bookings, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Booking])
    if err != nil {
        fmt.Printf("CollectRows error: %v", err)
    }

    return &bookings, nil
}

func (r *Repository) GetHotelBookingList(ctx context.Context, hotel string) (*[]domain.Booking, error) {
    query := `
        SELECT * FROM bookings WHERE hotel = $1
    `

    rows, err := r.pgx.Query(context.Background(), query, hotel)
    if err != nil {
        log.Printf("Error Querying the Table")
        fmt.Println("ERROR : ", err)
        return nil, err
    }
    defer rows.Close()

    
    bookings, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Booking])
    if err != nil {
        fmt.Printf("CollectRows error: %v", err)
    }

    return &bookings, nil
}