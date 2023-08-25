package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Variables globales
var DB *gorm.DB

// Función que estable conexión a la base de datos
func DBConnection() *gorm.DB {
	//Datos del .env para la conexión a postgres
	dsn := "host=" + os.Getenv("DB_HOST") +" user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_DB") + " port=" + os.Getenv("DB_PORT")
	//Se establece la conexión con la base de datos
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		log.Println("Conexión con DB postgres fallida")
	} else{
		log.Println("Inicio conexión con la base de datos de postgres")
	}
	return DB
}

/*// Función que cierra la conexión a la base de datos
func dbClose(db *gorm.DB) {
	if db.Error !=nil{
		log.Fatal(db.Error)
	} else{
		db.Close()
		log.Println("Fin conexión con la base de datos de postgres")
	}
}*/