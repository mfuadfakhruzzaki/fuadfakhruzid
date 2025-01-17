// models/profile.go
package models

type Profile struct {
	ID                string `json:"id,omitempty"` // Di-set saat GET dari Firestore
	FullName          string `json:"full_name"`
	Headline          string `json:"headline"`
	About             string `json:"about"`
	ProfilePictureURL string `json:"profile_picture_url"`
	CVURL             string `json:"cv_url"`
	UpdatedAt         string `json:"updated_at"`
}
