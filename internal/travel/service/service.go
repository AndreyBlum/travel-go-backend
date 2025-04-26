package service

import (
	"travel-go/backend/internal/travel"
)

type TravelService struct {
	Repo travel.Repository
}

func NewTravelService(repo travel.Repository) *TravelService {
	return &TravelService{Repo: repo}
}

func (service *TravelService) CreateTravel(user *travel.Travel) error {
	return service.Repo.Create(user)
}

func (service *TravelService) GetTravel() ([]travel.Travel, error) {
	return service.Repo.Get()
}
