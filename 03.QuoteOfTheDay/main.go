package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

var quotes = []string{
	"Success isn't given. It's taken by those who refuse to be average.",
	"Discipline is the bridge between dreams and becoming a weapon.",
	"A man without control is a slave — to women, food, and weakness.",
	"Wake up with purpose or sleep forever in mediocrity.",
	"The world bows to those who never quit.",
	"Hustle in silence. Let the empire make the noise.",
	"Weak men wait. Strong men build.",
	"If you're broke, stop playing. Start building.",
	"The grind doesn't care about your feelings.",
	"You're not tired. You're undisciplined.",
	"Money is attracted to strength, not excuses.",
	"Masculinity is forged in adversity.",
	"They want you weak. Stay dangerous.",
	"Don't chase women. Chase excellence — and they'll chase you.",
	"Every second you're lazy, someone else is getting rich.",
	"A man without ambition is already dead.",
	"Comfort is the enemy of greatness.",
	"Be so good they call you arrogant.",
	"Losers talk. Winners take the throne.",
	"You don't need motivation. You need to stop being soft.",
}

func HandleRoutes(app *gin.Engine) {
	app.GET("/quotes", func(c *gin.Context) {
		c.JSON(200, quotes)
	})
	app.GET("/quote", func(c *gin.Context) {
		randomIndex := rand.Intn(len(quotes))
		c.String(200, quotes[randomIndex])
	})
}

func main() {
	app := gin.Default()
	HandleRoutes(app)

	app.Run("localhost:3000")
}
