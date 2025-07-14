package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDB()
	app := gin.Default()
	SetupRoutes(app)

	app.Run(":8000")
}
