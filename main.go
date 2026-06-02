package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"hotel-booking/internal/config"
	"hotel-booking/internal/db"
	"hotel-booking/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()

	database := db.Connect(cfg)
	defer database.Close()

	r := routes.SetupRoutes()

	fmt.Printf("Server starting on port %s\n", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
