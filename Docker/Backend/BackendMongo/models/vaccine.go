package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vaccine struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"nombreVacuna" bson:"nombreVacuna,omitempty"`
	Date      time.Time          `json:"fecha" bson:"fecha,omitempty"`
	Dog       primitive.ObjectID `json:"id_perro" bson:"id_perro,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}