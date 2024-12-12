package main

import (
	"backend-challenge/config"
	"backend-challenge/controllers"
	"backend-challenge/database"
	"backend-challenge/repositories"
	"backend-challenge/routes"
	"backend-challenge/services"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = database.CreateTables(db)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	log.Println("Tabelas criadas com sucesso!")

	createFarmRepo := repositories.NewCreateFarmRepository(db)
	listFarmRepo := repositories.NewListFarmRepository(db)
	deleteFarmRepo := repositories.NewDeleteFarmRepository(db)

	createFarmService := services.NewCreateFarmService(createFarmRepo)
	listFarmService := services.NewListFarmService(listFarmRepo)
	deleteFarmService := services.NewDeleteFarmService(deleteFarmRepo)

	createFarmController := controllers.NewCreateFarmController(createFarmService)
	listFarmController := controllers.NewListFarmController(listFarmService)
	deleteFarmController := controllers.NewDeleteFarmController(deleteFarmService)

	router := routes.SetupRoutes(createFarmController, listFarmController, deleteFarmController)

	log.Printf("Server is running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
