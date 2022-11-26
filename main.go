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

	r.Run("localhost:" + os.Getenv("APP_PORT"))
}
