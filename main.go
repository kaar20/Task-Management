package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/kaar20.taskmanagemnt/Routes"
)

func main() {
	fmt.Println("Welcome To The Task Management System..............")
	// r := gin.New()
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.UsersRoute(router)
	routes.CategoriesRoute(router)
	routes.TasksRoute(router)
	routes.TaskCategories(router)
	router.Run(":" + port)

}
