package repositories

import (
	"backend-challenge/models"
	"database/sql"
	"fmt"
)

type ListFarmRepository struct {
	db *sql.DB
}

func NewListFarmRepository(db *sql.DB) *ListFarmRepository {
	return &ListFarmRepository{db: db}
}

func (r *ListFarmRepository) ListFarms(page, limit int, cropType, landArea string) ([]models.Farm, error) {
	query := `
		SELECT id, name, land_area, unit_of_measure, address, created_at, updated_at
		FROM farms
		WHERE TRUE
	`

	var args []interface{}
	if cropType != "" {
		query += ` AND id IN (SELECT farm_id FROM productions WHERE crop_type = $1)`
		args = append(args, cropType)
	}
	if landArea != "" {
		query += ` AND land_area = $2`
		args = append(args, landArea)
	}

	offset := (page - 1) * limit
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch farms: %v", err)
	}
	defer rows.Close()

	var farms []models.Farm

	for rows.Next() {
		var farm models.Farm
		err := rows.Scan(&farm.ID, &farm.Name, &farm.LandArea, &farm.UnitOfMeasure, &farm.Address, &farm.CreatedAt, &farm.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan farm: %v", err)
		}

		productionQuery := `
			SELECT id, crop_type, is_irrigated, is_insured, created_at, updated_at
			FROM productions
			WHERE farm_id = $1
		`
		productionRows, err := r.db.Query(productionQuery, farm.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch productions: %v", err)
		}

		for productionRows.Next() {
			var production models.Production
			err := productionRows.Scan(&production.ID, &production.CropType, &production.IsIrrigated, &production.IsInsured, &production.CreatedAt, &production.UpdatedAt)
			if err != nil {
				return nil, fmt.Errorf("failed to scan production: %v", err)
			}
			farm.Productions = append(farm.Productions, production)
		}
		productionRows.Close()

		farms = append(farms, farm)
	}

	if len(farms) == 0 {
		return []models.Farm{}, nil
	}

	return farms, nil
}
