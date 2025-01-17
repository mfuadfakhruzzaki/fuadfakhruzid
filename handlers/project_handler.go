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

// GET /projects
func GetAllProjects(c *gin.Context) {
	coll := config.GetCollection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	iter := coll.Documents(ctx)
	defer iter.Stop()

	var results []models.Project
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var proj models.Project
		if err := doc.DataTo(&proj); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		proj.ID = doc.Ref.ID
		results = append(results, proj)
	}

	c.JSON(http.StatusOK, results)
}

// GET /projects/:id
func GetProjectByID(c *gin.Context) {
	coll := config.GetCollection("projects")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var proj models.Project
	if err := docSnap.DataTo(&proj); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	proj.ID = docSnap.Ref.ID

	c.JSON(http.StatusOK, proj)
}

// POST /projects
func CreateProject(c *gin.Context) {
	coll := config.GetCollection("projects")
	var payload models.Project

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Tambahkan data ke Firestore; Firestore akan menghasilkan Document ID secara otomatis.
	docRef, _, err := coll.Add(ctx, map[string]interface{}{
		"title":       payload.Title,
		"description": payload.Description,
		"start_date":  payload.StartDate,
		"end_date":    payload.EndDate,
		"project_url": payload.ProjectURL,
		"tech_stack":  payload.TechStack,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inserted_id": docRef.ID})
}

// PUT /projects/:id
func UpdateProject(c *gin.Context) {
	coll := config.GetCollection("projects")
	id := c.Param("id")

	var payload models.Project
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updates := []firestore.Update{}
	if payload.Title != "" {
		updates = append(updates, firestore.Update{Path: "title", Value: payload.Title})
	}
	if payload.Description != "" {
		updates = append(updates, firestore.Update{Path: "description", Value: payload.Description})
	}
	if payload.StartDate != "" {
		updates = append(updates, firestore.Update{Path: "start_date", Value: payload.StartDate})
	}
	if payload.EndDate != "" {
		updates = append(updates, firestore.Update{Path: "end_date", Value: payload.EndDate})
	}
	if payload.ProjectURL != "" {
		updates = append(updates, firestore.Update{Path: "project_url", Value: payload.ProjectURL})
	}
	if payload.TechStack != "" {
		updates = append(updates, firestore.Update{Path: "tech_stack", Value: payload.TechStack})
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
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated"})
}

// DELETE /projects/:id
func DeleteProject(c *gin.Context) {
	coll := config.GetCollection("projects")
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	docRef := coll.Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		if err != nil && err.Error() == "rpc error: code = NotFound desc = Document not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
