package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/zhwei820/gostresser/handlers"
)

func StartApi() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	r.GET("/conf", handlers.ListConf)
	r.POST("/conf", handlers.CreateConf)
	r.PUT("/conf/:id", handlers.UpdateConf)
	r.GET("/conf/:id", handlers.DetailConf)
	r.DELETE("/conf/:id", handlers.RemoveConf)

	r.POST("/start/:id", handlers.Start)
	r.POST("/stop/:id", handlers.Stop)

	r.GET("/stat/", handlers.Stat)

	r.Run("0.0.0.0:8179") // listen and serve on 0.0.0.0:8080
}
