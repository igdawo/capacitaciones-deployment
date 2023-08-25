package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// Se registran las rutas del CRUD y funciones de Dog
func InitDogRoutes(r *gin.Engine) {

	//Grupo de rutas
	dogGroup := r.Group("/dog")
	{
		dogGroup.GET("/:id", controller.GetDogByID)
		dogGroup.GET("/", controller.GetAllDogs)
		dogGroup.GET("/dog-owner/:id", controller.GetOwnerByDogID)
		dogGroup.GET("/dog-vaccines/:id", controller.GetAllVaccinesByDogID)
		dogGroup.POST("/", controller.CreateDog)
		dogGroup.PUT("/:id", controller.UpdateDog)
		dogGroup.DELETE("/:id", controller.DeleteDog)
	}
}
