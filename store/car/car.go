package store

import (
	"context"
	"database/sql"

	"github.com/amit8889/car_managemnt/models"
)

type CarStore struct {
	DB *sql.DB
}

func New(db *sql.DB) CarStore {
	return CarStore{DB: db}
}
func (s CarStore) GetCarById(ctx context.Context, id int) (models.Car, error) {
	var car models.Car
	return car, nil

}
