package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhwei820/gostresser/worker"
)

// @Summary stat
// @Description stat
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /stat [get]
func Stat(c *gin.Context) {
	res := worker.StatReqs()
	c.JSON(200, res)
}
