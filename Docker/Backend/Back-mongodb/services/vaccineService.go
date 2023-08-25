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

// Función para crear un Vaccine en la base de datos
func CreateVaccineService(newVaccine models.Vaccine) (models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Vaccines.
	collection := dbConnection.GetCollection(CollectionNameVaccine)

	// Genera un nuevo ID
	newVaccine.ID = primitive.NewObjectID()

	// Inserta el Vaccine en la colección.
	_, err := collection.InsertOne(dbConnection.Context, newVaccine)
	//En caso de algun error
	if err != nil {
		return newVaccine, err
	}

	return newVaccine, err
}

// Función para obtener un Vaccine por id
func GetVaccineByIDService(vaccineID string) (models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Crea un objeto ID de MongoDB a partir del ID del Vaccine.
	var Vaccine models.Vaccine
	oid, err := primitive.ObjectIDFromHex(vaccineID)
	//En caso de algun error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return Vaccine, err
	}
	// Filtro para buscar el Vaccine por su ID.
	filter := bson.M{"_id": oid}

	// Obtiene el Vaccine.
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	err = collection.FindOne(dbConnection.Context, filter).Decode(&Vaccine)
	// En caso de algun error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Vaccine no encontrado")
			return Vaccine, err
		}
		return Vaccine, err
	}
	log.Println("Se encontró el Vaccine")
	// Devuelve el Vaccine encontrado
	return Vaccine, nil

}

func DeleteVaccineService(vaccineID string) error {
	// Crea un objeto ID de MongoDB a partir del ID del Vaccine.
	oid, err := primitive.ObjectIDFromHex(vaccineID)
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return errors.New("id invalido")
	}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()

	// Se elimina el Vaccine
	filter := bson.M{"_id": oid}
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	result, _ := collection.DeleteOne(dbConnection.Context, filter)
	// En caso de algun error
	if result.DeletedCount == 1 {
		return nil
	}
	return errors.New("el Vaccine pudo ser eliminado")
}

func UpdateVaccineService(updatedVaccine models.Vaccine, ID string) (models.Vaccine, error) {
	// Se crea el modelo del Vaccine a actualizar
	var resultVaccine models.Vaccine
	//Se crea el objectId de mongo
	oid, err := primitive.ObjectIDFromHex(ID)
	// En caso de algun error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return resultVaccine, err
	}
	// Se actualiza
	update := bson.M{"$set": updatedVaccine}
	filter := bson.M{"_id": oid}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Vaccines.
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	_, err = collection.UpdateOne(dbConnection.Context, filter, update)
	if err != nil {
		return resultVaccine, err
	}
	log.Println("Vaccine actualizado")
	return resultVaccine, nil
}

func GetAllVaccinesService() ([]models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	// Variable que contiene a todos los Vaccines
	var vaccines []models.Vaccine
	// Trae a todos los Vaccines desde la base de datos
	results, err := collection.Find(dbConnection.Context, bson.M{})
	if err != nil {
		return vaccines, errors.New("no fue posible traer a todos los Vaccines")
	}
	for results.Next(dbConnection.Context) {
		var singleVaccine models.Vaccine
		if err = results.Decode(&singleVaccine); err != nil {
			log.Println("Vaccine no se pudo añadir")
		}

		vaccines = append(vaccines, singleVaccine)
	}
	return vaccines, nil
}

// Función para obtener un Vaccine por id de un perro
func GetVaccineByDogService(perroID string) ([]models.Vaccine, error) {
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	// Variable que contiene a todos los Vaccines
	var vaccines []models.Vaccine
	// Se transforma el perroid a str
	oid, err := primitive.ObjectIDFromHex(perroID)
	//En caso de algun error
	if err != nil {
		log.Println("No fue posible convertir el ID")
		return vaccines, err
	}
	// Filtro para buscar el Vaccine por Dog_ID.
	filter := bson.M{"Perro_id": oid}
	// Trae a todos los Vaccines desde la base de datos
	results, err := collection.Find(dbConnection.Context, filter)
	if err != nil {
		return vaccines, errors.New("no fue posible traer a todos los Vaccines")
	}
	for results.Next(dbConnection.Context) {
		var singleVaccine models.Vaccine
		if err = results.Decode(&singleVaccine); err != nil {
			log.Println("Vaccine no se pudo añadir")
		}

		vaccines = append(vaccines, singleVaccine)
	}
	return vaccines, nil
}

// Funcion para actualizar todos Vaccines con un o unos atributos
func UpdateVaccineGeneralService(updatedVaccine models.Vaccine) (models.Vaccine, error) {
	// Se crea el modelo del Vaccine a actualizar
	var resultVaccine models.Vaccine
	// Se actualiza
	update := bson.M{"$set": updatedVaccine}
	// Se establece conexión
	dbConnection := config.NewDbConnection()
	// Define un defer para cerrar la conexión a la base de datos al finalizar la función.
	defer dbConnection.Close()
	// Obtiene la colección de Vaccines.
	collection := dbConnection.GetCollection(CollectionNameVaccine)
	_, err := collection.UpdateMany(dbConnection.Context, bson.M{}, update)
	if err != nil {
		return resultVaccine, err
	}
	log.Println("Vaccines actualizadas")
	return resultVaccine, nil
}
