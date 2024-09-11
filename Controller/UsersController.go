package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// route.GET("/users", controller.ListUsers())
// 	route.GET("/users/:id", controller.GetUser())
// 	route.POST("/users", controller.CreateUser())
// 	route.PUT("/users/:id", controller.UpdateUser())
// 	route.DELETE("/users/:id", controller.DeleteUser())

func ListUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "List of Users"})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User By ID"})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User Created"})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User Updated"})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "User Deleted"})
    }
}