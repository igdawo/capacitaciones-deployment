package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dog struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"Nombre" bson:"Nombre,omitempty"`
	Breed string             `json:"Raza" bson:"Raza,omitempty"`
	Color string             `json:"Color" bson:"Color,omitempty"`
	Age   int                `json:"Edad" bson:"Edad,omitempty"`
	Owner primitive.ObjectID `json:"Dueño_id" bson:"Dueño_id,omitempty"`
}
