package main

import "github.com/gin-gonic/gin"

func StartApi() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8179") // listen and serve on 0.0.0.0:8080
}
