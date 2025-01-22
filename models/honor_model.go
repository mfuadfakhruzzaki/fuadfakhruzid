package models

type Honor struct {
	ID          string `firestore:"id,omitempty" json:"id,omitempty"` // Akan diisi saat GET
	Title       string `firestore:"title" json:"title"`
	Issuer      string `firestore:"issuer" json:"issuer"`
	DateAwarded string `firestore:"date_awarded" json:"date_awarded"`
	Description string `firestore:"description" json:"description"`
}
