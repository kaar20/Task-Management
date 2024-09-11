package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/kaar20.taskmanagemnt/Controller"
)

func UsersRoute(route *gin.Engine) {
	route.GET("/users", controller.ListUsers())
	route.GET("/users/:id", controller.GetUser())
	route.POST("/users", controller.CreateUser())
	route.PUT("/users/:id", controller.UpdateUser())
	route.DELETE("/users/:id", controller.DeleteUser())
}
