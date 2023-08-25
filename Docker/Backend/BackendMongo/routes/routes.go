package routes

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes registra todas las rutas de la aplicación en el enrutador de la aplicación de gin
func InitRoutes(r *gin.Engine) {
	//Registra las rutas del grupo de gatos del archivo gato.go en package Routes
	//InitCatRoutes(r)
	//Registra las rutas del grupo de gatos del archivo gatoRouter.go
	InitAuthRoutes(r)
	//Registra las rutas del grupo de usuario del archivo usuarioRouter.go
	//InitUserRoutes(r)
	//Registra las rutas del grupo de dueños del archivo ownerRouter.go
	InitOwnerRoutes(r)
	//Registra las rutas del grupo de perros del archivo dogRouter.go
	InitDogRoutes(r)
	//Registra las rutas del grupo de vacunas del archivo vaccineRouter.go
	InitVaccineRoutes(r)
}
