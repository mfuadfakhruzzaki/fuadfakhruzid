package models

type Education struct {
	ID           string `firestore:"id,omitempty" json:"id,omitempty"`
	Institution  string `firestore:"institution" json:"institution"`
	Degree       string `firestore:"degree" json:"degree"`
	FieldOfStudy string `firestore:"field_of_study" json:"field_of_study"`
	Location     string `firestore:"location" json:"location"`
	StartDate    string `firestore:"start_date" json:"start_date"`
	EndDate      string `firestore:"end_date" json:"end_date"`
	Description  string `firestore:"description" json:"description"`
}
