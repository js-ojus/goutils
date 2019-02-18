package goutils

import (
	"errors"
	"io"
	"log"
	"os"
)

// `goutils` provides an assorted collection of utility types that
// provide various convenience methods.

// init initialises the default logger.
func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(os.Stderr)
}

// SetLogOutput can be used to change the logger's output to any
// `io.Writer`.
func SetLogOutput(w io.Writer) error {
	if w == nil {
		return errors.New("given writer is nil")
	}

	log.SetOutput(w)
	return nil
}

// KV is a simple map from strings to arbitrary values.
type KV map[string]interface{}
