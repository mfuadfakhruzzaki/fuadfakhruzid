package handlers

import (
	"context"
	"net/http"
	"time"

	"my-gin-mongo/config" // pastikan path-nya sesuai dengan struktur project Anda
	"my-gin-mongo/models" // pastikan model Certification sudah didefinisikan di sini

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// GET /certifications
func GetAllCertifications(c *gin.Context) {
	coll := config.GetCollection("certifications")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ambil semua dokumen dari koleksi
	iter := coll.Documents(ctx)
	defer iter.Stop()

	var results []models.Certification

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var cert models.Certification
		// Perlu meng-assign ID dari dokumen ke field ID pada model jika diperlukan
		if err := doc.DataTo(&cert); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Jika model Anda memiliki field ID, Anda bisa menyet nilainya:
		cert.ID = doc.Ref.ID

		results = append(results, cert)
	}

	c.JSON(http.StatusOK, results)
}

// GET /certifications/:id
func GetCertificationByID(c *gin.Context) {
	coll := config.GetCollection("certifications")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		// Jika dokumen tidak ditemukan, Firestore akan mengembalikan error
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Certification not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var cert models.Certification
	if err := docSnap.DataTo(&cert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Set ID dokumen
	cert.ID = docSnap.Ref.ID

	c.JSON(http.StatusOK, cert)
}

// POST /certifications
func CreateCertification(c *gin.Context) {
	coll := config.GetCollection("certifications")
	var payload models.Certification

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Buat context dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Tambahkan data ke Firestore. Firestore akan mengenerate Document ID secara otomatis.
	docRef, _, err := coll.Add(ctx, map[string]interface{}{
		"name":                  payload.Name,
		"issuing_organization":  payload.IssuingOrganization,
		"issue_date":            payload.IssueDate,
		"expiration_date":       payload.ExpirationDate,
		"description":           payload.Description,
		// Tambahkan field lain jika perlu
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": docRef.ID})
}

// PUT /certifications/:id
func UpdateCertification(c *gin.Context) {
	coll := config.GetCollection("certifications")
	id := c.Param("id")

	var payload models.Certification
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updateData := []firestore.Update{}

	if payload.Name != "" {
		updateData = append(updateData, firestore.Update{Path: "name", Value: payload.Name})
	}
	if payload.IssuingOrganization != "" {
		updateData = append(updateData, firestore.Update{Path: "issuing_organization", Value: payload.IssuingOrganization})
	}
	if payload.IssueDate != "" {
		updateData = append(updateData, firestore.Update{Path: "issue_date", Value: payload.IssueDate})
	}
	if payload.ExpirationDate != "" {
		updateData = append(updateData, firestore.Update{Path: "expiration_date", Value: payload.ExpirationDate})
	}
	if payload.Description != "" {
		updateData = append(updateData, firestore.Update{Path: "description", Value: payload.Description})
	}

	if len(updateData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Update(ctx, updateData)
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Certification not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certification updated"})
}

// DELETE /certifications/:id
func DeleteCertification(c *gin.Context) {
	coll := config.GetCollection("certifications")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Certification not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certification deleted"})
}
