package server

import (
	"database/sql"
	"fmt"
	"food_delivery/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

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

	fmt.Println("Server is running on port :8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}
