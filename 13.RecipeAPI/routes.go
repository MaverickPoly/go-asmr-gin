package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	api := app.Group("/api/recipes")

	api.GET("/", AllRecipes)
	api.POST("/", CreateRecipe)
	api.GET("/:id", SingleRecipe)
	api.DELETE("/:id", DeleteRecipe)
	api.PUT("/:id", UpdateRecipe)
}
