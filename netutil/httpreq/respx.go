package httpreq

import (
	"bytes"
	"net/http"
)

// Resp alias of RespX
type Resp = RespX

// RespX wrap http.Response and add some useful methods.
type RespX struct {
	*http.Response
	// CostTime for a request-response
	CostTime int64
	// body data buffer
	bodyBuf *bytes.Buffer
}

// WrapResp wrap http.Response to RespX
func WrapResp(hr *http.Response, err error) (*RespX, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewResp instance
func NewResp(hr *http.Response) *RespX { _ = "STUB: not implemented"; return nil }

// IsFail check status code is not equals to 200
func (r *RespX) IsFail() bool { _ = "STUB: not implemented"; return false }

// IsOk check status code is equals to 200
func (r *RespX) IsOk() bool { _ = "STUB: not implemented"; return false }

// IsSuccessful check status code is in 200-300
func (r *RespX) IsSuccessful() bool { _ = "STUB: not implemented"; return false }

// IsEmptyBody check response body is empty
func (r *RespX) IsEmptyBody() bool { _ = "STUB: not implemented"; return false }

// ContentType get response content type
func (r *RespX) ContentType() string { _ = "STUB: not implemented"; return "" }

// BodyString get body as string.
func (r *RespX) String() string { _ = "STUB: not implemented"; return "" }

//
// ------------------------ read body ------------------------
//

// ReadBody read body to buffer. allow reading body multiple times.
//
// NOTE: will close the response.Body
func (r *RespX) ReadBody() error { _ = "STUB: not implemented"; return nil }

// prof: assign memory before read

// NOTICE: must close resp body.

// BodyBuffer read body to buffer. NOTE: will close the response.Body
func (r *RespX) BodyBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

// BodyString get body as string.
func (r *RespX) BodyString() string { _ = "STUB: not implemented"; return "" }

// BindJSONOnOk body data on response status is in 200-300.
// If ptr is nil, will do nothing.
func (r *RespX) BindJSONOnOk(ptr any) error { _ = "STUB: not implemented"; return nil }

// BindJSON body data to a ptr, will don't check status code.
// If ptr is nil, will do nothing.
func (r *RespX) BindJSON(ptr any) error { _ = "STUB: not implemented"; return nil }

func (r *RespX) bindJSON(ptr any, checkStatus bool) error { _ = "STUB: not implemented"; return nil }

// CloseBuffer close body buffer
func (r *RespX) CloseBuffer() { _ = "STUB: not implemented"; return }

// CloseBody close resp body
func (r *RespX) CloseBody() error { _ = "STUB: not implemented"; return nil }

// SafeCloseBody close resp body, ignore error
func (r *RespX) SafeCloseBody() { _ = "STUB: not implemented"; return }
