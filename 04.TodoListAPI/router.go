package main

import "github.com/gin-gonic/gin"

func SetupRoutes(app *gin.Engine) {
	router := app.Group("/api/todos")

	router.GET("/", FetchAllTodos)
	router.POST("/", CreateTodo)
	router.GET("/:todoId", FetchTodo)
	router.DELETE("/:todoId", DeleteTodo)
	router.PUT("/:todoId", UpdateTodo)
}
