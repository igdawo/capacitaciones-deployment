package models

type Dog struct {
	ID    int    `gorm:"column:id" gorm:"primarykey" json:"id"`
	Name  string `gorm:"column:Nombre" json:"Nombre"`
	Breed string `gorm:"column:Raza"   json:"Raza"`
	Color string `gorm:"column:Color"  json:"Color"`
	Age   int    `gorm:"column:Edad"   json:"Edad"`
	Owner int    `gorm:"column:Dueño_id"  json:"Dueño_id"`
}

func (Dog) TableName() string {
	return "Perro"
}
