package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"
import (
	"unsafe"
)

type Address struct {
	ptr *C.SignalProtocolAddress
}

// C.signal_address_new()
func NewAddress(addr string, deviceID uint32) (*Address, error) {
	var (
		ptr   C.SignalMutPointerProtocolAddress
		caddr = (*C.int8_t)(unsafe.Pointer(
			C.CString(addr),
		))

		cdeviceID = C.uint32_t(deviceID)
	)
	defer C.free(unsafe.Pointer(caddr))

	err := convertError(
		// Call
		C.signal_address_new(
			&ptr,
			caddr,
			cdeviceID,
		),
	)

	if err != nil {
		return nil, err
	}

	return &Address{ptr: ptr.raw}, nil
}

// C.signal_address_get_name()
func (addr *Address) Name() (string, error) {
	var (
		out     C.SignalCStringPtr
		address = C.SignalConstPointerProtocolAddress{raw: addr.ptr}
	)

	namePtr, err := FfiCallFunc1(
		&out,
		address,
		Signal_address_get_name,
	)
	if err != nil {
		return "", err
	}

	name := C.GoString(
		(*C.char)(unsafe.Pointer(out)),
	)
	_ = namePtr

	return name, nil
}

// C.signal_address_get_device_id()
func (addr *Address) GetDeviceID() (uint32, error) {
	var (
		out     C.uint32_t
		address = C.SignalConstPointerProtocolAddress{raw: addr.ptr}
	)

	deviceIDPtr, err := FfiCallFunc1(
		&out,
		address,
		Signal_address_get_device_id,
	)
	if err != nil {
		return 0, err
	}

	return uint32(*deviceIDPtr), nil
}

// C.signal_address_clone()
func (addr *Address) Clone() (*Address, error) {
	var (
		out     C.SignalMutPointerProtocolAddress
		address = C.SignalConstPointerProtocolAddress{raw: addr.ptr}
	)

	_, err := FfiCallFunc1(
		&out,
		address,
		Signal_address_clone,
	)
	if err != nil {
		return nil, err
	}

	return &Address{ptr: out.raw}, nil
}

// C.signal_address_destroy()
func (addr *Address) Destroy() error {
	if addr.ptr == nil {
		return nil
	}

	var address = C.SignalMutPointerProtocolAddress{raw: addr.ptr}

	err := convertError(
		Signal_address_destroy(address),
	)

	addr.ptr = nil
	return err
}
