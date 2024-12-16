package store

import (
	"context"

	"github.com/amit8889/car_managemnt/models"
)

type CarStore interface {
	GetCarById(ctx context.Context, id int) (models.Car, error)
}
type EngineStore interface {
	GetEngineById(ctx context.Context, id int) (models.Engine, error)
}
