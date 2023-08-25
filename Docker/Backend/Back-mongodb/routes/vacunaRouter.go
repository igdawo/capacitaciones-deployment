package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// Se registran las rutas del CRUD y funciones de Vaccine
func InitVaccineRoutes(r *gin.Engine) {
	// Ya que las vistas no tiene login, se modifica la autenticacion de las rutas
	//Grupo de rutas
	vaccineGroup := r.Group("/vacuna")
	{
		vaccineGroup.GET("/:id", controller.GetVaccineByID)
		vaccineGroup.GET("/", controller.GetAllVaccines)
		//Solo usuarios logeados pueden crear o actualizar
		vaccineGroup.POST("/", controller.CreateVaccine)
		vaccineGroup.PUT("/:id", controller.UpdateVaccine)
		vaccineGroup.PUT("/general/", controller.UpdateVaccineGeneral)
		//Solo un Admin puede borrar
		vaccineGroup.DELETE("/:id", controller.DeleteVaccine)
	}
}
