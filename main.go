package main

import (
	"minity/config"
	"minity/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	router := gin.Default()

	//membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//mulai server dengan port 3000
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
