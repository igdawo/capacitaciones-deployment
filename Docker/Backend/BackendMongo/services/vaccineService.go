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
	CollectionNameVaccine = "Vacuna"
)

// Función para crear una vacuna e insertarla a la base de datos de mongodb
func CreateVaccineService(newVaccine models.Vaccine) (models.Vaccine, error) {
	//Se establece conexión con la base de datos mongo
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de vacunas.
	collection := dbConnection.GetCollection(CollectionNameVaccine)

	// Genera un nuevo ID único para el vacuna.
	newVaccine.ID = primitive.NewObjectID()

	// Inserta el vacuna en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newVaccine)
	//Si hubo un error
	if err != nil {
		return newVaccine, err
	}

	return newVaccine, err
}

// Función para obtener un vacuna por id
func GetVaccineByIDService(vaccineID string) (models.Vaccine, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del vacuna.
	var vaccine models.Vaccine
	oid, err := primitive.ObjectIDFromHex(vaccineID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return vaccine, err
	}
	// Crea un filtro para buscar el vacuna por su ID.
	filter := bson.M{"_id": oid}

	// Obtiene la colección de vacunas.
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&vaccine)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró ningún documento con el ID especificado.
			log.Println("Vacuna no encontrada")
			return vaccine, err
		}
		// Ocurrió un error durante la búsqueda.
		return vaccine, err
	}
	log.Println("Se encontró la Vacuna")
	// Devuelve el vacuna encontrado.
	return vaccine, nil

}

func DeleteVaccineService(vaccineID string) error {
	// Obtiene el ID del vacuna a partir del parámetro de la ruta.
	// Crea un objeto ID de MongoDB a partir del ID del vacuna.
	oid, err := primitive.ObjectIDFromHex(vaccineID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return errors.New("id invalido")
	}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Se elimina el vacuna
	filter := bson.M{"_id": oid}
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	// Elimina el vacuna de la colección.
	result, _ := collection.DeleteOne(dbConnection.Context, filter)
	// Si no hay error
	if result.DeletedCount == 1 {
		// Se pudo eliminar al vacuna
		return nil
	}
	// No se pudo eliminar al vacuna
	return errors.New("la vacuna no pudo ser eliminada")
}

func UpdateVaccineService(updatedVaccine models.Vaccine, userID string) (models.Vaccine, error) {
	// Se crea el modelo de la vacuna que se actualizara
	var resultVaccine models.Vaccine
	//Se crea el objectId de mongo
	oid, err := primitive.ObjectIDFromHex(userID)
	//Si ocurrio un error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return resultVaccine, err
	}
	update := bson.M{"$set": updatedVaccine}
	filter := bson.M{"_id": oid}
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de vacunas.
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	_, err = collection.UpdateOne(dbConnection.Context, filter, update)
	if err != nil {
		return resultVaccine, err
	}
	log.Println("Vacuna actualizada")
	return resultVaccine, nil
}

func GetAllVaccinesService() ([]models.Vaccine, error) {
	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	// Variable que contiene a todos los vacunas
	var vaccines []models.Vaccine
	// Trae a todos los vacunas desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return vaccines, errors.New("No fue posible traer a todas las Vacunas")
	}
	for results.Next(dbConnection.Context) {
		var singleVaccine models.Vaccine
		if err = results.Decode(&singleVaccine); err != nil {
			log.Println("La vacuna no se pudo añadir")
		}

		vaccines = append(vaccines, singleVaccine)
	}
	return vaccines, nil
}

// Función para retornar todas las vacunas de un perro a partir de su id
func GetAllVaccinesByDogIDService(dogID string) ([]models.Vaccine, error) {

	// Crea una nueva instancia a la conexión de base de datos
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()

	var vaccinesAux []models.Vaccine
	// Crea un objeto ID de MongoDB a partir del ID del perro.
	oid, err := primitive.ObjectIDFromHex(dogID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return vaccinesAux, err
	}
	// Crea un filtro para buscar las vacunas que contengan el ID del perro
	filter := bson.M{"id_Perro": oid}
	collection := dbConnection.GetCollection(CollectionNameVaccine)

	// Variable que contiene a todos los vacunas
	var vaccines []models.Vaccine
	// Trae todas los vacunas desde la base de datos que cumplan con el filtro
	results, err := collection.Find(dbConnection.Context, filter)
	if err != nil {
		return vaccines, errors.New("No fue posible traer todas las vacunas")
	}
	for results.Next(dbConnection.Context) {
		var singleVaccine models.Vaccine
		if err = results.Decode(&singleVaccine); err != nil {
			log.Println("La vacuna no se pudo añadir")
		}

		vaccines = append(vaccines, singleVaccine)
	}
	return vaccines, nil
}
