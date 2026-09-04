package ffi

/*
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
*/
import "C"

/*
   * Wrapper
args;funcs
1   140
2	288
3	75
4	57
5	34
6	22
7	20
8	8
9	3
10	2
11	0
12	2
13	1
*/

// statt out: promise, p, buffer, new_obj

/*
type FFI0[R any] func() R
type FFI1[A, R any] func(A) R
type FFI2[A, B, R any] func(A, B) R
type FFI3[A, B, C, R any] func(A, B, C) R
type FFI4[A, B, C, D, R any] func(A, B, C, D) R
*/

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

type FfiFunc8Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4) *C.SignalFfiError
type FfiFunc9Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4) *C.SignalFfiError
type FfiFunc10Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4) *C.SignalFfiError
type FfiFunc11Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4) *C.SignalFfiError
type FfiFunc12Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4) *C.SignalFfiError
type FfiFunc13Call[Out, Arg0, Arg1, Arg2, Arg3, Arg4 any] func(*Out, Arg0, Arg1, Arg2, Arg3, Arg4) *C.SignalFfiError

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

func FfiCallFun3[Out, Arg0, Arg1, Arg2 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	arg2 Arg2,
	f FfiFunc3Call[Out, Arg0, Arg1, Arg2],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1, arg2))
}

func FfiCallFun4[Out, Arg0, Arg1, Arg2, Arg3 any](
	out *Out,
	arg0 Arg0,
	arg1 Arg1,
	arg2 Arg2,
	arg3 Arg3,
	f FfiFunc4Call[Out, Arg0, Arg1, Arg2, Arg3],
) (*Out, error) {
	return out, convertError(f(out, arg0, arg1, arg2, arg3))
}

/*
	FfiCallFunc1(
		_, C.GoString("dsd"), signal_account_entropy_pool_derive_backup_key,
	)
*/
