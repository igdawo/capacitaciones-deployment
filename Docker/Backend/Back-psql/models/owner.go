package models

type Owner struct {
	ID     int    `gorm:"column:id" gorm:"primarykey" json:"id"`
	Name   string `gorm:"column:Nombre" json:"Nombre"`
	Gender string `gorm:"column:Sexo"   json:"Sexo"`
	Age    int    `gorm:"column:Edad"   json:"Edad"`
}

func (Owner) TableName() string {
	return "Dueño"
}
