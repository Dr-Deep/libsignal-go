package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"
import "unsafe"

/*
 C.SignalMutPointerBridgedStringMap

SignalFfiError* signal_bridged_string_map_clone(
SignalFfiError* signal_bridged_string_map_destroy(
SignalFfiError* signal_bridged_string_map_insert(
SignalFfiError* signal_bridged_string_map_new(
*/

type RemoteConfigMap struct {
	ptr *C.SignalBridgedStringMap
}

func (m *RemoteConfigMap) MutPointer() C.SignalMutPointerBridgedStringMap {
	return C.SignalMutPointerBridgedStringMap{raw: m.ptr}
}

func (m *RemoteConfigMap) ConstPointer() C.SignalConstPointerBridgedStringMap {
	return C.SignalConstPointerBridgedStringMap{raw: m.ptr}
}

// C.signal_bridged_string_map_new()
func NewRemoteConfigMap(initialCapacity uint32) (*RemoteConfigMap, error) {
	var (
		out               C.SignalMutPointerBridgedStringMap
		c_initialCapacity = C.uint32_t(initialCapacity)
	)

	err := convertError(
		C.signal_bridged_string_map_new(
			&out,
			c_initialCapacity,
		),
	)
	if err != nil {
		return nil, err
	}

	return &RemoteConfigMap{ptr: out.raw}, nil
}

// C.signal_bridged_string_map_insert()
func (m *RemoteConfigMap) Insert(key, value string) error {
	if m.ptr == nil {
		return nil
	}

	var (
		c_key   = StringToCString(key)
		c_value = StringToCString(value)
	)
	defer C.free(unsafe.Pointer(c_key))
	defer C.free(unsafe.Pointer(&c_value))

	err := convertError(
		C.signal_bridged_string_map_insert(
			m.MutPointer(),
			(*C.int8_t)(unsafe.Pointer(c_key)),
			(*C.int8_t)(unsafe.Pointer(c_value)),
		),
	)

	return err
}

// C.signal_bridged_string_map_clone()
func (m *RemoteConfigMap) Clone() (*RemoteConfigMap, error) {
	if m.ptr == nil {
		return nil, nil
	}

	var out C.SignalMutPointerBridgedStringMap
	err := convertError(
		C.signal_bridged_string_map_clone(
			&out,
			m.ConstPointer(),
		),
	)
	if err != nil {
		return nil, err
	}

	return &RemoteConfigMap{ptr: out.raw}, nil
}

// C.signal_bridged_string_map_destroy()
func (m *RemoteConfigMap) Destroy() {
	if m.ptr != nil {
		C.signal_bridged_string_map_destroy(
			m.MutPointer(),
		)
		m.ptr = nil
	}
}
