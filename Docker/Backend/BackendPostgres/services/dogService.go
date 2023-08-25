package services

import (
	"errors"
	"log"
	"reflect"
	"rest-template/config"
	"rest-template/models"
)

// Función para crear un Dog en la base de datos
func CreateDogService(newDog models.Dog) (models.Dog, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Variable que contiene a el ultimo Dog creado
	var DogLast models.Dog
	// Trae a el ultimo Dog creado
	resultAux := db.Last(&DogLast)
	//En caso de algun error
	if resultAux.Error != nil {
		DogLast.ID = 0
	}
	// Genera un nuevo ID, que es el de el ultimo Dog aumentado en 1
	newDog.ID = DogLast.ID + 1
	// Inserta el Dog en la colección.
	result := db.Create(&newDog)
	//En caso de algun error
	if result.Error != nil {
		return newDog, result.Error
	}
	return newDog, nil
}

// Función para obtener un Dog por id
func GetDogByIDService(DogID string) (models.Dog, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Crea un objeto ID de MongoDB a partir del ID del Dog.
	var Dog models.Dog
	// Obtiene el Dog.
	result := db.First(&Dog, DogID)
	//En caso de algun error
	if result.Error != nil {
		return Dog, result.Error
	}
	log.Println("Se encontró el Dog")
	// Devuelve el Dog encontrado
	return Dog, nil

}

// Función para Borrar un Dog por id
func DeleteDogService(DogID string) error {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Se elimina el Dog
	result := db.Delete(&models.Dog{}, DogID)
	// En caso de algun error
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return errors.New("Dog no encontrado")
	}
	return errors.New("Dogs eliminado")
}

// funcion para modificar un Dog
func UpdateDogService(updatedDog models.Dog, ID string) (models.Dog, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)

	// Se recorre el modelo updatedDog
	modelType := reflect.TypeOf(updatedDog)
	modelValue := reflect.ValueOf(updatedDog)
	// Para cada atributo se obtiene su nombre y valor
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i).Interface()
		// si el valor es no nulo, 0 o string vacio, se actualiza donde los id coincidan
		if value != 0 && value != "" {
			result := db.Model(&models.Dog{}).Where("id="+ID).Update(field.Name, value)
			// En caso de algun error
			if result.Error != nil {
				return updatedDog, result.Error
			} else if result.RowsAffected < 1 {
				return updatedDog, errors.New("Dog no encontrado")
			}
		}
	}

	log.Println("Dogs actualizado")
	return updatedDog, nil
}

// Getall de un Dog
func GetAllDogsService() ([]models.Dog, error) {
	// Se establece conexión
	db := config.DBConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer config.DBclose(db)
	// Variable que contiene a todos los Dogs
	var Dogs []models.Dog
	// Trae a todos los Dogs desde la base de datos
	result := db.Find(&Dogs)
	//En caso de algun error
	if result.Error != nil {
		return Dogs, result.Error
	}
	return Dogs, nil
}


