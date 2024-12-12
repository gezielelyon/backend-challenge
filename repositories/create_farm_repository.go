package repositories

import (
	"backend-challenge/models"
	"database/sql"
	"fmt"
)

type CreateFarmRepository struct {
	db *sql.DB
}

func NewCreateFarmRepository(db *sql.DB) *CreateFarmRepository {
	return &CreateFarmRepository{db: db}
}

func (r *CreateFarmRepository) CreateFarm(farm *models.Farm) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO farms (id, name, land_area, unit_of_measure, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = tx.Exec(query, farm.ID, farm.Name, farm.LandArea, farm.UnitOfMeasure, farm.Address)
	if err != nil {
		return fmt.Errorf("failed to insert farm: %v", err)
	}

	for _, production := range farm.Productions {
		query := `
			INSERT INTO productions (id, farm_id, crop_type, is_irrigated, is_insured, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		`
		_, err := tx.Exec(query, production.ID, farm.ID, production.CropType, production.IsIrrigated, production.IsInsured)
		if err != nil {
			return fmt.Errorf("failed to insert production: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
