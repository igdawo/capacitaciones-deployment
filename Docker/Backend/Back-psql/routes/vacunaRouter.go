package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// Se registran las rutas del CRUD y funciones de Vaccine
func InitVaccineRoutes(r *gin.Engine) {

	//Grupo de rutas
	vaccineGroup := r.Group("/vacuna")
	{
		vaccineGroup.GET("/:id", controller.GetVaccineByID)
		vaccineGroup.GET("/", controller.GetAllVaccines)
		vaccineGroup.POST("/", controller.CreateVaccine)
		vaccineGroup.PUT("/:id", controller.UpdateVaccine)
		vaccineGroup.PUT("/general/", controller.UpdateVaccineGeneral)
		vaccineGroup.DELETE("/:id", controller.DeleteVaccine)
	}
}
