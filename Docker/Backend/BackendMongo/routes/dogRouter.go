package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitDogRoutes(r *gin.Engine) {
	// Define a group of routes with a shared set of middleware
	// Se define un grupo de rutas
	dogGroup := r.Group("/dog")
	{
		//Solo un usuario logueado sin importar su rol, puede crear un gato
		dogGroup.POST("/", controller.CreateDog)
		dogGroup.GET("/:id", controller.GetDogByID)
		dogGroup.GET("/", controller.GetAllDogs)
		dogGroup.GET("/dog-owner/:id", controller.GetOwnerByDogID)
		dogGroup.GET("/dog-vaccines/:id", controller.GetAllVaccinesByDogID)
		//Solo un usuario o admin logueados pueden actualizar a un gato
		dogGroup.PUT("/:id", controller.UpdateDog)
		//Solo un Admin logueado puede eliminar a un gato
		dogGroup.DELETE("/:id", controller.DeleteDog)
	}
}
