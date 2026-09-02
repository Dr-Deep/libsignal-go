package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_hsm_enclave_client_complete_handshake(
SignalFfiError* signal_hsm_enclave_client_destroy(
SignalFfiError* signal_hsm_enclave_client_established_recv(
SignalFfiError* signal_hsm_enclave_client_established_send(
SignalFfiError* signal_hsm_enclave_client_initial_request(
SignalFfiError* signal_hsm_enclave_client_new(
*/
