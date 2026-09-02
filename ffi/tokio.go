package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

type TokioAsyncContext struct{}

// C.signal_tokio_async_context_new
func (ctx *TokioAsyncContext) NewContext()

// C.signal_tokio_async_context_destroy()
func (ctx *TokioAsyncContext) DestroyContext()

// C.signal_tokio_async_context_new()
func (ctx *TokioAsyncContext) CancelContext()
