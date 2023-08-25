package services

import (
	"errors"
	"log"
	"rest-template/config"
	"rest-template/models"
)

// Función para crear un dueño e insertarlo a la base de datos
func CreateOwnerService(newOwner models.Owner) (models.Owner, error) {
	//Se establece conexión con la base de datos
	dbConnection := config.DBConnection()
	//Inserta al dueño en la base de datos
	resultado := dbConnection.Create(newOwner)
	//Si hubo un error
	if resultado.Error != nil {
		return newOwner, resultado.Error
	}
	// Devuelve el dueño creado
	return newOwner, resultado.Error
}

// Función para obtener un dueño por ID
func GetOwnerByIDService(ownerID string) (models.Owner, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Variable para el dueño a encontrar
	var dueno models.Owner
	result := dbConnection.First(&dueno, ownerID)
	if result.Error != nil {
		log.Println("No fue posible convertir encontrar al dueño")
		return dueno, result.Error
	}
	log.Println("Dueño encontrado")
	// Devuelve el dueño encontrado
	return dueno, nil

}

// Función para eliminar a un dueño de la base de datos
func DeleteOwnerService(ownerID string) error {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	result := dbConnection.Delete(&models.Owner{}, ownerID)
	// Si no hay error
	if result.Error != nil {
		// Se pudo eliminar el dueño
		return result.Error
	}
	// Mensaje de que se puedo eliminar el dueño
	log.Println("Dueño eliminado")
	return nil
}

// Función para actualizar a un dueño según su ID
func UpdateOwnerService(updatedOwner models.Owner, ownerID string) (models.Owner, error) {
	// Se crea el modelo del dueño que se actualizará
	var resultOwner models.Owner
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Obtiene la colección de dueños
	result := dbConnection.First(&resultOwner,ownerID)
	// Si hay error
	if result.Error != nil {
		return resultOwner, result.Error
	}
	// Se actualizan los datos
	if updatedOwner.Name != ""{
		resultOwner.Name = updatedOwner.Name
	}
	if updatedOwner.Age != 0 {
		resultOwner.Age = updatedOwner.Age
	}
	if updatedOwner.Sex != ""{
		resultOwner.Sex = updatedOwner.Sex
	}
	// Se actualiza en la base de datos
	save := dbConnection.Save(&resultOwner)
	// Si hay error
	if save.Error !=nil{
		return resultOwner,save.Error
	}
	log.Println("Dueño actualizado")
	return resultOwner, nil
}

// Función para obtener todos los dueños
func GetAllOwnersService() ([]models.Owner, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.DBConnection()
	// Variable que contiene a todos los dueños
	var owners []models.Owner
	// Trae a todos los dueños desde la base de datos
	result := dbConnection.Find(&owners)
	if result.Error != nil {
		return owners, errors.New("no fue posible traer a todos los dueños")
	}
	return owners, nil
}