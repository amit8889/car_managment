package store

import (
	"context"

	"github.com/amit8889/car_managemnt/models"
	"github.com/google/uuid"
)

type CarStore interface {
	GetCarById(ctx context.Context, id string) (models.Car, error)
	GetAllCars(ctx context.Context, page int32) ([]models.Car, int32, error)
	CreateCar(ctx context.Context, car models.Car) error
	DeleteCarById(ctx context.Context, id uuid.UUID) error
}
type EngineStore interface {
	GetEngineById(ctx context.Context, id int) (models.Engine, error)
}
