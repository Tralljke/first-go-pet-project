package main

import (
	"first-go-pet-project/internal/app"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	application, err := app.NewApp("configs")
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	application.Start()
}
