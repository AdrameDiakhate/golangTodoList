package routes

import "github.com/gin-gonic/gin"

func InitRouting(r *gin.Engine) {
	setUserRoute(r)
	setHelloRoute(r)
	// setUserRoute(r)
	// setTaskRoute(r)
}
