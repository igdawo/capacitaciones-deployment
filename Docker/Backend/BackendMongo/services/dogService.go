package services

import (
	"errors"
	"log"
	"rest-template/config"
	"rest-template/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Se establece el nombre de la colección que se traera desde la base de datos
*/
const (
	CollectionNameDog = "perro"
)

// Función para crear un perro e insertarlo a la base de datos de mongodb
func CreateDogService(newDog models.Dog) (models.Dog, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función
	defer dbConnection.Close()
	// Obtiene la colección de perros.
	collection := dbConnection.GetCollection(CollectionNameDog)

	// Genera un nuevo ID único para el perro
	newDog.ID = primitive.NewObjectID()

	// Inserta el perro en la colección
	_, err := collection.InsertOne(dbConnection.Context, newDog)
	//Si hubo un error
	if err != nil {
		return newDog, err
	}

	return newDog, err
}

// Función para obtener un perro por id
func GetDogByIDService(dogID string) (models.Dog, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del perro
	var perro models.Dog
	oid, err := primitive.ObjectIDFromHex(dogID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return perro, err
	}
	// Crea un filtro para buscar el perro por su ID
	filter := bson.M{"_id": oid}

	// Obtiene la colección de perros.
	collection := dbConnection.GetCollection(CollectionNameDog)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&perro)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado.
			log.Println("Perro no encontrado")
			return perro, err
		}
		// Ocurrió un error durante la búsqueda.
		return perro, err
	}
	log.Println("Perro encontrado")
	// Devuelve el perro encontrado.
	return perro, nil
}

// Función para eliminar un perro de la base de datos de mongodb
func DeleteDogService(dogID string) error {
	// Obtiene el ID del perro a partir del parámetro de la ruta
	// Crea un objeto ID de MongoDB a partir del ID del perro
	oid, err := primitive.ObjectIDFromHex(dogID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return errors.New("id invalido")
	}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Se elimina el perro
	filter := bson.M{"_id": oid}
	collection := dbConnection.GetCollection(CollectionNameDog)
	// Elimina el perro de la colección
	result, _ := collection.DeleteOne(dbConnection.Context, filter)
	// Si no hay error
	if result.DeletedCount == 1 {
		// Se pudo eliminar el perro
		return nil
	}
	// No se pudo eliminar el perro
	return errors.New("no se pudo eliminar al perro de la BD")
}

// Función para actualizar un perro según su ID
func UpdateDogService(updatedDog models.Dog, userID string) (models.Dog, error) {
	// Se crea el modelo del perro que se actualizara
	var resultDog models.Dog
	//Se crea el objectId de mongo
	oid, err := primitive.ObjectIDFromHex(userID)
	//Si ocurrio un error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return resultDog, err
	}
	// Se actualiza la fecha de actualización
	update := bson.M{"$set": updatedDog}
	filter := bson.M{"_id": oid}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función
	defer dbConnection.Close()
	// Obtiene la colección de perros
	collection := dbConnection.GetCollection(CollectionNameDog)
	_, err = collection.UpdateOne(dbConnection.Context, filter, update)
	if err != nil {
		return resultDog, err
	}
	log.Println("Perro actualizado")
	return resultDog, nil
}

// Función para obtener todos los perros
func GetAllDogsService() ([]models.Dog, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameDog)
	// Variable que contiene a todos los perros
	var dogs []models.Dog
	// Trae a todos los perros desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return dogs, errors.New("no fue posible traer a todos los perros")
	}
	for results.Next(dbConnection.Context) {
		var singleDog models.Dog
		if err = results.Decode(&singleDog); err != nil {
			log.Println("Perro no se pudo añadir")
		}

		dogs = append(dogs, singleDog)
	}
	return dogs, nil
}

// Obtener el dueño de un perro
func GetOwnerByDogIDService(dogID string)(models.Owner, error){
	//Variable para guardar al dueño
	var owner models.Owner
	//Se obtiene al perro correspondiente a la ID ingresada
	dog,err := GetDogByIDService(dogID)
	//Si hay error
	if err != nil {
		return owner, err
	}
	//Se obtiene el id del dueño
	ownerID := dog.Owner.Hex()
	//Se obtiene el dueño correspondiente a esa ID
	owner,err = GetOwnerByIDService(ownerID)
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
		if dogID == vaccine.Dog.Hex(){
			vaccinesByDog = append(vaccinesByDog, vaccine)
		}
	}
		return vaccinesByDog, err
}