package routes

import (
	"github.com/gin-gonic/gin"
	Controller "github.com/kaar20.taskmanagemnt/Controller"
)

func TasksRoute(route *gin.Engine) {
	route.GET("/tasks", Controller.ListTasks())
	route.GET("/tasks/:id", Controller.GetTask())
	route.POST("/tasks", Controller.CreateTask())
	route.PUT("/tasks/:id", Controller.UpdateTask())
	route.DELETE("/tasks/:id", Controller.DeleteTask())
}
