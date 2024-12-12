package models

import (
	"time"

	"github.com/google/uuid"
)

type Production struct {
	ID          uuid.UUID `json:"id"`
	FarmID      uuid.UUID `json:"farm_id"`
	CropType    string    `json:"crop_type"`
	IsIrrigated bool      `json:"is_irrigated"`
	IsInsured   bool      `json:"is_insured"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
