package routes

import (
	"github.com/gin-gonic/gin"
	Controller "github.com/kaar20.taskmanagemnt/Controller"
)

func CategoriesRoute(route *gin.Engine) {
	route.GET("/categories", Controller.ListCategories())
	route.GET("/categories/:id", Controller.GetCategory())
	route.POST("/categories", Controller.CreateCategory())
	route.PUT("/categories/:id", Controller.UpdateCategory())
	route.DELETE("/categories/:id", Controller.DeleteCategory())

}
