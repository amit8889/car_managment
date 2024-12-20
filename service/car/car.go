package service

import (
	"context"

	"github.com/amit8889/car_managemnt/models"
	"github.com/amit8889/car_managemnt/store"
	"github.com/google/uuid"
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

func (s *CarService) GetCarById(ctx context.Context, id string) (models.Car, error) {
	return s.carStore.GetCarById(ctx, id)
}
func (s *CarService) GetAllCars(ctx context.Context, page int32) ([]models.Car, int32, error) {
	return s.carStore.GetAllCars(ctx, page)
}
func (s *CarService) CreateCar(ctx context.Context, car models.Car) error {
	return s.carStore.CreateCar(ctx, car)
}
func (s *CarService) DeleteCarById(ctx context.Context, carId uuid.UUID) error {
	return s.carStore.DeleteCarById(ctx, carId)
}

//validate req body also here
