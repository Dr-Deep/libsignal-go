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
SignalFfiError* signal_connection_manager_clear_proxy(
SignalFfiError* signal_connection_manager_destroy(
SignalFfiError* signal_connection_manager_new(
SignalFfiError* signal_connection_manager_on_network_change(
SignalFfiError* signal_connection_manager_set_censorship_circumvention_enabled(
SignalFfiError* signal_connection_manager_set_invalid_proxy(
SignalFfiError* signal_connection_manager_set_proxy(
SignalFfiError* signal_connection_manager_set_remote_config(
*/

type ConnectionManager struct {
	ptr *C.SignalConnectionManager
}

func (connMgr *ConnectionManager) MutPointer() C.SignalMutPointerConnectionManager {
	return C.SignalMutPointerConnectionManager{raw: connMgr.ptr}
}

func (connMgr *ConnectionManager) ConstPointer() C.SignalConstPointerConnectionManager {
	return C.SignalConstPointerConnectionManager{raw: connMgr.ptr}
}

/*
environment: 0 = production, 1 = staging (check SignalNet.Environment)
build_variant: usually 0 for production builds
remoteConfig can be nil
*/
func NewConnectionManager(environment uint8, userAgent string, remoteConfig any, buildVariant uint8) (*ConnectionManager, error) {
	var (
		out            C.SignalMutPointerConnectionManager
		c_environment  = C.uint8_t(environment)
		c_userAgent    = StringToCString(userAgent)
		c_remoteConfig C.SignalMutPointerBridgedStringMap
		c_buildVariant = C.uint8_t(buildVariant)
	)
	defer C.free(unsafe.Pointer(c_userAgent))

	if remoteConfig != nil {
		//c_remoteConfig = remoteConfig.MutPointer()
	}

	err := convertError(
		C.signal_connection_manager_new(
			&out,
			c_environment,
			(*C.int8_t)(unsafe.Pointer(c_userAgent)),
			c_remoteConfig,
			c_buildVariant,
		),
	)
	if err != nil {
		return nil, err
	}

	return &ConnectionManager{ptr: out.raw}, nil

}

// C.signal_connection_manager_set_remote_config()
func (connMgr *ConnectionManager) SetRemoteConfig(remoteCfg *RemoteConfigMap, buildVariant uint8) error {
	if connMgr.ptr == nil {
		return nil
	}

	var (
		c_remoteCfg    C.SignalMutPointerBridgedStringMap
		c_buildVariant = C.uint8_t(buildVariant)
	)
	if remoteCfg != nil {
		c_remoteCfg = remoteCfg.MutPointer()
	}

	err := convertError(
		C.signal_connection_manager_set_remote_config(
			connMgr.ConstPointer(),
			c_remoteCfg,
			c_buildVariant,
		),
	)

	return err
}

// C.signal_connection_manager_set_proxy()
func (connMgr *ConnectionManager) SetProxy(proxyCfg *ConnectionProxyConfig) error {
	if connMgr.ptr == nil {
		return nil
	}

	if proxyCfg == nil {
		return convertError(
			C.signal_connection_manager_set_invalid_proxy(
				connMgr.ConstPointer(),
			),
		)
	}

	err := convertError(
		C.signal_connection_manager_set_proxy(
			connMgr.ConstPointer(),
			proxyCfg.ConstPointer(),
		),
	)

	return err
}

// C.signal_connection_manager_set_invalid_proxy()
func (connMgr *ConnectionManager) SetInvalidProxy() error {
	if connMgr.ptr == nil {
		return nil
	}

	err := convertError(
		C.signal_connection_manager_set_invalid_proxy(
			connMgr.ConstPointer(),
		),
	)

	return err
}

// C.signal_connection_manager_set_censorship_circumvention_enabled()
func (connMgr *ConnectionManager) SetCensorshipCircumventionEnabled(value bool) error {
	if connMgr.ptr == nil {
		return nil
	}

	err := convertError(
		C.signal_connection_manager_set_censorship_circumvention_enabled(
			connMgr.ConstPointer(),
			C.bool(value),
		),
	)

	return err
}

// C.signal_connection_manager_on_network_change()
func (connMgr *ConnectionManager) OnNetworkChange() error {
	if connMgr.ptr == nil {
		return nil
	}

	err := convertError(
		C.signal_connection_manager_on_network_change(
			connMgr.ConstPointer(),
		),
	)

	return err
}

// C.signal_connection_manager_destroy()
func (connMgr *ConnectionManager) Destroy() {
	if connMgr.ptr != nil {
		C.signal_connection_manager_destroy(
			connMgr.MutPointer(),
		)
		connMgr.ptr = nil
	}
}

// C.signal_connection_manager_clear_proxy()
func (connMgr *ConnectionManager) ClearProxy() error {
	if connMgr.ptr == nil {
		return nil
	}

	err := convertError(
		C.signal_connection_manager_clear_proxy(
			connMgr.ConstPointer(),
		),
	)

	return err
}
