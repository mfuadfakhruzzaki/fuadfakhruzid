// models/experience.go
package models

type Experience struct {
	ID          string `json:"id,omitempty"` // Akan diisi saat GET
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}
