package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"
import (
	"unsafe"
)

/*
SignalFfiError* signal_connection_proxy_config_clone(
SignalFfiError* signal_connection_proxy_config_destroy(
SignalFfiError* signal_connection_proxy_config_new(
*/

type ConnectionProxyConfig struct {
	ptr *C.SignalConnectionProxyConfig
}

func (proxy *ConnectionProxyConfig) MutPointer() C.SignalMutPointerConnectionProxyConfig {
	return C.SignalMutPointerConnectionProxyConfig{raw: proxy.ptr}
}

func (proxy *ConnectionProxyConfig) ConstPointer() C.SignalConstPointerConnectionProxyConfig {
	return C.SignalConstPointerConnectionProxyConfig{raw: proxy.ptr}
}

// C.signal_connection_proxy_config_new()
/*
strings may be empty
port is the proxy port.
*/
func NewConnectionProxyConfig(scheme, host, username, password string, port int32) (*ConnectionProxyConfig, error) {
	var (
		out      C.SignalMutPointerConnectionProxyConfig
		c_scheme = StringToCString(scheme)
		c_host   = StringToCString(host)
		c_user   = StringToCString(username)
		c_pass   = StringToCString(password)
		c_port   = C.int32_t(port)
	)
	defer C.free(unsafe.Pointer(c_scheme))
	defer C.free(unsafe.Pointer(c_host))
	defer C.free(unsafe.Pointer(c_user))
	defer C.free(unsafe.Pointer(c_pass))

	err := convertError(
		C.signal_connection_proxy_config_new(
			&out,
			(*C.int8_t)(unsafe.Pointer(c_scheme)),
			(*C.int8_t)(unsafe.Pointer(c_host)),
			c_port,
			(*C.int8_t)(unsafe.Pointer(c_user)),
			(*C.int8_t)(unsafe.Pointer(c_pass)),
		),
	)

	if err != nil {
		return nil, err
	}

	return &ConnectionProxyConfig{ptr: out.raw}, nil
}

// C.signal_connection_proxy_config_clone()
func (proxy *ConnectionProxyConfig) Clone() (*ConnectionProxyConfig, error) {
	if proxy.ptr == nil {
		return nil, nil
	}

	var out C.SignalMutPointerConnectionProxyConfig
	err := convertError(
		C.signal_connection_proxy_config_clone(
			&out,
			proxy.ConstPointer(),
		),
	)
	if err != nil {
		return nil, err
	}

	return &ConnectionProxyConfig{ptr: out.raw}, nil
}

// C.signal_connection_proxy_config_destroy()
func (proxy *ConnectionProxyConfig) Destroy() {
	if proxy.ptr != nil {
		C.signal_connection_proxy_config_destroy(
			proxy.MutPointer(),
		)
		proxy.ptr = nil
	}
}
