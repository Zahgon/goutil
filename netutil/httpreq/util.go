package httpreq

import (
	"io"
	"net/http"
	"net/url"
)

// BasicAuthConf struct
type BasicAuthConf struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// IsValid value
func (ba *BasicAuthConf) IsValid() bool { _ = "STUB: not implemented"; return false }

// Value build to auth header "Authorization".
func (ba *BasicAuthConf) Value() string { _ = "STUB: not implemented"; return "" }

// String build to auth header "Authorization".
func (ba *BasicAuthConf) String() string { _ = "STUB: not implemented"; return "" }

// IsOK check response status code is 200
func IsOK(statusCode int) bool { _ = "STUB: not implemented"; return false }

// IsSuccessful check response status code is in 200-300
func IsSuccessful(statusCode int) bool { _ = "STUB: not implemented"; return false }

// IsRedirect check response status code is in [301, 302, 303, 307]
func IsRedirect(statusCode int) bool { _ = "STUB: not implemented"; return false }

// IsForbidden is this response forbidden(403)
func IsForbidden(statusCode int) bool { _ = "STUB: not implemented"; return false }

// IsNotFound is this response not found(404)
func IsNotFound(statusCode int) bool { _ = "STUB: not implemented"; return false }

// IsClientError check response is client error (400-500)
func IsClientError(statusCode int) bool { _ = "STUB: not implemented"; return false }

// IsServerError check response is server error (500-600)
func IsServerError(statusCode int) bool { _ = "STUB: not implemented"; return false }

// IsNoBodyMethod check
func IsNoBodyMethod(method string) bool { _ = "STUB: not implemented"; return false }

// IsValidMethod check method is valid
func IsValidMethod(method string) bool { _ = "STUB: not implemented"; return false }

// BuildBasicAuth returns the base64 encoded username:password for basic auth.
// Then set to header "Authorization".
//
// copied from net/http.
func BuildBasicAuth(username, password string) string { _ = "STUB: not implemented"; return "" }

// AddHeaders adds the key, value pairs from the given http.Header to the
// request. Values for existing keys are appended to the keys values.
func AddHeaders(req *http.Request, header http.Header) { _ = "STUB: not implemented"; return }

// SetHeaders sets the key, value pairs from the given http.Header to the
// request. Values for existing keys are overwritten.
func SetHeaders(req *http.Request, headers ...http.Header) { _ = "STUB: not implemented"; return }

// AddHeaderMap to request instance.
func AddHeaderMap(req *http.Request, headerMap map[string]string) {
	_ = "STUB: not implemented"
	return
}

// SetHeaderMap to request instance.
func SetHeaderMap(req *http.Request, headerMap map[string]string) {
	_ = "STUB: not implemented"
	return
}

// HeaderToStringMap convert
func HeaderToStringMap(rh http.Header) map[string]string { _ = "STUB: not implemented"; return nil }

// MakeQuery make query string, convert data to url.Values
func MakeQuery(data any) url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

// ToQueryValues convert string-map or any-map to url.Values
//
// data support:
//   - url.Values
//   - []byte
//   - string
//   - map[string][]string
//   - map[string]string
//   - map[string]any
func ToQueryValues(data any) url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

// use url.Values directly if we have it

// MergeURLValues merge url.Values by overwrite.
//
// values support: url.Values, map[string]string, map[string][]string
func MergeURLValues(uv url.Values, values ...any) url.Values {
	_ = "STUB: not implemented"
	return *new(url.Values)
}

// AppendQueryToURL appends the given query string to the given url.
func AppendQueryToURL(reqURL *url.URL, uv url.Values) error { _ = "STUB: not implemented"; return nil }

// url.Values format to a sorted "url encoded" string.
// e.g. "key=val&foo=bar"

// AppendQueryToURLString appends the given query data to the given url.
func AppendQueryToURLString(urlStr string, query url.Values) string {
	_ = "STUB: not implemented"
	return ""
}

// MakeBody make request body, convert data to io.Reader
func MakeBody(data any, cType string) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// ToRequestBody make request body, convert data to io.Reader
//
// Allow type for data:
//   - string
//   - []byte
//   - map[string]string
//   - map[string][]string/url.Values
//   - io.Reader(eg: bytes.Buffer, strings.Reader)
func ToRequestBody(data any, cType string) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// nobody

// encode body data to json

func toJSONReader(data any) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// close escape  &, <, >  TO  \u0026, \u003c, \u003e

// HeaderToString convert http Header to string
func HeaderToString(h http.Header) string { _ = "STUB: not implemented"; return "" }

// RequestToString convert http Request to string
func RequestToString(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// ResponseToString convert http Response to string
func ResponseToString(w *http.Response) string { _ = "STUB: not implemented"; return "" }

// ParseAccept header to strings. referred from gin framework
//
// eg: acceptHeader = "application/json, text/plain, */*"
func ParseAccept(acceptHeader string) []string { _ = "STUB: not implemented"; return nil }
