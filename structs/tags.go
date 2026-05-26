package structs

import (
	"errors"
	"reflect"

	"github.com/gookit/goutil/maputil"
)

// ErrNotAnStruct error
// var emptyStringMap = make(maputil.SMap)
var ErrNotAnStruct = errors.New("must input an struct value")

// ParseTags for parse struct tags.
func ParseTags(st any, tagNames []string) (map[string]maputil.SMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseReflectTags parse struct tags info.
func ParseReflectTags(rt reflect.Type, tagNames []string) (map[string]maputil.SMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TagValFunc handle func
type TagValFunc func(field, tagVal string) (maputil.SMap, error)

// TagParser struct
type TagParser struct {
	// TagNames want parsed tag names.
	TagNames []string
	// ValueFunc tag value parse func.
	ValueFunc TagValFunc

	// key: field name
	// value: tag map {tag-name: value string.}
	tags map[string]maputil.SMap
}

// Tags map data for struct fields
func (p *TagParser) Tags() map[string]maputil.SMap {
	_ = "STUB: not implemented"

	// NewTagParser instance
	return nil
}

func NewTagParser(tagNames ...string) *TagParser { _ = "STUB: not implemented"; return nil }

// Parse an struct value
func (p *TagParser) Parse(st any) error { _ = "STUB: not implemented"; return nil }

// ParseType parse a struct type value
func (p *TagParser) ParseType(rt reflect.Type) error { _ = "STUB: not implemented"; return nil }

// key is field name.

func (p *TagParser) parseType(rt reflect.Type, parent string) error {
	_ = "STUB: not implemented"
	return nil
}

// skip don't exported field

// eg: `json:"age"`
// eg: "name=int0;shorts=i;required=true;desc=int option message"

// field is struct.

// Info parse the give field, returns tag value info.
//
//	info, err := p.Info("Name", "json")
//	exportField := info.Get("name")
func (p *TagParser) Info(field, tag string) (maputil.SMap, error) {
	_ = "STUB: not implemented"
	return *new(maputil.SMap), nil
}

// parse tag value

/*************************************************************
 * some built in tag value parse func
 *************************************************************/

// ParseTagValueDefault parse like json tag value.
//
// see json.Marshal():
//
//	// JSON as key "myName", skipped if empty.
//	Field int `json:"myName,omitempty"`
//
//	// Field appears in JSON as key "Field" (the default), but skipped if empty.
//	Field int `json:",omitempty"`
//
//	// Field is ignored by this package.
//	Field int `json:"-"`
//
//	// Field appears in JSON as key "-".
//	Field int `json:"-,"`
//
//	Int64String int64 `json:",string"`
//
// Returns:
//
//	{
//		"name": "myName", // maybe is empty, on tag value is "-"
//		"omitempty": "true",
//		"string": "true",
//		// ... more custom bool settings.
//	}
func ParseTagValueDefault(field, tagVal string) (mp maputil.SMap, err error) {
	_ = "STUB: not implemented"
	return *new(maputil.SMap), nil
}

// valid field name

// ln > 1

// other settings: omitempty, string

// ParseTagValueQuick quick parse tag value string by sep(;)
func ParseTagValueQuick(tagVal string, defines []string) maputil.SMap {
	_ = "STUB: not implemented"
	return *new(maputil.SMap)
}

// ParseTagValueDefine parse tag value string by given defines.
//
// Examples:
//
//	eg: "desc;required;default;shorts"
//	type MyStruct {
//		Age int `flag:"int option message;;a,b"`
//	}
//	sepStr := ";"
//	defines := []string{"desc", "required", "default", "shorts"}
func ParseTagValueDefine(sep string, defines []string) TagValFunc {
	_ = "STUB: not implemented"
	return *new(TagValFunc)
}

// ParseTagValueNamed parse k-v tag value string. it's like INI format contents.
//
// Examples:
//
//	eg: "name=val0;shorts=i;required=true;desc=a message"
//	=>
//	{name: val0, shorts: i, required: true, desc: a message}
func ParseTagValueNamed(field, tagVal string, keys ...string) (mp maputil.SMap, err error) {
	_ = "STUB: not implemented"
	return *new(maputil.SMap), nil
}
