package config

import (
	"fmt"
	"regexp"
)

// Global Config .
type BaseConf struct {
	H2   bool // Enable HTTP/2.
	Cpus int  // Number of used cpu cores. (default for current machine is %d cores)

	DisableCompression bool   // Disable compression.
	DisableKeepAlives  bool   // Disable keep-alive, prevents re-use of TCP connections between different HTTP requests.
	DisableRedirects   bool   // Disable following of HTTP redirects
	ProxyAddr          string // HTTP Proxy address as host:port.
	Timeout            int    // Timeout for each request in seconds. Default is 20, use 0 for infinite.
	Duration           int    // Duration of a stress round (in second)
	Concurrency        int    // Num of Concurrent user
}

// Config of a single request.
type ReqConf struct {
	Url string

	Method  string      // HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
	Headers HeaderSlice // Custom HTTP header. You can specify as many as needed by repeating the flag.       For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
	Body    string      // HTTP request body.
	//BodyFile    string      // HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
	Accept      string // HTTP Accept header.
	ContentType string // Content-type, defaults to "text/html".
	AuthHeader  string // Basic authentication, username:password.
	HostHeader  string // HTTP Host header.

	ReqFreq int // How offten this request called in a stress test,  1~10
}

func ParseInputWithRegexp(input, regx string) ([]string, error) {
	re := regexp.MustCompile(regx)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 1 {
		return nil, fmt.Errorf("could not parse the provided input; input = %v", input)
	}
	return matches, nil
}

type HeaderSlice []string

func (h *HeaderSlice) String() string {
	return fmt.Sprintf("%s", *h)
}

func (h *HeaderSlice) Set(value string) error {
	*h = append(*h, value)
	return nil
}
