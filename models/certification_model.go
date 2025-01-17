// models/certification.go
package models

type Certification struct {
	ID                  string `json:"id,omitempty"` // Field ini tidak diisi saat Create, akan di-set saat GET
	Name                string `json:"name"`
	IssuingOrganization string `json:"issuing_organization"`
	IssueDate           string `json:"issue_date,omitempty"`
	ExpirationDate      string `json:"expiration_date,omitempty"`
	Description         string `json:"description,omitempty"`
}
