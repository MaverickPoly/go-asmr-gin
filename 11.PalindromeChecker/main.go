package main

import (
	"github.com/gin-gonic/gin"
)

func isPalindrome(text string) bool {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversedStr := string(runes)
	return reversedStr == text
}

func HandlePalindrome(c *gin.Context) {
	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{"msg": "Invalid string!"})
		return
	}
	text := string(raw)

	res := isPalindrome(text)

	c.JSON(200, gin.H{
		"text":          text,
		"is_palindrome": res,
	})
}

func main() {
	app := gin.Default()
	app.POST("/palindrome", HandlePalindrome)

	app.Run(":8000")
}
