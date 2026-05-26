package finder

import (
	"bytes"
)

// Matcher for match file path.
type Matcher interface {
	// Apply check find elem. return False will skip this file.
	Apply(elem Elem) bool
}

// MatcherFunc for match file info, return False will skip this file
type MatcherFunc func(elem Elem) bool

// Apply check file path. return False will skip this file.
func (fn MatcherFunc) Apply(elem Elem) bool {
	_ = "STUB: not implemented"

	// ------------------ Multi matcher wrapper ------------------
	return false
}

// MultiMatcher wrapper for multi matchers
type MultiMatcher struct {
	Before   Matcher
	Matchers []Matcher
}

// Add matchers
func (mf *MultiMatcher) Add(fls ...Matcher) { _ = "STUB: not implemented"; return }

// Apply check file path is match.
func (mf *MultiMatcher) Apply(el Elem) bool { _ = "STUB: not implemented"; return false }

// NewDirMatchers create a new dir matchers
func NewDirMatchers(fls ...Matcher) *MultiMatcher { _ = "STUB: not implemented"; return nil }

// NewFileMatchers create a new dir matchers
func NewFileMatchers(fls ...Matcher) *MultiMatcher { _ = "STUB: not implemented"; return nil }

// ------------------ Body Matcher ------------------

// BodyMatcher for match file contents.
type BodyMatcher interface {
	Apply(filePath string, body *bytes.Buffer) bool
}

// BodyMatcherFunc for match file contents.
type BodyMatcherFunc func(filePath string, body *bytes.Buffer) bool

// Apply for match file contents.
func (fn BodyMatcherFunc) Apply(filePath string, body *bytes.Buffer) bool {
	_ = "STUB: not implemented"
	return false

	// BodyMatchers multi body matchers as Matcher
}

type BodyMatchers struct {
	Matchers []BodyMatcher
}

// NewBodyMatchers create a new body matchers
//
// Usage:
//
//		bf := finder.NewBodyMatchers(
//			finder.BodyMatcherFunc(func(filePath string, buf *bytes.Buffer) bool {
//				// match file contents
//				return true
//			}),
//		)
//
//	 es := finder.NewFinder('path/to/dir').Add(bf).Elems()
//	 for el := range es {
//			fmt.Println(el.Path())
//	 }
func NewBodyMatchers(fls ...BodyMatcher) *BodyMatchers { _ = "STUB: not implemented"; return nil }

// AddMatcher add matchers
func (mf *BodyMatchers) AddMatcher(fls ...BodyMatcher) { _ = "STUB: not implemented"; return }

// Apply check file contents is match.
func (mf *BodyMatchers) Apply(el Elem) bool { _ = "STUB: not implemented"; return false }

// read file contents

// apply matchers
