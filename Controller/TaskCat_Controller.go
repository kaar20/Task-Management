package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaar20.taskmanagemnt/database"
	"github.com/kaar20.taskmanagemnt/db"
)

// route.GET("/task-cat", Controller.ListTaskCat())
// 	route.GET("/task-cat/:id", Controller.GetTaskCat())
// 	route.POST("/task-cat", Controller.CreateTaskCat())
// 	route.PUT("/task-cat/:id", Controller.UpdateTaskCat())
// 	route.DELETE("/task-cat/:id", Controller.DeleteTaskCat())

func ListTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "List of Task Categories",
		// })
		taskCatId := c.Param("id")
		taskIdConv, err := strconv.Atoi(taskCatId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
			return
		}
		taskId := int32(taskIdConv)
		query := db.New(database.Client)
		ctx := context.Background()
		taskCatList, err := query.GetCategoriesForTask(ctx, taskId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"Data": taskCatList,
		})
	}
}

func GetTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		catTaskID := c.Param("id")
		catTaskIdConv, err := strconv.Atoi(catTaskID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task Category ID"})
			return
		}
		catTaskId := int32(catTaskIdConv)
		query := db.New(database.Client)
		ctx := context.Background()
		taskCat, err := query.GetTasksForCategory(ctx, catTaskId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Data": taskCat,
		})
		// c.JSON(200, gin.H{
		// 	"message": "Task Category By ID",
		// })

	}
}

func CreateTaskCat() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "Task Category Created",
		// })

		var taskCat db.TaskCategory
		if err := c.BindJSON(&taskCat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		query := db.New(database.Client)
		ctx := context.Background()
		err := query.AddTaskCategory(ctx, db.AddTaskCategoryParams{
			// Name: taskCat.Name,
			TaskID:     taskCat.TaskID,
			CategoryID: taskCat.CategoryID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, taskCat)

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
