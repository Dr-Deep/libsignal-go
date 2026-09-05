package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_unauthenticated_chat_connection_account_exists(
SignalFfiError* signal_unauthenticated_chat_connection_backup_copy_media(
SignalFfiError* signal_unauthenticated_chat_connection_backup_delete_all(
SignalFfiError* signal_unauthenticated_chat_connection_backup_delete_media(
SignalFfiError* signal_unauthenticated_chat_connection_backup_get_cdn_credentials(
SignalFfiError* signal_unauthenticated_chat_connection_backup_get_media_backup_info(
SignalFfiError* signal_unauthenticated_chat_connection_backup_get_media_upload_form(
SignalFfiError* signal_unauthenticated_chat_connection_backup_get_message_backup_info(
SignalFfiError* signal_unauthenticated_chat_connection_backup_get_svrb_credentials(
SignalFfiError* signal_unauthenticated_chat_connection_backup_get_upload_form(
SignalFfiError* signal_unauthenticated_chat_connection_backup_list_media(
SignalFfiError* signal_unauthenticated_chat_connection_backup_refresh(
SignalFfiError* signal_unauthenticated_chat_connection_backup_set_public_key(
SignalFfiError* signal_unauthenticated_chat_connection_connect(
SignalFfiError* signal_unauthenticated_chat_connection_destroy(
SignalFfiError* signal_unauthenticated_chat_connection_disconnect(
SignalFfiError* signal_unauthenticated_chat_connection_get_pre_keys_access_key_auth(
SignalFfiError* signal_unauthenticated_chat_connection_get_pre_keys_group_auth(
SignalFfiError* signal_unauthenticated_chat_connection_get_pre_keys_unrestricted_auth(
SignalFfiError* signal_unauthenticated_chat_connection_info(
SignalFfiError* signal_unauthenticated_chat_connection_init_listener(
SignalFfiError* signal_unauthenticated_chat_connection_look_up_username_hash(
SignalFfiError* signal_unauthenticated_chat_connection_look_up_username_link(
SignalFfiError* signal_unauthenticated_chat_connection_send(
SignalFfiError* signal_unauthenticated_chat_connection_send_message(
SignalFfiError* signal_unauthenticated_chat_connection_send_multi_recipient_message(
SignalFfiError* signal_unauthenticated_chat_connection_send_raw_grpc(
SignalFfiError* signal_unauthenticated_chat_connection_submit_call_quality_survey(
*/

type UnauthenticatedChatConnection struct {
	ptr *C.SignalUnauthenticatedChatConnection
}

func (chat *UnauthenticatedChatConnection) MutPointer() C.SignalMutPointerUnauthenticatedChatConnection {
	return C.SignalMutPointerUnauthenticatedChatConnection{raw: chat.ptr}
}

func (chat *UnauthenticatedChatConnection) ConstPointer() C.SignalConstPointerUnauthenticatedChatConnection {
	return C.SignalConstPointerUnauthenticatedChatConnection{raw: chat.ptr}
}

func (chat *UnauthenticatedChatConnection) Destroy() {
	if chat.ptr != nil {
		C.signal_unauthenticated_chat_connection_destroy(
			chat.MutPointer(),
		)
		chat.ptr = nil
	}
}
