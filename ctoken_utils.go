package goutils

import (
	"encoding/base64"
)

// EncodeCtoken accepts a string continuation token, base64-encodes it,
// and answers the final string.
func EncodeCtoken(ctoken string) string {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(ctoken)))
	base64.StdEncoding.Encode(buf, []byte(ctoken))
	return string(buf)
}

// DecodeCtoken converts the encoded continuation token back into its
// original form.
func DecodeCtoken(ctoken string) (string, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(ctoken)))
	_, err := base64.StdEncoding.Decode(buf, []byte(ctoken))
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
