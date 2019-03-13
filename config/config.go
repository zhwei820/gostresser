package config

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"regexp"
)

// Global Config .
type BaseConf struct {
	Id bson.ObjectId `json:"_id" bson:"_id"`

	Name string `json:"name" bson:"name"` // name.
	H2   bool   `json:"h2" bson:"h2"`     // Enable HTTP/2.
	Cpus int    `json:"cpus" bson:"cpus"` // Number of used cpu cores. (default for current machine is %d cores)

	DisableCompression bool       `json:"disablecompression" bson:"disablecompression"` // Disable compression.
	DisableKeepAlives  bool       `json:"disablekeepalives" bson:"disablekeepalives"`   // Disable keep-alive, prevents re-use of TCP connections between different HTTP requests.
	DisableRedirects   bool       `json:"disableredirects" bson:"disableredirects"`     // Disable following of HTTP redirects
	ProxyAddr          string     `json:"proxyaddr" bson:"proxyaddr"`                   // HTTP Proxy address as host:port.
	Timeout            int        `json:"timeout" bson:"timeout"`                       // Timeout for each request in seconds. Default is 20, use 0 for infinite.
	Duration           int        `json:"duration" bson:"duration"`                     // Duration of a stress round (in second)
	Concurrency        int        `json:"concurrency" bson:"concurrency"`               // Num of Concurrent user
	ReqConfs           []*ReqConf `json:"reqconfs" bson:"reqconfs"`
}

// Each of your data model that needs to be persisted should implment gmgo.Document interface
func (user BaseConf) CollectionName() string {
	return "baseconf"
}

// Config of a single request.
type ReqConf struct {
	Url string `json:"url" bson:"url"` //

	Method      string      `json:"method" bson:"method"`           // HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
	Headers     HeaderSlice `json:"headers" bson:"headers"`         // Custom HTTP header. You can specify as many as needed by repeating the flag.       For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
	Body        string      `json:"body" bson:"body"`               // HTTP request body.
	Accept      string      `json:"accept" bson:"accept"`           // HTTP Accept header.
	ContentType string      `json:"contenttype" bson:"contenttype"` // Content-type, defaults to "text/html".
	AuthHeader  string      `json:"authheader" bson:"authheader"`   // Basic authentication, username:password.
	HostHeader  string      `json:"hostheader" bson:"hostheader"`   // HTTP Host header.

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
