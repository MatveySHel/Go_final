package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/MatveyShel/Go_final/hotels/domain"
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

func (r *Repository) CreateNewHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
	query := `
        INSERT INTO hotels (name, price, city) VALUES (@name, @price, @city) RETURNING id;
    `
    // Define the named arguments for the query.
    args := pgx.NamedArgs{
        "name"	:   hotel.Name,
        "price" :   hotel.Price,
        "city"  :   hotel.City,
    }
    
    // Execute the query with named arguments to insert the book details into the database.
    err := r.pgx.QueryRow(context.Background(), query, args).Scan(&hotel.ID)
    if err != nil {
        log.Println("Error Adding Hotel")
        return nil, err
    }

	return hotel, nil
}


func (r *Repository) GetHotelsList(ctx context.Context) (*[]domain.Hotel, error) {
    // Define the SQL query to select all books.
    query := `
        SELECT * FROM hotels
    `
    // Execute the query to fetch all book details from the database.
    rows, err := r.pgx.Query(context.Background(), query)
    if err != nil {
        log.Printf("Error Querying the Table")
        return nil, err
    }
    defer rows.Close()

    hotels, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Hotel])
    if err != nil {
        fmt.Printf("CollectRows error: %v", err)
    }
    
    return &hotels, nil
}

func (r *Repository) EditHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
    query := `
        UPDATE hotels SET name = $1, price = $2, city = $3 WHERE name = $4 
    `
    _, err := r.pgx.Exec(context.Background(), query, hotel.Name, hotel.Price, hotel.City, hotel.Name)
    if err != nil {
        log.Printf("Error Querying the Table")
        return nil, err
    }
    
    return hotel, nil
}

func (r *Repository) GetPrice(ctx context.Context, hotel string) (int, error) {
    query := `
        SELECT price FROM hotels WHERE name = $1
    `
    var price int
    err := r.pgx.QueryRow(context.Background(), query, hotel).Scan(&price)
    if err != nil {
        log.Printf("Error Querying the Table")
        return 0, err
    }
    
    return price, nil
}