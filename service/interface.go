package service

import (
	"context"

	"github.com/amit8889/car_managemnt/models"
)

type CarServiceInterface interface {
	GetCarById(ctx context.Context, id int) (models.Car, error)
}

type EngineServiceInterface interface {
	GetEngineById(ctx context.Context, id int) (models.Engine, error)
}
