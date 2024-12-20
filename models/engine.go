package models

import (
	"github.com/google/uuid"
)

type Engine struct {
	EngineID     uuid.UUID `json:"engine_id"`
	Displacement float64   `json:"displacement" validate:"required,gt=0"`
	NoOfCylinder int64     `json:"noOfCylinder" validate:"required,min=1"`
	CarRang      float64   `json:"carRange" validate:"required,gt=0"`
}
