package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"
import (
	"unsafe"
)

/*
// C data with explicit length to Go []byte
func C.GoBytes(unsafe.Pointer, C.int) []byte
*/

//  runtime.Pinner.
// runtime/cgo.Handle

func StringToCString(s string) *C.char {
	return C.CString(s)
}

func CStringToString(s *C.char) string {
	if s == nil {
		return ""
	}

	return C.GoString(s)
}

func SignalCStringPtrToString(s C.SignalCStringPtr) string {
	if s == nil {
		return ""
	}

	return CStringToString(
		(*C.char)(unsafe.Pointer(s)),
	)
}
