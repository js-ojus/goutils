package goutils

import (
	"encoding/json"
)

// Log provides a structure for log entries.  When written, these are
// converted to JSON.
type Log map[string]interface{}

// String converts the map of fields to JSON.
func (l Log) String() string {
	buf, _ := json.Marshal(l)
	return string(buf)
}

// MessageObj is a holder of a status and a message.  This is defined to
// ease generation of JSON response messages.
type MessageObj struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// JSONForMessage answers a JSON-encoded object with the given status
// and message.
func JSONForMessage(ok bool, s string) string {
	obj := MessageObj{Status: "OK", Message: s}
	if !ok {
		obj.Status = "Error"
	}
	buf, _ := json.Marshal(obj)
	return string(buf)
}

// JSONForError answers a JSON-encoded message object with the given
// error.
func JSONForError(err error) string {
	return JSONForMessage(false, err.Error())
}

// JSONForMap answers a JSON-encoded message object with keys and values
// from the given map.
func JSONForMap(ok bool, obj map[string]interface{}) string {
	if ok {
		obj["status"] = "OK"
	} else {
		obj["status"] = "Error"
	}
	buf, err := json.Marshal(obj)
	if err != nil {
		return JSONForError(err)
	}

	return string(buf)
}

// JSONForKV answers a JSON-encoded message object with the given
// interface as the value of the specified key.
func JSONForKV(ok bool, key string, value interface{}) string {
	obj := map[string]interface{}{key: value}
	return JSONForMap(ok, obj)
}
