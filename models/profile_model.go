package models

type Profile struct {
	ID                string `json:"id,omitempty"`
	FullName          string `json:"full_name"`
	Headline          string `json:"headline"`
	About             string `json:"about"`
	ProfilePictureURL string `json:"profile_picture_url"`
	CVURL             string `json:"cv_url"`
	UpdatedAt         string `json:"updated_at"`
	Location          string `json:"location"`
	ContactNumber     string `json:"contact_number"`
	ContactEmail      string `json:"contact_email"`
	Website           string `json:"website"`
}
