package routes

import (
	"rest-template/controller"

	"github.com/gin-gonic/gin"
)

// InitRoutes registra las rutas junto a las funciones que ejecutan
func InitOwnerRoutes(r *gin.Engine) {
	// Define a group of routes with a shared set of middleware
	// Se define un grupo de rutas
	ownerGroup := r.Group("/owner")
	{
		ownerGroup.POST("/", controller.CreateOwner)
		ownerGroup.GET("/:id", controller.GetOwnerByID)
		ownerGroup.GET("/", controller.GetAllOwners)
		ownerGroup.PUT("/:id", controller.UpdateOwner)
		ownerGroup.DELETE("/:id", controller.DeleteOwner)
	}
}