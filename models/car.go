package models

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required,min=3,max=5"`
	Year      string    `json:"year" validate:"required,yearValidation"`
	Brand     string    `json:"brand" validate:"required"`
	FuleType  string    `json:"fule_type" validate:"required,oneof=petrol diesel electric hybrid"`
	Engine    Engine    `json:"engine" validate:"required"`
	Price     float64   `json:"price" validate:"required,gt=0"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
