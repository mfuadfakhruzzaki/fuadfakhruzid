package models

type Profile struct {
	ID                string `firestore:"id,omitempty" json:"id,omitempty"`
	FullName          string `firestore:"full_name" json:"full_name"`
	Headline          string `firestore:"headline" json:"headline"`
	About             string `firestore:"about" json:"about"`
	ProfilePictureURL string `firestore:"profile_picture_url" json:"profile_picture_url"`
	CVURL             string `firestore:"cv_url" json:"cv_url"`
	UpdatedAt         string `firestore:"updated_at" json:"updated_at"`
	Location          string `firestore:"location" json:"location"`
	ContactNumber     string `firestore:"contact_number" json:"contact_number"`
	ContactEmail      string `firestore:"contact_email" json:"contact_email"`
	Website           string `firestore:"website" json:"website"`
}
