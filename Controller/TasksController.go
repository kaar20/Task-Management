package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kaar20.taskmanagemnt/database"
	"github.com/kaar20.taskmanagemnt/db"
)

// route.GET("/tasks", Controller.ListTasks())
//
//	route.GET("/tasks/:id", Controller.GetTask())
//	route.POST("/tasks", Controller.CreateTask())
//	route.PUT("/tasks/:id", Controller.UpdateTask())
//	route.DELETE("/tasks/:id", Controller.DeleteTask())
// var validate = validator.New()

func ListTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		query := db.New(database.Client)
		taskList, err := query.TasksList(ctx)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, taskList)
	}
}

func GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Task By ID",
		})
	}

}
func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task db.Task
		query := db.New(database.Client)
		ctx := context.Background()

		// Bind JSON body to task struct
		if err := c.BindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		_, err := query.GetUserByID(ctx, 1)
		if err != nil {
			c.JSON(400, gin.H{"error": "User does not exist"})
			return
		}

		// Create the task
		err = query.CreateTask(ctx, db.CreateTaskParams{
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			Priority:    task.Priority,
			DueDate:     task.DueDate,
			UserID:      task.UserID,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, task)
	}
}

func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Task Updated",
		})
	}
}

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Task Deleted",
		})

	}
}
