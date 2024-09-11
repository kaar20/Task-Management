package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/kaar20.taskmanagemnt/database"
	"github.com/kaar20.taskmanagemnt/db"
)

var validate = validator.New()

func ListCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{"message": "List of categories"})

		ctx := context.Background()
		query := db.New(database.Client)

		categoriesList, err := query.ListCategories(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categoriesList)

	}
}

func GetCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		catId := c.Param("id")
		catIdConv, err := strconv.Atoi(catId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
			return
		}
		categoryId := int32(catIdConv)
		ctx := context.Background()
		query := db.New(database.Client)

		getCategory, err := query.GetCategoryByID(ctx, categoryId)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, gin.H{
			"Data": getCategory,
		})

	}

}

func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {

		var category db.Category
		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validation = validate.Struct(category)
		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.Error()})
			return
		}
		var query = db.New(database.Client)
		var ctx = context.Background()
		newCat := query.CreateCategory(ctx, category.Name)
		if newCat != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": newCat.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"success": "New Category Created successfully",
		})

	}
}

func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		catId := c.Param("id")
		catIdConv, er := strconv.Atoi(catId)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
			return
		}
		categoryId := int32(catIdConv)
		var category db.Category
		if err := c.BindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var validation = validate.Struct(category)
		if validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.Error()})
			return
		}
		var query = db.New(database.Client)
		var ctx = context.Background()
		Error := query.UpdateCategory(ctx, db.UpdateCategoryParams{
			ID:   categoryId,
			Name: category.Name,
		})
		if Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "Category Updated Successfully",
		})
	}
}

func DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"Message": "Deleted Successfully",
		// })
		catId := c.Param("id")
		catIdConv, err := strconv.Atoi(catId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
			return
		}
		categoryId := int32(catIdConv)
		var query = db.New(database.Client)
		var ctx = context.Background()
		Error := query.DeleteCategory(ctx, categoryId)
		if Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "Category Deleted Successfully",
		})
	}
}
