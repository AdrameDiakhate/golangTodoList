package main

import (
	"fmt"
	"os"

	"github.com/AdrameDiakhate/golangTodoList.git/configs"
	"github.com/AdrameDiakhate/golangTodoList.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routes.InitRouting(r)
	fmt.Println("Hello world!")
	configs.DBConnection()

	r.Run("localhost:" + os.Getenv("APP_PORT"))
}
