package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dog struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"nombre" bson:"nombre,omitempty"`
	Breed     string             `json:"raza" bson:"raza,omitempty"`
	Age       int                `json:"edad" bson:"edad,omitempty"`
	Owner     primitive.ObjectID `json:"id_dueno" bson:"id_dueno,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}