package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Yoas-Hutapea/Microservice_09/api/handlers"
	"github.com/Yoas-Hutapea/Microservice_09/api/middleware"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Create instances of the services and middleware
	authService := services.NewAuthService()
	pendudukService := services.NewPendudukService()
	kegiatanService := services.NewKegiatanService()
	perangkatService := services.NewPerangkatDesaService()
	validationMiddleware := middleware.NewValidationMiddleware()

	// Create instances of the handlers and apply the validation middleware
	authHandler := handlers.NewAuthHandler(authService)
	pendudukHandler := handlers.NewPendudukHandler(pendudukService)
	kegiatanHandler := handlers.NewKegiatanHandler(kegiatanService)
	perangkatHandler := handlers.NewPerangkatDesaHandler(perangkatService)

	// Register the routes and apply the validation middleware where needed
	router.HandleFunc("/login", validationMiddleware.ValidateUserInput(authHandler.Login)).Methods("POST")
	router.HandleFunc("/penduduk", validationMiddleware.ValidateUserInput(pendudukHandler.AddPenduduk)).Methods("POST")
	router.HandleFunc("/penduduk/{nik}", validationMiddleware.ValidateUserInput(pendudukHandler.UpdatePenduduk)).Methods("PUT")
	router.HandleFunc("/kegiatan", validationMiddleware.ValidateUserInput(kegiatanHandler.AddKegiatan)).Methods("POST")
	router.HandleFunc("/kegiatan/{id}", validationMiddleware.ValidateUserInput(kegiatanHandler.UpdateKegiatan)).Methods("PUT")
	router.HandleFunc("/perangkat", validationMiddleware.ValidateUserInput(perangkatHandler.AddPerangkatDesa)).Methods("POST")
	router.HandleFunc("/perangkat/{id}", validationMiddleware.ValidateUserInput(perangkatHandler.UpdatePerangkatDesa)).Methods("PUT")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
