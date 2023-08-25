package models

type Owner struct {
	ID   int    `gorm:"column:id"     json:"id"`
	Name string `gorm:"column:nombre" json:"nombre"`
	Age  int    `gorm:"column:sexo"   json:"sexo"`
	Sex  string `gorm:"column:edad"   json:"edad"`
}