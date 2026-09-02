package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_aes256_ctr32_destroy(
SignalFfiError* signal_aes256_ctr32_new(
SignalFfiError* signal_aes256_ctr32_process(
*/
