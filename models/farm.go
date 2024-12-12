package models

import (
	"time"

	"github.com/google/uuid"
)

type Farm struct {
	ID            uuid.UUID    `json:"id"`
	Name          string       `json:"name"`
	LandArea      float64      `json:"land_area"`
	UnitOfMeasure string       `json:"unit_of_measure"`
	Address       string       `json:"address"`
	Productions   []Production `json:"productions"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}
