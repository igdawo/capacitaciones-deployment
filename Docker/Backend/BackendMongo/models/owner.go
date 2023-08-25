package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string             `json:"Nombre" bson:"Nombre,omitempty"`
	Age  int                `json:"Edad" bson:"Edad,omitempty"`
	Sex  string             `json:"Sexo" bson:"Sexo,omitempty"`
}
