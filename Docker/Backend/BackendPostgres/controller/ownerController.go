package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Create de Owner
func CreateOwner(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	var owner models.Owner

	//En caso de algun error
	if err := ctx.ShouldBindJSON(&owner); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Se llama al servicio que crea el Owner
	createdOwner, err := services.CreateOwnerService(owner)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el Owner"})
		return
	}
	// Devuelve el mensaje de confirmación
	log.Println("Dueño creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdOwner)

}

// Get de Owner por id
func GetOwnerByID(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	ownerID := ctx.Param("id")
	// Se busca al Owner
	resultOwner, err := services.GetOwnerByIDService(ownerID)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al Dueño"})
		return
	}
	// Devuelve el Owner
	ctx.JSON(http.StatusOK, resultOwner)
}

// Delete de Owner
func DeleteOwner(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	ownerID := ctx.Param("id")
	// Se elimina al Owner segun el id
	err := services.DeleteOwnerService(ownerID)
	// Si hubo un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusOK, "Dueño eliminado")
}

// Put de Owner
func UpdateOwner(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	var updatedOwner models.Owner
	//Si hubo un error
	if err := ctx.ShouldBindJSON(&updatedOwner); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos del Dueños"})
		return
	}
	ownerID := ctx.Param("id")
	//Se actualiza el Owner
	updatedOwner, err := services.UpdateOwnerService(updatedOwner, ownerID)
	//En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar el Dueños"})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusCreated, "Owner actualizado")
}

// Get all de Owner
func GetAllOwners(ctx *gin.Context) {
	// Se consiguen a los Owners de la base de datos
	resultOwners, err := services.GetAllOwnersService()
	//En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los Dueños"})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusCreated, resultOwners)
}
