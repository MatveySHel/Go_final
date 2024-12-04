package app

import (
	"context"

	"github.com/ArtemShamro/Go_Booking/booking/domain"
)

type Repository interface {
	CreateNewBooking(ctx context.Context, hotel domain.Booking) (*domain.Booking, error)
	GetClientBookingList(ctx context.Context, client string) (*[]domain.Booking, error)
	GetHotelBookingList(ctx context.Context, hotel string) (*[]domain.Booking, error)
}