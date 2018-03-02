package rfc2822

import (
	"fmt"
	"testing"
)

func TestEncodeRFC2047(t *testing.T) {
	s := EncodeRFC2047(UTF8, []byte("テスト"))
	expected := "=?UTF-8?B?44OG44K544OI?="

	if s != expected {
		t.Errorf("Failed encode: expected %s, got %s", expected, s)
	}
}
