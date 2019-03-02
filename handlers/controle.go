package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhwei820/gostresser/config"
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

	res, _ := config.BaseConfManager().FindOne(id)
	c.JSON(200, res)
}

// @Summary End
// @Description End
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /end/{id} [post]
func End(c *gin.Context) {
	id := c.Param("id")

	res, _ := config.BaseConfManager().FindOne(id)
	c.JSON(200, res)
}
