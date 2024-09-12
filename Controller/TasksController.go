package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

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
		taskId := c.Param("id")
		taskIdConv, err := strconv.Atoi(taskId)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Task ID"})
			return
		}
		taskIdV := int32(taskIdConv)
		ctx := context.Background()
		query := db.New(database.Client)
		task, err := query.GetTaskByID(ctx, taskIdV)
		if err != nil {
			c.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(200, task)
		// c.JSON(200, gin.H{
		// 	"message": "Task By ID",
		// })
	}
}

// LocalTask mirrors the structure of db.Task but uses pointers for nullable fields
type LocalTask struct {
	ID          int32   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	Priority    *string `json:"priority"`
	DueDate     *string `json:"due_date"`
	UserID      *int32  `json:"user_id"`
}

// Custom UnmarshalJSON for LocalTask
func (t *LocalTask) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID          int32   `json:"id"`
		Title       string  `json:"title"`
		Description *string `json:"description"`
		Status      *string `json:"status"`
		Priority    *string `json:"priority"`
		DueDate     *string `json:"due_date"`
		UserID      *int32  `json:"user_id"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.ID = aux.ID
	t.Title = aux.Title

	// Handle nullable Description field
	t.Description = aux.Description

	// Handle nullable Status field
	t.Status = aux.Status

	// Handle nullable Priority field
	t.Priority = aux.Priority

	// Handle nullable DueDate field
	t.DueDate = aux.DueDate

	// Handle nullable UserID field
	t.UserID = aux.UserID

	return nil
}

// Convert string to time.Time, returning an empty time if conversion fails
func parseDate(s *string) (time.Time, bool) {
	if s == nil {
		return time.Time{}, false
	}
	date, err := time.Parse("2006-01-02", *s) // Adjust format as needed
	if err != nil {
		return time.Time{}, false
	}
	return date, true
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var localTask LocalTask
		query := db.New(database.Client)
		ctx := context.Background()

		// Bind JSON body to localTask struct
		if err := c.BindJSON(&localTask); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Convert LocalTask to Task
		dueDate, validDueDate := parseDate(localTask.DueDate)
		task := db.Task{
			ID:          localTask.ID,
			Title:       localTask.Title,
			Description: sql.NullString{String: coalesceString(localTask.Description), Valid: localTask.Description != nil},
			Status:      sql.NullString{String: coalesceString(localTask.Status), Valid: localTask.Status != nil},
			Priority:    sql.NullString{String: coalesceString(localTask.Priority), Valid: localTask.Priority != nil},
			DueDate:     sql.NullTime{Time: dueDate, Valid: validDueDate},
			UserID:      sql.NullInt32{Int32: coalesceInt32(localTask.UserID), Valid: localTask.UserID != nil},
		}

		// Check if user exists
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

// Helper functions to coalesce values
func coalesceString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func coalesceInt32(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId := c.Param("id")
		taskIdConv, err := strconv.Atoi(taskId)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Task ID"})
			return
		}
		taskID := int32(taskIdConv)
		var task db.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		ctx := context.Background()
		query := db.New(database.Client)
		err = query.UpdateTask(ctx, db.UpdateTaskParams{
			ID:          taskID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			Priority:    task.Priority,
			DueDate:     task.DueDate,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, task)

	}
}

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId := c.Param("id")
		taskIdConv, err := strconv.Atoi(taskId)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Task ID"})
			return
		}
		taskID := int32(taskIdConv)
		ctx := context.Background()
		query := db.New(database.Client)
		err = query.DeleteTask(ctx, taskID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, gin.H{"Message": "Deleted task Sucessfully"})
		// c.JSON(200, gin.H{
		// 	"message": "Task Deleted",
		// })

	}
}
