package models

import "time"

type User struct {
	ID             string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string    `bson:"name" json:"name"`
	Email          string    `bson:"email" json:"email"`
	PasswordHash   string    `bson:"password_hash" json:"-"`
	Role           string    `bson:"role" json:"role"` // e.g., freelancer/client
	LinkedIn       string    `bson:"linkedin,omitempty" json:"linkedin,omitempty"`
	PhotoURL       string    `bson:"photo_url,omitempty" json:"photo_url,omitempty"`
	ResetToken     string    `bson:"reset_token,omitempty" json:"-"`
	ResetExpiresAt time.Time `bson:"reset_expires_at,omitempty" json:"-"`
	CreatedAt      time.Time `bson:"created_at,omitempty" json:"created_at"`
}
