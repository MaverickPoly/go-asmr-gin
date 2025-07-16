package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRecipeById(c *gin.Context, recipe *Recipe) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return 400, errors.New("invalid recipe id")
	}

	if err := DB.Find(&recipe, id).Error; err != nil {
		return 500, fmt.Errorf("failed to get recipe: %v", err.Error())
	}

	if recipe.ID == 0 {
		return 404, fmt.Errorf("recipe with id %v not found", id)
	}

	return 200, nil
}
