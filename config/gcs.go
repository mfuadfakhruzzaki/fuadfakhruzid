package config

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
)

var GCSClient *storage.Client

func InitGCS() error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
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
