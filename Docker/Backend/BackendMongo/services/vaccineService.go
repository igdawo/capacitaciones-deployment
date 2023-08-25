package services

import (
	"errors"
	"log"
	"rest-template/config"
	"rest-template/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Se establece el nombre de la colección que se traera desde la base de datos
*/
const (
	CollectionNameVaccine = "vacuna"
)

// Función para crear una vacuna e insertarla a la base de datos de mongodb
func CreateVaccineService(newVaccine models.Vaccine) (models.Vaccine, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función
	defer dbConnection.Close()
	// Obtiene la colección de vacunas
	collection := dbConnection.GetCollection(CollectionNameVaccine)

	// Genera un nuevo ID único para la vacuna
	newVaccine.ID = primitive.NewObjectID()
	// Establece la fecha de creación y actualización de la vacuna
	newVaccine.CreatedAt = time.Now()
	newVaccine.UpdatedAt = time.Now()

	// Inserta la vacuna en la colección
	_, err := collection.InsertOne(dbConnection.Context, newVaccine)
	//Si hubo un error
	if err != nil {
		return newVaccine, err
	}

	return newVaccine, err
}

// Función para obtener una vacuna por ID
func GetVaccineByIDService(vaccineID string) (models.Vaccine, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID de la vacuna
	var vacuna models.Vaccine
	oid, err := primitive.ObjectIDFromHex(vaccineID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return vacuna, err
	}
	// Crea un filtro para buscar la vacuna por su ID
	filter := bson.M{"_id": oid}

	// Obtiene la colección de vacunas
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&vacuna)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado
			log.Println("Vacuna no encontrada")
			return vacuna, err
		}
		// Ocurrió un error durante la búsqueda
		return vacuna, err
	}
	log.Println("Vacuna encontrada")
	// Devuelve el vacuna encontrada
	return vacuna, nil

}

// Función para eliminar una vacuna de la base de datos de mongodb
func DeleteVaccineService(vaccineID string) error {
	// Obtiene el ID de la vacuna a partir del parámetro de la ruta
	// Crea un objeto ID de MongoDB a partir del ID de la vacuna
	oid, err := primitive.ObjectIDFromHex(vaccineID)
	if err != nil {
		log.Println("No fue posible convertir el ID.")
		return errors.New("id invalido")
	}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Se elimina la vacuna
	filter := bson.M{"_id": oid}
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	// Elimina la vacuna de la colección
	result, _ := collection.DeleteOne(dbConnection.Context, filter)
	// Si no hay error
	if result.DeletedCount == 1 {
		// Se pudo eliminar la vacuna
		return nil
	}
	// No se pudo eliminar la vacuna
	return errors.New("no se pudo eliminar la vacuna de la BD")
}

// Función para actualizar una vacuna según su ID
func UpdateVaccineService(updatedVaccine models.Vaccine, userID string) (models.Vaccine, error) {
	// Se crea el modelo de la vacuna que se actualizará
	var resultVaccine models.Vaccine
	//Se crea el objectId de mongo
	oid, err := primitive.ObjectIDFromHex(userID)
	//Si ocurrio un error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return resultVaccine, err
	}
	// Se actualiza la fecha de actualización
	resultVaccine.UpdatedAt = time.Now()
	update := bson.M{"$set": updatedVaccine}
	filter := bson.M{"_id": oid}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función
	defer dbConnection.Close()
	// Obtiene la colección de vacunas
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	_, err = collection.UpdateOne(dbConnection.Context, filter, update)
	if err != nil {
		return resultVaccine, err
	}
	log.Println("Vacuna actualizado")
	return resultVaccine, nil
}

// Función para obtener todas las vacunas
func GetAllVaccinesService() ([]models.Vaccine, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	// Variable que contiene a todas las vacunas
	var vaccines []models.Vaccine
	// Trae a todos las vacunas desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return vaccines, errors.New("no fue posible traer todas las vacunas")
	}
	for results.Next(dbConnection.Context) {
		var singleVaccine models.Vaccine
		if err = results.Decode(&singleVaccine); err != nil {
			log.Println("Vacuna no se pudo añadir")
		}

		vaccines = append(vaccines, singleVaccine)
	}
	return vaccines, nil
}