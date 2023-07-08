package server

import (
	"database/sql"
	"fmt"
	"food_delivery/config"
	"food_delivery/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {

	//fmt.Println(cfg)
	connStr := "postgres://postgres:password@localhost:5432/db_fooddelivery?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	userHandler := handler.NewUserHandler(db)
	r.HandleFunc("/users", userHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/users", userHandler.Create).Methods(http.MethodPost)

	addressHandler := handler.NewAddressHandler(db)
	r.HandleFunc("/address", addressHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/address", addressHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/address/{id:[0-9]+}", addressHandler.Delete).Methods(http.MethodDelete)

	fmt.Println("Server is running on port :8080")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
