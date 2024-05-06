package config

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-by-sjb/e-commerce-go-api/db"
	"github.com/joho/godotenv"
)

type Config struct {
	*db.Config
	ServerPort string
}

var Envs = initConfig()

func initConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &Config{
		Config: &db.Config{
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASSWORD", "postgres"),
			DBName:     getEnv("DB_NAME", "e-commerce-go-api"),
		},
		ServerPort: getEnv("SERVER_PORT", "3000"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Printf("Loaded \"%s\" env successfully ✅.\n", key)
		return value
	}
	fmt.Printf("Failed to \"%s\" env read 👎. Using fallback env\n", key)
	return fallback
}
