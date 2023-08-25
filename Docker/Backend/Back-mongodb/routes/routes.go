package routes

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes registra todas las rutas de la aplicación en el enrutador de la aplicación de gin
func InitRoutes(r *gin.Engine) {
	//Registra las rutas del archivo authRouter.go
	InitAuthRoutes(r)
	//Registra las rutas del archivo ownerRouter.go
	InitOwnerRoutes(r)
	//Registra las rutas del archivo dogRouter.go
	InitDogRoutes(r)
	//Registra las rutas del archivo vaccineRouter.go
	InitVaccineRoutes(r)
	//Registra las rutas del archivo userRouter.go
	InitUserRoutes(r)
}
