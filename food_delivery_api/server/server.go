package server

import (
	"fmt"
	"food_delivery/config"
	"food_delivery/handler"
	"food_delivery/middleware"
	"food_delivery/repository"
	"food_delivery/repository/db"
	"food_delivery/service"
	_ "github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {
	redisClient, err := db.NewRedisClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//refactor name db to postgresDB due to conflict

	postgresClient, err := db.NewPostgreSQLDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// create router

	r := mux.NewRouter()

	// user handler

	UserRepository := repository.NewUserRepository(postgresClient)
	UserService := service.NewUserService(UserRepository)
	UserHandler := handler.NewUserHandler(UserService, cfg)

	AuthMiddleware := middleware.NewAuthMiddleware(redisClient, cfg)

	r.HandleFunc("/users", UserHandler.GetAll).Methods("GET")

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.Use(AuthMiddleware.ValidateAccessToken)
	userRouter.HandleFunc("/profile", UserHandler.GetUserProfile).Methods("GET")
	userRouter.HandleFunc("/profile", UserHandler.UpdateUserProfile).Methods("PUT")

	// auth handler
	AuthService := service.NewAuthService(UserService, redisClient, cfg)
	AuthHandler := handler.NewAuthHandler(AuthService, cfg)

	r.HandleFunc("/user/profile/password", AuthHandler.ChangePassword).Methods("POST")

	r.HandleFunc("/auth/login", AuthHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/auth/register", AuthHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/auth/reset-password", AuthHandler.InitiatePasswordReset).Methods(http.MethodPost)
	r.HandleFunc("/auth/submit-code", AuthHandler.SubmitResetCode).Methods(http.MethodPost)

	refreshRouter := r.PathPrefix("/auth/refresh").Subrouter()
	refreshRouter.HandleFunc("", AuthHandler.GetTokenPair).Methods(http.MethodPost)
	refreshRouter.Use(AuthMiddleware.ValidateRefreshToken)

	logoutRouter := r.PathPrefix("/auth/logout").Subrouter()
	logoutRouter.HandleFunc("", AuthHandler.Logout).Methods(http.MethodPost)
	logoutRouter.Use(AuthMiddleware.ValidateAccessToken)

	// supplier handler

	SupplierRepository := repository.NewSupplierRepository(postgresClient)
	SupplierService := service.NewSupplierService(SupplierRepository)
	SupplierHandler := handler.NewSupplierHandler(SupplierService)

	r.HandleFunc("/suppliers", SupplierHandler.GetAllSuppliers).Methods(http.MethodGet)
	r.HandleFunc("/supplier/{id}", SupplierHandler.GetSupplierByID).Methods(http.MethodGet)
	r.HandleFunc("/suppliers/{type}", SupplierHandler.GetSupplierByType).Methods(http.MethodGet)
	r.HandleFunc("/suppliers/category/{id}", SupplierHandler.GetSuppliersByCategoryID).Methods(http.MethodGet)

	// category handler
	CategoryHandler := handler.NewCategoryHandler(postgresClient)
	r.HandleFunc("/categories", CategoryHandler.GetAll).Methods(http.MethodGet)

	// product handler

	ProductHandler := handler.NewProductHandler(postgresClient)
	r.HandleFunc("/products", ProductHandler.GetAll).Methods(http.MethodGet)

	fmt.Println("Server started at port", cfg.Port)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
