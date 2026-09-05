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
		//c_remoteConfig = C.SignalMutPointerBridgedStringMap{raw: remoteConfig.ptr}
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

func (ConnMgr *ConnectionManager) SetRemoteConfig()
func (ConnMgr *ConnectionManager) SetProxy()
func (ConnMgr *ConnectionManager) SetInvalidProxy()
func (ConnMgr *ConnectionManager) SetCensorshipCircumventionEnabled()
func (ConnMgr *ConnectionManager) OnNetworkChange()

func (ConnMgr *ConnectionManager) Destroy() {
	if ConnMgr.ptr != nil {
		C.signal_connection_manager_destroy(
			C.SignalMutPointerConnectionManager{raw: ConnMgr.ptr},
		)
		ConnMgr.ptr = nil
	}
}

func (ConnMgr *ConnectionManager) ClearProxy()
