package routes

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes registra todas las rutas de la aplicación en el enrutador de la aplicación de gin
func InitRoutes(r *gin.Engine) {

	//Registra las rutas del grupo de Perros del archivo usuarioRouter.go
	InitDogRoutes(r)
	//Registra las rutas del grupo de Dueños del archivo usuarioRouter.go
	InitOwnerRoutes(r)
	//Registra las rutas del grupo de Vacunas del archivo usuarioRouter.go
	InitVaccineRoutes(r)
}
