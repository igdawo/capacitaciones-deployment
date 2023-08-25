package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"Nombre" bson:"Nombre,omitempty"`
	Gender string             `json:"Sexo" bson:"Sexo,omitempty"`
	Age    int                `json:"Edad" bson:"Edad,omitempty"`
}
