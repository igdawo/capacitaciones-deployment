package services

import (
	"errors"
	"log"
	"reflect"
	"rest-template/config"
	"rest-template/models"
)

// Función para crear un Vaccine en la base de datos
func CreateVaccineService(newVaccine models.Vaccine) (models.Vaccine, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Variable que contiene a el ultimo Vaccine creado
	var VaccineLast models.Vaccine
	// Trae a el ultimo Vaccine creado
	resultAux := db.Last(&VaccineLast)
	//En caso de algun error
	if resultAux.Error != nil {
		VaccineLast.ID = 0
	}
	// Genera un nuevo ID, que es el de el ultimo Vaccine aumentado en 1
	newVaccine.ID = VaccineLast.ID + 1
	// Inserta el Vaccine en la colección.
	result := db.Create(&newVaccine)
	//En caso de algun error
	if result.Error != nil {
		return newVaccine, result.Error
	}
	return newVaccine, nil
}

// Función para obtener un Vaccine por id
func GetVaccineByIDService(VaccineID string) (models.Vaccine, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Crea un objeto ID de MongoDB a partir del ID del Vaccine.
	var Vaccine models.Vaccine
	// Obtiene el Vaccine.
	result := db.First(&Vaccine, VaccineID)
	//En caso de algun error
	if result.Error != nil {
		return Vaccine, result.Error
	}
	log.Println("Se encontró el Vaccine")
	// Devuelve el Vaccine encontrado
	return Vaccine, nil

}

// Función para Borrar un Vaccine por id
func DeleteVaccineService(VaccineID string) error {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Se elimina el Vaccine
	result := db.Delete(&models.Vaccine{}, VaccineID)
	// En caso de algun error
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return errors.New("Vaccine no encontrado")
	}
	return errors.New("Vaccines eliminado")
}

// funcion para modificar un Vaccine
func UpdateVaccineService(updatedVaccine models.Vaccine, ID string) (models.Vaccine, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)

	// Se recorre el modelo updatedVaccine
	modelType := reflect.TypeOf(updatedVaccine)
	modelValue := reflect.ValueOf(updatedVaccine)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualiza donde los id coincidan
		if value != 0 && value != "" {
			result := db.Model(&models.Vaccine{}).Where("id="+ID).Update(field.Name, value)
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
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Variable que contiene a todos los Vaccines
	var Vaccines []models.Vaccine
	// Trae a todos los Vaccines desde la base de datos
	result := db.Find(&Vaccines)
	//En caso de algun error
	if result.Error != nil {
		return Vaccines, result.Error
	}
	return Vaccines, nil
}

// Funcion para actualizar todos Vaccines con un o unos atributos
func UpdateVaccineGeneralService(updatedVaccine models.Vaccine) (models.Vaccine, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)

	// Se recorre el modelo updatedVaccine
	modelType := reflect.TypeOf(updatedVaccine)
	modelValue := reflect.ValueOf(updatedVaccine)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualizan todas las filas
		if value != 0 && value != "" {
			result := db.Model(&models.Vaccine{}).Where("1=1").Update(field.Name, value)
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

// Función para obtener un Vaccine por id de un perro
func GetVaccineByDogService(perroID string) ([]models.Vaccine, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)

	// Variable que contiene a todos los Vaccines
	var vaccines []models.Vaccine

	// Filtro para buscar el Vaccine por Dog_ID.
	filterUser := models.Vaccine{Dog: perroID}
	//obtiene las vacunas
	db.Find(&vaccines, filterUser)
	return vaccines, nil
}
