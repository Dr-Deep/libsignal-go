package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_tokio_async_context_cancel(
SignalFfiError* signal_tokio_async_context_destroy(
SignalFfiError* signal_tokio_async_context_new(
*/

type TokioAsyncContext struct {
	ptr *C.SignalTokioAsyncContext
}

func (ctx *TokioAsyncContext) MutPointer() C.SignalMutPointerTokioAsyncContext {
	return C.SignalMutPointerTokioAsyncContext{raw: ctx.ptr}
}

func (ctx *TokioAsyncContext) ConstPointer() C.SignalConstPointerTokioAsyncContext {
	return C.SignalConstPointerTokioAsyncContext{raw: ctx.ptr}
}

// C.signal_tokio_async_context_new()
func NewTokioAsyncContext() (*TokioAsyncContext, error) {
	var (
		out C.SignalMutPointerTokioAsyncContext
	)

	err := convertError(
		C.signal_tokio_async_context_new(
			&out,
		),
	)
	if err != nil {
		return nil, err
	}

	return &TokioAsyncContext{ptr: out.raw}, nil
}

// C.signal_tokio_async_context_cancel()
func (ctx *TokioAsyncContext) Cancel(cancellationID uint64) error {
	if ctx.ptr == nil {
		return nil
	}

	err := convertError(
		C.signal_tokio_async_context_cancel(
			ctx.ConstPointer(),
			C.uint64_t(cancellationID),
		),
	)

	return err
}

// C.signal_tokio_async_context_destroy()
func (ctx *TokioAsyncContext) Destroy() {
	if ctx.ptr != nil {
		C.signal_tokio_async_context_destroy(
			ctx.MutPointer(),
		)
		ctx.ptr = nil
	}
}
