// config.go
package config

import (
	"log"

	"github.com/joho/godotenv"
)


func InitConfig() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, relying on environment variables.")
    }
}