package main

import (
	"first-go-pet-project/backend/internal/app"
	"first-go-pet-project/backend/internal/routes"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	application, err := app.NewApp("backend/configs")
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	// Регистрируем маршрут для health
	application.Router.HandleFunc("/api/health", routes.HealthHandler)

	application.Start()
}
