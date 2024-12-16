package service

import (
	"context"

	"github.com/amit8889/car_managemnt/models"
	"github.com/amit8889/car_managemnt/store"
)

type CarService struct {
	carStore store.CarStore
}

// Dipendency injection
func NewCarService(carStore store.CarStore) *CarService {
	return &CarService{
		carStore: carStore,
	}
}

func (s *CarService) GetCarById(ctx context.Context, id int) (models.Car, error) {
	return s.carStore.GetCarById(ctx, id)
}

//validate req body also here
