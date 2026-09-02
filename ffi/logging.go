package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

type SignalLogLevel int

const (
	SignalLogLevelError SignalLogLevel = 1
	SignalLogLevelWarn  SignalLogLevel = 2
	SignalLogLevelInfo  SignalLogLevel = 3
	SignalLogLevelDebug SignalLogLevel = 4
	SignalLogLevelTrace SignalLogLevel = 5
)

// C.signal_init_logger()
func SignalInitLogger() bool
