package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AllRecipes(c *gin.Context) {
	recipes := make([]Recipe, 0)

	if err := DB.Find(&recipes).Error; err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Failed to fetch all recipes: %v", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "Fetched all recipes successfully!",
		"data": recipes,
	})
}

func CreateRecipe(c *gin.Context) {
	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid body!",
		})
		return
	}
	if recipe.Name == "" {
		c.JSON(400, gin.H{
			"error": "Recipe name is missing!",
		})
		return
	}

	if err := DB.Create(&recipe).Error; err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Failed to create recipe: %v", err.Error()),
		})
	}

	c.JSON(201, gin.H{
		"msg":  "Created recipe successfully!",
		"data": recipe,
	})
}

func SingleRecipe(c *gin.Context) {
	var recipe Recipe
	status, err := GetRecipeById(c, &recipe)

	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{
		"msg":  "Fetched recipe successfully",
		"data": recipe,
	})
}

func DeleteRecipe(c *gin.Context) {
	var recipe Recipe
	status, err := GetRecipeById(c, &recipe)

	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Delete(&recipe).Error; err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Failed to delete recipe: %v", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "Deleted recipe successfully!",
		"data": recipe,
	})
}

func UpdateRecipe(c *gin.Context) {
	var dbRecipe Recipe
	status, err := GetRecipeById(c, &dbRecipe)

	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	var recipe Recipe
	if err := c.ShouldBindBodyWithJSON(&recipe); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid body!",
		})
		return
	}

	if err := DB.Model(&dbRecipe).Updates(&recipe).Error; err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Failed to update recipe: %v", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":  "Updated recipe successfully!",
		"data": dbRecipe,
	})
}
