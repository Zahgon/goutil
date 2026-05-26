package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// some data.
type (
	// M short name for a string-map
	M map[string]string
	// MD simple request data
	MD struct {
		// Headers headers
		Headers M
		// Body reader. eg: strings.NewReader("name=inhere")
		Body io.Reader
		// BodyString quick adds string body.
		BodyString string
		// BeforeSend callback
		BeforeSend func(req *http.Request)
	}
)

// NewHTTPRequest quick create request for http testing
// Usage:
//
//	req := NewHttpRequest("GET", "/path", nil)
//
//	// with data 1
//	body := strings.NewReader("string ...")
//	req := NewHttpRequest("POST", "/path", &MD{
//		Body: body,
//		Headers: M{"x-head": "val"}
//	})
//
//	// with data 2
//	req := NewHttpRequest("POST", "/path", &MD{
//		BodyString: "data string",
//		Headers: M{"x-head": "val"}
//	})
func NewHTTPRequest(method, path string, data *MD) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

// create fake request

// MockRequest mock an HTTP Request
//
// Usage:
//
//	handler := router.New()
//	res := MockRequest(handler, "GET", "/path", nil)
//
//	// with data 1
//	body := strings.NewReader("string ...")
//	res := MockRequest(handler, "POST", "/path", &MD{
//		Body: body,
//		HeaderM: M{"x-head": "val"}
//	})
//
//	// with data 2
//	res := MockRequest(handler, "POST", "/path", &MD{
//		BodyString: "data string",
//		HeaderM: M{"x-head": "val"}
//	})
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder {
	_ = "STUB: not implemented"
	// w.Result() will return http.Response
	return nil
}

// s := httptest.NewServer()

// EchoReply http response data reply model
type EchoReply struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
	Method string `json:"method"`
	// Query data
	Query map[string]any `json:"query,omitempty"`
	// Headers data.
	//
	// If value is one elem, will return string, otherwise will return []string
	//
	// Example:
	// 	map[string]any{
	//		"Connection": "close",
	//		"Vary": []string{"Accept-Encoding", "Accept-Encoding"},
	//	}
	Headers map[string]any `json:"headers,omitempty"`
	// Form data.
	//
	// If value is one elem, will return string, otherwise will return []string
	Form map[string]any `json:"form,omitempty"`
	// Body data string from request body
	Body string `json:"body,omitempty"`
	// JSON data on Content-Type: application/json
	//  - 通常是 map[string]any 类型数据
	JSON any `json:"json,omitempty"`
	// Files data.
	Files map[string]any `json:"files,omitempty"`
}

// ContentType get content type
func (r *EchoReply) ContentType() string { _ = "STUB: not implemented"; return "" }

// JSONMap assert JSON data to map[string]any
func (r *EchoReply) JSONMap() map[string]any { _ = "STUB: not implemented"; return nil }

// HeaderString get header value as string
func (r *EchoReply) HeaderString(name string) string { _ = "STUB: not implemented"; return "" }

// EchoServer for testing http request.
type EchoServer struct {
	*httptest.Server
}

// HostAddr get host address. eg: 127.0.0.1:8999
func (s *EchoServer) HostAddr() string { _ = "STUB: not implemented"; return "" }

// HTTPHost get http host address. eg: http://127.0.0.1:8999
func (s *EchoServer) HTTPHost() string { _ = "STUB: not implemented"; return "" }

// PrintHttpHost print host address to console
func (s *EchoServer) PrintHttpHost() string { _ = "STUB: not implemented"; return "" }

// HandleRequest handle request
func (s *EchoServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// eg. GET /404

// eg. GET /500

// custom reply status code. eg: /status-{code}

// 405 eg: "GET /post"

// default: 200 ok

// w.Header().Set("Connection", "close")

// MockHttpServer create an echo server for testing. alias of NewEchoServer
func MockHttpServer() *EchoServer { _ = "STUB: not implemented"; return nil }

// NewEchoServer create an echo server for testing.
//
// Usage on testing:
//
//	var testSrvAddr string
//
//	func TestMain(m *testing.M) {
//		// create server
//		s := testutil.NewEchoServer()
//		defer s.Close()
//		testSrvAddr = s.PrintHttpHost()
//
//		m.Run()
//	}
//
//	// in a test case ...
//	res := http.Get(testSrvAddr + "/get/some-one")
//	rpl := testutil.ParseRespToReply(res)
//	// assert ...
func NewEchoServer() *EchoServer { _ = "STUB: not implemented"; return nil }

// BuildEchoReply build reply body data
func BuildEchoReply(r *http.Request) *EchoReply {
	_ = "STUB: not implemented"
	// get headers
	return nil
}

// get query args

// get form data

// get form files

// get body data

// defer r.Body.Close()

// try to parse json

/*
// HTTP tool for testing
var HTTP = &HTTPTool{}

// HTTPTool http tool for testing
type HTTPTool struct {
}

func (ht *HTTPTool) ParseRespToReply(r *http.Response) *EchoReply {
	return ParseRespToReply(r)
}

func (ht *HTTPTool) ParseBodyToReply(bd io.ReadCloser) *EchoReply {
	return ParseBodyToReply(bd)
}
*/

// ParseRespToReply parse http response to reply
func ParseRespToReply(w *http.Response) *EchoReply { _ = "STUB: not implemented"; return nil }

// ParseBodyToReply parse http body to reply
func ParseBodyToReply(bd io.ReadCloser) *EchoReply { _ = "STUB: not implemented"; return nil }

func stringsMapToAnyMap(ssMp map[string][]string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
