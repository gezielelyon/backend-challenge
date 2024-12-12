package controllers

import (
	"backend-challenge/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ListFarmController struct {
	service services.ListFarmServiceInterface
}

func NewListFarmController(service services.ListFarmServiceInterface) *ListFarmController {
	return &ListFarmController{service: service}
}

func (c *ListFarmController) Handle(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	cropType := r.URL.Query().Get("crop_type")
	landArea := r.URL.Query().Get("land_area")

	page := 1
	limit := 10

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}
	}

	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	}

	farms, err := c.service.ListFarms(page, limit, cropType, landArea)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(farms)
}
