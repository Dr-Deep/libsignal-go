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

type ProtocolAddress struct {
	ptr *C.SignalProtocolAddress
}

func NewProtocolAddress(name string, deviceID uint32) (*ProtocolAddress, error) {
	var (
		c_name     = StringToCString(name)
		c_deviceID = C.uint32_t(deviceID)
		out        C.SignalMutPointerProtocolAddress
	)
	defer C.free(unsafe.Pointer(c_name))

	err := convertError(
		C.signal_address_new(
			&out,
			(*C.int8_t)(unsafe.Pointer(c_name)),
			c_deviceID,
		),
	)
	if err != nil {
		return nil, err
	}

	return &ProtocolAddress{ptr: out.raw}, nil
}

func (addr *ProtocolAddress) Name() (string, error) {
	var (
		out C.SignalCStringPtr
	)

	err := convertError(
		C.signal_address_get_name(
			&out,
			C.SignalConstPointerProtocolAddress{raw: addr.ptr},
		),
	)
	if err != nil {
		return "", err
	}

	var name = SignalCStringPtrToString(out)
	C.signal_free_string(out)

	return name, nil
}

func (addr *ProtocolAddress) DeviceID() (uint32, error) {
	var (
		out C.uint32_t
	)

	err := convertError(
		C.signal_address_get_device_id(
			&out,
			C.SignalConstPointerProtocolAddress{raw: addr.ptr},
		),
	)
	if err != nil {
		return 0, err
	}

	return uint32(out), nil
}

func (addr *ProtocolAddress) Destroy() {
	if addr.ptr != nil {
		C.signal_address_destroy(
			C.SignalMutPointerProtocolAddress{raw: addr.ptr},
		)
		addr.ptr = nil
	}
}

// !
func (addr *ProtocolAddress) Clone() {}
