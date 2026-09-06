package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

//export goPromiseCompleteOwnedBuffer
func goPromiseCompleteOwnedBuffer(err *C.SignalFfiError, buf *C.SignalOwnedBuffer, userData unsafe.Pointer) {
	h := cgo.Handle(uintptr(userData))
	promise, ok := h.Value().(*Promise[[]byte])

	// invalid handle
	if !ok {
		if err != nil {
			C.signal_error_free(err)
		}

		if buf != nil && buf.base != nil {
			C.signal_free_buffer(buf.base, buf.length)
		}

		return
	}

	if err != nil {
		promise.complete(nil, convertError(err))
		return
	}

	var out []byte
	if buf != nil && buf.base != nil && buf.length > 0 {
		out = C.GoBytes(
			unsafe.Pointer(buf.base), C.int(buf.length),
		)

		C.signal_free_buffer(buf.base, buf.length)
	}

	promise.complete(out, nil)
}
