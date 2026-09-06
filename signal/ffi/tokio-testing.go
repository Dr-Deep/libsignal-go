package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++

extern void promise_complete_owned_buffer_cgo(
    SignalFfiError *err,
    const SignalOwnedBuffer *buf,
    const void *userData
);
*/
import "C"
import (
	"context"
	"unsafe"
)

/*
SignalFfiError* signal_testing_tokio_async_context_future_success_bytes(
SignalFfiError* signal_testing_tokio_async_context_new_single_threaded(
SignalFfiError* signal_testing_tokio_async_future(
*/

func TokioTesting_Future() {}

func TokioTesting_FutureSuccessBytes(ctx context.Context, tokio *TokioAsyncContext, count int32) ([]byte, error) {
	var (
		c_promise          C.SignalCPromiseOwnedBuffer
		promise, h_promise = NewPromise[[]byte](tokio)
	)
	c_promise.complete = (C.SignalType_FunctionPointer_void_SignalType_MutPointer_SignalFfiError_SignalType_ConstPointer_SignalOwnedBuffer_SignalType_ConstPointer_void)(
		unsafe.Pointer(C.promise_complete_owned_buffer_cgo),
	)
	c_promise.context = h_promise

	err := convertError(
		C.signal_testing_tokio_async_context_future_success_bytes(
			&c_promise,
			tokio.ConstPointer(),
			C.int32_t(count),
		),
	)

	if err != nil {
		promise.handle.Delete()
		return nil, err
	}

	promise.SetCancellationID(
		uint64(c_promise.cancellation_id),
	)

	return promise.Await(ctx)
}

func TokioTesting_NewSingleThreaded() {}
