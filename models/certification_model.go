package models

type Certification struct {
	ID                  string `firestore:"id,omitempty" json:"id,omitempty"` // Field ini tidak diisi saat Create, akan di-set saat GET
	Name                string `firestore:"name" json:"name"`
	IssuingOrganization string `firestore:"issuing_organization" json:"issuing_organization"`
	IssueDate           string `firestore:"issue_date,omitempty" json:"issue_date,omitempty"`
	ExpirationDate      string `firestore:"expiration_date,omitempty" json:"expiration_date,omitempty"`
	Description         string `firestore:"description,omitempty" json:"description,omitempty"`
}
