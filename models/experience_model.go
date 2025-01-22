package models

type ExperiencesWrapper struct {
	Experiences []Experience `firestore:"experiences" json:"experiences"`
}

type Experience struct {
	ID             string `firestore:"id,omitempty" json:"id,omitempty"` // Akan diisi saat GET
	Title          string `firestore:"title" json:"title"`
	Company        string `firestore:"company" json:"company"`
	Location       string `firestore:"location" json:"location"`
	StartDate      string `firestore:"start_date" json:"start_date"`
	EndDate        string `firestore:"end_date" json:"end_date"`
	Description    string `firestore:"description" json:"description"`
	ExperienceType string `firestore:"experience_type" json:"experience_type"`
	DateRange      string `firestore:"date_range" json:"date_range"`
}
