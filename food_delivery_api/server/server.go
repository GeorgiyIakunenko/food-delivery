package server

import (
	"context"
	"database/sql"
	"fmt"
	"food_delivery/config"
	"food_delivery/handler"
	"food_delivery/repository"
	"food_delivery/service"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// open redis connection

	redisConnStr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	redisDb := redis.NewClient(&redis.Options{
		Addr: redisConnStr,
	})

	pong, err := redisDb.Ping(redisDb.Context()).Result()
	fmt.Println(pong, err)

	// redis test

	value, err := redisDb.Get(context.Background(), "user:1:token").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key does not exist")
		} else {
			fmt.Println("Failed to retrieve data:", err)
		}
		return
	}
	fmt.Println("Value:", value)

	// create router

	r := mux.NewRouter()

	// user handler

	UserRepository := repository.NewUserRepository(db)
	UserService := service.NewUserService(UserRepository)
	UserHandler := handler.NewUserHandler(UserService, cfg)

	r.HandleFunc("/users", UserHandler.GetAll).Methods("GET")
	r.HandleFunc("/user/profile", UserHandler.GetUserProfile).Methods("GET")

	// auth handler
	AuthHandler := handler.NewAuthHandler(UserService, cfg)

	r.HandleFunc("/login", AuthHandler.Login).Methods("POST")
	r.HandleFunc("/register", AuthHandler.Register).Methods("POST")
	r.HandleFunc("/refresh", AuthHandler.GetTokenPair).Methods("POST")

	// supplier handler

	// category handler
	CategoryHandler := handler.NewCategoryHandler(db)
	r.HandleFunc("/categories", CategoryHandler.GetAll).Methods(http.MethodGet)

	// product handler

	ProductHandler := handler.NewProductHandler(db)
	r.HandleFunc("/products", ProductHandler.GetAll).Methods(http.MethodGet)

	fmt.Println("Server started at port", cfg.Port)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
