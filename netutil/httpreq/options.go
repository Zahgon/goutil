package httpreq

import (
	"context"
	"io"
	"net/http"
)

// Options alias of Option
type Options = Option

// Option struct
type Option struct {
	cli  *Client
	sent bool
	// Timeout for request. unit: ms
	Timeout int
	// Method for request
	Method string
	// HeaderMap data. eg: traceid
	HeaderMap map[string]string
	// ContentType header
	ContentType string

	// Logger for request
	Logger ReqLogger
	// Context for request
	Context context.Context

	// Data for request. can be used on any request method.
	//
	// type allow:
	// 	string, []byte, io.Reader, map[string]string, ...
	Data any
	// Body data for request. used on POST, PUT, PATCH method.
	//
	// eg: strings.NewReader("name=inhere")
	Body io.Reader
}

// OptionFn option func type
type OptionFn func(opt *Option)

// OptOrNew create a new Option if opt is nil
func OptOrNew(opt *Option) *Option { _ = "STUB: not implemented"; return nil }

// NewOpt create a new Option and with option func
func NewOpt(fns ...OptionFn) *Option { _ = "STUB: not implemented"; return nil }

// NewOption create a new Option and set option func
func NewOption(fns []OptionFn) *Option { _ = "STUB: not implemented"; return nil }

// WithOptionFn set option func
func (o *Option) WithOptionFn(fns ...OptionFn) *Option { _ = "STUB: not implemented"; return nil }

// WithOptionFns set option func
func (o *Option) WithOptionFns(fns []OptionFn) *Option { _ = "STUB: not implemented"; return nil }

// WithClient set client
func (o *Option) WithClient(cli *Client) *Option { _ = "STUB: not implemented"; return nil }

// Copy option for new request, use for repeat send request
func (o *Option) Copy() *Option { _ = "STUB: not implemented"; return nil }

// WithMethod set method
func (o *Option) WithMethod(method string) *Option { _ = "STUB: not implemented"; return nil }

// WithContentType set content type
func (o *Option) WithContentType(ct string) *Option { _ = "STUB: not implemented"; return nil }

// WithHeaderMap set header map
func (o *Option) WithHeaderMap(m map[string]string) *Option { _ = "STUB: not implemented"; return nil }

// WithHeader set header
func (o *Option) WithHeader(key, val string) *Option { _ = "STUB: not implemented"; return nil }

// WithData with custom data
func (o *Option) WithData(data any) *Option { _ = "STUB: not implemented"; return nil }

// AnyBody with custom body.
//
// Allow type:
//   - string, []byte, map[string][]string/url.Values, io.Reader(eg: bytes.Buffer, strings.Reader)
func (o *Option) AnyBody(data any) *Option { _ = "STUB: not implemented"; return nil }

// WithBody with custom body
func (o *Option) WithBody(r io.Reader) *Option { _ = "STUB: not implemented"; return nil }

// BytesBody with custom bytes body
func (o *Option) BytesBody(bs []byte) *Option { _ = "STUB: not implemented"; return nil }

// FormBody with custom form body data
func (o *Option) FormBody(data any) *Option { _ = "STUB: not implemented"; return nil }

// WithJSON with custom JSON body
func (o *Option) WithJSON(data any) *Option { _ = "STUB: not implemented"; return nil }

// JSONBytesBody with custom bytes body, and set JSON content type
func (o *Option) JSONBytesBody(bs []byte) *Option { _ = "STUB: not implemented"; return nil }

// StringBody with custom string body
func (o *Option) StringBody(s string) *Option { _ = "STUB: not implemented"; return nil }

//
// send request with options
//

// Get send GET request and return http response
func (o *Option) Get(url string, fns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Post send POST request and return http response
func (o *Option) Post(url string, data any, fns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put send PUT request and return http response
func (o *Option) Put(url string, data any, fns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete send DELETE request and return http response
func (o *Option) Delete(url string, fns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send request and return http response
func (o *Option) Send(method, url string, fns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustSend request. will panic on error
func (o *Option) MustSend(method, url string, fns ...OptionFn) *http.Response {
	_ = "STUB: not implemented"
	return nil
}
