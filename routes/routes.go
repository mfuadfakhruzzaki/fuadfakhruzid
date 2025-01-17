package routes

import (
	"my-gin-mongo/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    // Profile
    r.GET("/profiles", handlers.GetAllProfiles)
    r.GET("/profiles/:id", handlers.GetProfileByID)
    r.POST("/profiles", handlers.CreateProfile)
    r.PUT("/profiles/:id", handlers.UpdateProfile)
    r.DELETE("/profiles/:id", handlers.DeleteProfile)

    // Education
    r.GET("/educations", handlers.GetAllEducations)
    r.GET("/educations/:id", handlers.GetEducationByID)
    r.POST("/educations", handlers.CreateEducation)
    r.PUT("/educations/:id", handlers.UpdateEducation)
    r.DELETE("/educations/:id", handlers.DeleteEducation)

    // Experience
    r.GET("/experiences", handlers.GetAllExperiences)
    r.GET("/experiences/:id", handlers.GetExperienceByID)
    r.POST("/experiences", handlers.CreateExperience)
    r.PUT("/experiences/:id", handlers.UpdateExperience)
    r.DELETE("/experiences/:id", handlers.DeleteExperience)

    // Certification
    r.GET("/certifications", handlers.GetAllCertifications)
    r.GET("/certifications/:id", handlers.GetCertificationByID)
    r.POST("/certifications", handlers.CreateCertification)
    r.PUT("/certifications/:id", handlers.UpdateCertification)
    r.DELETE("/certifications/:id", handlers.DeleteCertification)

    // Project
    r.GET("/projects", handlers.GetAllProjects)
    r.GET("/projects/:id", handlers.GetProjectByID)
    r.POST("/projects", handlers.CreateProject)
    r.PUT("/projects/:id", handlers.UpdateProject)
    r.DELETE("/projects/:id", handlers.DeleteProject)

    // Honor
    r.GET("/honors", handlers.GetAllHonors)
    r.GET("/honors/:id", handlers.GetHonorByID)
    r.POST("/honors", handlers.CreateHonor)
    r.PUT("/honors/:id", handlers.UpdateHonor)
    r.DELETE("/honors/:id", handlers.DeleteHonor)

    // Contact
    r.GET("/contact", handlers.GetAllContact)
    r.POST("/contact", handlers.CreateContact)

    // Upload
    r.POST("/profiles/:id/upload-picture", handlers.UploadProfilePicture)
    r.POST("/profiles/:id/upload-cv", handlers.UploadCV)

}
