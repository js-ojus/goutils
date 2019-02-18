package goutils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

// Report holds the specifics of an error that has to be reported back
// to the user.  Since the ouput is potentially visible to an end-user,
// caution should be exercised in forming this object.
type Report struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    KV     `json:"data,omitempty"`
}

// String answers the JSON-formatted string of this report object.
func (r Report) String() string {
	buf, _ := json.Marshal(r)
	return string(buf)
}

// RequestEnvelope holds the application layer method and an opaque
// request-specific JSON body.
//
// Examples:
//     {"method": "GET", "body": {"id": 1234}}
//     {"method": "POST", "body": {...}}
//     {"method": "DELETE", "body": {"id": 1234}}
//
// All client requests MUST follow this envelope structure when making
// API requests that do not involve blobs.
//
// Similarly, responses sent by the API server hold the application
// layer response status and one of the following types of JSON bodies:
//     - ("OK", informational message or acknowledgement),
//     - ("Error", error code, error message, a map of important parameters), and
//     - ("OK", an opaque response-specific JSON body).
//
// All clients MUST check the top-level status before parsing the
// response body.
type RequestEnvelope struct {
	Method string          `json:"method"`
	Body   json.RawMessage `json:"body"`
}

// OpenEnvelope opens the envelope by reading the full body of the
// request, and then reading it into an instance of `RequestEnvelope`.
func OpenEnvelope(r io.Reader) (*RequestEnvelope, *Report) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, &Report{Code: 1001, Data: KV{"ioErr": err.Error()}}
	}
	var env RequestEnvelope
	err = json.Unmarshal(buf, &env)
	if err != nil {
		return nil, &Report{Code: 1002, Data: KV{"ioErr": err.Error()}}
	}

	return &env, nil
}

// SendSuccess prepares and writes an informational JSON response based
// on the given message.
func SendSuccess(w io.Writer, msg string) {
	r := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{"OK", msg}
	buf, err := json.Marshal(r)
	if err != nil {
		log.Println(NewError("json").Add("msg", err.Error()))
		io.WriteString(w, Report{Code: 102, Message: Errors[102]}.String())
		return
	}

	w.Write(buf)
}

// SendError prepares and writes an error JSON response based on the
// given values.
func SendError(w io.Writer, rep Report) {
	if rep.Message == "" {
		if rep.Code > 0 {
			rep.Message = Errors[rep.Code]
		}
	}
	r := struct {
		Status string `json:"status"`
		Report
	}{"Error", rep}
	buf, err := json.Marshal(r)
	if err != nil {
		log.Println(NewError("json").Add("msg", err.Error()).Add("reportCode", rep.Code))
		io.WriteString(w, Report{Code: 102, Message: Errors[102]}.String())
		return
	}

	w.Write(buf)
}

// SendResult prepares and writes a result-carrying response, using the
// given `json.RawMessage`.
func SendResult(w io.Writer, body interface{}) {
	r := struct {
		Status string      `json:"status"`
		Body   interface{} `json:"body"`
	}{"OK", body}
	buf, err := json.Marshal(r)
	if err != nil {
		log.Println(NewError("json").Add("msg", err.Error()))
		io.WriteString(w, Report{Code: 102, Message: Errors[102]}.String())
		return
	}

	w.Write(buf)
}
