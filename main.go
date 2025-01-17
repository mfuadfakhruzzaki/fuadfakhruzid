package main

import (
	"fmt"
	"log"
	"os"

	"my-gin-mongo/config"
	"my-gin-mongo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    // Inisialisasi Firestore (dan baca environment variables)
    if err := config.InitFirestore(); err != nil {
        log.Fatalf("Error initializing Firestore: %v", err)
    }

    r := gin.Default()

    // Register Routes
    routes.RegisterRoutes(r)

    // Baca PORT dari env
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default
    }

    addr := fmt.Sprintf(":%s", port)
    log.Printf("Starting server on port %s", port)
    if err := r.Run(addr); err != nil {
        log.Fatal(err)
    }
}
