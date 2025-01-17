// models/project.go
package models

type Project struct {
	ID          string `json:"id,omitempty"` // Akan di-set saat GET dari Firestore
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ProjectURL  string `json:"project_url"`
	TechStack   string `json:"tech_stack"`
}
