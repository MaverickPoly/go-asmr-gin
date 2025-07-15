package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDB()
	app := gin.Default()

	NoteRoutes(app)

	app.Run("localhost:8080")
}
