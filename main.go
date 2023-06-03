package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/Yoas-Hutapea/Microservice_09/api/handlers"
	"github.com/Yoas-Hutapea/Microservice_09/api/middleware"
	"github.com/Yoas-Hutapea/Microservice_09/api/repositories"
	"github.com/Yoas-Hutapea/Microservice_09/api/services"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Connect to the first MySQL database
	dbUser, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_users")
	if err != nil {
		log.Fatal(err)
	}
	defer dbUser.Close()

	// Connect to the second MySQL database
	dbPerangkat, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_perangkat")
	if err != nil {
		log.Fatal(err)
	}
	defer dbPerangkat.Close()

	// Connect to the third MySQL database
	dbPengumuman, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_pengumuman")
	if err != nil {
		log.Fatal(err)
	}
	defer dbPengumuman.Close()

	// Connect to the fourth MySQL database
	dbKegiatan, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_kegiatan")
	if err != nil {
		log.Fatal(err)
	}
	defer dbKegiatan.Close()
	// Create an instance of the Repository
	UserRepository := repositories.NewUserRepository(dbUser)
	perangkatRepository :=  repositories.NewPerangkatDesaRepository(dbPerangkat)
	pengumumanRepository := repositories.NewPengumumanRepository(dbPengumuman)
	kegiatanRepository := repositories.NewKegiatanRepository(dbKegiatan)

	// Create instances of the services and middleware
	authService := services.NewAuthService(UserRepository)
	pendudukService := services.NewPendudukService(UserRepository)
	kegiatanService := services.NewKegiatanService(kegiatanRepository)
	perangkatService := services.NewPerangkatDesaService(perangkatRepository)
	pengumumanService := services.NewPengumumanService(pengumumanRepository)
	validationMiddleware := middleware.NewValidationMiddleware()

	// Create instances of the handlers and apply the validation middleware
	authHandler := handlers.NewAuthHandler(authService)
	pendudukHandler := handlers.NewPendudukHandler(pendudukService)
	kegiatanHandler := handlers.NewKegiatanHandler(kegiatanService)
	perangkatHandler := handlers.NewPerangkatHandler(perangkatService)
	pengumumanHandler := handlers.NewPengumumanHandler(pengumumanService)

	// Register the routes and apply the validation middleware where needed
	router.HandleFunc("/login", validationMiddleware.ValidateUserInput(authHandler.Login)).Methods("POST")
	router.HandleFunc("/penduduk", validationMiddleware.ValidateUserInput(pendudukHandler.AddUser)).Methods("POST")
	router.HandleFunc("/penduduk/{nik}", validationMiddleware.ValidateUserInput(pendudukHandler.UpdateUser)).Methods("PUT")
	router.HandleFunc("/kegiatan", validationMiddleware.ValidateUserInput(kegiatanHandler.AddKegiatan)).Methods("POST")
	router.HandleFunc("/kegiatan/{id}", validationMiddleware.ValidateUserInput(kegiatanHandler.UpdateKegiatan)).Methods("PUT")
	router.HandleFunc("/perangkat", validationMiddleware.ValidateUserInput(perangkatHandler.AddPerangkat)).Methods("POST")
	router.HandleFunc("/perangkat/{id}", validationMiddleware.ValidateUserInput(perangkatHandler.UpdatePerangkatDesa)).Methods("PUT")
	router.HandleFunc("/pengumuman", validationMiddleware.ValidateUserInput(pengumumanHandler.AddPengumuman)).Methods("POST")
	router.HandleFunc("/pengumuman/{id}", validationMiddleware.ValidateUserInput(pengumumanHandler.UpdatePengumuman)).Methods("PUT")

	// Start the HTTP server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
