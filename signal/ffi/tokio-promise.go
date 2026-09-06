package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++

extern void goPromiseCompleteOwnedBuffer(SignalFfiError *err, const SignalOwnedBuffer *buf, const void *userData);

void promise_complete_owned_buffer_cgo(          // ← no "static"
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

type PromiseResult[T any] struct {
	value T
	err   error
}

type Promise[T any] struct {
	channel   chan PromiseResult[T]
	cancelID  uint64
	tokioCtx  *TokioAsyncContext
	handle    cgo.Handle // keep alive across C callbacks
	completed bool
}

// returns promise, C context Pointer (SignalCPromise.context)
func NewPromise[T any](tokio *TokioAsyncContext) (*Promise[T], unsafe.Pointer) {
	p := &Promise[T]{
		channel:  make(chan PromiseResult[T], 1),
		tokioCtx: tokio,
	}

	p.handle = cgo.NewHandle(p)
	return p, unsafe.Pointer(uintptr(p.handle)) //?
}

// called once from C
func (p *Promise[T]) complete(value T, err error) {
	if p.completed {
		return
	}

	p.completed = true
	p.channel <- PromiseResult[T]{value: value, err: err}
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
