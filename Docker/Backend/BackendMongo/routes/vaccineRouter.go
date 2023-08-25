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
		//Solo un usuario logueado sin importar su rol, puede crear un gato
		vaccineGroup.POST("/", controller.CreateVaccine)
		vaccineGroup.GET("/:id", controller.GetVaccineByID)
		vaccineGroup.GET("/", controller.GetAllVaccines)
		//Solo un usuario o admin logueados pueden actualizar a un gato
		vaccineGroup.PUT("/:id", controller.UpdateVaccine)
		//Solo un Admin logueado puede eliminar a un gato
		vaccineGroup.DELETE("/:id", controller.DeleteVaccine)
	}
}
