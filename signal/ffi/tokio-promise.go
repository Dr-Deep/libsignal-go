package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++

extern void goPromiseCompleteOwnedBuffer(SignalFfiError *err, const SignalOwnedBuffer *buf, const void *userData);

static void promise_complete_owned_buffer_cgo(
    SignalFfiError *err,
    const SignalOwnedBuffer *buf,
    const void *userData
) {
    goPromiseCompleteOwnedBuffer(err, buf, userData);
}
*/
import "C"
import (
	"context"
	"runtime/cgo"
	"unsafe"
)

type result[T any] struct {
	value T
	err   error
}

type Promise[T any] struct {
	channel   chan result[T]
	cancelID  uint64
	tokioCtx  *TokioAsyncContext
	handle    cgo.Handle // keeps alive across C callbacks
	completed bool
}

// returns C context Pointer (SignalCPromise.context)
func newPromise[T any](tokio *TokioAsyncContext) (*Promise[T], unsafe.Pointer) {
	p := &Promise[T]{
		channel:  make(chan result[T], 1),
		tokioCtx: tokio,
	}
	p.handle = cgo.NewHandle(p)
	return p, unsafe.Pointer(p.handle)
}

// called once from C
func (p *Promise[T]) complete(value T, err error) {
	if p.completed {
		return
	}
	p.completed = true
	p.channel <- result[T]{value: value, err: err}
	close(p.channel)
	p.handle.Delete()
}

// block until promise resolves or cancel
func (p *Promise[T]) Await(ctx context.Context) (T, error) {
	select {
	case r := <-p.channel:
		return r.value, r.err

	case <-ctx.Done():
		if p.tokioCtx != nil && p.cancelID != 0 {
			_ = p.tokioCtx.Cancel(p.cancelID)
		}

		var zero T
		return zero, ctx.Err()
	}
}

// SignalCPromise (the cancellation_id is written by libsignal).
func (p *Promise[T]) SetCancellationID(id uint64) {
	p.cancelID = id
}
