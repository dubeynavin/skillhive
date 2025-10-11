package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TeamMember struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Role     string             `bson:"role" json:"role"`
	Photo    string             `bson:"photo" json:"photo"`
	LinkedIn string             `bson:"linkedin" json:"linkedin"`
	Bio      string             `bson:"bio" json:"bio"`
}
