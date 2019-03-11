package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhwei820/gostresser/stat"
)

// @Summary stat
// @Description stat
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /stat [get]
func Stat(c *gin.Context) {
	res := stat.StatReqs()
	c.JSON(200, res)
}
