package repositories

import (
	"database/sql"
	"fmt"
)

type DeleteFarmRepository struct {
	db *sql.DB
}

func NewDeleteFarmRepository(db *sql.DB) *DeleteFarmRepository {
	return &DeleteFarmRepository{db: db}
}

func (r *DeleteFarmRepository) DeleteFarm(farmID string) error {
	query := `
		DELETE FROM farms
		WHERE id = $1
	`
	_, err := r.db.Exec(query, farmID)
	if err != nil {
		return fmt.Errorf("failed to delete farm: %v", err)
	}

	return nil
}
