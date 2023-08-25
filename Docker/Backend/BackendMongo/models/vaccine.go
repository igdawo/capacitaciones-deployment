package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vaccine struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Date        time.Time          `json:"Fecha" bson:"Fecha,omitempty"`
	VaccineName string             `json:"NombreVacuna" bson:"NombreVacuna,omitempty"`
	Dog         primitive.ObjectID `json:"idPerro" bson:"id_Perro,omitempty"`
}
