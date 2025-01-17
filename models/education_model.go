// models/education.go
package models

type Education struct {
	ID           string `json:"id,omitempty"` // akan diisi saat GET
	Institution  string `json:"institution"`
	Degree       string `json:"degree"`
	FieldOfStudy string `json:"field_of_study"`
	Description  string `json:"description"`
	StartYear    int    `json:"start_year"`
	EndYear      int    `json:"end_year"`
}
