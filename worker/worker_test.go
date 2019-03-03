package worker

import (
	. "github.com/zhwei820/gostresser/config"
	"runtime"
)

func init() {

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
		ReqConfs:           []ReqConf{reqConf},
	}
	Run(&baseConf)
}
