package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/Yoas-Hutapea/Microservice_09/auth/handlers"
	"github.com/Yoas-Hutapea/Microservice_09/auth/repositories"
	"github.com/Yoas-Hutapea/Microservice_09/auth/services"
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
	authService := services.NewAuthService(userRepository)

	// Create an instance of the handler
	authHandler := handlers.NewAuthHandler(authService)

	// Register the routes
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/penduduk/{nik}", authHandler.GetUserByNIK).Methods("GET")
	router.HandleFunc("/penduduk", authHandler.GetAllUsers).Methods("GET")

	// Start the HTTP server
	log.Println("Auth service is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
