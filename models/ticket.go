package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Category  string             `json:"category,omitempty" bson:"category,omitempty"`
	Body      string             `json:"body,omitempty" bson:"body,omitempty"`
	UserID    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
