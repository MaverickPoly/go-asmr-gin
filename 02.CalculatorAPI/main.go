package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	SetupRoutes(app)

	app.Run("localhost:8080")
}
