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
Se establecen los nombres de la colección que se traeran desde la base de datos
*/
const (
	CollectionNameDog = "Perro"
)

// Función para crear un Dog en la base de datos
func CreateDogService(newDog models.Dog) (models.Dog, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Dogs.
	collection := dbConnection.GetCollection(CollectionNameDog)

	// Genera un nuevo ID
	newDog.ID = primitive.NewObjectID()

	// Inserta el Dog en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newDog)
	//En caso de algun error
	if err != nil {
		return newDog, err
	}

	return newDog, err
}

// Función para obtener un Dog por id
func GetDogByIDService(dogID string) (models.Dog, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del Dog.
	var Dog models.Dog
	oid, err := primitive.ObjectIDFromHex(dogID)
	//En caso de algun error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return Dog, err
	}
	// Filtro para buscar el Dog por su ID.
	filter := bson.M{"_id": oid}

	// Obtiene el Dog.
	collection := dbConnection.GetCollection(CollectionNameDog)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&Dog)
	// En caso de algun error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Dog no encontrado")
			return Dog, err
		}
		return Dog, err
	}
	log.Println("Se encontró el Dog")
	// Devuelve el Dog encontrado
	return Dog, nil

}

func DeleteDogService(dogID string) error {
	// Crea un objeto ID de MongoDB a partir del ID del Dog.
	oid, err := primitive.ObjectIDFromHex(dogID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return errors.New("id invalido")
	}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()

	// Se elimina el Dog
	filter := bson.M{"_id": oid}
	collection := dbConnection.GetCollection(CollectionNameDog)
	result, _ := collection.DeleteOne(dbConnection.Context, filter)
	// En caso de algun error
	if result.DeletedCount == 1 {
		return nil
	}
	return errors.New("el Dog no pudo ser eliminado")
}

func UpdateDogService(updatedDog models.Dog, ID string) (models.Dog, error) {
	// Se crea el modelo del Dog a actualizar
	var resultDog models.Dog
	//Se crea el objectId de mongo
	oid, err := primitive.ObjectIDFromHex(ID)
	// En caso de algun error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return resultDog, err
	}
	// Se actualiza
	update := bson.M{"$set": updatedDog}
	filter := bson.M{"_id": oid}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Dogs.
	collection := dbConnection.GetCollection(CollectionNameDog)
	_, err = collection.UpdateOne(dbConnection.Context, filter, update)
	if err != nil {
		return resultDog, err
	}
	log.Println("Dog actualizado")
	return resultDog, nil
}

func GetAllDogsService() ([]models.Dog, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameDog)
	// Variable que contiene a todos los Dogs
	var dogs []models.Dog
	// Trae a todos los Dogs desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return dogs, errors.New("no fue posible traer a todos los Dogs")
	}
	for results.Next(dbConnection.Context) {
		var singleDog models.Dog
		if err = results.Decode(&singleDog); err != nil {
			log.Println("Dog no se pudo añadir")
		}

		dogs = append(dogs, singleDog)
	}
	return dogs, nil
}

// Funcion para actualizar todos Dogs con un o unos atributos
func UpdateDogGeneralService(updatedDog models.Dog) (models.Dog, error) {
	// Se crea el modelo del Dog a actualizar
	var resultDog models.Dog
	// Se actualiza
	update := bson.M{"$set": updatedDog}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Dogs.
	collection := dbConnection.GetCollection(CollectionNameDog)
	_, err := collection.UpdateMany(dbConnection.Context, bson.M{}, update)
	if err != nil {
		return resultDog, err
	}
	log.Println("Dog actualizado")
	return resultDog, nil
}
