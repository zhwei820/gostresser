package worker

import (
	"fmt"
	. "github.com/zhwei820/gostresser/config"
	"github.com/zhwei820/gostresser/hey/requester"
	"github.com/zhwei820/gostresser/utils"
	"net/http"
	gourl "net/url"
	"runtime"
	"strconv"
	"time"
)

var Workers = make(map[string]*requester.Work)

const (
	headerRegexp = `^([\w-]+):\s*(.+)`
	authRegexp   = `^(.+):([^\s].+)`
	heyUA        = "hey/0.0.1"
)

func workerRun(baseConf *BaseConf, reqConf ReqConf, worker_key string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v\n", err)
			fmt.Printf("%s\n", utils.Stack(3))
		}
	}()
	runtime.GOMAXPROCS(baseConf.Cpus)

	// set content-type
	header := make(http.Header)
	header.Set("Content-Type", reqConf.ContentType)

	// set any other additional repeatable headers
	for _, h := range reqConf.Headers {
		match, _ := ParseInputWithRegexp(h, headerRegexp)

		header.Set(match[1], match[2])
	}

	if reqConf.Accept != "" {
		header.Set("Accept", reqConf.Accept)
	}

	// set basic auth if set
	var username, password string
	if reqConf.AuthHeader != "" {
		match, _ := ParseInputWithRegexp(reqConf.AuthHeader, authRegexp)

		username, password = match[1], match[2]
	}

	var bodyAll []byte
	if reqConf.Body != "" {
		bodyAll = []byte(reqConf.Body)
	}

	var proxyURL *gourl.URL
	if baseConf.ProxyAddr != "" {
		proxyURL, _ = gourl.Parse(baseConf.ProxyAddr)
	}

	req, _ := http.NewRequest(reqConf.Method, reqConf.Url, nil)

	req.ContentLength = int64(len(bodyAll))
	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}

	// set host header if set
	if reqConf.HostHeader != "" {
		req.Host = reqConf.HostHeader
	}

	ua := req.UserAgent()
	if ua == "" {
		ua = heyUA
	} else {
		ua += " " + heyUA
	}
	header.Set("User-Agent", ua)
	req.Header = header

	Workers[worker_key] = &requester.Work{
		Request:            req,
		RequestBody:        bodyAll,
		N:                  1000000,              // big num
		C:                  baseConf.Concurrency, // update by reqFreq
		QPS:                float64(100000),      // big num
		Timeout:            baseConf.Timeout,
		DisableCompression: baseConf.DisableCompression,
		DisableKeepAlives:  baseConf.DisableKeepAlives,
		DisableRedirects:   baseConf.DisableRedirects,
		H2:                 baseConf.H2,
		ProxyAddr:          proxyURL,
		Output:             "",
	}
	Workers[worker_key].Init()

	if baseConf.Duration > 0 {
		go func() {
			time.Sleep(time.Duration(baseConf.Duration) * time.Second)
			Workers[worker_key].Stop()
		}()
	}
	Workers[worker_key].Run()

}

func Run(baseConf *BaseConf) {
	for ii, reqConf := range baseConf.ReqConfs {
		worker_key := baseConf.Id.String() + strconv.Itoa(ii)
		go workerRun(baseConf, reqConf, worker_key)
		time.Sleep(100 * time.Millisecond)
	}
}

func Stop(baseConf *BaseConf) {
	for ii := range baseConf.ReqConfs {
		worker_key := baseConf.Id.String() + strconv.Itoa(ii)
		Workers[worker_key].Stop()
	}
}
