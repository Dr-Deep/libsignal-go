package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

type TokioAsyncContext struct {
	ptr *TokioAsyncContext
}

/*
SignalFfiError* signal_testing_tokio_async_context_future_success_bytes(
SignalFfiError* signal_testing_tokio_async_context_new_single_threaded(
SignalFfiError* signal_testing_tokio_async_future(
SignalFfiError* signal_tokio_async_context_cancel(
SignalFfiError* signal_tokio_async_context_destroy(
SignalFfiError* signal_tokio_async_context_new(
*/

// C.signal_tokio_async_context_new
func (ctx *TokioAsyncContext) NewContext()

// C.signal_tokio_async_context_destroy()
func (ctx *TokioAsyncContext) DestroyContext()

// C.signal_tokio_async_context_new()
func (ctx *TokioAsyncContext) CancelContext()
