package engine

import (
	"context"
	"database/sql"

	"github.com/amit8889/car_managemnt/models"
)

type EngineStore struct {
	DB *sql.DB
}

func New(db *sql.DB) EngineStore {
	return EngineStore{DB: db}
}

func (s EngineStore) GetEngineById(ctx context.Context, id int) (models.Engine, error) {
	var engine models.Engine
	return engine, nil

}
