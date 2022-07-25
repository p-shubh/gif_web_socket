package main

import (
	"github.com/gin-gonic/gin"
)

type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
	H string `json:"h"`
}

type GIF struct {
	Results []struct {
		Media []struct {
			Gif struct {
				URL string `json:"url"`
			} `json:"gif"`
		} `json:"media"`
	} `json:"results"`
}

func main() {
	//create gin app
	app := gin.Default()
	//set route
	app.GET("/", getHandler)
	app.GET("/ws", RegisterClient)
	//run app
	app.Run(":8080")
}

func getHandler(c *gin.Context) {
	res := gin.H{
		"message": "Hello World",
	}
	c.JSON(200, res)
	c.Header("Content-Type", "application/json")
	return
}
