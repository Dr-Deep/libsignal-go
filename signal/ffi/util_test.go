package ffi

import "testing"

func TestCString(t *testing.T) {
	var (
		str   = "teststring"
		c_str = StringToCString(str)
	)

	if str != CStringToString(c_str) {
		t.Fatalf("'%s' != '%s'", str, CStringToString(c_str))
	}
}
