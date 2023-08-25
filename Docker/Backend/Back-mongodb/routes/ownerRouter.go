package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// Se registran las rutas del CRUD y funciones de Owner
func InitOwnerRoutes(r *gin.Engine) {
	// Ya que las vistas no tiene login, se modifica la autenticacion de las rutas
	//Grupo de rutas
	ownerGroup := r.Group("/dueno")
	{
		ownerGroup.GET("/:id", controller.GetOwnerByID)
		ownerGroup.GET("/", controller.GetAllOwners)
		//Solo usuarios logeados pueden crear o actualizar
		ownerGroup.POST("/", controller.CreateOwner)
		ownerGroup.PUT("/:id", controller.UpdateOwner)
		ownerGroup.PUT("/general/", controller.UpdateOwnerGeneral)
		//Solo un Admin puede borrar
		ownerGroup.DELETE("/:id", controller.DeleteOwner)
	}
}
