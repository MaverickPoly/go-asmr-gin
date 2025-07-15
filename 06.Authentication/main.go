package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectDB()
	app := gin.Default()
	RegisterRoutes(app)

	app.Run(":6000")
}
