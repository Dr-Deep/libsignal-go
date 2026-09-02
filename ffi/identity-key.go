package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_identitykey_verify_alternate_identity(
SignalFfiError* signal_identitykeypair_deserialize(
SignalFfiError* signal_identitykeypair_serialize(
SignalFfiError* signal_identitykeypair_sign_alternate_identity(
*/
