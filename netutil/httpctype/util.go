package httpctype

// ToKind name match
func ToKind(cType, defaultType string) string { _ = "STUB: not implemented"; return "" }

// ToKindWithFunc match base kind name by content-type, with a fallback func
func ToKindWithFunc(cType string, fbFunc func(cType string) string) string {
	_ = "STUB: not implemented"
	// JSON body request: "application/json"
	return ""
}

// basic POST form data binding. content type: "application/x-www-form-urlencoded"

// contains file uploaded form: "multipart/form-data" "multipart/mixed"
// strings.HasPrefix(mediaType, "multipart/")

// XML body request: "text/xml"
