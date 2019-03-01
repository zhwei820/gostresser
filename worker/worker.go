package worker

import (
	. "github.com/zhwei820/gostresser/config"
	"github.com/zhwei820/gostresser/hey/requester"
	"net/http"
	gourl "net/url"
	"runtime"
	"time"
)

const (
	headerRegexp = `^([\w-]+):\s*(.+)`
	authRegexp   = `^(.+):([^\s].+)`
	heyUA        = "hey/0.0.1"
)

func Run() {
	baseConf := BaseConf{
		H2:                 false,
		Cpus:               runtime.GOMAXPROCS(-1),
		DisableCompression: false,
		DisableKeepAlives:  false,
		DisableRedirects:   false,
		ProxyAddr:          "",
		Duration:           10,
		Timeout:            10,
		Concurrency:        20,
	}
	reqConf := ReqConf{
		Url:     "https://www.baidu.com/",
		Method:  "Get",
		Headers: HeaderSlice{"Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"},
		Body:    "",
		//BodyFile:"",
		Accept:      "",
		ContentType: "",
		AuthHeader:  "",
		HostHeader:  "",
		ReqFreq:     5,
	}

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

	w := &requester.Work{
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
	w.Init()

	if baseConf.Duration > 0 {
		go func() {
			time.Sleep(time.Duration(baseConf.Duration) * time.Second)
			w.Stop()
		}()
	}
	w.Run()
}
