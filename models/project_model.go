package models

type Project struct {
	ID          string `firestore:"id,omitempty" json:"id,omitempty"` // Akan di-set saat GET dari Firestore
	Title       string `firestore:"title" json:"title"`
	Description string `firestore:"description" json:"description"`
	StartDate   string `firestore:"start_date" json:"start_date"`
	EndDate     string `firestore:"end_date" json:"end_date"`
	ProjectURL  string `firestore:"project_url" json:"project_url"`
	TechStack   string `firestore:"tech_stack" json:"tech_stack"`
}
