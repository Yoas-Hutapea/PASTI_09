package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/handlers"
	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/repositories"
	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/services"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_kegiatan")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create an instance of the Repository
	kegiatanRepository := repositories.NewKegiatanRepository(db)

	// Create an instance of the service
	kegiatanService := services.NewKegiatanService(kegiatanRepository)

	// Create an instance of the handler
	kegiatanHandler := handlers.NewKegiatanHandler(kegiatanService)

	// Register the routes
	router.HandleFunc("/kegiatan", kegiatanHandler.AddKegiatan).Methods("POST")
	router.HandleFunc("/kegiatan/{id}", kegiatanHandler.UpdateKegiatan).Methods("PUT")

	// Start the HTTP server
	log.Println("Kegiatan service is running on http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}
