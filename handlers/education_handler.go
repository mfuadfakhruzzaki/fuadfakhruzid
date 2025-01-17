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

// GET /educations
func GetAllEducations(c *gin.Context) {
	coll := config.GetCollection("educations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	iter := coll.Documents(ctx)
	defer iter.Stop()

	var results []models.Education
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var edu models.Education
		if err := doc.DataTo(&edu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		edu.ID = doc.Ref.ID
		results = append(results, edu)
	}

	c.JSON(http.StatusOK, results)
}

// GET /educations/:id
func GetEducationByID(c *gin.Context) {
	coll := config.GetCollection("educations")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		// Jika dokumen tidak ditemukan, Firestore akan mengembalikan error
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var edu models.Education
	if err := docSnap.DataTo(&edu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	edu.ID = docSnap.Ref.ID

	c.JSON(http.StatusOK, edu)
}

// POST /educations
func CreateEducation(c *gin.Context) {
	coll := config.GetCollection("educations")
	var payload models.Education

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simpan data ke Firestore; Firestore akan meng-generate document ID secara otomatis.
	docRef, _, err := coll.Add(ctx, map[string]interface{}{
		"institution":   payload.Institution,
		"degree":        payload.Degree,
		"field_of_study": payload.FieldOfStudy,
		"description":   payload.Description,
		"start_year":    payload.StartYear,
		"end_year":      payload.EndYear,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": docRef.ID})
}

// PUT /educations/:id
func UpdateEducation(c *gin.Context) {
	coll := config.GetCollection("educations")
	id := c.Param("id")

	var payload models.Education
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Kumpulkan field yang akan diupdate
	updates := []firestore.Update{}
	if payload.Institution != "" {
		updates = append(updates, firestore.Update{Path: "institution", Value: payload.Institution})
	}
	if payload.Degree != "" {
		updates = append(updates, firestore.Update{Path: "degree", Value: payload.Degree})
	}
	if payload.FieldOfStudy != "" {
		updates = append(updates, firestore.Update{Path: "field_of_study", Value: payload.FieldOfStudy})
	}
	if payload.Description != "" {
		updates = append(updates, firestore.Update{Path: "description", Value: payload.Description})
	}
	if payload.StartYear != 0 {
		updates = append(updates, firestore.Update{Path: "start_year", Value: payload.StartYear})
	}
	if payload.EndYear != 0 {
		updates = append(updates, firestore.Update{Path: "end_year", Value: payload.EndYear})
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
		// Jika dokumen tidak ditemukan
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education updated"})
}

// DELETE /educations/:id
func DeleteEducation(c *gin.Context) {
	coll := config.GetCollection("educations")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Education not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Education deleted"})
}
