package handlers

import (
	"context"
	"net/http"
	"time"

	"my-gin-mongo/config"
	"my-gin-mongo/models"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// GET /contact
// Ambil semua pesan contact menggunakan Firestore
func GetAllContact(c *gin.Context) {
	coll := config.GetCollection("contacts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	iter := coll.Documents(ctx)
	defer iter.Stop()

	var results []models.Contact
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var contact models.Contact
		if err := doc.DataTo(&contact); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Set dokumen ID ke field ID (jika model memiliki field tersebut)
		contact.ID = doc.Ref.ID
		results = append(results, contact)
	}

	c.JSON(http.StatusOK, results)
}

// POST /contact
// Simpan satu pesan contact menggunakan Firestore
func CreateContact(c *gin.Context) {
	coll := config.GetCollection("contacts")
	var payload models.Contact

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Atur created_at jika diperlukan, misalnya:
	payload.CreatedAt = time.Now().Format(time.RFC3339)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Tambahkan data ke Firestore; Firestore akan mengenerate document ID secara otomatis.
	docRef, _, err := coll.Add(ctx, map[string]interface{}{
		"name":       payload.Name,
		"email":      payload.Email,
		"message":    payload.Message,
		"created_at": payload.CreatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"inserted_id": docRef.ID,
	})
}
