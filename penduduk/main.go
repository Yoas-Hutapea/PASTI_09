package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/Yoas-Hutapea/Microservice_09/penduduk/handlers"
	"github.com/Yoas-Hutapea/Microservice_09/penduduk/repositories"
	"github.com/Yoas-Hutapea/Microservice_09/penduduk/services"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_users")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create an instance of the Repository
	userRepository := repositories.NewUserRepository(db)

	// Create an instance of the service
	pendudukService := services.NewPendudukService(userRepository)

	// Create an instance of the handler
	pendudukHandler := handlers.NewPendudukHandler(pendudukService)

	// Register the routes
	router.HandleFunc("/penduduk", pendudukHandler.AddUser).Methods("POST")
	router.HandleFunc("/penduduk/{id}", pendudukHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/penduduk/{id}", pendudukHandler.DeleteUser).Methods("DELETE")

	// Start the HTTP server
	log.Println("Penduduk service is running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
