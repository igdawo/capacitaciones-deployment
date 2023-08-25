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
	CollectionNameOwner = "Dueño"
)

// Función para crear un Owner en la base de datos
func CreateOwnerService(newOwner models.Owner) (models.Owner, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Owners.
	collection := dbConnection.GetCollection(CollectionNameOwner)

	// Genera un nuevo ID
	newOwner.ID = primitive.NewObjectID()

	// Inserta el Owner en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newOwner)
	//En caso de algun error
	if err != nil {
		return newOwner, err
	}

	return newOwner, err
}

// Función para obtener un Owner por id
func GetOwnerByIDService(ownerID string) (models.Owner, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del Owner.
	var Owner models.Owner
	oid, err := primitive.ObjectIDFromHex(ownerID)
	//En caso de algun error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return Owner, err
	}
	// Filtro para buscar el Owner por su ID.
	filter := bson.M{"_id": oid}

	// Obtiene el Owner.
	collection := dbConnection.GetCollection(CollectionNameOwner)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&Owner)
	// En caso de algun error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Dueños no encontrado")
			return Owner, err
		}
		return Owner, err
	}
	log.Println("Se encontró el Owner")
	// Devuelve el Owner encontrado
	return Owner, nil

}

func DeleteOwnerService(ownerID string) error {
	// Crea un objeto ID de MongoDB a partir del ID del Owner.
	oid, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return errors.New("id invalido")
	}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()

	// Se elimina el Owner
	filter := bson.M{"_id": oid}
	collection := dbConnection.GetCollection(CollectionNameOwner)
	result, _ := collection.DeleteOne(dbConnection.Context, filter)
	// En caso de algun error
	if result.DeletedCount == 1 {
		return nil
	}
	return errors.New("el Dueños no pudo ser eliminado")
}

func UpdateOwnerService(updatedOwner models.Owner, ID string) (models.Owner, error) {
	// Se crea el modelo del Owner a actualizar
	var resultOwner models.Owner
	//Se crea el objectId de mongo
	oid, err := primitive.ObjectIDFromHex(ID)
	// En caso de algun error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return resultOwner, err
	}
	// Se actualiza
	update := bson.M{"$set": updatedOwner}
	filter := bson.M{"_id": oid}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Owners.
	collection := dbConnection.GetCollection(CollectionNameOwner)
	_, err = collection.UpdateOne(dbConnection.Context, filter, update)
	if err != nil {
		return resultOwner, err
	}
	log.Println("Dueños actualizado")
	return resultOwner, nil
}

func GetAllOwnersService() ([]models.Owner, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameOwner)
	// Variable que contiene a todos los D uenos
	var owners []models.Owner
	// Trae a todos los Owners desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return owners, errors.New("no fue posible traer a todos los Dueños")
	}
	for results.Next(dbConnection.Context) {
		var singleOwner models.Owner
		if err = results.Decode(&singleOwner); err != nil {
			log.Println("Dueño no se pudo añadir")
		}

		owners = append(owners, singleOwner)
	}
	return owners, nil
}

// Funcion para actualizar todos Owners con un o unos atributos
func UpdateOwnerGeneralService(updatedOwner models.Owner) (models.Owner, error) {
	// Se crea el modelo del Owner a actualizar
	var resultOwner models.Owner
	// Se actualiza
	update := bson.M{"$set": updatedOwner}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Owners.
	collection := dbConnection.GetCollection(CollectionNameOwner)
	_, err := collection.UpdateMany(dbConnection.Context, bson.M{}, update)
	if err != nil {
		return resultOwner, err
	}
	log.Println("Dueños actualizado")
	return resultOwner, nil
}
