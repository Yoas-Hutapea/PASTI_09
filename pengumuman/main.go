package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/Yoas-Hutapea/Microservice_09/perangkat/handlers"
	"github.com/Yoas-Hutapea/Microservice_09/perangkat/repositories"
	"github.com/Yoas-Hutapea/Microservice_09/perangkat/services"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_perangkat")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create an instance of the Repository
	perangkatRepository := repositories.NewPerangkatDesaRepository(db)

	// Create an instance of the service
	perangkatService := services.NewPerangkatDesaService(perangkatRepository)

	// Create an instance of the handler
	perangkatHandler := handlers.NewPerangkatHandler(perangkatService)

	// Register the routes
	router.HandleFunc("/perangkat", perangkatHandler.AddPerangkat).Methods("POST")
	router.HandleFunc("/perangkat/{id}", perangkatHandler.UpdatePerangkatDesa).Methods("PUT")

	// Start the HTTP server
	log.Println("Perangkat service is running on http://localhost:8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}
