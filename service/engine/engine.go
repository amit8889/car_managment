package service

import (
	"context"

	"github.com/amit8889/car_managemnt/models"
	"github.com/amit8889/car_managemnt/store"
)

type EngineService struct {
	engineService store.EngineStore
}

// Dipendency injection
func NewEngineService(engineService store.EngineStore) *EngineService {
	return &EngineService{
		engineService: engineService,
	}
}

func (e *EngineService) GetEngineById(ctx context.Context, id int) (models.Engine, error) {
	return e.engineService.GetEngineById(ctx, id)
}

//validate req body also here
