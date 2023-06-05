package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/handlers"
	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/repositories"
	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/services"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_pengumuman")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create an instance of the Repository
	pengumumanRepository := repositories.NewPengumumanRepository(db)

	// Create an instance of the service
	pengumumanService := services.NewPengumumanService(pengumumanRepository)

	// Create an instance of the handler
	pengumumanHandler := handlers.NewPengumumanHandler(pengumumanService)

	// Register the routes
	router.HandleFunc("/pengumuman", pengumumanHandler.AddPengumuman).Methods("POST")
	router.HandleFunc("/pengumuman/{id}", pengumumanHandler.UpdatePengumuman).Methods("PUT")

	// Start the HTTP server
	log.Println("Pengumuman service is running on http://localhost:8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}
