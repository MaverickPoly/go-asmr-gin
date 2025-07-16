package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	ID  primitive.ObjectID `bson:"_id,omitempty"`
	URL string             `bson:"url"`
}

func main() {
	ConnectDB()
	app := gin.Default()

	app.POST("/create", CreateURL)
	app.GET("/:id", GetEndpoint)

	app.Run(":8000")
}
