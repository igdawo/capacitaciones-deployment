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

// Función para crear un owner e insertarlo a la base de datos de mongodb
func CreateOwnerService(newOwner models.Owner) (models.Owner, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de dueños.
	collection := dbConnection.GetCollection(CollectionNameOwner)

	// Genera un nuevo ID único para el dueño.
	newOwner.ID = primitive.NewObjectID()

	// Inserta el dueño en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newOwner)
	//Si hubo un error
	if err != nil {
		return newOwner, err
	}

	return newOwner, err
}

// Función para obtener un owner por id
func GetOwnerByIDService(ownerID string) (models.Owner, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del dueño.
	var owner models.Owner
	oid, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return owner, err
	}
	// Crea un filtro para buscar el dueño por su ID.
	filter := bson.M{"_id": oid}

	// Obtiene la colección de dueños.
	collection := dbConnection.GetCollection(CollectionNameOwner)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&owner)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado.
			log.Println("Dueño no encontrado")
			return owner, err
		}
		// Ocurrió un error durante la búsqueda.
		return owner, err
	}
	log.Println("Se encontró el Dueño")
	// Devuelve el owner encontrado.
	return owner, nil

}

func DeleteOwnerService(ownerID string) error {
	// Obtiene el ID del dueño a partir del parámetro de la ruta.
	// Crea un objeto ID de MongoDB a partir del ID del owner.
	oid, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return errors.New("id invalido")
	}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Se elimina el owner
	filter := bson.M{"_id": oid}
	collection := dbConnection.GetCollection(CollectionNameOwner)
	// Elimina el owner de la colección.
	result, _ := collection.DeleteOne(dbConnection.Context, filter)
	// Si no hay error
	if result.DeletedCount == 1 {
		// Se pudo eliminar al dueño
		return nil
	}
	// No se pudo eliminar al dueño
	return errors.New("el dueño no pudo ser eliminado")
}

func UpdateOwnerService(updatedOwner models.Owner, userID string) (models.Owner, error) {
	// Se crea el modelo del owner que se actualizara
	var resultOwner models.Owner
	//Se crea el objectId de mongo
	oid, err := primitive.ObjectIDFromHex(userID)
	//Si ocurrio un error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return resultOwner, err
	}
	// Se actualiza la fecha de actualización

	update := bson.M{"$set": updatedOwner}
	filter := bson.M{"_id": oid}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de dueños.
	collection := dbConnection.GetCollection(CollectionNameOwner)
	_, err = collection.UpdateOne(dbConnection.Context, filter, update)
	if err != nil {
		return resultOwner, err
	}
	log.Println("dueño actualizado")
	return resultOwner, nil
}

func GetAllOwnersService() ([]models.Owner, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameOwner)
	// Variable que contiene a todos los dueños
	var owners []models.Owner
	// Trae a todos los dueños desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return owners, errors.New("No fue posible traer a todos los Dueños")
	}
	for results.Next(dbConnection.Context) {
		var singleOwner models.Owner
		if err = results.Decode(&singleOwner); err != nil {
			log.Println("El dueño no se pudo añadir")
		}

		owners = append(owners, singleOwner)
	}
	return owners, nil
}
