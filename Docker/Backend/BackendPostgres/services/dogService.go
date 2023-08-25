package services

import (
	"errors"
	"log"
	"rest-template/config"
	"rest-template/models"
	"strconv"
)

// Función para crear un perro e insertarlo a la base de datos de postgres
func CreateDogService(newDog models.Dog) (models.Dog, error) {
	//Se establece conexión con la base de datos
	dbConnection := config.DBConnection()
	//Inserta al perro en la base de datos
	resultado:= dbConnection.Create(&newDog)
	//Si hubo un error
	if resultado.Error != nil {
		return newDog, resultado.Error
	}
	return newDog, nil
}

// Función para obtener un perro por id
func GetDogByIDService(dogID string) (models.Dog, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Variable para el perro a encontrar
	var perro models.Dog
	// Se busca al perro correspondiente a ese ID
	resultado:= dbConnection.First(&perro,dogID)
	// Si hay error
	if resultado.Error != nil {
		log.Println("No fue posible encontrar un perro con ese ID")
		return perro, resultado.Error
	}
	log.Println("Perro encontrado")
	// Devuelve el perro encontrado.
	return perro, nil
}

// Función para eliminar un perro de la base de datos de mongodb
func DeleteDogService(dogID string) error {
	// Se establece conexión con la base de datos postgres
	dbConnection := config.DBConnection()
	// Trata de eliminar al perro
	resultado := dbConnection.Delete(&models.Dog{}, dogID)
	// Si hay error
	if resultado.Error != nil {
		return resultado.Error
	}
	return errors.New("perro eliminado")
	
}

// Función para obtener todos los perros
func GetAllDogsService() ([]models.Dog, error) {
	// Se establece conexión con la base de datos postgres
	dbConnection := config.DBConnection()
	// Variable que contiene a todos los perros
	var perros []models.Dog
	// Trae a todos los perros desde la base de datos
	resultado := dbConnection.Find(&perros)
	if resultado.Error != nil {
		return perros, errors.New("no fue posible traer a todos los perros")
	}
	return perros, nil
}

//Funciones especiales solicitadas

// Función para actualizar un perro según su ID
func UpdateDogService(updatedDog models.Dog, dogID string) (models.Dog, error) {
	// Se establece conexión con la base de datos postgres
	dbConnection := config.DBConnection()
	// Se crea el modelo del perro que se actualizara
	var resultDog models.Dog
	// Se busca al perro con el ID ingresado
	resultado := dbConnection.First(&resultDog, dogID)
	// Si hay error
	if resultado.Error != nil {
		return resultDog, resultado.Error
	}
	// Se actualizan los datos
	if updatedDog.Name != ""{
		resultDog.Name = updatedDog.Name
	}
	if updatedDog.Breed != ""{
		resultDog.Breed = updatedDog.Breed
	}
	if updatedDog.Color != ""{
		resultDog.Color = updatedDog.Color
	}
	if updatedDog.Age != 0 {
		resultDog.Age = updatedDog.Age
	}
	// Se actualiza en la base de datos
	save := dbConnection.Save(&resultDog)
	// Si hay error
	if save.Error != nil {
		return resultDog, save.Error
	}
	log.Println("Perro actualizado")
	return resultDog, nil
}

// Obtener el dueño de un perro
func GetOwnerByDogIDService(dogID string)(models.Owner, error){
	//Variable para guardar al dueño
	var ownerVacio models.Owner
	//Se obtiene al perro correspondiente a la ID ingresada
	perro,err := GetDogByIDService(dogID)
	//Si hay error
	if err != nil {
		return ownerVacio, err
	}
	//Se obtiene el id del dueño
	ownerID := strconv.Itoa(perro.Owner)
	//Se obtiene el dueño correspondiente a esa ID
	owner,err := GetOwnerByIDService(ownerID)
	//Si hay error
	if err != nil {
		return owner, err
	}
	log.Println("Dueño encontrado")
	// Devuelve el dueño encontrado
	return owner, nil
}

// Obtener todas las vacunas de un perro
func GetAllVaccinesByDogIDService(dogID string)([]models.Vaccine, error){
	//Variable para guardar las vacunas
	var vaccinesByDog []models.Vaccine
	//Listado de todas las vacunas
	vaccines,err := GetAllVaccinesService()
	//Si hay error
	if err != nil {
		return vaccinesByDog, err
	}
	for _, vaccine := range vaccines {
		if dogID == strconv.Itoa(vaccine.Dog){
			vaccinesByDog = append(vaccinesByDog, vaccine)
		}
	}
		return vaccinesByDog, err
}