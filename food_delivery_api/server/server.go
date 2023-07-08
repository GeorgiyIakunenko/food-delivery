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
	//connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	//cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)

	connStr := "postgres://postgres:password@localhost:5432/food_delivery_v2?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	// user handler

	UserHandler := handler.NewUserHandler(db)
	r.HandleFunc("/users", UserHandler.GetAll).Methods("GET")

	// supplier handler

	SupplierHandler := handler.NewSupplierHandler(db)
	r.HandleFunc("/suppliers", SupplierHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/supplier/{id}", SupplierHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/suppliers/{category}", SupplierHandler.GetByCategory).Methods(http.MethodGet)

	// category handler

	CategoryHandler := handler.NewCategoryHandler(db)
	r.HandleFunc("/categories", CategoryHandler.GetAll).Methods(http.MethodGet)

	fmt.Println("Server started at port", cfg.Port)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
