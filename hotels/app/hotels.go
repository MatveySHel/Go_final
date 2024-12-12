package app

import (
	"context"
	"github.com/MatveyShel/Go_final/hotels/domain"
)

type Service struct{
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo : repo,
	}
}

func (s *Service) CreateNewHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
	return s.repo.CreateNewHotel(ctx, hotel)
}

func (s *Service) EditHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
	return s.repo.EditHotel(ctx, hotel)
}

func (s *Service) GetHotelsList(ctx context.Context) (*[]domain.Hotel, error) {
	return s.repo.GetHotelsList(ctx)
}

func (s *Service) GetPrice(ctx context.Context, hotel string) (int, error) {
	return s.repo.GetPrice(ctx, hotel)
}