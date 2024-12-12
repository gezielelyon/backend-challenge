package controllers

import (
	"backend-challenge/services"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type DeleteFarmController struct {
	service services.DeleteFarmServiceInterface
}

func NewDeleteFarmController(service services.DeleteFarmServiceInterface) *DeleteFarmController {
	return &DeleteFarmController{service: service}
}

func (c *DeleteFarmController) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	farmID := vars["id"]

	if _, err := uuid.Parse(farmID); err != nil {
		http.Error(w, fmt.Sprintf("Invalid farm ID: %v", err), http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteFarm(farmID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
