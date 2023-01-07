package main

import (
	"github.com/AdrameDiakhate/golangTodoList.git/configs"
	"github.com/AdrameDiakhate/golangTodoList.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routes.InitRouting(r)
	// configs.DBConnection()
	// r.Run(configs.APP_LISTENING_ADDRESS + configs.APP_PORT)
	r.Run(configs.APP_LISTENING_ADDRESS + ":" + configs.APP_PORT)

}
