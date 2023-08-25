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
		//Solo un usuario logueado sin importar su rol, puede crear un gato
		ownerGroup.POST("/", controller.CreateOwner)
		ownerGroup.GET("/:id", controller.GetOwnerByID)
		ownerGroup.GET("/", controller.GetAllOwners)
		//Solo un usuario o admin logueados pueden actualizar a un gato
		ownerGroup.PUT("/:id", controller.UpdateOwner)
		//Solo un Admin logueado puede eliminar a un gato
		ownerGroup.DELETE("/:id", controller.DeleteOwner)
	}
}
