package handlers

import (
	"context"
	"net/http"
	"time"

	"my-gin-mongo/config"
	"my-gin-mongo/models"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// GET /honors
func GetAllHonors(c *gin.Context) {
	coll := config.GetCollection("honors")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	iter := coll.Documents(ctx)
	defer iter.Stop()

	var results []models.Honor
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var honor models.Honor
		if err := doc.DataTo(&honor); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		honor.ID = doc.Ref.ID
		results = append(results, honor)
	}

	c.JSON(http.StatusOK, results)
}

// GET /honors/:id
func GetHonorByID(c *gin.Context) {
	coll := config.GetCollection("honors")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		// Jika dokumen tidak ditemukan
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Honor not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var honor models.Honor
	if err := docSnap.DataTo(&honor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	honor.ID = docSnap.Ref.ID

	c.JSON(http.StatusOK, honor)
}

// POST /honors
func CreateHonor(c *gin.Context) {
	coll := config.GetCollection("honors")
	var payload models.Honor

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Tambahkan data ke Firestore; Firestore akan menghasilkan Document ID secara otomatis.
	docRef, _, err := coll.Add(ctx, map[string]interface{}{
		"title":        payload.Title,
		"issuer":       payload.Issuer,
		"date_awarded": payload.DateAwarded,
		"description":  payload.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": docRef.ID})
}

// PUT /honors/:id
func UpdateHonor(c *gin.Context) {
	coll := config.GetCollection("honors")
	id := c.Param("id")

	var payload models.Honor
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Kumpulkan update field ke dalam slice []firestore.Update
	updates := []firestore.Update{}
	if payload.Title != "" {
		updates = append(updates, firestore.Update{Path: "title", Value: payload.Title})
	}
	if payload.Issuer != "" {
		updates = append(updates, firestore.Update{Path: "issuer", Value: payload.Issuer})
	}
	if payload.DateAwarded != "" {
		updates = append(updates, firestore.Update{Path: "date_awarded", Value: payload.DateAwarded})
	}
	if payload.Description != "" {
		updates = append(updates, firestore.Update{Path: "description", Value: payload.Description})
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Update(ctx, updates)
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Honor not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Honor updated"})
}

// DELETE /honors/:id
func DeleteHonor(c *gin.Context) {
	coll := config.GetCollection("honors")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Honor not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Honor deleted"})
}
