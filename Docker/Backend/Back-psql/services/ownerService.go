package services

import (
	"errors"
	"log"
	"reflect"
	"rest-template/config"
	"rest-template/models"
)

// Función para crear un Owner en la base de datos
func CreateOwnerService(newOwner models.Owner) (models.Owner, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Variable que contiene a el ultimo owner creado
	var ownerLast models.Owner
	// Trae a el ultimo owner creado
	resultAux := db.Last(&ownerLast)
	//En caso de algun error
	if resultAux.Error != nil {
		ownerLast.ID = 0
	}
	// Genera un nuevo ID, que es el de el ultimo owner aumentado en 1
	newOwner.ID = ownerLast.ID + 1
	// Inserta el Owner en la colección.
	result := db.Create(&newOwner)
	//En caso de algun error
	if result.Error != nil {
		return newOwner, result.Error
	}
	return newOwner, nil
}

// Función para obtener un Owner por id
func GetOwnerByIDService(ownerID string) (models.Owner, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Crea un objeto ID de MongoDB a partir del ID del Owner.
	var Owner models.Owner
	// Obtiene el Owner.
	result := db.First(&Owner, ownerID)
	//En caso de algun error
	if result.Error != nil {
		return Owner, result.Error
	}
	log.Println("Se encontró el Owner")
	// Devuelve el Owner encontrado
	return Owner, nil

}

// Función para Borrar un Owner por id
func DeleteOwnerService(ownerID string) error {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Se elimina el Owner
	result := db.Delete(&models.Owner{}, ownerID)
	// En caso de algun error
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return errors.New("Owner no encontrado")
	}
	return errors.New("Owners eliminado")
}

// funcion para modificar un Owner
func UpdateOwnerService(updatedOwner models.Owner, ID string) (models.Owner, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)

	// Se recorre el modelo updatedOwner
	modelType := reflect.TypeOf(updatedOwner)
	modelValue := reflect.ValueOf(updatedOwner)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualiza donde los id coincidan
		if value != 0 && value != "" {
			result := db.Model(&models.Owner{}).Where("id="+ID).Update(field.Name, value)
			// En caso de algun error
			if result.Error != nil {
				return updatedOwner, result.Error
			} else if result.RowsAffected < 1 {
				return updatedOwner, errors.New("Owner no encontrado")
			}
		}
	}

	log.Println("Owners actualizado")
	return updatedOwner, nil
}

// Getall de un Owner
func GetAllOwnersService() ([]models.Owner, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Variable que contiene a todos los Owners
	var owners []models.Owner
	// Trae a todos los Owners desde la base de datos
	result := db.Find(&owners)
	//En caso de algun error
	if result.Error != nil {
		return owners, result.Error
	}
	return owners, nil
}

// Funcion para actualizar todos Owners con un o unos atributos
func UpdateOwnerGeneralService(updatedOwner models.Owner) (models.Owner, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)

	// Se recorre el modelo updatedOwner
	modelType := reflect.TypeOf(updatedOwner)
	modelValue := reflect.ValueOf(updatedOwner)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualizan todas las filas
		if value != 0 && value != "" {
			result := db.Model(&models.Owner{}).Where("1=1").Update(field.Name, value)
			// En caso de algun error
			if result.Error != nil {
				return updatedOwner, result.Error
			} else if result.RowsAffected < 1 {
				return updatedOwner, errors.New("Owner no encontrado")
			}
		}
	}

	log.Println("Owners actualizado")
	return updatedOwner, nil
}
