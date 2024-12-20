package service

import (
	"context"

	"github.com/amit8889/car_managemnt/models"
	"github.com/google/uuid"
)

type CarServiceInterface interface {
	GetCarById(ctx context.Context, id string) (models.Car, error)
	GetAllCars(ctx context.Context, page int32) ([]models.Car, int32, error)
	CreateCar(ctx context.Context, car models.Car) error
	DeleteCarById(ctx context.Context, carId uuid.UUID) error
}

type EngineServiceInterface interface {
	GetEngineById(ctx context.Context, id int) (models.Engine, error)
}
