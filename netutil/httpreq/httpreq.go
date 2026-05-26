// Package httpreq provide an simple http requester and some useful util functions.
package httpreq

import (
	"net/http"
	"sync"
	"time"
)

// ValidMethods valid http methods
var ValidMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodHead,
	http.MethodOptions,
	http.MethodTrace,
}

// AfterSendFn callback func
type AfterSendFn func(resp *http.Response, err error)

// Doer interface for an http client.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// DoerFunc implements the Doer
type DoerFunc func(req *http.Request) (*http.Response, error)

// Do send request and return response.
func (do DoerFunc) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"

	// ReqLogger request logger interface
	return nil, nil
}

type ReqLogger interface {
	Infof(format string, args ...any)
	Errorf(format string, args ...any)
}

const defaultTimeout = 500 * time.Millisecond

var (
	// global lock
	_gl = sync.Mutex{}

	// client cache map
	cs = map[int]*Client{}
)

// NewClient create a new http client and cache it.
//
// Note: timeout unit is millisecond
func NewClient(timeout int) *Client { _ = "STUB: not implemented"; return nil }

// MustResp check error and return response
func MustResp(r *http.Response, err error) *http.Response { _ = "STUB: not implemented"; return nil }

// MustRespX check error and create a new RespX instance
func MustRespX(r *http.Response, err error) *RespX { _ = "STUB: not implemented"; return nil }

// WithJSONType set request content type to JSON
func WithJSONType(opt *Option) { _ = "STUB: not implemented"; return }

// WithData set request data, will auto convert to body data or query string
func WithData(data any) OptionFn { _ = "STUB: not implemented"; return *new(OptionFn) }
