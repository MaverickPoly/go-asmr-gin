package main

import (
	"fmt"
	"os"

	"14.ContactListAPI/config"
	"14.ContactListAPI/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	PORT := os.Getenv("PORT")

	config.ConnectDB()
	app := gin.Default()
	routes.RegisterRoutes(app)

	app.Run(fmt.Sprintf(":%v", PORT))
}
