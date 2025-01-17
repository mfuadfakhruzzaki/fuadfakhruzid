// models/honor.go
package models

type Honor struct {
	ID          string `json:"id,omitempty"` // Akan diisi saat GET
	Title       string `json:"title"`
	Issuer      string `json:"issuer"`
	DateAwarded string `json:"date_awarded"`
	Description string `json:"description"`
}
