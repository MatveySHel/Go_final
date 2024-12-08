package app

import (
	"context"

	"github.com/MatveyShel/Go_final/booking/domain"
)

type Service struct{
	repo Repository
	msger RpcMessager
}

func NewService(repo Repository, msger RpcMessager) *Service {
	return &Service{
		repo : repo,
		msger : msger,
	}
}

func (s *Service) CreateNewBooking(ctx context.Context, hotel domain.Booking) (*domain.Booking, error) {
	return s.repo.CreateNewBooking(ctx, hotel)
}

func (s *Service) GetClientBookingList(ctx context.Context, client string) (*[]domain.Booking, error) {
	return s.repo.GetClientBookingList(ctx, client)
}

func (s *Service) GetHotelBookingList(ctx context.Context, hotel string) (*[]domain.Booking, error) {
	return s.repo.GetHotelBookingList(ctx, hotel)
}

func (s *Service) GetPrice(ctx context.Context, hotel string) (int, error) {
	return s.msger.AskPrice(ctx, hotel)
}