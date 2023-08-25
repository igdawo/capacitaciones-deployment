package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitVaccineRoutes(r *gin.Engine) {
	// Define a group of routes with a shared set of middleware
	// Se define un grupo de rutas
	vaccineGroup := r.Group("/vaccine")
	{
		vaccineGroup.POST("/", controller.CreateVaccine)
		vaccineGroup.GET("/:id", controller.GetVaccineByID)
		vaccineGroup.GET("/", controller.GetAllVaccines)
		vaccineGroup.PUT("/:id", controller.UpdateVaccine)
		vaccineGroup.DELETE("/:id", controller.DeleteVaccine)
	}
}