package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vaccine struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string             `json:"NombreVacuna" bson:"NombreVacuna,omitempty"`
	Dog  primitive.ObjectID `json:"Perro_id" bson:"Perro_id,omitempty"`
	Date time.Time          `json:"Fecha" bson:"Fecha,omitempty"`
}
