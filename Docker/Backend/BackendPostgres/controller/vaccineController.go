package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Create de Vaccine
func CreateVaccine(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	var vaccine models.Vaccine

	//En caso de algun error
	if err := ctx.ShouldBindJSON(&vaccine); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Se llama al servicio que crea el Vaccine
	createdVaccine, err := services.CreateVaccineService(vaccine)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear el Vaccine"})
		return
	}
	// Devuelve el mensaje de confirmación
	log.Println("Dueño creado en la base de datos")
	ctx.JSON(http.StatusCreated, createdVaccine)

}

// Get de Vaccine por id
func GetVaccineByID(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	vaccineID := ctx.Param("id")
	// Se busca al Vaccine
	resultVaccine, err := services.GetVaccineByIDService(vaccineID)
	// En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener al Vaccine"})
		return
	}
	// Devuelve el Vaccine
	ctx.JSON(http.StatusOK, resultVaccine)
}

// Delete de Vaccine
func DeleteVaccine(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	vaccineID := ctx.Param("id")
	// Se elimina al Vaccine segun el id
	err := services.DeleteVaccineService(vaccineID)
	// Si hubo un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusOK, "Vaccine eliminado")
}

// Put de Vaccine
func UpdateVaccine(ctx *gin.Context) {
	// Se obtiene los datos del cuerpo de la peticion
	var updatedVaccine models.Vaccine
	//Si hubo un error
	if err := ctx.ShouldBindJSON(&updatedVaccine); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos del Vaccine"})
		return
	}
	vaccineID := ctx.Param("id")
	//Se actualiza el Vaccine
	updatedVaccine, err := services.UpdateVaccineService(updatedVaccine, vaccineID)
	//En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar la Vaccine"})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusCreated, "Vaccine actualizado")
}

// Get all de Vaccine
func GetAllVaccines(ctx *gin.Context) {
	// Se consiguen a los Vaccines de la base de datos
	resultVaccines, err := services.GetAllVaccinesService()
	//En caso de algun error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los Vaccines"})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusCreated, resultVaccines)
}
