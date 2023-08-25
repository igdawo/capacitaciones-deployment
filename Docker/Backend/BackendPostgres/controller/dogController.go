package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite crear un perro
func CreateDog(ctx *gin.Context) {
	// Obtiene los datos del perro a partir del cuerpo de la solicitud HTTP.
	var dog models.Dog
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&dog); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el perro en la base de datos Mongo
	createdDog, err := services.CreateDogService(dog)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el perro"})
		return
	}
	log.Println("Perro creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdDog)

}

// Servicio que permite retornar un perro buscandolo por su id
func GetDogByID(ctx *gin.Context) {
	// Obtiene el ID del perro a partir del parámetro de la ruta
	dogID := ctx.Param("id")
	// Se busca al perro en la base de datos
	resultDog, err := services.GetDogByIDService(dogID)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al perro"})
		return
	}
	// Devuelve el perro encontrado.
	ctx.JSON(http.StatusOK, resultDog)
}

// Servicio que permite eliminar a un perro de la base de datos
func DeleteDog(ctx *gin.Context) {
	dogID := ctx.Param("id")
	err := services.DeleteDogService(dogID)
	// Si hubo un error
	if err != nil {
		// Ocurrió un error durante la búsqueda.
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusOK, "Perro eliminado")
}

// Servicio que permite actualizar a un perro
func UpdateDog(ctx *gin.Context) {
	//Se crea un modelo de perro
	var updatedDog models.Dog
	//Se guardan los datos de la petición http en el modelo
	//Si hubo un error
	if err := ctx.ShouldBindJSON(&updatedDog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos del perro"})
		return
	}
	//Se consigue el id del perro a actualizar
	dogID := ctx.Param("id")
	//Se actualiza el perro en la base de datos
	updatedDog, err := services.UpdateDogService(updatedDog, dogID)
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar al perro"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, "Perro actualizado")
}

// Servicio para conseguir a todos los perros de la base de datos
func GetAllDogs(ctx *gin.Context) {
	// Se consiguen a los perros de la base de datos
	resultDogs, err := services.GetAllDogsService()
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los perros"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, resultDogs)
}

//Servicio para conseguir el dueño de un perro de la base de datos
func GetOwnerByDogID(ctx *gin.Context){
	// Obtiene el ID del perro a partir del parámetro de la ruta
	dogID := ctx.Param("id")
	//Se consigue el dueño correspondiente a ese perro
	owner,err := services.GetOwnerByDogIDService(dogID)
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al dueño de este perro"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, owner)
}

//Servicio para conseguir las vacunas de un perro
func GetAllVaccinesByDogID(ctx *gin.Context){
	// Obtiene el ID del perro a partir del parámetro de la ruta
	dogID := ctx.Param("id")
	//Se consiguen las vacunas correspondientes a ese perro
	vaccines,err := services.GetAllVaccinesByDogIDService(dogID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener las vacunas de este perro"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, vaccines)
}