package goutils

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type _Error struct {
	Location string `json:"location"`
	Type     string `json:"type"`
	Data     KV     `json:"data"`
	Inner    *Error `json:"inner,omitempty"`
}

// Error is an enhanced error object that can hold information about the
// location of error, type of error, associated data, and optionally an
// inner error that it wraps.
type Error struct {
	_err _Error
}

// NewError creates a new error context of the given error type.
func NewError(t string) *Error {
	t = strings.TrimSpace(t)
	if t == "" {
		return nil
	}
	var e Error
	e._err.Type = t
	return &e
}

// Wrap creates a new error context of the given error type, with the
// given error retained as that from the underlying layer.
func Wrap(t string, inner *Error) *Error {
	e := NewError(t)
	if e == nil || inner == nil {
		return nil
	}
	e._err.Inner = inner
	return e
}

// Add can be used to insert associated data into the context of this
// error, as string keys and any values.
//
// N.B. When you have to add your own associated data to an existing
// error, you must wrap the underlying error before adding your data.
// The trace and the context could get corrupted otherwise.
func (e *Error) Add(k string, v interface{}) *Error {
	if e._err.Location == "" {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			comps := strings.SplitAfter(file, "/")
			idx := len(comps) - 3
			if idx < 0 {
				idx = 0
			}
			e._err.Location = fmt.Sprintf("%s:%d", strings.Join(comps[idx:], ""), line)
		}
	}

	k = strings.TrimSpace(k)
	if k == "" || v == nil {
		return e
	}

	if e._err.Data == nil {
		e._err.Data = make(map[string]interface{})
	}
	e._err.Data[k] = v
	return e
}

// Get retrieves the value associated with the given key in the context
// of this error.  It answers `nil` in case the key could not be found.
func (e *Error) Get(k string) interface{} {
	v := e._err.Data[k]
	if v != nil {
		return v
	}
	if e._err.Inner != nil {
		return e._err.Inner.Get(k)
	}
	return nil
}

// Location answers the source "file:line" of the first addition of data
// to this error object.
func (e *Error) Location() string {
	return e._err.Location
}

// Type answers the user-defined type of this error object.
func (e *Error) Type() string {
	return e._err.Type
}

// Data answers the user-defined associated data captured in this error
// object.
func (e *Error) Data() json.RawMessage {
	buf, err := json.Marshal(e._err.Data)
	if err != nil {
		return nil
	}
	return buf
}

// Inner answers the wrapped underlying error of this error object.
func (e *Error) Inner() *Error {
	return e._err.Inner
}

// MarshalJSON marshals this error object recursively, unwrapping any
// inner error at each level.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e._err)
}

// Error answers a JSON-formatted representation of this error object.
func (e *Error) Error() string {
	buf, err := json.Marshal(e._err)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}
