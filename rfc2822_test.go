package rfc2822

import (
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	s := Encode(UTF8, []byte("テスト"))
	expected := "=?UTF-8?B?44OG44K544OI?="

	if s != expected {
		t.Errorf("Failed encode: expected %s, got %s", expected, s)
	}
}

func TestFormatDate(t *testing.T) {
	s := FormatDate(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	expected := "Tue, 10 Nov 2009 23:00:00 +0000"

	if s != expected {
		t.Errorf("Failed format date: expected %s, got %s", expected, s)
	}
}
