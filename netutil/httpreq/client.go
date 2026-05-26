package httpreq

import (
	"io"
	"net/http"
)

// Client a simple http request client.
type Client struct {
	client Doer
	// default config for request
	method  string
	baseURL string
	timeout int // unit: ms
	// custom set default headers
	headerMap map[string]string

	// before send callback
	beforeSend func(req *http.Request)
	afterSend  AfterSendFn
}

// New instance with base URL and use http.Client as default http client
func New(baseURL ...string) *Client { _ = "STUB: not implemented"; return nil }

// NewWithTimeout new instance use http.Client and with custom timeout(ms)
func NewWithTimeout(ms int) *Client { _ = "STUB: not implemented"; return nil }

// NewWithDoer instance with a custom http client
func NewWithDoer(d Doer) *Client { _ = "STUB: not implemented"; return nil }

// init map

// Doer get the http client driver
func (h *Client) Doer() Doer {
	_ = "STUB: not implemented"

	// SetClient custom set http client doer
	return *new(Doer)
}

func (h *Client) SetClient(c Doer) *Client { _ = "STUB: not implemented"; return nil }

// SetTimeout set default timeout for http client doer
func (h *Client) SetTimeout(ms int) *Client { _ = "STUB: not implemented"; return nil }

// SetMaxIdleConns Set the maximum number of idle connections.
func (h *Client) SetMaxIdleConns(maxIdleConns, maxIdleConnsPerHost int) {
	_ = "STUB: not implemented"
	return
}

// BaseURL set request base URL
func (h *Client) BaseURL(baseURL string) *Client { _ = "STUB: not implemented"; return nil }

// DefaultMethod set default request method
func (h *Client) DefaultMethod(method string) *Client { _ = "STUB: not implemented"; return nil }

// ContentType set default content-Type header.
func (h *Client) ContentType(cType string) *Client { _ = "STUB: not implemented"; return nil }

// DefaultHeader set default header for all requests
func (h *Client) DefaultHeader(key, val string) *Client { _ = "STUB: not implemented"; return nil }

// DefaultHeaderMap set default headers for all requests
func (h *Client) DefaultHeaderMap(kvMap map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// OnBeforeSend add callback before send.
func (h *Client) OnBeforeSend(fn func(req *http.Request)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// OnAfterSend add callback after send.
func (h *Client) OnAfterSend(fn AfterSendFn) *Client { _ = "STUB: not implemented"; return nil }

//
// build request options
//

// WithOption with custom request options
func (h *Client) WithOption(optFns ...OptionFn) *Option { _ = "STUB: not implemented"; return nil }

func optWithClient(cli *Client) *Option { _ = "STUB: not implemented"; return nil }

// WithData with custom request data
func (h *Client) WithData(data any) *Option { _ = "STUB: not implemented"; return nil }

// WithBody with custom body
func (h *Client) WithBody(r io.Reader) *Option { _ = "STUB: not implemented"; return nil }

// BytesBody with custom bytes body
func (h *Client) BytesBody(bs []byte) *Option { _ = "STUB: not implemented"; return nil }

// StringBody with custom string body
func (h *Client) StringBody(s string) *Option { _ = "STUB: not implemented"; return nil }

// FormBody with custom form data body
func (h *Client) FormBody(data any) *Option { _ = "STUB: not implemented"; return nil }

// JSONBody with custom JSON data body
func (h *Client) JSONBody(data any) *Option { _ = "STUB: not implemented"; return nil }

// JSONBytesBody with custom bytes body, and set JSON content type
func (h *Client) JSONBytesBody(bs []byte) *Option { _ = "STUB: not implemented"; return nil }

// AnyBody with custom body.
//
// Allow type:
//   - string, []byte, map[string][]string/url.Values, io.Reader(eg: bytes.Buffer, strings.Reader)
func (h *Client) AnyBody(data any) *Option { _ = "STUB: not implemented"; return nil }

//
// ------------ send request with options ------------
//

// Get send GET request with options, return http response
func (h *Client) Get(url string, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Post send POST request with options, return http response
func (h *Client) Post(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostJSON send JSON POST request with options, return http response
func (h *Client) PostJSON(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put send PUT request with options, return http response
func (h *Client) Put(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete send DELETE request with options, return http response
func (h *Client) Delete(url string, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send request with option func, return http response
func (h *Client) Send(method, url string, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustSend request, will panic on error
func (h *Client) MustSend(method, url string, optFns ...OptionFn) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

// SendWithOpt request and return http response
func (h *Client) SendWithOpt(url string, opt *Option) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create request

// SendRequest send request and return http response
func (h *Client) SendRequest(req *http.Request, opt *Option) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// send request and return http response
}

func (h *Client) sendRequest(req *http.Request, opt *Option) (*http.Response, error) {
	_ = "STUB: not implemented"
	// apply default headers
	return nil, nil
}

// apply options

// - apply header map

// if timeout changed, create new client
