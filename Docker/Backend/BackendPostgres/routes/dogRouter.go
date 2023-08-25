package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitDogRoutes(r *gin.Engine) {
	// Se define un grupo de rutas
	dogGroup := r.Group("/dog")
	{
		dogGroup.POST("/", controller.CreateDog)
		dogGroup.GET("/:id", controller.GetDogByID)
		dogGroup.GET("/", controller.GetAllDogs)
		dogGroup.PUT("/:id", controller.UpdateDog)
		dogGroup.DELETE("/:id", controller.DeleteDog)
		dogGroup.GET("/owner/:id", controller.GetOwnerByDogID)
		dogGroup.GET("/vaccines/:id", controller.GetAllVaccinesByDogID)
	}
}