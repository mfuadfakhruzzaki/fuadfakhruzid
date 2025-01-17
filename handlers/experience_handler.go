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

// GET /experiences
func GetAllExperiences(c *gin.Context) {
	coll := config.GetCollection("experiences")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	iter := coll.Documents(ctx)
	defer iter.Stop()

	var results []models.Experience
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var exp models.Experience
		if err := doc.DataTo(&exp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Set document ID ke field ID model
		exp.ID = doc.Ref.ID
		results = append(results, exp)
	}

	c.JSON(http.StatusOK, results)
}

// GET /experiences/:id
func GetExperienceByID(c *gin.Context) {
	coll := config.GetCollection("experiences")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var exp models.Experience
	if err := docSnap.DataTo(&exp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	exp.ID = docSnap.Ref.ID

	c.JSON(http.StatusOK, exp)
}

// POST /experiences
func CreateExperience(c *gin.Context) {
	coll := config.GetCollection("experiences")
	var payload models.Experience

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simpan data ke Firestore; Firestore akan meng-generate document ID secara otomatis.
	docRef, _, err := coll.Add(ctx, map[string]interface{}{
		"title":       payload.Title,
		"company":     payload.Company,
		"location":    payload.Location,
		"start_date":  payload.StartDate,
		"end_date":    payload.EndDate,
		"description": payload.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": docRef.ID})
}

// PUT /experiences/:id
func UpdateExperience(c *gin.Context) {
	coll := config.GetCollection("experiences")
	id := c.Param("id")

	var payload models.Experience
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updates := []firestore.Update{}
	if payload.Title != "" {
		updates = append(updates, firestore.Update{Path: "title", Value: payload.Title})
	}
	if payload.Company != "" {
		updates = append(updates, firestore.Update{Path: "company", Value: payload.Company})
	}
	if payload.Location != "" {
		updates = append(updates, firestore.Update{Path: "location", Value: payload.Location})
	}
	if payload.StartDate != "" {
		updates = append(updates, firestore.Update{Path: "start_date", Value: payload.StartDate})
	}
	if payload.EndDate != "" {
		updates = append(updates, firestore.Update{Path: "end_date", Value: payload.EndDate})
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
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience updated"})
}

// DELETE /experiences/:id
func DeleteExperience(c *gin.Context) {
	coll := config.GetCollection("experiences")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Experience not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted"})
}
