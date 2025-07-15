package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Visibility int    `json:"visibility"`
	Name       string `json:"name"`
	Timezone   int    `json:"timezone"`
	ID         int    `json:"id"`
}

func SetupRoutes(app *gin.Engine) {
	client := resty.New()
	app.GET("/weather", func(c *gin.Context) {
		API_KEY := os.Getenv("API_KEY")
		if API_KEY == "" {
			panic("API_KEY environment variable not set")
		}

		cityName := c.DefaultQuery("city", "")
		if cityName == "" {
			c.JSON(400, gin.H{"msg": "City name is missing!"})
			return
		}

		res, err := client.R().
			Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&appid=%v", cityName, API_KEY))
		if err != nil {
			c.JSON(500, gin.H{
				"error": fmt.Sprintf("Error fetching weather: %v", err.Error()),
			})
			return
		}

		var weatherData WeatherData
		if err := json.Unmarshal(res.Body(), &weatherData); err != nil {
			fmt.Println("Error parsing weather response:", err.Error())
			c.JSON(500, gin.H{
				"error": "Error parsing weather response.",
			})
			return
		}

		c.JSON(200, gin.H{
			"temp":     weatherData.Main.Temp,
			"humidity": weatherData.Main.Humidity,
			"pressure": weatherData.Main.Pressure,

			"description": weatherData.Weather[0].Description,
			"icon":        weatherData.Weather[0].Icon,

			"wind_speed": weatherData.Wind.Speed,
			"wind_deg":   weatherData.Wind.Deg,

			"visibility": weatherData.Visibility,
			"name":       weatherData.Name,
			"timezone":   weatherData.Timezone,
			"id":         weatherData.ID,
		})
	})
}
