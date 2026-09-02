package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_username_candidates_from(
SignalFfiError* signal_username_hash(
SignalFfiError* signal_username_hash_from_parts(
SignalFfiError* signal_username_link_create(
SignalFfiError* signal_username_link_decrypt_username(
SignalFfiError* signal_username_proof(
SignalFfiError* signal_username_verify(
*/
