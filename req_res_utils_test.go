package goutils

import (
	"bytes"
	"testing"
)

func TestSendSuccess001(t *testing.T) {
	buf := new(bytes.Buffer)
	out := `{"status":"OK","message":"Hello!"}`

	SendSuccess(buf, "Hello!")
	if string(buf.Bytes()) != out {
		t.Errorf("expected: '%s', got: '%s'\n", out, buf)
	}
}

func TestSendError001(t *testing.T) {
	buf := new(bytes.Buffer)
	out := `{"status":"Error","code":1001,"message":"Test error 1001"}`

	r := Report{Code: 1001, Message: "Test error 1001"}
	SendError(buf, r)
	if string(buf.Bytes()) != out {
		t.Errorf("expected: '%s', got: '%s'\n", out, buf)
	}
}

func TestSendError002(t *testing.T) {
	buf := new(bytes.Buffer)
	out := `{"status":"Error","code":1001,"message":"Test error 1001","data":{"mode":"test"}}`

	r := Report{Code: 1001, Message: "Test error 1001", Data: map[string]interface{}{"mode": "test"}}
	SendError(buf, r)
	if string(buf.Bytes()) != out {
		t.Errorf("expected: '%s', got: '%s'\n", out, buf)
	}
}

func TestSendResult001(t *testing.T) {
	buf := new(bytes.Buffer)
	out := `{"status":"OK","body":{"mode":"test","participants":100}}`

	obj := struct {
		Mode         string `json:"mode"`
		Participants int    `json:"participants"`
	}{"test", 100}
	SendResult(buf, obj)
	if string(buf.Bytes()) != out {
		t.Errorf("expected: '%s', got: '%s'\n", out, buf)
	}
}
