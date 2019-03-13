package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhwei820/gostresser/config"
	"log"
	"os"
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

	for _, item := range baseConf.ReqConfs {
		cmdID := baseConf.Id.Hex() + item.Url
		cmd1 := exec.Command("sh", "-c", getStartCmd(*item, *baseConf, cmdID))
		cmd1.Stdin = os.Stdin
		cmd1.Stdout = os.Stdout
		cmd1.Stderr = os.Stderr

		err := cmd1.Start()
		if err != nil {
			log.Fatal(err)
		}
		ChildPids[cmdID] = cmd1.Process.Pid

	}

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
	for _, item := range baseConf.ReqConfs {

		stopcmd := fmt.Sprintf("kill -9 %v", ChildPids[baseConf.Id.Hex()+item.Url])
		println(stopcmd)
		cmd1 := exec.Command("sh", "-c", stopcmd)
		err := cmd1.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(200, "ok")
}

func getStartCmd(conf config.ReqConf, baseConf config.BaseConf, id string) string {
	//	`
	//  -n  Number of requests to run. Default is 200.
	//  -c  Number of requests to run concurrently. Total number of requests cannot
	//      be smaller than the concurrency level. Default is 50.
	//  -q  Rate limit, in queries per second (QPS). Default is no rate limit.
	//  -z  Duration of application to send requests. When duration is reached,
	//      application stops and exits. If duration is specified, n is ignored.
	//      Examples: -z 10s -z 3m.
	//  -o  Output type. If none provided, a summary is printed.
	//      "csv" is the only supported alternative. Dumps the response
	//      metrics in comma-separated values format.
	//
	//  -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
	//  -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
	//      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
	//  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
	//  -A  HTTP Accept header.
	//  -d  HTTP request body.
	//  -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
	//  -T  Content-type, defaults to "text/html".
	//  -a  Basic authentication, username:password.
	//  -x  HTTP Proxy address as host:port.
	//  -h2 Enable HTTP/2.
	//
	//  -host HTTP Host header.
	//
	//  -disable-compression  Disable compression.
	//  -disable-keepalive    Disable keep-alive, prevents re-use of TCP
	//                        connections between different HTTP requests.
	//  -disable-redirects    Disable following of HTTP redirects
	//  -cpus                 Number of used cpu cores.
	//                        (default for current machine is 4 cores)
	//
	//`
	s := fmt.Sprintf("./worker/worker -n 10000000 -id %v -c %v ", id, baseConf.Concurrency)
	if conf.ReqFreq != 0 {
		s += fmt.Sprintf(" -q %v", conf.ReqFreq)
	}
	if baseConf.Duration != 0 {
		s += fmt.Sprintf(" -z %vs", baseConf.Duration)
	}
	if conf.Method != "" {
		s += fmt.Sprintf(" -m %v", conf.Method)
	}
	//if conf.Headers != nil {
	//	s += fmt.Sprintf(" -z %v", conf.Headers)
	//}
	if baseConf.Timeout != 0 {
		s += fmt.Sprintf(" -t %v", baseConf.Timeout)
	}
	if conf.Accept != "" {
		s += fmt.Sprintf(" -A %v", conf.Accept)
	}
	if conf.Body != "" {
		s += fmt.Sprintf(" -d %v", conf.Body)
	}
	if conf.ContentType != "" {
		s += fmt.Sprintf(" -T %v", conf.ContentType)
	}
	if conf.Body != "" {
		s += fmt.Sprintf(" -d %v", conf.Body)
	}
	if conf.AuthHeader != "" {
		s += fmt.Sprintf(" -a %v", conf.AuthHeader)
	}
	if conf.HostHeader != "" {
		s += fmt.Sprintf(" -host %v", conf.HostHeader)
	}
	if baseConf.DisableCompression {
		s += fmt.Sprintf(" -disable-compression %v", baseConf.DisableCompression)
	}
	if baseConf.DisableKeepAlives {
		s += fmt.Sprintf(" -disable-keepalive %v", baseConf.DisableKeepAlives)
	}
	if baseConf.DisableRedirects {
		s += fmt.Sprintf(" -disable-redirects %v", baseConf.DisableRedirects)
	}

	return s + " " + conf.Url
}
