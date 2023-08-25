package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// Se registran las rutas del CRUD y funciones de Owner
func InitOwnerRoutes(r *gin.Engine) {

	//Grupo de rutas
	ownerGroup := r.Group("/owner")
	{
		ownerGroup.GET("/:id", controller.GetOwnerByID)
		ownerGroup.GET("/", controller.GetAllOwners)
		ownerGroup.POST("/", controller.CreateOwner)
		ownerGroup.PUT("/:id", controller.UpdateOwner)
		ownerGroup.DELETE("/:id", controller.DeleteOwner)
	}
}
