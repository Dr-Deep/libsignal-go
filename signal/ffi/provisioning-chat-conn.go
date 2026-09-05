package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_provisioning_chat_connection_connect(
SignalFfiError* signal_provisioning_chat_connection_destroy(
SignalFfiError* signal_provisioning_chat_connection_disconnect(
SignalFfiError* signal_provisioning_chat_connection_info(
SignalFfiError* signal_provisioning_chat_connection_init_listener(
*/
