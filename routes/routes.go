package routes

import (
	"backend-challenge/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(
	createFarmController *controllers.CreateFarmController,
	listFarmController *controllers.ListFarmController,
	deleteFarmController *controllers.DeleteFarmController,
) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/farms", createFarmController.Handle).Methods(http.MethodPost)

	router.HandleFunc("/farms", listFarmController.Handle).Methods(http.MethodGet)

	router.HandleFunc("/farms/{id}", deleteFarmController.Handle).Methods(http.MethodDelete)

	return router
}
