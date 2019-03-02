package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/narup/gmgo"
	"github.com/swaggo/swag/example/basic/web"
	"github.com/zhwei820/gostresser/config"
	"net/http"
)

func init() {
	_ = web.APIError{}

}

// @Summary ListConf
// @Description ListConf
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /conf [get]
func ListConf(c *gin.Context) {
	res, _ := config.BaseConfManager().FindAll(gmgo.Q{})
	c.JSON(200, res)
}

// @Summary CreateConf
// @Description CreateConf
// @Accept  json
// @Produce  json
// @Param  account body config.BaseConf true "create BaseConf"
// @Success 200 {string} string	"ok"
// @Router /conf [post]
func CreateConf(c *gin.Context) {
	var baseConf config.BaseConf
	if err := c.ShouldBindJSON(&baseConf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, _ := config.BaseConfManager().Create(baseConf)
	c.JSON(200, res)
}

// @Summary UpdateConf
// @Description UpdateConf
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Param  account body config.BaseConf true "create BaseConf"
// @Success 200 {string} string	"ok"
// @Router /conf/{id} [put]
func UpdateConf(c *gin.Context) {
	id := c.Param("id")
	var baseConf config.BaseConf
	if err := c.ShouldBindJSON(&baseConf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	baseConf.Id = bson.ObjectIdHex(id)

	res, _ := config.BaseConfManager().Save(baseConf)
	c.JSON(200, res)
}

// @Summary DetailConf
// @Description DetailConf
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /conf/{id} [get]
func DetailConf(c *gin.Context) {
	id := c.Param("id")

	res, _ := config.BaseConfManager().FindOne(id)
	c.JSON(200, res)
}

// @Summary RemoveConf
// @Description RemoveConf
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Router /conf/{id} [delete]
func RemoveConf(c *gin.Context) {
	id := c.Param("id")

	res, _ := config.BaseConfManager().RemoveOne(id)
	c.JSON(200, res)
}
