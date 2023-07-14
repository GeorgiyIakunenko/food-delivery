package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                   string
	AccessSecret           string
	RefreshSecret          string
	AccessLifetimeMinutes  int
	RefreshLifetimeMinutes int
	DbUser                 string
	DbPassword             string
	DbPort                 string
	DbName                 string
	DbHost                 string
	RedisHost              string
	RedisPort              string
	Email                  string
	EmailPassword          string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessMin, err := strconv.Atoi(os.Getenv("ACCESS_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("Error importing .env file")
	}
	refreshMin, err := strconv.Atoi(os.Getenv("REFRESH_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("Error importing .env file")
	}

	return &Config{
		Port:                   os.Getenv("PORT"),
		AccessSecret:           os.Getenv("ACCESS_SECRET"),
		AccessLifetimeMinutes:  accessMin,
		RefreshSecret:          os.Getenv("REFRESH_SECRET"),
		RefreshLifetimeMinutes: refreshMin,
		DbUser:                 os.Getenv("DB_USER"),
		DbHost:                 os.Getenv("DB_HOST"),
		DbPassword:             os.Getenv("DB_PASSWORD"),
		DbPort:                 os.Getenv("DB_PORT"),
		DbName:                 os.Getenv("DB_NAME"),
		RedisHost:              os.Getenv("REDIS_HOST"),
		RedisPort:              os.Getenv("REDIS_PORT"),
		Email:                  os.Getenv("EMAIL"),
		EmailPassword:          os.Getenv("EMAIL_PASSWORD"),
	}

}
