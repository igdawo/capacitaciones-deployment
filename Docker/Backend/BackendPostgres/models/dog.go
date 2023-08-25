package models

type Dog struct {
	ID    int    `gorm:"column:id"     json:"id"`
	Name  string `gorm:"column:nombre" json:"nombre"`
	Breed string `gorm:"column:raza"   json:"raza"`
	Color string `gorm:"column:color"  json:"color"`
	Age   int    `gorm:"column:edad"   json:"edad"`
	Owner int    `gorm:"column:id_dueno"  json:"id_dueno"`
}