package routes

import (
	"08.BookManagement/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(app *gin.Engine) {
	router := app.Group("/api")

	router.GET("/books", controllers.FetchAllBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FetchBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	router.PUT("/books/:id", controllers.UpdateBook)
}
