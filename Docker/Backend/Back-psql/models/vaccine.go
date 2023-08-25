package models

import (
	"time"
)

type Vaccine struct {
	ID   int       `gorm:"column:id" gorm:"primarykey" json:"id"`
	Name string    `gorm:"column:NombreVacuna" json:"NombreVacuna"`
	Dog  string    `gorm:"column:Perro_id" json:"Perro_id"`
	Date time.Time `gorm:"column:Fecha" json:"Fecha" `
}

func (Vaccine) TableName() string {
	return "Vacuna"
}
