package goutils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// RequestEnvelope holds the application layer method and an opaque body.
type RequestEnvelope struct {
	Method string          `json:"method"`
	Body   json.RawMessage `json:"body"`
}

// OpenEnvelope opens the envelope by reading the full body of the HTTP
// request, and then reading it into an instance of `RequestEnvelope`.
func OpenEnvelope(r *http.Request) (*RequestEnvelope, error) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("corrupt request body : " + err.Error())
	}
	var env RequestEnvelope
	err = json.Unmarshal(buf, &env)
	if err != nil {
		return nil, errors.New("request body contains invalid JSON : " + err.Error())
	}

	return &env, nil
}
