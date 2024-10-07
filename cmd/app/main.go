package main

import (
	"fmt"
	"log"
	"golang_with_couchdb2/internal/adapters/db"
	"golang_with_couchdb2/internal/usecases/interactors"
	"golang_with_couchdb2/internal/delivery"
	"golang_with_couchdb2/internal/router"
)

func main() {
	couchDBURL := "http://admin:admin123@localhost:5983"
	productRepo := db.NewCouchDBProductRepository(couchDBURL)
	productInteractor := interactors.NewProductInteractor(productRepo)
	productHandler := delivery.NewProductHandler(productInteractor)

	// Initialize router with product handler
	r := router.InitRoutes(productHandler)

	fmt.Println("Server running on port 8080")
	log.Fatal(r.Run(":8080"))
}
