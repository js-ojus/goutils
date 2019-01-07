package goutils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Errors is the global map of errors.  It serves as the central place
// to define error codes and their textual messages.
var Errors map[int]string

//

func init() {
	Errors = map[int]string{
		// Internal system errors.
		101: "Internal database error.",
		102: "Internal system error.",

		// API / transport errors.
		1001: "Corrupt HTTP request body.",
		1002: "Corrupt request envelope.",
		1003: "Request body contains invalid JSON.",
		1004: "Unhandled request method.",

		// Generic application logic errors.
		1101: "Missing mandatory arguments.",
		1102: "Uniqueness constraint violation.",
		1103: "Invalid continuation token format.",
		1104: "Invalid continuation token.",
		1105: "Referential integrity violation.",
		1106: "Empty result set; record could not be found.",
    }
}

// KV is a simple map from strings to arbitrary values.
type KV map[string]interface{}

// Report holds the specifics of an error that has to be reported back
// to the user.  Since the ouput is potentially visible to an end-user,
// caution should be exercised in forming this object.
type Report struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    KV     `json:"data,omitempty"`
}

// String answers the JSON-formatted string of this report object.
func (r Report) String() string {
	buf, _ := json.Marshal(r)
	return string(buf)
}

// Log specifies the level of the message: "INFO", "WARN", "ERRO" or
// "FATA".  It also holds specific information that should be logged to
// aid the developer.
type Log struct {
	Level string `json:"level"`
	Data  KV     `json:"data,omitempty"`
}

// String answers the JSON-formatted string of this log object.
func (l Log) String() string {
	buf, _ := json.Marshal(l)
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

// OpenEnvelope opens the envelope by reading the full body of the HTTP
// request, and then reading it into an instance of `RequestEnvelope`.
func OpenEnvelope(r *http.Request) (*RequestEnvelope, *Report) {
	buf, err := ioutil.ReadAll(r.Body)
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
		log.Print(Log{Level: "ERRO", Data: KV{"jsonErr": err.Error(), "msg": msg}})
		io.WriteString(w, Report{Code: 102, Message: Errors[102]}.String())
		return
	}

	w.Write(buf)
}

// SendError prepares and writes an error JSON response based on the
// given values.
func SendError(w io.Writer, eobj *Report) {
	if eobj.Message == "" {
		if eobj.Code > 0 {
			eobj.Message = Errors[eobj.Code]
		}
	}
	r := struct {
		Status string `json:"status"`
		*Report
	}{"Error", eobj}
	buf, err := json.Marshal(r)
	if err != nil {
		log.Print(Log{Level: "ERRO", Data: KV{"jsonErr": err.Error()}})
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
		log.Print(Log{Level: "ERRO", Data: KV{"jsonErr": err.Error()}})
		io.WriteString(w, Report{Code: 102, Message: Errors[102]}.String())
		return
	}

	w.Write(buf)
}
