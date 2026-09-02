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

// C.signal_address_clone()
func (*Address) Clone() SignalError

// C. signal_address_destroy()
func (*Address) Destroy() SignalError

// C.signal_address_get_name()
func (addr *Address) Name() (string, error) {
	var (
		caddrptr = C.SignalConstPointerProtocolAddress{raw: addr.ptr}
		ptr      *C.char
	)

	err := convertError(
		// Call
		C.signal_address_get_name(
			&ptr,
			caddrptr,
		),
	)

	if err != nil {
		return "", err
	}

	//
	name := C.GoString(ptr)
	C.signal_free_string(ptr)

	return name, nil
}

// C.signal_address_get_device_id()
func (*Address) GetDeviceID() (uint32, SignalError)
