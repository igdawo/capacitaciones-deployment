package services

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"rest-template/config"
	"rest-template/models"
	"strconv"
)

// Función para crear un Vaccine en la base de datos
func CreateVaccineService(newVaccine models.Vaccine) (models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(dbConnection)
	// Variable que contiene a el ultimo Vaccine creado
	var vaccineLast models.Vaccine
	// Trae a el ultimo Vaccine creado
	resultAux := dbConnection.Last(&vaccineLast)
	//En caso de algun error
	if resultAux.Error != nil {
		vaccineLast.ID = 0
	}
	// Genera un nuevo ID, que es el de el ultimo Vaccine aumentado en 1
	newVaccine.ID = vaccineLast.ID + 1
	// Inserta el Vaccine en la colección.
	result := dbConnection.Create(&newVaccine)
	//En caso de algun error
	if result.Error != nil {
		return newVaccine, result.Error
	}
	return newVaccine, nil
}

// Función para obtener un Vaccine por id
func GetVaccineByIDService(vaccineID string) (models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(dbConnection)
	// Crea un objeto ID de MongoDB a partir del ID del Vaccine.
	var Vaccine models.Vaccine
	// Obtiene el Vaccine.
	result := dbConnection.First(&Vaccine, vaccineID)
	//En caso de algun error
	if result.Error != nil {
		return Vaccine, result.Error
	}
	log.Println("Se encontró la Vacuna")
	// Devuelve el Vaccine encontrado
	return Vaccine, nil

}

// Función para Borrar un Vaccine por id
func DeleteVaccineService(vaccineID string) error {
	// Se establece conexión
	dbConnection := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(dbConnection)
	// Se elimina el Vaccine
	result := dbConnection.Delete(&models.Vaccine{}, vaccineID)
	// En caso de algun error
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return errors.New("Vacuna no encontrada")
	}
	return errors.New("Vacuna eliminada")
}

// funcion para modificar un Vaccine
func UpdateVaccineService(updatedVaccine models.Vaccine, ID string) (models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(dbConnection)

	// Se recorre el modelo updatedVaccine
	modelType := reflect.TypeOf(updatedVaccine)
	modelValue := reflect.ValueOf(updatedVaccine)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualiza donde los id coincidan
		if value != 0 && value != "" {
			result := dbConnection.Model(&models.Vaccine{}).Where("id="+ID).Update(field.Name, value)
			// En caso de algun error
			if result.Error != nil {
				return updatedVaccine, result.Error
			} else if result.RowsAffected < 1 {
				return updatedVaccine, errors.New("Vaccine no encontrado")
			}
		}
	}

	log.Println("Vaccines actualizado")
	return updatedVaccine, nil
}

// Getall de un Vaccine
func GetAllVaccinesService() ([]models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(dbConnection)
	// Variable que contiene a todos los Vaccines
	var Vaccines []models.Vaccine
	// Trae a todos los Vaccines desde la base de datos
	result := dbConnection.Find(&Vaccines)
	//En caso de algun error
	if result.Error != nil {
		return Vaccines, result.Error
	}
	return Vaccines, nil
}

// Función para obtener un Vaccine por id de un perro
func GetVaccineByDogService(perroID string) ([]models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(dbConnection)

	// Variable que contiene a todos los Vaccines
	var vaccines []models.Vaccine

	// Filtro para buscar el Vaccine por Dog_ID.
	number, err := strconv.Atoi(perroID)
	if err != nil {
		fmt.Println("Error al convertir la cadena a entero:", err)
	}
	filterUser := models.Vaccine{Dog: number}
	//obtiene las vacunas
	dbConnection.Find(&vaccines, filterUser)
	return vaccines, nil
}

// Función para retornar todas las vacunas de un perro a partir de su id
func GetAllVaccinesByDogIDService(dogID string) ([]models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(dbConnection)

	// Variable que contiene a todos los Vaccines
	var vaccines []models.Vaccine

	// Filtro para buscar el Vaccine por Dog_ID.
	number, err := strconv.Atoi(dogID)
	if err != nil {
		fmt.Println("Error al convertir el ID:", err)
	}
	filterUser := models.Vaccine{Dog: number}
	//obtiene las vacunas
	dbConnection.Find(&vaccines, filterUser)
	return vaccines, nil
}
