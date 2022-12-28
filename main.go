package main

import (
	"os"

	"github.com/AdrameDiakhate/golangTodoList.git/configs"
	"github.com/AdrameDiakhate/golangTodoList.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routes.InitRouting(r)
	configs.DBConnection()
	APP_PORT := os.Getenv("APP_PORT")
	r.Run("0.0.0.0:" + APP_PORT)
}
