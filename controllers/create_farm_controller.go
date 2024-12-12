package controllers

import (
	"backend-challenge/models"
	"backend-challenge/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CreateFarmController struct {
	service services.CreateFarmServiceInterface
}

func NewCreateFarmController(service services.CreateFarmServiceInterface) *CreateFarmController {
	return &CreateFarmController{service: service}
}

func (c *CreateFarmController) Handle(w http.ResponseWriter, r *http.Request) {
	var farm models.Farm
	if err := json.NewDecoder(r.Body).Decode(&farm); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validateFarm(&farm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateFarm(&farm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(farm)
}

func validateFarm(farm *models.Farm) error {
	if strings.TrimSpace(farm.Name) == "" {
		return fmt.Errorf("Farm name is required")
	}

	if farm.LandArea <= 0 {
		return fmt.Errorf("Land area must be greater than zero")
	}

	if strings.TrimSpace(farm.UnitOfMeasure) == "" {
		return fmt.Errorf("Unit of measure is required")
	}

	if err := validateUnitOfMeasure(farm.UnitOfMeasure); err != nil {
		return err
	}

	if strings.TrimSpace(farm.Address) == "" {
		return fmt.Errorf("Address is required")
	}

	for _, production := range farm.Productions {
		if err := validateProduction(&production); err != nil {
			return err
		}
	}

	return nil
}

func validateProduction(production *models.Production) error {
	if strings.TrimSpace(production.CropType) == "" {
		return fmt.Errorf("Crop type is required")
	}

	if err := validateCropType(production.CropType); err != nil {
		return err
	}

	if production.IsIrrigated != true && production.IsIrrigated != false {
		return fmt.Errorf("Irrigation status must be a valid boolean value (true or false)")
	}

	if production.IsInsured != true && production.IsInsured != false {
		return fmt.Errorf("Insurance status must be a valid boolean value (true or false)")
	}

	return nil
}

func validateUnitOfMeasure(unit string) error {
	validUnits := []string{"acre", "hectare"}

	for _, validUnit := range validUnits {
		if strings.ToLower(unit) == validUnit {
			return nil
		}
	}
	return fmt.Errorf("Invalid unit of measure, allowed values are: acre, hectare")
}

func validateCropType(cropType string) error {
	validCrops := []string{"RICE", "BEANS", "CORN", "COFFEE", "SOYBEANS"}

	for _, validCrop := range validCrops {
		if strings.ToUpper(cropType) == validCrop {
			return nil
		}
	}
	return fmt.Errorf("Invalid crop type, allowed values are: RICE, BEANS, CORN, COFFEE, SOYBEANS")
}
