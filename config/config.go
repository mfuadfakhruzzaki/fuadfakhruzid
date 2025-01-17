// config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func InitConfig() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, relying on environment variables.")
    }
}

func LoadEnv() {
	// Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Fungsi untuk mendapatkan nilai variabel environment
func GetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}