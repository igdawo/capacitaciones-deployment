package controller

import (
	"log"
	"net/http"
	"rest-template/models"
	"rest-template/services"

	"github.com/gin-gonic/gin"
)

// Servicio que permite Crear una vacuna
func CreateVaccine(ctx *gin.Context) {
	// Obtiene los datos de la vacuna a partir del cuerpo de la solicitud HTTP.
	var vaccine models.Vaccine
	//Si ocurrio un error
	if err := ctx.ShouldBindJSON(&vaccine); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Se llama al servicio que crea una vacuna en la base de datos Mongo
	createdVaccine, err := services.CreateVaccineService(vaccine)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear la Vacuna"})
		return
	}
	log.Println("Vacuna creada en la base de datos")
	ctx.JSON(http.StatusCreated, createdVaccine)

}

// Servicio que permite retornar una vacuna buscandola por su id
func GetVaccineByID(ctx *gin.Context) {
	// Obtiene el ID de la vacuna a partir del parámetro de la ruta.
	vaccineID := ctx.Param("id")
	// Se busca la vacuna en la base de datos
	resultVaccine, err := services.GetVaccineByIDService(vaccineID)
	// Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener la Vacuna"})
		return
	}
	// Devuelve la vacuna encontrada.
	ctx.JSON(http.StatusOK, resultVaccine)
}

// Servicio que permite eliminar una vacuna de la base de datos
func DeleteVaccine(ctx *gin.Context) {
	vaccineID := ctx.Param("id")
	err := services.DeleteVaccineService(vaccineID)
	// Si hubo un error
	if err != nil {
		// Ocurrió un error durante la búsqueda.
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Devuelve el mensaje de confirmación
	ctx.JSON(http.StatusOK, "Vacuna eliminada")
}

// Servicio que permite actualizar una vacuna
func UpdateVaccine(ctx *gin.Context) {
	//Se crea un modelo de la vacuna
	var updateVaccine models.Vaccine
	//Se guardan los datos de la petición http en el modelo
	//Si hubo un error
	if err := ctx.ShouldBindJSON(&updateVaccine); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar los datos de la Vacuna"})
		return
	}
	//Se consigue el id de la vacuna a actualizar
	vaccineID := ctx.Param("id")
	//Se actualiza la vacuna en la base de datos
	updateVaccine, err := services.UpdateVaccineService(updateVaccine, vaccineID)
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al actualizar la Vacuna"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, "Vacuna actualizada")
}

// Servicio para conseguir todas las vacunas de la base de datos
func GetAllVaccines(ctx *gin.Context) {
	// Se consiguen las vacunas de la base de datos
	resultVaccines, err := services.GetAllVaccinesService()
	//Si ocurrio un error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener las Vacunas"})
		return
	}
	//Se envia la respuesta http
	ctx.JSON(http.StatusCreated, resultVaccines)
}
