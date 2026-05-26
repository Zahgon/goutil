package httpreq

import "net/http"

// default standard client instance, with 500 ms timeout
var std = NewClient(500)

// Std instance
func Std() *Client {
	_ = "STUB: not implemented"

	// SetTimeout set default timeout(ms) for std client
	//
	// Note: timeout unit is millisecond
	return nil
}

func SetTimeout(ms int) { _ = "STUB: not implemented"; return }

// Config std http client
func Config(fn func(hc *http.Client)) { _ = "STUB: not implemented"; return }

//
// send request by default client
//

// Get quick sends a GET request by default client
func Get(url string, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Post quick sends a POST request by default client
}

func Post(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PostJSON quick sends a POST request by default client, with JSON content type
func PostJSON(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put quick send a PUT request by default client
func Put(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete quick send a DELETE request by default client
func Delete(url string, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send quick send a request by default client
func Send(method, url string, optFns ...OptionFn) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustSend quick send a request by default client
func MustSend(method, url string, optFns ...OptionFn) *http.Response {
	_ = "STUB: not implemented"
	return nil
}

// SendRequest quick send a request by default client
func SendRequest(req *http.Request, opt *Option) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
