// models/contact.go
package models

type Contact struct {
	ID        string `json:"id,omitempty"` // akan diisi saat GET
	Name      string `json:"name"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at,omitempty"`
}
