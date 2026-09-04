package ffi

/*
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
*/
import "C"

/*
FfiCallFunc1(

	_, C.GoString("dsd"), signal_account_entropy_pool_derive_backup_key,

)
*/
// https://pkg.go.dev/runtime/cgo@go1.17#Handle

// statt out: promise, p, buffer, new_obj

/*
 * signal_func(*out, *arg1) *SignalFfiError => func(arg1) (out, error)
 */

type FfiFunc0Call[Out any] func(*Out) *C.SignalFfiError
type FfiFunc1Call[Out, Arg0 any] func(*Out, Arg0) *C.SignalFfiError
type FfiFunc2Call[Out, Arg0, Arg1 any] func(*Out, Arg0, Arg1) *C.SignalFfiError
type FfiFunc3Call[Out, Arg0, Arg1, Arg2 any] func(*Out, Arg0, Arg1, Arg2) *C.SignalFfiError
type FfiFunc4Call[Out, Arg0, Arg1, Arg2, Arg3 any] func(*Out, Arg0, Arg1, Arg2, Arg3) *C.SignalFfiError
type FfiFunc5Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4) *C.SignalFfiError
type FfiFunc6Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5) *C.SignalFfiError
type FfiFunc7Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6) *C.SignalFfiError
type FfiFunc8Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7) *C.SignalFfiError
type FfiFunc9Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8) *C.SignalFfiError
type FfiFunc10Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9) *C.SignalFfiError
type FfiFunc11Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9, Arg10 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9, Arg10) *C.SignalFfiError
type FfiFunc12Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9, Arg10, Arg11 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9, Arg10, Arg11) *C.SignalFfiError
type FfiFunc13Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9, Arg10, Arg11, Arg12 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6, Arg7, Arg8, Arg9, Arg10, Arg11, Arg12) *C.SignalFfiError

/*
 * Wrapper
 */

func FfiCallFunc0[Out any](
	out *Out,
	f FfiFunc0Call[Out],
) (*Out, error) {
	return out, convertError(f(out))
}

func FfiCallFunc1[Out, Arg0 any](
	out *Out,
	arg0 Arg0,
	f FfiFunc1Call[Out, Arg0],
) (*Out, error) {
	return out, convertError(f(out, arg0))
}

func FfiCallFunc2[Out, Arg0, Arg1 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	f FfiFunc2Call[Out, Arg0, Arg1],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1))
}

func FfiCallFunc3[Out, Arg0, Arg1, Arg2 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	arg2 Arg2,
	f FfiFunc3Call[Out, Arg0, Arg1, Arg2],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1, arg2))
}

func FfiCallFunc4[Out, Arg0, Arg1, Arg2, Arg3 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	arg2 Arg2,
	arg3 Arg3,
	f FfiFunc4Call[Out, Arg0, Arg1, Arg2, Arg3],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1, arg2, arg3))
}

func FfiCallFunc5[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	arg2 Arg2,
	arg3 Arg3,
	arg4 Arg4,
	f FfiFunc5Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1, arg2, arg3, arg4))
}

func FfiCallFunc6[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	arg2 Arg2,
	arg3 Arg3,
	arg4 Arg4,
	arg5 Arg5,
	f FfiFunc6Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1, arg2, arg3, arg4, arg5))
}

func FfiCallFunc7[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	arg2 Arg2,
	arg3 Arg3,
	arg4 Arg4,
	arg5 Arg5,
	arg6 Arg6,
	f FfiFunc7Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4, Arg5, Arg6],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1, arg2, arg3, arg4, arg5, arg6))
}
