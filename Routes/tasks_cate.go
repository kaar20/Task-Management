package routes

import (
	"github.com/gin-gonic/gin"
	Controller "github.com/kaar20.taskmanagemnt/Controller"
)

func TaskCategories(route *gin.Engine) {
	route.GET("/task-cat", Controller.ListTaskCat())
	route.GET("/task-cat/:id", Controller.GetTaskCat())
	route.POST("/task-cat", Controller.CreateTaskCat())
	route.PUT("/task-cat/:id", Controller.UpdateTaskCat())
	route.DELETE("/task-cat/:id", Controller.DeleteTaskCat())

}
