package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"nombre" bson:"nombre,omitempty"`
	Age       int                `json:"edad" bson:"edad,omitempty"`
	Sex       string             `json:"sexo" bson:"sexo,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}