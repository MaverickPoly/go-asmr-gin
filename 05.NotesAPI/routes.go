package main

import "github.com/gin-gonic/gin"

func NoteRoutes(app *gin.Engine) {
	router := app.Group("/api/v1/notes")

	router.GET("/", FetchAllNotes)
	router.POST("/", CreateNote)
	router.GET("/:id", FetchNoteById)
	router.DELETE("/:id", DeleteNote)
	router.PUT("/:id", UpdateNote)
}
