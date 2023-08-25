package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// Se registran las rutas del CRUD y funciones de Dog
func InitDogRoutes(r *gin.Engine) {

	//Grupo de rutas
	dogGroup := r.Group("/perro")
	{
		dogGroup.GET("/:id", controller.GetDogByID)
		dogGroup.GET("/", controller.GetAllDogs)
		dogGroup.GET("/dueno/:id", controller.GetOwnerbyDogID)
		dogGroup.GET("/vacunas/:id", controller.GetVaccinesDogs)
		dogGroup.POST("/", controller.CreateDog)
		dogGroup.PUT("/:id", controller.UpdateDog)
		dogGroup.PUT("/general/", controller.UpdateDogGeneral)
		dogGroup.DELETE("/:id", controller.DeleteDog)
	}
}
