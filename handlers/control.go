package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhwei820/gostresser/config"
	"github.com/zhwei820/gostresser/worker"
)

// @Summary Start
// @Description Start
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /start/{id} [post]
func Start(c *gin.Context) {
	id := c.Param("id")

	baseConf, _ := config.BaseConfManager().FindOne(id)
	worker.Run(baseConf)
	c.JSON(200, "ok")
}

// @Summary End
// @Description End
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /end/{id} [post]
func Stop(c *gin.Context) {
	id := c.Param("id")

	baseConf, _ := config.BaseConfManager().FindOne(id)
	worker.Stop(baseConf)
	c.JSON(200, "ok")
}

// @Summary ShutDownServer anyway
// @Description ShutDownServer anyway
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /shutdownserver/ [post]
func ShutDownServer(c *gin.Context) {

	worker.ShutDownServer()
	c.JSON(200, "ok")
}
