package main

import (
	"fmt"
	"log"
	"os"

	"my-gin-mongo/config"
	"my-gin-mongo/routes"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware menambahkan header CORS pada tiap response
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Daftar domain yang diizinkan
		allowedOrigins := []string{
			"https://www.fuadfakhruz.id",
			"https://admin.fuadfakhruz.id",
		}

		origin := c.Request.Header.Get("Origin")
		allowOrigin := ""

		// Periksa apakah asal request ada dalam daftar domain yang diizinkan
		for _, o := range allowedOrigins {
			if o == origin {
				allowOrigin = o
				break
			}
		}

		if allowOrigin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Jika request method OPTIONS, langsung balas dengan status No Content (204)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Memuat environment variables
	config.LoadEnv()

	// Inisialisasi Firestore (dan membaca environment variables)
	if err := config.InitFirestore(); err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}

	// Inisialisasi Google Cloud Storage
	if err := config.InitGCS(); err != nil {
		log.Fatalf("Failed to initialize Google Cloud Storage: %v", err)
	}

	// Inisialisasi router dengan Gin
	r := gin.Default()

	// Daftarkan middleware CORS
	r.Use(CORSMiddleware())

	// Register routes dari package routes
	routes.RegisterRoutes(r)

	// Ambil nilai PORT dari environment, jika tidak ada gunakan nilai default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on port %s", port)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
