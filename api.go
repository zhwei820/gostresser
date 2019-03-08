package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/zhwei820/gostresser/handlers"

	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func StartApi() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.GET("/conf", handlers.ListConf)
	r.POST("/conf/", handlers.CreateConf)
	r.PUT("/conf/:id/", handlers.UpdateConf)
	r.GET("/conf/:id/", handlers.DetailConf)
	r.DELETE("/conf/:id/", handlers.RemoveConf)

	r.POST("/start/:id/", handlers.Start)
	r.POST("/stop/:id/", handlers.Stop)
	r.POST("/shutdownserver/", handlers.ShutDownServer)

	r.GET("/stat/", handlers.Stat)

	//r.Run("0.0.0.0:8179") // listen and serve on 0.0.0.0:8080

	if err := endless.ListenAndServe("0.0.0.0:8179", r); err != nil {
		log.Fatalln(err)
	}
}
