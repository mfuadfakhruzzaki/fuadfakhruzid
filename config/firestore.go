// config/firestore.go
package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

// InitFirestore menginisialisasi koneksi ke Firebase dan Firestore
func InitFirestore() error {
	// Load file .env (jika ada)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file tidak ditemukan, menggunakan environment variables")
	}

	// Pastikan environment variable GOOGLE_APPLICATION_CREDENTIALS ter-set,
	// atau secara langsung gunakan file kredensial yang berada di folder config.
	// Misalnya: config/serviceAccountKey.json
	credFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credFile == "" {
		// Jika tidak diset, gunakan path default
		credFile = "config/serviceAccountKey.json"
	}

	opt := option.WithCredentialsFile(credFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error inisialisasi Firebase App: %v", err)
		return err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Error inisialisasi Firestore: %v", err)
		return err
	}

	FirestoreClient = client
	log.Println("Berhasil terhubung ke Firestore!")
	return nil
}

// GetCollection mengembalikan reference ke koleksi tertentu dalam Firestore
func GetCollection(name string) *firestore.CollectionRef {
	return FirestoreClient.Collection(name)
}
