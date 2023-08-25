package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite crear un dueño
func CreateOwner(ctx *gin.Context) {
	// Obtiene los datos del dueño a partir del cuerpo de la solicitud HTTP.
	var owner models.Owner
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&owner); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea el dueño en la base de datos Mongo
	createdOwner, err := services.CreateOwnerService(owner)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el dueño"})
		return
	}
	log.Println("Dueño creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdOwner)

}

// Servicio que permite retornar un dueño buscandolo por su id
func GetOwnerByID(ctx *gin.Context) {
	// Obtiene el ID del perro a partir del parámetro de la ruta.
	ownerID := ctx.Param("id")
	// Se busca al dueño en la base de datos
	resultOwner, err := services.GetOwnerByIDService(ownerID)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al owner"})
		return
	}
	// Devuelve el dueño encontrado.
	ctx.JSON(http.StatusOK, resultOwner)
}

// Servicio que permite eliminar a un dueño de la base de datos
func DeleteOwner(ctx *gin.Context) {
	ownerID := ctx.Param("id")
	err := services.DeleteOwnerService(ownerID)
	// Si hubo un error
	if err != nil {
		// Ocurrió un error durante la búsqueda.
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusOK, "Dueño eliminado")
}

// Servicio que permite actualizar a un dueño
func UpdateOwner(ctx *gin.Context) {
	//Se crea un modelo de dueño
	var updatedOwner models.Owner
	//Se guardan los datos de la petición http en el modelo
	//Si hubo un error
	if err := ctx.ShouldBindJSON(&updatedOwner); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos del dueño"})
		return
	}
	//Se consigue el id del dueño a actualizar
	ownerID := ctx.Param("id")
	//Se actualiza el dueño en la base de datos
	updatedOwner, err := services.UpdateOwnerService(updatedOwner, ownerID)
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar al dueño"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, "Dueño actualizado")
}

// Servicio para conseguir a todos los dueños de la base de datos
func GetAllOwners(ctx *gin.Context) {
	// Se consiguen a los dueños de la base de datos
	resultOwners, err := services.GetAllOwnersService()
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los dueños"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, resultOwners)
}
