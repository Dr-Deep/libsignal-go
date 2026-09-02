package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_fingerprint_clone(
SignalFfiError* signal_fingerprint_compare(
SignalFfiError* signal_fingerprint_destroy(
SignalFfiError* signal_fingerprint_display_string(
SignalFfiError* signal_fingerprint_new(
SignalFfiError* signal_fingerprint_scannable_encoding(
*/
