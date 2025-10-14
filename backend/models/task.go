package models

import "time"

type Task struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	PostedBy    string    `bson:"posted_by" json:"posted_by"`
	CreatedAt   time.Time `bson:"created_at,omitempty" json:"created_at"`
}
