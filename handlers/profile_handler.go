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

func ApiAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is alive"})
}

// GET /profiles
func GetAllProfiles(c *gin.Context) {
	coll := config.GetCollection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	iter := coll.Documents(ctx)
	defer iter.Stop()

	var results []models.Profile
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var profile models.Profile
		if err := doc.DataTo(&profile); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		profile.ID = doc.Ref.ID
		results = append(results, profile)
	}

	c.JSON(http.StatusOK, results)
}

// GET /profiles/:id
func GetProfileByID(c *gin.Context) {
	coll := config.GetCollection("profiles")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var profile models.Profile
	if err := docSnap.DataTo(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	profile.ID = docSnap.Ref.ID

	c.JSON(http.StatusOK, profile)
}

// POST /profiles
func CreateProfile(c *gin.Context) {
	coll := config.GetCollection("profiles")
	var payload models.Profile

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simpan data ke Firestore; Firestore meng-generate Document ID secara otomatis.
	docRef, _, err := coll.Add(ctx, map[string]interface{}{
		"full_name":           payload.FullName,
		"headline":            payload.Headline,
		"about":               payload.About,
		"profile_picture_url": payload.ProfilePictureURL,
		"cv_url":              payload.CVURL,
		"updated_at":          payload.UpdatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": docRef.ID})
}

// PUT /profiles/:id
func UpdateProfile(c *gin.Context) {
	coll := config.GetCollection("profiles")
	id := c.Param("id")

	var payload models.Profile
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Kumpulkan field yang akan diupdate ke dalam slice []firestore.Update
	updates := []firestore.Update{}
	if payload.FullName != "" {
		updates = append(updates, firestore.Update{Path: "full_name", Value: payload.FullName})
	}
	if payload.Headline != "" {
		updates = append(updates, firestore.Update{Path: "headline", Value: payload.Headline})
	}
	if payload.About != "" {
		updates = append(updates, firestore.Update{Path: "about", Value: payload.About})
	}
	if payload.ProfilePictureURL != "" {
		updates = append(updates, firestore.Update{Path: "profile_picture_url", Value: payload.ProfilePictureURL})
	}
	if payload.CVURL != "" {
		updates = append(updates, firestore.Update{Path: "cv_url", Value: payload.CVURL})
	}
	if payload.UpdatedAt != "" {
		updates = append(updates, firestore.Update{Path: "updated_at", Value: payload.UpdatedAt})
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
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}

// DELETE /profiles/:id
func DeleteProfile(c *gin.Context) {
	coll := config.GetCollection("profiles")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted"})
}
