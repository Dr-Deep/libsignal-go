package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_chat_connection_info_description(
SignalFfiError* signal_chat_connection_info_destroy(
SignalFfiError* signal_chat_connection_info_ip_version(
SignalFfiError* signal_chat_connection_info_local_port(
*/

type IpType uint8

const (
	IpTypeUnknown IpType = 0
	IpTypeV4      IpType = 1
	IpTypeV6      IpType = 2
)

type ChatConnectionInfo struct {
	ptr *C.SignalChatConnectionInfo
}

func (info *ChatConnectionInfo) MutPointer() C.SignalMutPointerChatConnectionInfo {
	return C.SignalMutPointerChatConnectionInfo{raw: info.ptr}
}

func (info *ChatConnectionInfo) ConstPointer() C.SignalConstPointerChatConnectionInfo {
	return C.SignalConstPointerChatConnectionInfo{raw: info.ptr}
}

// C.signal_chat_connection_info_description()
func (info *ChatConnectionInfo) Description() (string, error) {
	if info.ptr == nil {
		return "", nil
	}

	var (
		out C.SignalCStringPtr
	)

	err := convertError(
		C.signal_chat_connection_info_description(
			&out,
			info.ConstPointer(),
		),
	)

	if err != nil {
		return "", err
	}

	descr := SignalCStringPtrToString(out)
	C.signal_free_string(out)
	return descr, nil
}

// C.signal_chat_connection_info_ip_version()
func (info *ChatConnectionInfo) IpVersion() (IpType, error) {
	if info.ptr == nil {
		return IpTypeUnknown, nil
	}

	var (
		out C.uint8_t
	)

	err := convertError(
		C.signal_chat_connection_info_ip_version(
			&out,
			info.ConstPointer(),
		),
	)

	if err != nil {
		return IpTypeUnknown, err
	}

	return IpType(out), nil
}

// C.signal_chat_connection_info_local_port()
func (info *ChatConnectionInfo) LocalPort() (uint16, error) {
	if info.ptr == nil {
		return 0, nil
	}

	var (
		out C.uint16_t
	)

	err := convertError(
		C.signal_chat_connection_info_local_port(
			&out,
			info.ConstPointer(),
		),
	)

	if err != nil {
		return 0, err
	}

	return uint16(out), nil
}

// C.signal_chat_connection_info_destroy()
func (info *ChatConnectionInfo) Destroy() {
	if info.ptr != nil {
		C.signal_chat_connection_info_destroy(
			info.MutPointer(),
		)
		info.ptr = nil
	}
}
