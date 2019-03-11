package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhwei820/gostresser/config"
	"log"
	"os/exec"
)

var ChildPids = make(map[string]int)

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

	cmd1 := exec.Command("sh", "-c", "./worker/worker -n 30000 -c 10 http://127.0.0.1:8000")
	err := cmd1.Run()
	if err != nil {
		log.Fatal(err)
	}

	ChildPids[baseConf.Id.String()+baseConf.ReqConfs[0].Url] = cmd1.Process.Pid
	out, _ := cmd1.Output()
	fmt.Printf("%s\n", out)
	fmt.Printf("%s\n", out)
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
	stopcmd := fmt.Sprintf("kill -9 %v", ChildPids[baseConf.Id.String()+baseConf.ReqConfs[0].Url])

	cmd1 := exec.Command("sh", "-c", stopcmd)
	err := cmd1.Start()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, "ok")
}

func getStartCmd() string {
	return ""
}

func getStopCmd() string {
	return ""
}
