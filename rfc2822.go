package rfc2822

import (
	"encoding/base64"
	"time"
)

const (
	CRLF = "\r\n"

	ISO2022 = "ISO-2022-JP"
	UTF8    = "UTF-8"

	DateFormat = "Mon, 02 Jan 2006 15:04:05 -0700"
)

type Message struct {
	Headers map[string]string
	Body    string
}

func (m *Message) Create() string {
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

// FormatDate formats time.Time in RFC 2822 date format,
// "Mon, 02 Jan 2006 15:04:05 -0700".
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}
