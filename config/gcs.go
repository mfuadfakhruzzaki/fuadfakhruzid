package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var GCSClient *storage.Client

func InitGCS() error {
	// Ambil path JSON Key dari environment variable
	keyFile := os.Getenv("GCS_KEY_FILE")
	if keyFile == "" {
		log.Fatal("GCS_KEY_FILE not set in .env")
	}

	// Inisialisasi Google Cloud Storage Client
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(keyFile))
	if err != nil {
		return err
	}

	GCSClient = client
	log.Println("Google Cloud Storage client initialized")
	return nil
}

func GetBucket(bucketName string) *storage.BucketHandle {
	return GCSClient.Bucket(bucketName)
}
