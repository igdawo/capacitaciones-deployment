package config

import (
	"database/sql"
	"log"
	"os"
	"rest-template/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var db *sql.DB
var err error

// Funcion para la conectar a la base de datos
func DBConnection() *gorm.DB {
	// extrae los datos del .env
	var DBenv = "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_DB") + " port=" + os.Getenv("DB_PORT")
	DB, err = gorm.Open(postgres.Open(DBenv), &gorm.Config{
		QueryFields: true,
	})
	// En caso de algun error
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection successful")
	}
	DB.AutoMigrate(&models.Owner{})
	DB.AutoMigrate(&models.Dog{})
	DB.AutoMigrate(&models.Vaccine{})
	return DB
}

// Funcion para cerrar la conexion
// El cerrar la conexion crea problemas con la vista, por ello se anula la funcion.
func DBclose(*gorm.DB) {
	// db, err = DB.DB()
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	db.Close()
	// 	log.Println("Database close successful")
	// }
}
