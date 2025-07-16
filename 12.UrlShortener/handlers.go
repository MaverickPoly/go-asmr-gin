package main

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateURL(c *gin.Context) {
	var url Url

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(400, gin.H{"error": "Invalid body!"})
		return
	}

	if url.URL == "" || !strings.Contains(url.URL, "http") {
		c.JSON(400, gin.H{"error": "Invalid body!"})
		return
	}

	res, err := URLCollection.InsertOne(context.TODO(), url)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to insert url!"})
		return
	}

	c.JSON(200, gin.H{"msg": "Successfully created url", "id": res.InsertedID})
}

func GetEndpoint(c *gin.Context) {
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format!"})
		return
	}

	var url Url
	if err := URLCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&url); err != nil {
		c.JSON(404, gin.H{"error": "Endpoint not found!"})
		return
	}

	c.Redirect(301, url.URL)
}
