package rfc2822

import (
	"encoding/base64"
)

const (
	CRLF = "\r\n"

	ISO2022 = "ISO-2022-JP"
	UTF8    = "UTF-8"
)

type Message struct {
	Headers map[string]string
	Body    string
}

func (m *Message) Parse() string {
	message := ""
	for k, v := range m.Headers {
		message += k + ": " + v + CRLF
	}
	message += CRLF + m.Body
	return message
}

// EncodeRFC2047 encode not-ASCII header,
// following RFC 2047 (base64).
func EncodeRFC2047(charset string, b []byte) string {
	s := base64.StdEncoding.EncodeToString(b)
	return "=?" + charset + "?B?" + s + "?="
}
