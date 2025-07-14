package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchAllTodos(c *gin.Context) {
	todos := make([]Todo, 0)

	if err := DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "Error fetching all todos!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": "Fetched todos successfully!",
		"data":   todos,
	})
}

func CreateTodo(c *gin.Context) {
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid Body!",
		})
		return
	}

	if err := DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": "Failed to create todo!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": "Created todo successfully!",
		"data":   todo,
	})
}

func FetchTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("todoId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid todo Id!",
		})
		return
	}

	var todo Todo
	if err := DB.Find(&todo, todoId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": fmt.Sprintf("Failed to fetch todo with id %v", todoId),
		})
		return
	}
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": fmt.Sprintf("Todo with id %v not found!", todoId),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": fmt.Sprintf("Fetched todo with id %v successfully!", todoId),
		"data":   todo,
	})
}

func DeleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("todoId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid todo Id!",
		})
		return
	}

	var todo Todo
	if err := DB.Find(&todo, todoId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": fmt.Sprintf("Todo with id %v not found!", todoId),
		})
		return
	}
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": fmt.Sprintf("Todo with id %v not found!", todoId),
		})
		return
	}

	if err := DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": fmt.Sprintf("Failed to delete todo with id %v!", todoId),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": fmt.Sprintf("Deleted todo with id %v successfully!", todoId),
		"data":   todo,
	})
}

func UpdateTodo(c *gin.Context) {
	// Get ID
	todoId, err := strconv.Atoi(c.Param("todoId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid todo Id!",
		})
		return
	}

	// Check if todo exists
	var todo Todo
	if err := DB.Find(&todo, todoId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": fmt.Sprintf("Todo with id %v not found!", todoId),
		})
		return
	}
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": fmt.Sprintf("Todo with id %v not found!", todoId),
		})
		return
	}

	// Parse Body
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Invalid Body!",
		})
		return
	}

	// Update
	if err := DB.Model(todo).Updates(newTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"detail": fmt.Sprintf("Failed to update todo with id %v!", todoId),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"detail": fmt.Sprintf("Updated todo with id %v successfully!", todoId),
		"data":   newTodo,
	})
}
