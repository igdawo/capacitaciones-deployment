package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Create de Dog
func CreateDog(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	var dog models.Dog

	//En caso de algun error
	if err := ctx.ShouldBindJSON(&dog); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Se llama al servicio que crea el Dog
	createdDog, err := services.CreateDogService(dog)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el Dog"})
		return
	}
	// Devuelve el mensaje de confirmación
	log.Println("Dueño creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdDog)

}

// Get de Dog por id
func GetDogByID(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	dogID := ctx.Param("id")
	// Se busca al Dog
	resultDog, err := services.GetDogByIDService(dogID)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al Dog"})
		return
	}
	// Devuelve el Dog
	ctx.JSON(http.StatusOK, resultDog)
}

// Delete de Dog
func DeleteDog(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	dogID := ctx.Param("id")
	// Se elimina al Dog segun el id
	err := services.DeleteDogService(dogID)
	// Si hubo un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusOK, "Dog eliminado")
}

// Put de Dog
func UpdateDog(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	var updatedDog models.Dog
	//Si hubo un error
	if err := ctx.ShouldBindJSON(&updatedDog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos del Dog"})
		return
	}
	dogID := ctx.Param("id")
	//Se actualiza el Dog
	updatedDog, err := services.UpdateDogService(updatedDog, dogID)
	//En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar al Dog"})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusCreated, "Dog actualizado")
}

// Get all de Dog
func GetAllDogs(ctx *gin.Context) {
	// Se consiguen a los Dogs de la base de datos
	resultDogs, err := services.GetAllDogsService()
	//En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los Dogs"})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusCreated, resultDogs)
}

// Get Owner de un dog
func GetOwnerbyDogID(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	dogID := ctx.Param("id")
	// Se busca al Dog
	resultDog, err := services.GetDogByIDService(dogID)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al Dog"})
		return
	}

	// Se busca al Owner
	resultOwner, err := services.GetOwnerByIDService(resultDog.Owner.Hex())
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al Owner"})
		return
	}
	// Devuelve el Owner
	ctx.JSON(http.StatusOK, resultOwner)
}

// Get vaccines de un dog
func GetVaccinesDogs(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	dogID := ctx.Param("id")
	// Se busca Las vaccines del dog
	resultDogs, err := services.GetVaccineByDogService(dogID)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al Dog"})
		return
	}
	ctx.JSON(http.StatusCreated, resultDogs)
}

// Put General de Dog
func UpdateDogGeneral(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	var updatedDog models.Dog
	//Si hubo un error
	if err := ctx.ShouldBindJSON(&updatedDog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos de los Dogs"})
		return
	}
	//Se actualiza el Dog
	updatedDog, err := services.UpdateDogGeneralService(updatedDog)
	//En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar los Dogs"})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusCreated, "Dogs actualizados")
}
