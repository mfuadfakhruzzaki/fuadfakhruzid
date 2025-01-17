package handlers

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"my-gin-mongo/config"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func uploadFileToGCS(file multipart.File, fileHeader *multipart.FileHeader, objectName string) (string, error) {
	ctx := context.Background()
	bucketName := config.GetEnv("GCS_BUCKET_NAME", "default-bucket-name") // Ambil bucket name dari .env
	bucket := config.GetBucket(bucketName)

	// Buat objek di GCS
	obj := bucket.Object(objectName)
	writer := obj.NewWriter(ctx)
	defer writer.Close()

	// Salin isi file ke objek GCS
	fileBytes := make([]byte, fileHeader.Size)
	if _, err := file.Read(fileBytes); err != nil {
		return "", err
	}
	if _, err := writer.Write(fileBytes); err != nil {
		return "", err
	}

	// Set ACL (opsional)
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", err
	}

	// URL publik GCS
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	return url, nil
}


func UploadProfilePicture(c *gin.Context) {
	// Ambil file dari form-data
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}
	defer file.Close()

	// Buat nama unik untuk file
	objectName := fmt.Sprintf("profile_pictures/%d_%s", time.Now().Unix(), fileHeader.Filename)

	// Upload ke GCS
	url, err := uploadFileToGCS(file, fileHeader, objectName)
	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Ambil dokumen pertama dari koleksi profiles
	coll := config.GetCollection("profiles")
	ctx := context.Background()
	iter := coll.Limit(1).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		c.JSON(http.StatusNotFound, gin.H{"error": "No profile found"})
		return
	} else if err != nil {
		log.Printf("Failed to get profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	// Update field profile_picture_url
	_, err = coll.Doc(doc.Ref.ID).Update(ctx, []firestore.Update{
		{Path: "profile_picture_url", Value: url},
	})
	if err != nil {
		log.Printf("Failed to update Firestore: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL to Firestore"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile picture uploaded and saved successfully", "url": url})
}


func UploadCV(c *gin.Context) {
	// Ambil file dari form-data
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}
	defer file.Close()

	// Buat nama unik untuk file
	objectName := fmt.Sprintf("cv_uploads/%d_%s", time.Now().Unix(), fileHeader.Filename)

	// Upload ke GCS
	url, err := uploadFileToGCS(file, fileHeader, objectName)
	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Ambil dokumen pertama dari koleksi profiles
	coll := config.GetCollection("profiles")
	ctx := context.Background()
	iter := coll.Limit(1).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		c.JSON(http.StatusNotFound, gin.H{"error": "No profile found"})
		return
	} else if err != nil {
		log.Printf("Failed to get profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	// Update field cv_url
	_, err = coll.Doc(doc.Ref.ID).Update(ctx, []firestore.Update{
		{Path: "cv_url", Value: url},
	})
	if err != nil {
		log.Printf("Failed to update Firestore: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL to Firestore"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CV uploaded and saved successfully", "url": url})
}

