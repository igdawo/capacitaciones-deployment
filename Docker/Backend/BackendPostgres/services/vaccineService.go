package services

import (
	"errors"
	"log"
	"rest-template/config"
	"rest-template/models"
	"time"
)

// Función para crear una vacuna e insertarla a la base de datos
func CreateVaccineService(newVaccine models.Vaccine) (models.Vaccine, error) {
	//Se establece conexión con la base de datos
	dbConnection := config.DBConnection()
	//Inserta la vacuna en la base de datos
	resultado := dbConnection.Create(newVaccine)
	//Si hubo un error
	if resultado.Error != nil {
		return newVaccine, resultado.Error
	}
	// Devuelve la vacuna creada
	return newVaccine, resultado.Error
}

// Función para obtener una vacuna por ID
func GetVaccineByIDService(vaccineID string) (models.Vaccine, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Variable para la vacuna a encontrar
	var vacuna models.Vaccine
	// Búsqueda de la vacuna
	result := dbConnection.First(&vacuna, vaccineID)
	if result.Error != nil {
		log.Println("No fue posible convertir encontrar la vacuna")
		return vacuna, result.Error
	}
	log.Println("Vacuna encontrada")
	// Devuelve la vacuna encontrada
	return vacuna, nil

}

// Función para eliminar una vacuna de la base de datos
func DeleteVaccineService(vaccineID string) error {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Elimina la vacuna
	result := dbConnection.Delete(&models.Vaccine{}, vaccineID)
	// Si no hay error
	if result.Error != nil {
		// Se pudo eliminar la vacuna
		return result.Error
	}
	// Mensaje de que se puedo eliminar la vacuna
	log.Println("Vacuna eliminada")
	return errors.New("vacuna eliminada")
}

// Función para actualizar a un vacuna según su ID
func UpdateVaccineService(updatedVaccine models.Vaccine, VaccineID string) (models.Vaccine, error) {
	// Se crea el modelo del vacuna que se actualizará
	var resultVaccine models.Vaccine
	// Se actualiza la fecha de actualización
	resultVaccine.UpdatedAt = time.Now()
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Obtiene la colección de vacunas
	result := dbConnection.First(&resultVaccine,VaccineID)
	// Si hay error
	if result.Error != nil {
		return resultVaccine, result.Error
	}
	// Se actualizan los datos
	if updatedVaccine.Name != ""{
		resultVaccine.Name = updatedVaccine.Name
	}
	if updatedVaccine.Date.IsZero() {
		resultVaccine.Date = updatedVaccine.Date
	}
	if updatedVaccine.Dog != 0{
		resultVaccine.Dog = updatedVaccine.Dog
	}
	// Se actualiza en la base de datos
	save := dbConnection.Save(&resultVaccine)
	// Si hay error
	if save.Error != nil {
		return resultVaccine, save.Error
	}
	log.Println("Vacuna actualizada")
	return resultVaccine, nil
}

// Función para obtener todas las vacunas
func GetAllVaccinesService() ([]models.Vaccine, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Variable que contiene a todas las vacunas
	var vaccines []models.Vaccine
	// Trae a todos los vacunas a la variable
	result := dbConnection.Find(&vaccines)
	if result.Error != nil {
		return vaccines, errors.New("no fue posible traer a todos los vacunas")
	}
	return vaccines, nil
}