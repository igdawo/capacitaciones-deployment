package models

import (
	"time"
)

type Vaccine struct {
	ID   int       `gorm:"column:id" gorm:"primarykey" json:"_id"`
	Name string    `gorm:"column:NombreVacuna" json:"NombreVacuna"`
	Date time.Time `gorm:"column:Fecha" json:"Fecha" `
	Dog  int       `gorm:"column:idPerro" json:"idPerro"`
}

func (Vaccine) TableName() string {
	return "Vacuna"
}
