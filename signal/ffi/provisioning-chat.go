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

type ProvisioningChatConnection struct {
	ptr *C.SignalProvisioningChatConnection
}

func (chat *ProvisioningChatConnection) MutPointer() C.SignalMutPointerProvisioningChatConnection {
	return C.SignalMutPointerProvisioningChatConnection{raw: chat.ptr}
}

func (chat *ProvisioningChatConnection) ConstPointer() C.SignalConstPointerProvisioningChatConnection {
	return C.SignalConstPointerProvisioningChatConnection{raw: chat.ptr}
}

// C.signal_provisioning_chat_connection_info()
func (chat *ProvisioningChatConnection) Info() (*ChatConnectionInfo, error) {
	if chat.ptr == nil {
		return nil, nil
	}

	var out C.SignalMutPointerChatConnectionInfo
	err := convertError(
		C.signal_provisioning_chat_connection_info(
			&out,
			chat.ConstPointer(),
		),
	)

	if err != nil {
		return nil, err
	}

	return &ChatConnectionInfo{ptr: out.raw}, nil
}

// C.signal_provisioning_chat_connection_destroy()
func (chat *ProvisioningChatConnection) Destroy() {
	if chat.ptr != nil {
		C.signal_provisioning_chat_connection_destroy(
			chat.MutPointer(),
		)
		chat.ptr = nil
	}
}
