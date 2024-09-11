package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// route.GET("/task-cat", Controller.ListTaskCat())
// 	route.GET("/task-cat/:id", Controller.GetTaskCat())
// 	route.POST("/task-cat", Controller.CreateTaskCat())
// 	route.PUT("/task-cat/:id", Controller.UpdateTaskCat())
// 	route.DELETE("/task-cat/:id", Controller.DeleteTaskCat())

func ListTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List of Task Categories",
		})

	}
}

func GetTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Task Category By ID",
		})

	}
}

func CreateTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Task Category Created",
		})

	}

}

func UpdateTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Task Category Updated",
		})

	}
}

func DeleteTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{

			"message": "Task Category Deleted",
		})
	}
}
