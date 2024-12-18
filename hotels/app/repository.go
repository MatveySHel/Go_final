package app

import (
	"context"

	"github.com/MatveyShel/Go_final/hotels/domain"
)

type Repository interface {
	CreateNewHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error)
	GetHotelsList(ctx context.Context) (*[]domain.Hotel, error)
	EditHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error)
	GetPrice(ctx context.Context, hotel string) (int, error)
}