package models

type Dog struct {
	ID    int    `gorm:"column:id" gorm:"primarykey" json:"_id"`
	Name  string `gorm:"column:Nombre" json:"Nombre"`
	Breed string `gorm:"column:Raza"   json:"Raza"`
	Color string `gorm:"column:Color"  json:"Color"`
	Age   int    `gorm:"column:Edad"   json:"Edad"`
	Owner int    `gorm:"column:idDueño"  json:"idDueño"`
}

func (Dog) TableName() string {
	return "Perro"
}
