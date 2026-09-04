package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

// SignalFfiError* signal_account_entropy_pool_derive_backup_key(SignalType_FixedArray32_uint8_t* out, const int8_t* account_entropy);
func Signal_account_entropy_pool_derive_backup_key(
	out *C.SignalType_FixedArray32_uint8_t,
	account_entropy *C.int8_t,
) *C.SignalFfiError {
	return C.signal_account_entropy_pool_derive_backup_key(out, account_entropy)
}

// SignalFfiError* signal_account_entropy_pool_derive_svr_key(SignalType_FixedArray32_uint8_t* out, const int8_t* account_entropy);
func Signal_account_entropy_pool_derive_svr_key(
	out *C.SignalType_FixedArray32_uint8_t,
	account_entropy *C.int8_t,
) *C.SignalFfiError {
	return C.signal_account_entropy_pool_derive_svr_key(out, account_entropy)
}

// SignalFfiError* signal_account_entropy_pool_generate(SignalCStringPtr* out);
func Signal_account_entropy_pool_generate(
	out *C.SignalCStringPtr,
) *C.SignalFfiError {
	return C.signal_account_entropy_pool_generate(out)
}

// SignalFfiError* signal_account_entropy_pool_is_valid(bool* out, const int8_t* account_entropy);
func Signal_account_entropy_pool_is_valid(
	out *C.bool,
	account_entropy *C.int8_t,
) *C.SignalFfiError {
	return C.signal_account_entropy_pool_is_valid(out, account_entropy)
}

// SignalFfiError* signal_address_clone(SignalMutPointerProtocolAddress* new_obj, SignalConstPointerProtocolAddress obj);
func Signal_address_clone(
	new_obj *C.SignalMutPointerProtocolAddress,
	obj C.SignalConstPointerProtocolAddress,
) *C.SignalFfiError {
	return C.signal_address_clone(new_obj, obj)
}

// SignalFfiError* signal_address_destroy(SignalMutPointerProtocolAddress p);
func Signal_address_destroy(
	p C.SignalMutPointerProtocolAddress,
) *C.SignalFfiError {
	return C.signal_address_destroy(p)
}

// SignalFfiError* signal_address_get_device_id(uint32_t* out, SignalConstPointerProtocolAddress obj);
func Signal_address_get_device_id(
	out *C.uint32_t,
	obj C.SignalConstPointerProtocolAddress,
) *C.SignalFfiError {
	return C.signal_address_get_device_id(out, obj)
}

// SignalFfiError* signal_address_get_name(SignalCStringPtr* out, SignalConstPointerProtocolAddress obj);
func Signal_address_get_name(
	out *C.SignalCStringPtr,
	obj C.SignalConstPointerProtocolAddress,
) *C.SignalFfiError {
	return C.signal_address_get_name(out, obj)
}

// SignalFfiError* signal_address_new(SignalMutPointerProtocolAddress* out, const int8_t* name, uint32_t device_id);
func Signal_address_new(
	out *C.SignalMutPointerProtocolAddress,
	name *C.int8_t,
	device_id C.uint32_t,
) *C.SignalFfiError {
	return C.signal_address_new(out, name, device_id)
}

// SignalFfiError* signal_aes256_ctr32_destroy(SignalMutPointerAes256Ctr32 p);
func Signal_aes256_ctr32_destroy(
	p C.SignalMutPointerAes256Ctr32,
) *C.SignalFfiError {
	return C.signal_aes256_ctr32_destroy(p)
}

// SignalFfiError* signal_aes256_ctr32_new(SignalMutPointerAes256Ctr32* out, SignalBorrowedBuffer key, SignalBorrowedBuffer nonce, uint32_t initial_ctr);
func Signal_aes256_ctr32_new(
	out *C.SignalMutPointerAes256Ctr32,
	key C.SignalBorrowedBuffer,
	nonce C.SignalBorrowedBuffer,
	initial_ctr C.uint32_t,
) *C.SignalFfiError {
	return C.signal_aes256_ctr32_new(out, key, nonce, initial_ctr)
}

// SignalFfiError* signal_aes256_ctr32_process(SignalMutPointerAes256Ctr32 ctr, SignalBorrowedMutableBuffer data, uint32_t offset, uint32_t length);
func Signal_aes256_ctr32_process(
	ctr C.SignalMutPointerAes256Ctr32,
	data C.SignalBorrowedMutableBuffer,
	offset C.uint32_t,
	length C.uint32_t,
) *C.SignalFfiError {
	return C.signal_aes256_ctr32_process(ctr, data, offset, length)
}

// SignalFfiError* signal_aes256_gcm_decryption_destroy(SignalMutPointerAes256GcmDecryption p);
func Signal_aes256_gcm_decryption_destroy(
	p C.SignalMutPointerAes256GcmDecryption,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_decryption_destroy(p)
}

// SignalFfiError* signal_aes256_gcm_decryption_new(SignalMutPointerAes256GcmDecryption* out, SignalBorrowedBuffer key, SignalBorrowedBuffer nonce, SignalBorrowedBuffer associated_data);
func Signal_aes256_gcm_decryption_new(
	out *C.SignalMutPointerAes256GcmDecryption,
	key C.SignalBorrowedBuffer,
	nonce C.SignalBorrowedBuffer,
	associated_data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_decryption_new(out, key, nonce, associated_data)
}

// SignalFfiError* signal_aes256_gcm_decryption_update(SignalMutPointerAes256GcmDecryption gcm, SignalBorrowedMutableBuffer data, uint32_t offset, uint32_t length);
func Signal_aes256_gcm_decryption_update(
	gcm C.SignalMutPointerAes256GcmDecryption,
	data C.SignalBorrowedMutableBuffer,
	offset C.uint32_t,
	length C.uint32_t,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_decryption_update(gcm, data, offset, length)
}

// SignalFfiError* signal_aes256_gcm_decryption_verify_tag(bool* out, SignalMutPointerAes256GcmDecryption gcm, SignalBorrowedBuffer tag);
func Signal_aes256_gcm_decryption_verify_tag(
	out *C.bool,
	gcm C.SignalMutPointerAes256GcmDecryption,
	tag C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_decryption_verify_tag(out, gcm, tag)
}

// SignalFfiError* signal_aes256_gcm_encryption_compute_tag(SignalOwnedBuffer* out, SignalMutPointerAes256GcmEncryption gcm);
func Signal_aes256_gcm_encryption_compute_tag(
	out *C.SignalOwnedBuffer,
	gcm C.SignalMutPointerAes256GcmEncryption,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_encryption_compute_tag(out, gcm)
}

// SignalFfiError* signal_aes256_gcm_encryption_destroy(SignalMutPointerAes256GcmEncryption p);
func Signal_aes256_gcm_encryption_destroy(
	p C.SignalMutPointerAes256GcmEncryption,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_encryption_destroy(p)
}

// SignalFfiError* signal_aes256_gcm_encryption_new(SignalMutPointerAes256GcmEncryption* out, SignalBorrowedBuffer key, SignalBorrowedBuffer nonce, SignalBorrowedBuffer associated_data);
func Signal_aes256_gcm_encryption_new(
	out *C.SignalMutPointerAes256GcmEncryption,
	key C.SignalBorrowedBuffer,
	nonce C.SignalBorrowedBuffer,
	associated_data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_encryption_new(out, key, nonce, associated_data)
}

// SignalFfiError* signal_aes256_gcm_encryption_update(SignalMutPointerAes256GcmEncryption gcm, SignalBorrowedMutableBuffer data, uint32_t offset, uint32_t length);
func Signal_aes256_gcm_encryption_update(
	gcm C.SignalMutPointerAes256GcmEncryption,
	data C.SignalBorrowedMutableBuffer,
	offset C.uint32_t,
	length C.uint32_t,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_encryption_update(gcm, data, offset, length)
}

// SignalFfiError* signal_aes256_gcm_siv_decrypt(SignalOwnedBuffer* out, SignalConstPointerAes256GcmSiv aes_gcm_siv, SignalBorrowedBuffer ctext, SignalBorrowedBuffer nonce, SignalBorrowedBuffer associated_data);
func Signal_aes256_gcm_siv_decrypt(
	out *C.SignalOwnedBuffer,
	aes_gcm_siv C.SignalConstPointerAes256GcmSiv,
	ctext C.SignalBorrowedBuffer,
	nonce C.SignalBorrowedBuffer,
	associated_data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_siv_decrypt(out, aes_gcm_siv, ctext, nonce, associated_data)
}

// SignalFfiError* signal_aes256_gcm_siv_destroy(SignalMutPointerAes256GcmSiv p);
func Signal_aes256_gcm_siv_destroy(
	p C.SignalMutPointerAes256GcmSiv,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_siv_destroy(p)
}

// SignalFfiError* signal_aes256_gcm_siv_encrypt(SignalOwnedBuffer* out, SignalConstPointerAes256GcmSiv aes_gcm_siv_obj, SignalBorrowedBuffer ptext, SignalBorrowedBuffer nonce, SignalBorrowedBuffer associated_data);
func Signal_aes256_gcm_siv_encrypt(
	out *C.SignalOwnedBuffer,
	aes_gcm_siv_obj C.SignalConstPointerAes256GcmSiv,
	ptext C.SignalBorrowedBuffer,
	nonce C.SignalBorrowedBuffer,
	associated_data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_siv_encrypt(out, aes_gcm_siv_obj, ptext, nonce, associated_data)
}

// SignalFfiError* signal_aes256_gcm_siv_new(SignalMutPointerAes256GcmSiv* out, SignalBorrowedBuffer key);
func Signal_aes256_gcm_siv_new(
	out *C.SignalMutPointerAes256GcmSiv,
	key C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_aes256_gcm_siv_new(out, key)
}

// SignalFfiError* signal_auth_credential_presentation_check_valid_contents(SignalBorrowedBuffer presentation_bytes);
func Signal_auth_credential_presentation_check_valid_contents(
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_auth_credential_presentation_check_valid_contents(presentation_bytes)
}

// SignalFfiError* signal_auth_credential_presentation_get_pni_ciphertext(SignalType_FixedArray65_uint8_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_auth_credential_presentation_get_pni_ciphertext(
	out *C.SignalType_FixedArray65_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_auth_credential_presentation_get_pni_ciphertext(out, presentation_bytes)
}

// SignalFfiError* signal_auth_credential_presentation_get_redemption_time(uint64_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_auth_credential_presentation_get_redemption_time(
	out *C.uint64_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_auth_credential_presentation_get_redemption_time(out, presentation_bytes)
}

// SignalFfiError* signal_auth_credential_presentation_get_uuid_ciphertext(SignalType_FixedArray65_uint8_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_auth_credential_presentation_get_uuid_ciphertext(
	out *C.SignalType_FixedArray65_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_auth_credential_presentation_get_uuid_ciphertext(out, presentation_bytes)
}

// SignalFfiError* signal_auth_credential_with_pni_check_valid_contents(SignalBorrowedBuffer bytes);
func Signal_auth_credential_with_pni_check_valid_contents(
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_auth_credential_with_pni_check_valid_contents(bytes)
}

// SignalFfiError* signal_auth_credential_with_pni_response_check_valid_contents(SignalBorrowedBuffer bytes);
func Signal_auth_credential_with_pni_response_check_valid_contents(
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_auth_credential_with_pni_response_check_valid_contents(bytes)
}

// SignalFfiError* signal_authenticated_chat_connection_clear_push_token(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat);
func Signal_authenticated_chat_connection_clear_push_token(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_clear_push_token(promise, async_runtime, chat)
}

// SignalFfiError* signal_authenticated_chat_connection_clear_registration_lock(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat);
func Signal_authenticated_chat_connection_clear_registration_lock(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_clear_registration_lock(promise, async_runtime, chat)
}

// SignalFfiError* signal_authenticated_chat_connection_confirm_username(SignalCPromiseUuid* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, const int8_t* username, SignalBorrowedBuffer username_ciphertext, int64_t rng);
func Signal_authenticated_chat_connection_confirm_username(
	promise *C.SignalCPromiseUuid,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	username *C.int8_t,
	username_ciphertext C.SignalBorrowedBuffer,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_confirm_username(promise, async_runtime, chat, username, username_ciphertext, rng)
}

// SignalFfiError* signal_authenticated_chat_connection_connect(SignalCPromiseMutPointerAuthenticatedChatConnection* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerConnectionManager connection_manager, const int8_t* username, const int8_t* password, bool receive_stories, SignalBorrowedBytestringArray languages);
func Signal_authenticated_chat_connection_connect(
	promise *C.SignalCPromiseMutPointerAuthenticatedChatConnection,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	connection_manager C.SignalConstPointerConnectionManager,
	username *C.int8_t,
	password *C.int8_t,
	receive_stories C.bool,
	languages C.SignalBorrowedBytestringArray,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_connect(promise, async_runtime, connection_manager, username, password, receive_stories, languages)
}

// SignalFfiError* signal_authenticated_chat_connection_delete_username_hash(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat);
func Signal_authenticated_chat_connection_delete_username_hash(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_delete_username_hash(promise, async_runtime, chat)
}

// SignalFfiError* signal_authenticated_chat_connection_delete_username_link(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat);
func Signal_authenticated_chat_connection_delete_username_link(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_delete_username_link(promise, async_runtime, chat)
}

// SignalFfiError* signal_authenticated_chat_connection_destroy(SignalMutPointerAuthenticatedChatConnection p);
func Signal_authenticated_chat_connection_destroy(
	p C.SignalMutPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_destroy(p)
}

// SignalFfiError* signal_authenticated_chat_connection_disconnect(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat);
func Signal_authenticated_chat_connection_disconnect(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_disconnect(promise, async_runtime, chat)
}

// SignalFfiError* signal_authenticated_chat_connection_get_devices(SignalCPromiseOwnedBufferOfMaxAlignedLinkedDeviceInternalFfiResult* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat);
func Signal_authenticated_chat_connection_get_devices(
	promise *C.SignalCPromiseOwnedBufferOfMaxAlignedLinkedDeviceInternalFfiResult,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_get_devices(promise, async_runtime, chat)
}

// SignalFfiError* signal_authenticated_chat_connection_get_upload_form(SignalCPromiseFfiUploadForm* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, uint64_t upload_length);
func Signal_authenticated_chat_connection_get_upload_form(
	promise *C.SignalCPromiseFfiUploadForm,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	upload_length C.uint64_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_get_upload_form(promise, async_runtime, chat, upload_length)
}

// SignalFfiError* signal_authenticated_chat_connection_info(SignalMutPointerChatConnectionInfo* out, SignalConstPointerAuthenticatedChatConnection chat);
func Signal_authenticated_chat_connection_info(
	out *C.SignalMutPointerChatConnectionInfo,
	chat C.SignalConstPointerAuthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_info(out, chat)
}

// SignalFfiError* signal_authenticated_chat_connection_init_listener(SignalConstPointerAuthenticatedChatConnection chat, SignalConstPointerFfiChatListenerStruct listener);
func Signal_authenticated_chat_connection_init_listener(
	chat C.SignalConstPointerAuthenticatedChatConnection,
	listener C.SignalConstPointerFfiChatListenerStruct,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_init_listener(chat, listener)
}

// SignalFfiError* signal_authenticated_chat_connection_preconnect(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerConnectionManager connection_manager);
func Signal_authenticated_chat_connection_preconnect(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	connection_manager C.SignalConstPointerConnectionManager,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_preconnect(promise, async_runtime, connection_manager)
}

// SignalFfiError* signal_authenticated_chat_connection_redeem_backup_receipt(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, const SignalType_FixedArray329_uint8_t* presentation);
func Signal_authenticated_chat_connection_redeem_backup_receipt(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	presentation *C.SignalType_FixedArray329_uint8_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_redeem_backup_receipt(promise, async_runtime, chat, presentation)
}

// SignalFfiError* signal_authenticated_chat_connection_remove_device(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, uint8_t device_id);
func Signal_authenticated_chat_connection_remove_device(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	device_id C.uint8_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_remove_device(promise, async_runtime, chat, device_id)
}

// SignalFfiError* signal_authenticated_chat_connection_reserve_username_hash(SignalCPromisec_uchar32* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, SignalBorrowedSliceOfc_uchar32 username_hashes);
func Signal_authenticated_chat_connection_reserve_username_hash(
	promise *C.SignalCPromisec_uchar32,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	username_hashes C.SignalBorrowedSliceOfc_uchar32,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_reserve_username_hash(promise, async_runtime, chat, username_hashes)
}

// SignalFfiError* signal_authenticated_chat_connection_send(SignalCPromiseFfiChatResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, SignalConstPointerHttpRequest http_request, uint32_t timeout_millis);
func Signal_authenticated_chat_connection_send(
	promise *C.SignalCPromiseFfiChatResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	http_request C.SignalConstPointerHttpRequest,
	timeout_millis C.uint32_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_send(promise, async_runtime, chat, http_request, timeout_millis)
}

// SignalFfiError* signal_authenticated_chat_connection_send_message(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, const SignalType_FixedArray17_uint8_t* destination, uint64_t timestamp, SignalBorrowedSliceOfu32 device_ids, SignalBorrowedSliceOfu32 registration_ids, SignalBorrowedSliceOfConstPointerCiphertextMessage contents, bool online_only, bool is_urgent);
func Signal_authenticated_chat_connection_send_message(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	destination *C.SignalType_FixedArray17_uint8_t,
	timestamp C.uint64_t,
	device_ids C.SignalBorrowedSliceOfu32,
	registration_ids C.SignalBorrowedSliceOfu32,
	contents C.SignalBorrowedSliceOfConstPointerCiphertextMessage,
	online_only C.bool,
	is_urgent C.bool,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_send_message(promise, async_runtime, chat, destination, timestamp, device_ids, registration_ids, contents, online_only, is_urgent)
}

// SignalFfiError* signal_authenticated_chat_connection_send_raw_grpc(SignalCPromiseOwnedBuffer* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, const int8_t* service, const int8_t* method, SignalBorrowedBuffer payload);
func Signal_authenticated_chat_connection_send_raw_grpc(
	promise *C.SignalCPromiseOwnedBuffer,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	service *C.int8_t,
	method *C.int8_t,
	payload C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_send_raw_grpc(promise, async_runtime, chat, service, method, payload)
}

// SignalFfiError* signal_authenticated_chat_connection_send_sync_message(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, uint64_t timestamp, SignalBorrowedSliceOfu32 device_ids, SignalBorrowedSliceOfu32 registration_ids, SignalBorrowedSliceOfConstPointerCiphertextMessage contents, bool is_urgent);
func Signal_authenticated_chat_connection_send_sync_message(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	timestamp C.uint64_t,
	device_ids C.SignalBorrowedSliceOfu32,
	registration_ids C.SignalBorrowedSliceOfu32,
	contents C.SignalBorrowedSliceOfConstPointerCiphertextMessage,
	is_urgent C.bool,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_send_sync_message(promise, async_runtime, chat, timestamp, device_ids, registration_ids, contents, is_urgent)
}

// SignalFfiError* signal_authenticated_chat_connection_set_device_name(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, uint8_t device_id, SignalBorrowedBuffer encrypted_name);
func Signal_authenticated_chat_connection_set_device_name(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	device_id C.uint8_t,
	encrypted_name C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_set_device_name(promise, async_runtime, chat, device_id, encrypted_name)
}

// SignalFfiError* signal_authenticated_chat_connection_set_discoverable_by_phone_number(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, bool discoverable);
func Signal_authenticated_chat_connection_set_discoverable_by_phone_number(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	discoverable C.bool,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_set_discoverable_by_phone_number(promise, async_runtime, chat, discoverable)
}

// SignalFfiError* signal_authenticated_chat_connection_set_push_token_apns(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, const int8_t* apns_token);
func Signal_authenticated_chat_connection_set_push_token_apns(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	apns_token *C.int8_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_set_push_token_apns(promise, async_runtime, chat, apns_token)
}

// SignalFfiError* signal_authenticated_chat_connection_set_registration_lock(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, const SignalType_FixedArray32_uint8_t* svr_key);
func Signal_authenticated_chat_connection_set_registration_lock(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	svr_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_set_registration_lock(promise, async_runtime, chat, svr_key)
}

// SignalFfiError* signal_authenticated_chat_connection_set_registration_recovery_password(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, const SignalType_FixedArray32_uint8_t* svr_key);
func Signal_authenticated_chat_connection_set_registration_recovery_password(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	svr_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_set_registration_recovery_password(promise, async_runtime, chat, svr_key)
}

// SignalFfiError* signal_authenticated_chat_connection_set_username_link(SignalCPromiseUuid* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerAuthenticatedChatConnection chat, SignalBorrowedBuffer username_ciphertext, bool keep_link_handle);
func Signal_authenticated_chat_connection_set_username_link(
	promise *C.SignalCPromiseUuid,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerAuthenticatedChatConnection,
	username_ciphertext C.SignalBorrowedBuffer,
	keep_link_handle C.bool,
) *C.SignalFfiError {
	return C.signal_authenticated_chat_connection_set_username_link(promise, async_runtime, chat, username_ciphertext, keep_link_handle)
}

// SignalFfiError* signal_avatar_upload_credential_check_valid_contents(SignalBorrowedBuffer credential_bytes);
func Signal_avatar_upload_credential_check_valid_contents(
	credential_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_check_valid_contents(credential_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_get_cm(SignalType_FixedArray32_uint8_t* out, SignalBorrowedBuffer credential_bytes);
func Signal_avatar_upload_credential_get_cm(
	out *C.SignalType_FixedArray32_uint8_t,
	credential_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_get_cm(out, credential_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_get_redemption_time(uint64_t* out, SignalBorrowedBuffer credential_bytes);
func Signal_avatar_upload_credential_get_redemption_time(
	out *C.uint64_t,
	credential_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_get_redemption_time(out, credential_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_present_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer credential_bytes, SignalBorrowedBuffer server_params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_avatar_upload_credential_present_deterministic(
	out *C.SignalOwnedBuffer,
	credential_bytes C.SignalBorrowedBuffer,
	server_params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_present_deterministic(out, credential_bytes, server_params_bytes, randomness)
}

// SignalFfiError* signal_avatar_upload_credential_presentation_check_valid_contents(SignalBorrowedBuffer presentation_bytes);
func Signal_avatar_upload_credential_presentation_check_valid_contents(
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_presentation_check_valid_contents(presentation_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_presentation_get_cm(SignalType_FixedArray32_uint8_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_avatar_upload_credential_presentation_get_cm(
	out *C.SignalType_FixedArray32_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_presentation_get_cm(out, presentation_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_presentation_get_redemption_time(uint64_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_avatar_upload_credential_presentation_get_redemption_time(
	out *C.uint64_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_presentation_get_redemption_time(out, presentation_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_presentation_verify(SignalBorrowedBuffer presentation_bytes, uint64_t current_time, SignalBorrowedBuffer server_params_bytes);
func Signal_avatar_upload_credential_presentation_verify(
	presentation_bytes C.SignalBorrowedBuffer,
	current_time C.uint64_t,
	server_params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_presentation_verify(presentation_bytes, current_time, server_params_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_request_check_valid_contents(SignalBorrowedBuffer request_bytes);
func Signal_avatar_upload_credential_request_check_valid_contents(
	request_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_request_check_valid_contents(request_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_request_context_check_valid_contents(SignalBorrowedBuffer context_bytes);
func Signal_avatar_upload_credential_request_context_check_valid_contents(
	context_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_request_context_check_valid_contents(context_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_request_context_get_request(SignalOwnedBuffer* out, SignalBorrowedBuffer context_bytes);
func Signal_avatar_upload_credential_request_context_get_request(
	out *C.SignalOwnedBuffer,
	context_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_request_context_get_request(out, context_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_request_context_new(SignalOwnedBuffer* out, const SignalType_FixedArray17_uint8_t* aci, SignalBorrowedBuffer zk_credential_key_pair_bytes, uint64_t rotation_id, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_avatar_upload_credential_request_context_new(
	out *C.SignalOwnedBuffer,
	aci *C.SignalType_FixedArray17_uint8_t,
	zk_credential_key_pair_bytes C.SignalBorrowedBuffer,
	rotation_id C.uint64_t,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_request_context_new(out, aci, zk_credential_key_pair_bytes, rotation_id, randomness)
}

// SignalFfiError* signal_avatar_upload_credential_request_context_receive_response(SignalOwnedBuffer* out, SignalBorrowedBuffer context_bytes, SignalBorrowedBuffer response_bytes, uint64_t current_time, SignalBorrowedBuffer params_bytes);
func Signal_avatar_upload_credential_request_context_receive_response(
	out *C.SignalOwnedBuffer,
	context_bytes C.SignalBorrowedBuffer,
	response_bytes C.SignalBorrowedBuffer,
	current_time C.uint64_t,
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_request_context_receive_response(out, context_bytes, response_bytes, current_time, params_bytes)
}

// SignalFfiError* signal_avatar_upload_credential_request_issue_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer request_bytes, const SignalType_FixedArray17_uint8_t* aci, SignalBorrowedBuffer zk_credential_key_pub_bytes, uint64_t rotation_id, uint64_t redemption_time, SignalBorrowedBuffer params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_avatar_upload_credential_request_issue_deterministic(
	out *C.SignalOwnedBuffer,
	request_bytes C.SignalBorrowedBuffer,
	aci *C.SignalType_FixedArray17_uint8_t,
	zk_credential_key_pub_bytes C.SignalBorrowedBuffer,
	rotation_id C.uint64_t,
	redemption_time C.uint64_t,
	params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_request_issue_deterministic(out, request_bytes, aci, zk_credential_key_pub_bytes, rotation_id, redemption_time, params_bytes, randomness)
}

// SignalFfiError* signal_avatar_upload_credential_response_check_valid_contents(SignalBorrowedBuffer response_bytes);
func Signal_avatar_upload_credential_response_check_valid_contents(
	response_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_avatar_upload_credential_response_check_valid_contents(response_bytes)
}

// SignalFfiError* signal_backup_auth_credential_check_valid_contents(SignalBorrowedBuffer params_bytes);
func Signal_backup_auth_credential_check_valid_contents(
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_check_valid_contents(params_bytes)
}

// SignalFfiError* signal_backup_auth_credential_get_backup_id(SignalType_FixedArray16_uint8_t* out, SignalBorrowedBuffer credential_bytes);
func Signal_backup_auth_credential_get_backup_id(
	out *C.SignalType_FixedArray16_uint8_t,
	credential_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_get_backup_id(out, credential_bytes)
}

// SignalFfiError* signal_backup_auth_credential_get_backup_level(uint8_t* out, SignalBorrowedBuffer credential_bytes);
func Signal_backup_auth_credential_get_backup_level(
	out *C.uint8_t,
	credential_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_get_backup_level(out, credential_bytes)
}

// SignalFfiError* signal_backup_auth_credential_get_type(uint8_t* out, SignalBorrowedBuffer credential_bytes);
func Signal_backup_auth_credential_get_type(
	out *C.uint8_t,
	credential_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_get_type(out, credential_bytes)
}

// SignalFfiError* signal_backup_auth_credential_present_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer credential_bytes, SignalBorrowedBuffer server_params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_backup_auth_credential_present_deterministic(
	out *C.SignalOwnedBuffer,
	credential_bytes C.SignalBorrowedBuffer,
	server_params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_present_deterministic(out, credential_bytes, server_params_bytes, randomness)
}

// SignalFfiError* signal_backup_auth_credential_presentation_check_valid_contents(SignalBorrowedBuffer presentation_bytes);
func Signal_backup_auth_credential_presentation_check_valid_contents(
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_presentation_check_valid_contents(presentation_bytes)
}

// SignalFfiError* signal_backup_auth_credential_presentation_verify(SignalBorrowedBuffer presentation_bytes, uint64_t now, SignalBorrowedBuffer server_params_bytes);
func Signal_backup_auth_credential_presentation_verify(
	presentation_bytes C.SignalBorrowedBuffer,
	now C.uint64_t,
	server_params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_presentation_verify(presentation_bytes, now, server_params_bytes)
}

// SignalFfiError* signal_backup_auth_credential_request_check_valid_contents(SignalBorrowedBuffer request_bytes);
func Signal_backup_auth_credential_request_check_valid_contents(
	request_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_request_check_valid_contents(request_bytes)
}

// SignalFfiError* signal_backup_auth_credential_request_context_check_valid_contents(SignalBorrowedBuffer context_bytes);
func Signal_backup_auth_credential_request_context_check_valid_contents(
	context_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_request_context_check_valid_contents(context_bytes)
}

// SignalFfiError* signal_backup_auth_credential_request_context_get_request(SignalOwnedBuffer* out, SignalBorrowedBuffer context_bytes);
func Signal_backup_auth_credential_request_context_get_request(
	out *C.SignalOwnedBuffer,
	context_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_request_context_get_request(out, context_bytes)
}

// SignalFfiError* signal_backup_auth_credential_request_context_new(SignalOwnedBuffer* out, const SignalType_FixedArray32_uint8_t* backup_key, SignalUuid uuid);
func Signal_backup_auth_credential_request_context_new(
	out *C.SignalOwnedBuffer,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	uuid C.SignalUuid,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_request_context_new(out, backup_key, uuid)
}

// SignalFfiError* signal_backup_auth_credential_request_context_receive_response(SignalOwnedBuffer* out, SignalBorrowedBuffer context_bytes, SignalBorrowedBuffer response_bytes, uint64_t expected_redemption_time, SignalBorrowedBuffer params_bytes);
func Signal_backup_auth_credential_request_context_receive_response(
	out *C.SignalOwnedBuffer,
	context_bytes C.SignalBorrowedBuffer,
	response_bytes C.SignalBorrowedBuffer,
	expected_redemption_time C.uint64_t,
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_request_context_receive_response(out, context_bytes, response_bytes, expected_redemption_time, params_bytes)
}

// SignalFfiError* signal_backup_auth_credential_request_issue_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer request_bytes, uint64_t redemption_time, uint8_t backup_level, uint8_t credential_type, SignalBorrowedBuffer params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_backup_auth_credential_request_issue_deterministic(
	out *C.SignalOwnedBuffer,
	request_bytes C.SignalBorrowedBuffer,
	redemption_time C.uint64_t,
	backup_level C.uint8_t,
	credential_type C.uint8_t,
	params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_request_issue_deterministic(out, request_bytes, redemption_time, backup_level, credential_type, params_bytes, randomness)
}

// SignalFfiError* signal_backup_auth_credential_response_check_valid_contents(SignalBorrowedBuffer response_bytes);
func Signal_backup_auth_credential_response_check_valid_contents(
	response_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_backup_auth_credential_response_check_valid_contents(response_bytes)
}

// SignalFfiError* signal_backup_key_derive_backup_id(SignalType_FixedArray16_uint8_t* out, const SignalType_FixedArray32_uint8_t* backup_key, const SignalType_FixedArray17_uint8_t* aci);
func Signal_backup_key_derive_backup_id(
	out *C.SignalType_FixedArray16_uint8_t,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	aci *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_backup_key_derive_backup_id(out, backup_key, aci)
}

// SignalFfiError* signal_backup_key_derive_ec_key(SignalMutPointerPrivateKey* out, const SignalType_FixedArray32_uint8_t* backup_key, const SignalType_FixedArray17_uint8_t* aci);
func Signal_backup_key_derive_ec_key(
	out *C.SignalMutPointerPrivateKey,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	aci *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_backup_key_derive_ec_key(out, backup_key, aci)
}

// SignalFfiError* signal_backup_key_derive_local_backup_metadata_key(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray32_uint8_t* backup_key);
func Signal_backup_key_derive_local_backup_metadata_key(
	out *C.SignalType_FixedArray32_uint8_t,
	backup_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_backup_key_derive_local_backup_metadata_key(out, backup_key)
}

// SignalFfiError* signal_backup_key_derive_media_encryption_key(SignalType_FixedArray64_uint8_t* out, const SignalType_FixedArray32_uint8_t* backup_key, const SignalType_FixedArray15_uint8_t* media_id);
func Signal_backup_key_derive_media_encryption_key(
	out *C.SignalType_FixedArray64_uint8_t,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	media_id *C.SignalType_FixedArray15_uint8_t,
) *C.SignalFfiError {
	return C.signal_backup_key_derive_media_encryption_key(out, backup_key, media_id)
}

// SignalFfiError* signal_backup_key_derive_media_id(SignalType_FixedArray15_uint8_t* out, const SignalType_FixedArray32_uint8_t* backup_key, const int8_t* media_name);
func Signal_backup_key_derive_media_id(
	out *C.SignalType_FixedArray15_uint8_t,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	media_name *C.int8_t,
) *C.SignalFfiError {
	return C.signal_backup_key_derive_media_id(out, backup_key, media_name)
}

// SignalFfiError* signal_backup_key_derive_thumbnail_transit_encryption_key(SignalType_FixedArray64_uint8_t* out, const SignalType_FixedArray32_uint8_t* backup_key, const SignalType_FixedArray15_uint8_t* media_id);
func Signal_backup_key_derive_thumbnail_transit_encryption_key(
	out *C.SignalType_FixedArray64_uint8_t,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	media_id *C.SignalType_FixedArray15_uint8_t,
) *C.SignalFfiError {
	return C.signal_backup_key_derive_thumbnail_transit_encryption_key(out, backup_key, media_id)
}

// SignalFfiError* signal_backup_restore_response_destroy(SignalMutPointerBackupRestoreResponse p);
func Signal_backup_restore_response_destroy(
	p C.SignalMutPointerBackupRestoreResponse,
) *C.SignalFfiError {
	return C.signal_backup_restore_response_destroy(p)
}

// SignalFfiError* signal_backup_restore_response_get_forward_secrecy_token(SignalType_FixedArray32_uint8_t* out, SignalConstPointerBackupRestoreResponse response);
func Signal_backup_restore_response_get_forward_secrecy_token(
	out *C.SignalType_FixedArray32_uint8_t,
	response C.SignalConstPointerBackupRestoreResponse,
) *C.SignalFfiError {
	return C.signal_backup_restore_response_get_forward_secrecy_token(out, response)
}

// SignalFfiError* signal_backup_restore_response_get_next_backup_secret_data(SignalOwnedBuffer* out, SignalConstPointerBackupRestoreResponse response);
func Signal_backup_restore_response_get_next_backup_secret_data(
	out *C.SignalOwnedBuffer,
	response C.SignalConstPointerBackupRestoreResponse,
) *C.SignalFfiError {
	return C.signal_backup_restore_response_get_next_backup_secret_data(out, response)
}

// SignalFfiError* signal_backup_store_response_destroy(SignalMutPointerBackupStoreResponse p);
func Signal_backup_store_response_destroy(
	p C.SignalMutPointerBackupStoreResponse,
) *C.SignalFfiError {
	return C.signal_backup_store_response_destroy(p)
}

// SignalFfiError* signal_backup_store_response_get_forward_secrecy_token(SignalType_FixedArray32_uint8_t* out, SignalConstPointerBackupStoreResponse response);
func Signal_backup_store_response_get_forward_secrecy_token(
	out *C.SignalType_FixedArray32_uint8_t,
	response C.SignalConstPointerBackupStoreResponse,
) *C.SignalFfiError {
	return C.signal_backup_store_response_get_forward_secrecy_token(out, response)
}

// SignalFfiError* signal_backup_store_response_get_next_backup_secret_data(SignalOwnedBuffer* out, SignalConstPointerBackupStoreResponse response);
func Signal_backup_store_response_get_next_backup_secret_data(
	out *C.SignalOwnedBuffer,
	response C.SignalConstPointerBackupStoreResponse,
) *C.SignalFfiError {
	return C.signal_backup_store_response_get_next_backup_secret_data(out, response)
}

// SignalFfiError* signal_backup_store_response_get_opaque_metadata(SignalOwnedBuffer* out, SignalConstPointerBackupStoreResponse response);
func Signal_backup_store_response_get_opaque_metadata(
	out *C.SignalOwnedBuffer,
	response C.SignalConstPointerBackupStoreResponse,
) *C.SignalFfiError {
	return C.signal_backup_store_response_get_opaque_metadata(out, response)
}

// SignalFfiError* signal_bridged_string_map_clone(SignalMutPointerBridgedStringMap* new_obj, SignalConstPointerBridgedStringMap obj);
func Signal_bridged_string_map_clone(
	new_obj *C.SignalMutPointerBridgedStringMap,
	obj C.SignalConstPointerBridgedStringMap,
) *C.SignalFfiError {
	return C.signal_bridged_string_map_clone(new_obj, obj)
}

// SignalFfiError* signal_bridged_string_map_destroy(SignalMutPointerBridgedStringMap p);
func Signal_bridged_string_map_destroy(
	p C.SignalMutPointerBridgedStringMap,
) *C.SignalFfiError {
	return C.signal_bridged_string_map_destroy(p)
}

/*
// SignalFfiError* signal_bridged_string_map_insert(SignalMutPointerBridgedStringMap map, const int8_t* key, const int8_t* value);
func Signal_bridged_string_map_insert(
    map C.SignalMutPointerBridgedStringMap,
    key *C.int8_t,
    value *C.int8_t,
) *C.SignalFfiError {
	return C.signal_bridged_string_map_insert(map, key, value)
    }
*/

// SignalFfiError* signal_bridged_string_map_new(SignalMutPointerBridgedStringMap* out, uint32_t initial_capacity);
func Signal_bridged_string_map_new(
	out *C.SignalMutPointerBridgedStringMap,
	initial_capacity C.uint32_t,
) *C.SignalFfiError {
	return C.signal_bridged_string_map_new(out, initial_capacity)
}

// SignalFfiError* signal_call_link_auth_credential_check_valid_contents(SignalBorrowedBuffer credential_bytes);
func Signal_call_link_auth_credential_check_valid_contents(
	credential_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_check_valid_contents(credential_bytes)
}

// SignalFfiError* signal_call_link_auth_credential_present_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer credential_bytes, const SignalType_FixedArray17_uint8_t* user_id, uint64_t redemption_time, SignalBorrowedBuffer server_params_bytes, SignalBorrowedBuffer call_link_params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_call_link_auth_credential_present_deterministic(
	out *C.SignalOwnedBuffer,
	credential_bytes C.SignalBorrowedBuffer,
	user_id *C.SignalType_FixedArray17_uint8_t,
	redemption_time C.uint64_t,
	server_params_bytes C.SignalBorrowedBuffer,
	call_link_params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_present_deterministic(out, credential_bytes, user_id, redemption_time, server_params_bytes, call_link_params_bytes, randomness)
}

// SignalFfiError* signal_call_link_auth_credential_presentation_check_valid_contents(SignalBorrowedBuffer presentation_bytes);
func Signal_call_link_auth_credential_presentation_check_valid_contents(
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_presentation_check_valid_contents(presentation_bytes)
}

// SignalFfiError* signal_call_link_auth_credential_presentation_get_user_id(SignalType_FixedArray65_uint8_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_call_link_auth_credential_presentation_get_user_id(
	out *C.SignalType_FixedArray65_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_presentation_get_user_id(out, presentation_bytes)
}

// SignalFfiError* signal_call_link_auth_credential_presentation_verify(SignalBorrowedBuffer presentation_bytes, uint64_t now, SignalBorrowedBuffer server_params_bytes, SignalBorrowedBuffer call_link_params_bytes);
func Signal_call_link_auth_credential_presentation_verify(
	presentation_bytes C.SignalBorrowedBuffer,
	now C.uint64_t,
	server_params_bytes C.SignalBorrowedBuffer,
	call_link_params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_presentation_verify(presentation_bytes, now, server_params_bytes, call_link_params_bytes)
}

// SignalFfiError* signal_call_link_auth_credential_response_check_valid_contents(SignalBorrowedBuffer response_bytes);
func Signal_call_link_auth_credential_response_check_valid_contents(
	response_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_response_check_valid_contents(response_bytes)
}

// SignalFfiError* signal_call_link_auth_credential_response_issue_deterministic(SignalOwnedBuffer* out, const SignalType_FixedArray17_uint8_t* user_id, uint64_t redemption_time, SignalBorrowedBuffer params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_call_link_auth_credential_response_issue_deterministic(
	out *C.SignalOwnedBuffer,
	user_id *C.SignalType_FixedArray17_uint8_t,
	redemption_time C.uint64_t,
	params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_response_issue_deterministic(out, user_id, redemption_time, params_bytes, randomness)
}

// SignalFfiError* signal_call_link_auth_credential_response_receive(SignalOwnedBuffer* out, SignalBorrowedBuffer response_bytes, const SignalType_FixedArray17_uint8_t* user_id, uint64_t redemption_time, SignalBorrowedBuffer params_bytes);
func Signal_call_link_auth_credential_response_receive(
	out *C.SignalOwnedBuffer,
	response_bytes C.SignalBorrowedBuffer,
	user_id *C.SignalType_FixedArray17_uint8_t,
	redemption_time C.uint64_t,
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_auth_credential_response_receive(out, response_bytes, user_id, redemption_time, params_bytes)
}

// SignalFfiError* signal_call_link_public_params_check_valid_contents(SignalBorrowedBuffer params_bytes);
func Signal_call_link_public_params_check_valid_contents(
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_public_params_check_valid_contents(params_bytes)
}

// SignalFfiError* signal_call_link_secret_params_check_valid_contents(SignalBorrowedBuffer params_bytes);
func Signal_call_link_secret_params_check_valid_contents(
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_secret_params_check_valid_contents(params_bytes)
}

// SignalFfiError* signal_call_link_secret_params_decrypt_user_id(SignalType_FixedArray17_uint8_t* out, SignalBorrowedBuffer params_bytes, const SignalType_FixedArray65_uint8_t* user_id);
func Signal_call_link_secret_params_decrypt_user_id(
	out *C.SignalType_FixedArray17_uint8_t,
	params_bytes C.SignalBorrowedBuffer,
	user_id *C.SignalType_FixedArray65_uint8_t,
) *C.SignalFfiError {
	return C.signal_call_link_secret_params_decrypt_user_id(out, params_bytes, user_id)
}

// SignalFfiError* signal_call_link_secret_params_derive_from_root_key(SignalOwnedBuffer* out, SignalBorrowedBuffer root_key);
func Signal_call_link_secret_params_derive_from_root_key(
	out *C.SignalOwnedBuffer,
	root_key C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_secret_params_derive_from_root_key(out, root_key)
}

// SignalFfiError* signal_call_link_secret_params_encrypt_user_id(SignalType_FixedArray65_uint8_t* out, SignalBorrowedBuffer params_bytes, const SignalType_FixedArray17_uint8_t* user_id);
func Signal_call_link_secret_params_encrypt_user_id(
	out *C.SignalType_FixedArray65_uint8_t,
	params_bytes C.SignalBorrowedBuffer,
	user_id *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_call_link_secret_params_encrypt_user_id(out, params_bytes, user_id)
}

// SignalFfiError* signal_call_link_secret_params_get_public_params(SignalOwnedBuffer* out, SignalBorrowedBuffer params_bytes);
func Signal_call_link_secret_params_get_public_params(
	out *C.SignalOwnedBuffer,
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_call_link_secret_params_get_public_params(out, params_bytes)
}

// SignalFfiError* signal_cds2_client_state_new(SignalMutPointerSgxClientState* out, SignalBorrowedBuffer mrenclave, SignalBorrowedBuffer attestation_msg, uint64_t current_timestamp);
func Signal_cds2_client_state_new(
	out *C.SignalMutPointerSgxClientState,
	mrenclave C.SignalBorrowedBuffer,
	attestation_msg C.SignalBorrowedBuffer,
	current_timestamp C.uint64_t,
) *C.SignalFfiError {
	return C.signal_cds2_client_state_new(out, mrenclave, attestation_msg, current_timestamp)
}

// SignalFfiError* signal_cdsi_lookup_complete(SignalCPromiseFfiCdsiLookupResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerCdsiLookup lookup);
func Signal_cdsi_lookup_complete(
	promise *C.SignalCPromiseFfiCdsiLookupResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	lookup C.SignalConstPointerCdsiLookup,
) *C.SignalFfiError {
	return C.signal_cdsi_lookup_complete(promise, async_runtime, lookup)
}

// SignalFfiError* signal_cdsi_lookup_destroy(SignalMutPointerCdsiLookup p);
func Signal_cdsi_lookup_destroy(
	p C.SignalMutPointerCdsiLookup,
) *C.SignalFfiError {
	return C.signal_cdsi_lookup_destroy(p)
}

// SignalFfiError* signal_cdsi_lookup_new(SignalCPromiseMutPointerCdsiLookup* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerConnectionManager connection_manager, const int8_t* username, const int8_t* password, SignalConstPointerLookupRequest request);
func Signal_cdsi_lookup_new(
	promise *C.SignalCPromiseMutPointerCdsiLookup,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	connection_manager C.SignalConstPointerConnectionManager,
	username *C.int8_t,
	password *C.int8_t,
	request C.SignalConstPointerLookupRequest,
) *C.SignalFfiError {
	return C.signal_cdsi_lookup_new(promise, async_runtime, connection_manager, username, password, request)
}

// SignalFfiError* signal_cdsi_lookup_token(SignalOwnedBuffer* out, SignalConstPointerCdsiLookup lookup);
func Signal_cdsi_lookup_token(
	out *C.SignalOwnedBuffer,
	lookup C.SignalConstPointerCdsiLookup,
) *C.SignalFfiError {
	return C.signal_cdsi_lookup_token(out, lookup)
}

// SignalFfiError* signal_chat_connection_info_description(SignalCStringPtr* out, SignalConstPointerChatConnectionInfo connection_info);
func Signal_chat_connection_info_description(
	out *C.SignalCStringPtr,
	connection_info C.SignalConstPointerChatConnectionInfo,
) *C.SignalFfiError {
	return C.signal_chat_connection_info_description(out, connection_info)
}

// SignalFfiError* signal_chat_connection_info_destroy(SignalMutPointerChatConnectionInfo p);
func Signal_chat_connection_info_destroy(
	p C.SignalMutPointerChatConnectionInfo,
) *C.SignalFfiError {
	return C.signal_chat_connection_info_destroy(p)
}

// SignalFfiError* signal_chat_connection_info_ip_version(uint8_t* out, SignalConstPointerChatConnectionInfo connection_info);
func Signal_chat_connection_info_ip_version(
	out *C.uint8_t,
	connection_info C.SignalConstPointerChatConnectionInfo,
) *C.SignalFfiError {
	return C.signal_chat_connection_info_ip_version(out, connection_info)
}

// SignalFfiError* signal_chat_connection_info_local_port(uint16_t* out, SignalConstPointerChatConnectionInfo connection_info);
func Signal_chat_connection_info_local_port(
	out *C.uint16_t,
	connection_info C.SignalConstPointerChatConnectionInfo,
) *C.SignalFfiError {
	return C.signal_chat_connection_info_local_port(out, connection_info)
}

// SignalFfiError* signal_ciphertext_message_destroy(SignalMutPointerCiphertextMessage p);
func Signal_ciphertext_message_destroy(
	p C.SignalMutPointerCiphertextMessage,
) *C.SignalFfiError {
	return C.signal_ciphertext_message_destroy(p)
}

// SignalFfiError* signal_ciphertext_message_from_plaintext_content(SignalMutPointerCiphertextMessage* out, SignalConstPointerPlaintextContent m);
func Signal_ciphertext_message_from_plaintext_content(
	out *C.SignalMutPointerCiphertextMessage,
	m C.SignalConstPointerPlaintextContent,
) *C.SignalFfiError {
	return C.signal_ciphertext_message_from_plaintext_content(out, m)
}

// SignalFfiError* signal_ciphertext_message_serialize(SignalOwnedBuffer* out, SignalConstPointerCiphertextMessage obj);
func Signal_ciphertext_message_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerCiphertextMessage,
) *C.SignalFfiError {
	return C.signal_ciphertext_message_serialize(out, obj)
}

// SignalFfiError* signal_ciphertext_message_type(uint8_t* out, SignalConstPointerCiphertextMessage msg);
func Signal_ciphertext_message_type(
	out *C.uint8_t,
	msg C.SignalConstPointerCiphertextMessage,
) *C.SignalFfiError {
	return C.signal_ciphertext_message_type(out, msg)
}

// SignalFfiError* signal_connection_manager_clear_proxy(SignalConstPointerConnectionManager connection_manager);
func Signal_connection_manager_clear_proxy(
	connection_manager C.SignalConstPointerConnectionManager,
) *C.SignalFfiError {
	return C.signal_connection_manager_clear_proxy(connection_manager)
}

// SignalFfiError* signal_connection_manager_destroy(SignalMutPointerConnectionManager p);
func Signal_connection_manager_destroy(
	p C.SignalMutPointerConnectionManager,
) *C.SignalFfiError {
	return C.signal_connection_manager_destroy(p)
}

// SignalFfiError* signal_connection_manager_new(SignalMutPointerConnectionManager* out, uint8_t environment, const int8_t* user_agent, SignalMutPointerBridgedStringMap remote_config, uint8_t build_variant);
func Signal_connection_manager_new(
	out *C.SignalMutPointerConnectionManager,
	environment C.uint8_t,
	user_agent *C.int8_t,
	remote_config C.SignalMutPointerBridgedStringMap,
	build_variant C.uint8_t,
) *C.SignalFfiError {
	return C.signal_connection_manager_new(out, environment, user_agent, remote_config, build_variant)
}

// SignalFfiError* signal_connection_manager_on_network_change(SignalConstPointerConnectionManager connection_manager);
func Signal_connection_manager_on_network_change(
	connection_manager C.SignalConstPointerConnectionManager,
) *C.SignalFfiError {
	return C.signal_connection_manager_on_network_change(connection_manager)
}

// SignalFfiError* signal_connection_manager_set_censorship_circumvention_enabled(SignalConstPointerConnectionManager connection_manager, bool enabled);
func Signal_connection_manager_set_censorship_circumvention_enabled(
	connection_manager C.SignalConstPointerConnectionManager,
	enabled C.bool,
) *C.SignalFfiError {
	return C.signal_connection_manager_set_censorship_circumvention_enabled(connection_manager, enabled)
}

// SignalFfiError* signal_connection_manager_set_invalid_proxy(SignalConstPointerConnectionManager connection_manager);
func Signal_connection_manager_set_invalid_proxy(
	connection_manager C.SignalConstPointerConnectionManager,
) *C.SignalFfiError {
	return C.signal_connection_manager_set_invalid_proxy(connection_manager)
}

// SignalFfiError* signal_connection_manager_set_proxy(SignalConstPointerConnectionManager connection_manager, SignalConstPointerConnectionProxyConfig proxy);
func Signal_connection_manager_set_proxy(
	connection_manager C.SignalConstPointerConnectionManager,
	proxy C.SignalConstPointerConnectionProxyConfig,
) *C.SignalFfiError {
	return C.signal_connection_manager_set_proxy(connection_manager, proxy)
}

// SignalFfiError* signal_connection_manager_set_remote_config(SignalConstPointerConnectionManager connection_manager, SignalMutPointerBridgedStringMap remote_config, uint8_t build_variant);
func Signal_connection_manager_set_remote_config(
	connection_manager C.SignalConstPointerConnectionManager,
	remote_config C.SignalMutPointerBridgedStringMap,
	build_variant C.uint8_t,
) *C.SignalFfiError {
	return C.signal_connection_manager_set_remote_config(connection_manager, remote_config, build_variant)
}

// SignalFfiError* signal_connection_proxy_config_clone(SignalMutPointerConnectionProxyConfig* new_obj, SignalConstPointerConnectionProxyConfig obj);
func Signal_connection_proxy_config_clone(
	new_obj *C.SignalMutPointerConnectionProxyConfig,
	obj C.SignalConstPointerConnectionProxyConfig,
) *C.SignalFfiError {
	return C.signal_connection_proxy_config_clone(new_obj, obj)
}

// SignalFfiError* signal_connection_proxy_config_destroy(SignalMutPointerConnectionProxyConfig p);
func Signal_connection_proxy_config_destroy(
	p C.SignalMutPointerConnectionProxyConfig,
) *C.SignalFfiError {
	return C.signal_connection_proxy_config_destroy(p)
}

// SignalFfiError* signal_connection_proxy_config_new(SignalMutPointerConnectionProxyConfig* out, const int8_t* scheme, const int8_t* host, int32_t port, const int8_t* username, const int8_t* password);
func Signal_connection_proxy_config_new(
	out *C.SignalMutPointerConnectionProxyConfig,
	scheme *C.int8_t,
	host *C.int8_t,
	port C.int32_t,
	username *C.int8_t,
	password *C.int8_t,
) *C.SignalFfiError {
	return C.signal_connection_proxy_config_new(out, scheme, host, port, username, password)
}

// SignalFfiError* signal_copy_backup_media_stream_cancel(SignalConstPointerCopyBackupMediaStream stream);
func Signal_copy_backup_media_stream_cancel(
	stream C.SignalConstPointerCopyBackupMediaStream,
) *C.SignalFfiError {
	return C.signal_copy_backup_media_stream_cancel(stream)
}

// SignalFfiError* signal_copy_backup_media_stream_destroy(SignalMutPointerCopyBackupMediaStream p);
func Signal_copy_backup_media_stream_destroy(
	p C.SignalMutPointerCopyBackupMediaStream,
) *C.SignalFfiError {
	return C.signal_copy_backup_media_stream_destroy(p)
}

// SignalFfiError* signal_copy_backup_media_stream_next(SignalCPromiseCopyBackupMediaNextChunkFfiResult* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerCopyBackupMediaStream stream);
func Signal_copy_backup_media_stream_next(
	promise *C.SignalCPromiseCopyBackupMediaNextChunkFfiResult,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	stream C.SignalConstPointerCopyBackupMediaStream,
) *C.SignalFfiError {
	return C.signal_copy_backup_media_stream_next(promise, async_runtime, stream)
}

// SignalFfiError* signal_create_call_link_credential_check_valid_contents(SignalBorrowedBuffer params_bytes);
func Signal_create_call_link_credential_check_valid_contents(
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_check_valid_contents(params_bytes)
}

// SignalFfiError* signal_create_call_link_credential_present_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer credential_bytes, SignalBorrowedBuffer room_id, const SignalType_FixedArray17_uint8_t* user_id, SignalBorrowedBuffer server_params_bytes, SignalBorrowedBuffer call_link_params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_create_call_link_credential_present_deterministic(
	out *C.SignalOwnedBuffer,
	credential_bytes C.SignalBorrowedBuffer,
	room_id C.SignalBorrowedBuffer,
	user_id *C.SignalType_FixedArray17_uint8_t,
	server_params_bytes C.SignalBorrowedBuffer,
	call_link_params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_present_deterministic(out, credential_bytes, room_id, user_id, server_params_bytes, call_link_params_bytes, randomness)
}

// SignalFfiError* signal_create_call_link_credential_presentation_check_valid_contents(SignalBorrowedBuffer presentation_bytes);
func Signal_create_call_link_credential_presentation_check_valid_contents(
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_presentation_check_valid_contents(presentation_bytes)
}

// SignalFfiError* signal_create_call_link_credential_presentation_verify(SignalBorrowedBuffer presentation_bytes, SignalBorrowedBuffer room_id, uint64_t now, SignalBorrowedBuffer server_params_bytes, SignalBorrowedBuffer call_link_params_bytes);
func Signal_create_call_link_credential_presentation_verify(
	presentation_bytes C.SignalBorrowedBuffer,
	room_id C.SignalBorrowedBuffer,
	now C.uint64_t,
	server_params_bytes C.SignalBorrowedBuffer,
	call_link_params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_presentation_verify(presentation_bytes, room_id, now, server_params_bytes, call_link_params_bytes)
}

// SignalFfiError* signal_create_call_link_credential_request_check_valid_contents(SignalBorrowedBuffer request_bytes);
func Signal_create_call_link_credential_request_check_valid_contents(
	request_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_request_check_valid_contents(request_bytes)
}

// SignalFfiError* signal_create_call_link_credential_request_context_check_valid_contents(SignalBorrowedBuffer context_bytes);
func Signal_create_call_link_credential_request_context_check_valid_contents(
	context_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_request_context_check_valid_contents(context_bytes)
}

// SignalFfiError* signal_create_call_link_credential_request_context_get_request(SignalOwnedBuffer* out, SignalBorrowedBuffer context_bytes);
func Signal_create_call_link_credential_request_context_get_request(
	out *C.SignalOwnedBuffer,
	context_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_request_context_get_request(out, context_bytes)
}

// SignalFfiError* signal_create_call_link_credential_request_context_new_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer room_id, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_create_call_link_credential_request_context_new_deterministic(
	out *C.SignalOwnedBuffer,
	room_id C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_request_context_new_deterministic(out, room_id, randomness)
}

// SignalFfiError* signal_create_call_link_credential_request_context_receive_response(SignalOwnedBuffer* out, SignalBorrowedBuffer context_bytes, SignalBorrowedBuffer response_bytes, const SignalType_FixedArray17_uint8_t* user_id, SignalBorrowedBuffer params_bytes);
func Signal_create_call_link_credential_request_context_receive_response(
	out *C.SignalOwnedBuffer,
	context_bytes C.SignalBorrowedBuffer,
	response_bytes C.SignalBorrowedBuffer,
	user_id *C.SignalType_FixedArray17_uint8_t,
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_request_context_receive_response(out, context_bytes, response_bytes, user_id, params_bytes)
}

// SignalFfiError* signal_create_call_link_credential_request_issue_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer request_bytes, const SignalType_FixedArray17_uint8_t* user_id, uint64_t timestamp, SignalBorrowedBuffer params_bytes, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_create_call_link_credential_request_issue_deterministic(
	out *C.SignalOwnedBuffer,
	request_bytes C.SignalBorrowedBuffer,
	user_id *C.SignalType_FixedArray17_uint8_t,
	timestamp C.uint64_t,
	params_bytes C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_request_issue_deterministic(out, request_bytes, user_id, timestamp, params_bytes, randomness)
}

// SignalFfiError* signal_create_call_link_credential_response_check_valid_contents(SignalBorrowedBuffer response_bytes);
func Signal_create_call_link_credential_response_check_valid_contents(
	response_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_create_call_link_credential_response_check_valid_contents(response_bytes)
}

// SignalFfiError* signal_decrypt_message(SignalOwnedBuffer* out, SignalConstPointerSignalMessage message, SignalConstPointerProtocolAddress protocol_address, SignalConstPointerProtocolAddress local_address, SignalConstPointerFfiSessionStoreStruct session_store, SignalConstPointerFfiIdentityKeyStoreStruct identity_key_store);
func Signal_decrypt_message(
	out *C.SignalOwnedBuffer,
	message C.SignalConstPointerSignalMessage,
	protocol_address C.SignalConstPointerProtocolAddress,
	local_address C.SignalConstPointerProtocolAddress,
	session_store C.SignalConstPointerFfiSessionStoreStruct,
	identity_key_store C.SignalConstPointerFfiIdentityKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_decrypt_message(out, message, protocol_address, local_address, session_store, identity_key_store)
}

// SignalFfiError* signal_decrypt_pre_key_message(SignalOwnedBuffer* out, SignalConstPointerPreKeySignalMessage message, SignalConstPointerProtocolAddress protocol_address, SignalConstPointerProtocolAddress local_address, SignalConstPointerFfiSessionStoreStruct session_store, SignalConstPointerFfiIdentityKeyStoreStruct identity_key_store, SignalConstPointerFfiPreKeyStoreStruct prekey_store, SignalConstPointerFfiSignedPreKeyStoreStruct signed_prekey_store, SignalConstPointerFfiKyberPreKeyStoreStruct kyber_prekey_store);
func Signal_decrypt_pre_key_message(
	out *C.SignalOwnedBuffer,
	message C.SignalConstPointerPreKeySignalMessage,
	protocol_address C.SignalConstPointerProtocolAddress,
	local_address C.SignalConstPointerProtocolAddress,
	session_store C.SignalConstPointerFfiSessionStoreStruct,
	identity_key_store C.SignalConstPointerFfiIdentityKeyStoreStruct,
	prekey_store C.SignalConstPointerFfiPreKeyStoreStruct,
	signed_prekey_store C.SignalConstPointerFfiSignedPreKeyStoreStruct,
	kyber_prekey_store C.SignalConstPointerFfiKyberPreKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_decrypt_pre_key_message(out, message, protocol_address, local_address, session_store, identity_key_store, prekey_store, signed_prekey_store, kyber_prekey_store)
}

// SignalFfiError* signal_decryption_error_message_clone(SignalMutPointerDecryptionErrorMessage* new_obj, SignalConstPointerDecryptionErrorMessage obj);
func Signal_decryption_error_message_clone(
	new_obj *C.SignalMutPointerDecryptionErrorMessage,
	obj C.SignalConstPointerDecryptionErrorMessage,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_clone(new_obj, obj)
}

// SignalFfiError* signal_decryption_error_message_deserialize(SignalMutPointerDecryptionErrorMessage* out, SignalBorrowedBuffer data);
func Signal_decryption_error_message_deserialize(
	out *C.SignalMutPointerDecryptionErrorMessage,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_deserialize(out, data)
}

// SignalFfiError* signal_decryption_error_message_destroy(SignalMutPointerDecryptionErrorMessage p);
func Signal_decryption_error_message_destroy(
	p C.SignalMutPointerDecryptionErrorMessage,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_destroy(p)
}

// SignalFfiError* signal_decryption_error_message_extract_from_serialized_content(SignalMutPointerDecryptionErrorMessage* out, SignalBorrowedBuffer bytes);
func Signal_decryption_error_message_extract_from_serialized_content(
	out *C.SignalMutPointerDecryptionErrorMessage,
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_extract_from_serialized_content(out, bytes)
}

// SignalFfiError* signal_decryption_error_message_for_original_message(SignalMutPointerDecryptionErrorMessage* out, SignalBorrowedBuffer original_bytes, uint8_t original_type, uint64_t original_timestamp, uint32_t original_sender_device_id);
func Signal_decryption_error_message_for_original_message(
	out *C.SignalMutPointerDecryptionErrorMessage,
	original_bytes C.SignalBorrowedBuffer,
	original_type C.uint8_t,
	original_timestamp C.uint64_t,
	original_sender_device_id C.uint32_t,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_for_original_message(out, original_bytes, original_type, original_timestamp, original_sender_device_id)
}

// SignalFfiError* signal_decryption_error_message_get_device_id(uint32_t* out, SignalConstPointerDecryptionErrorMessage obj);
func Signal_decryption_error_message_get_device_id(
	out *C.uint32_t,
	obj C.SignalConstPointerDecryptionErrorMessage,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_get_device_id(out, obj)
}

// SignalFfiError* signal_decryption_error_message_get_ratchet_key(SignalMutPointerPublicKey* out, SignalConstPointerDecryptionErrorMessage m);
func Signal_decryption_error_message_get_ratchet_key(
	out *C.SignalMutPointerPublicKey,
	m C.SignalConstPointerDecryptionErrorMessage,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_get_ratchet_key(out, m)
}

// SignalFfiError* signal_decryption_error_message_get_timestamp(uint64_t* out, SignalConstPointerDecryptionErrorMessage obj);
func Signal_decryption_error_message_get_timestamp(
	out *C.uint64_t,
	obj C.SignalConstPointerDecryptionErrorMessage,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_get_timestamp(out, obj)
}

// SignalFfiError* signal_decryption_error_message_serialize(SignalOwnedBuffer* out, SignalConstPointerDecryptionErrorMessage obj);
func Signal_decryption_error_message_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerDecryptionErrorMessage,
) *C.SignalFfiError {
	return C.signal_decryption_error_message_serialize(out, obj)
}

// SignalFfiError* signal_delete_backup_media_stream_cancel(SignalConstPointerDeleteBackupMediaStream stream);
func Signal_delete_backup_media_stream_cancel(
	stream C.SignalConstPointerDeleteBackupMediaStream,
) *C.SignalFfiError {
	return C.signal_delete_backup_media_stream_cancel(stream)
}

// SignalFfiError* signal_delete_backup_media_stream_destroy(SignalMutPointerDeleteBackupMediaStream p);
func Signal_delete_backup_media_stream_destroy(
	p C.SignalMutPointerDeleteBackupMediaStream,
) *C.SignalFfiError {
	return C.signal_delete_backup_media_stream_destroy(p)
}

// SignalFfiError* signal_delete_backup_media_stream_next(SignalCPromiseDeleteBackupMediaNextChunkFfiResult* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerDeleteBackupMediaStream stream);
func Signal_delete_backup_media_stream_next(
	promise *C.SignalCPromiseDeleteBackupMediaNextChunkFfiResult,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	stream C.SignalConstPointerDeleteBackupMediaStream,
) *C.SignalFfiError {
	return C.signal_delete_backup_media_stream_next(promise, async_runtime, stream)
}

// SignalFfiError* signal_device_transfer_generate_certificate(SignalOwnedBuffer* out, SignalBorrowedBuffer private_key, const int8_t* name, uint32_t days_to_expire);
func Signal_device_transfer_generate_certificate(
	out *C.SignalOwnedBuffer,
	private_key C.SignalBorrowedBuffer,
	name *C.int8_t,
	days_to_expire C.uint32_t,
) *C.SignalFfiError {
	return C.signal_device_transfer_generate_certificate(out, private_key, name, days_to_expire)
}

// SignalFfiError* signal_device_transfer_generate_private_key(SignalOwnedBuffer* out);
func Signal_device_transfer_generate_private_key(
	out *C.SignalOwnedBuffer,
) *C.SignalFfiError {
	return C.signal_device_transfer_generate_private_key(out)
}

// SignalFfiError* signal_device_transfer_generate_private_key_with_format(SignalOwnedBuffer* out, uint8_t key_format);
func Signal_device_transfer_generate_private_key_with_format(
	out *C.SignalOwnedBuffer,
	key_format C.uint8_t,
) *C.SignalFfiError {
	return C.signal_device_transfer_generate_private_key_with_format(out, key_format)
}

// SignalFfiError* signal_donation_permit_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_donation_permit_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_check_valid_contents(buffer)
}

// SignalFfiError* signal_donation_permit_derived_key_pair_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_donation_permit_derived_key_pair_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_derived_key_pair_check_valid_contents(buffer)
}

// SignalFfiError* signal_donation_permit_derived_key_pair_for_expiration(SignalOwnedBuffer* out, uint64_t timestamp, SignalConstPointerServerSecretParams root);
func Signal_donation_permit_derived_key_pair_for_expiration(
	out *C.SignalOwnedBuffer,
	timestamp C.uint64_t,
	root C.SignalConstPointerServerSecretParams,
) *C.SignalFfiError {
	return C.signal_donation_permit_derived_key_pair_for_expiration(out, timestamp, root)
}

// SignalFfiError* signal_donation_permit_expiration(uint64_t* out, SignalBorrowedBuffer donation_permit);
func Signal_donation_permit_expiration(
	out *C.uint64_t,
	donation_permit C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_expiration(out, donation_permit)
}

// SignalFfiError* signal_donation_permit_request_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_donation_permit_request_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_request_check_valid_contents(buffer)
}

// SignalFfiError* signal_donation_permit_request_context_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_donation_permit_request_context_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_request_context_check_valid_contents(buffer)
}

// SignalFfiError* signal_donation_permit_request_context_new_deterministic(SignalOwnedBuffer* out, int32_t count, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_donation_permit_request_context_new_deterministic(
	out *C.SignalOwnedBuffer,
	count C.int32_t,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_donation_permit_request_context_new_deterministic(out, count, randomness)
}

// SignalFfiError* signal_donation_permit_request_context_receive(SignalBytestringArray* out, SignalBorrowedBuffer context, SignalBorrowedBuffer response, SignalConstPointerServerPublicParams public_params, uint64_t now);
func Signal_donation_permit_request_context_receive(
	out *C.SignalBytestringArray,
	context C.SignalBorrowedBuffer,
	response C.SignalBorrowedBuffer,
	public_params C.SignalConstPointerServerPublicParams,
	now C.uint64_t,
) *C.SignalFfiError {
	return C.signal_donation_permit_request_context_receive(out, context, response, public_params, now)
}

// SignalFfiError* signal_donation_permit_request_context_request(SignalOwnedBuffer* out, SignalBorrowedBuffer ctx);
func Signal_donation_permit_request_context_request(
	out *C.SignalOwnedBuffer,
	ctx C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_request_context_request(out, ctx)
}

// SignalFfiError* signal_donation_permit_request_len(int32_t* out, SignalBorrowedBuffer donation_permit_request);
func Signal_donation_permit_request_len(
	out *C.int32_t,
	donation_permit_request C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_request_len(out, donation_permit_request)
}

// SignalFfiError* signal_donation_permit_response_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_donation_permit_response_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_response_check_valid_contents(buffer)
}

// SignalFfiError* signal_donation_permit_response_default_expiration(uint64_t* out, uint64_t current_time);
func Signal_donation_permit_response_default_expiration(
	out *C.uint64_t,
	current_time C.uint64_t,
) *C.SignalFfiError {
	return C.signal_donation_permit_response_default_expiration(out, current_time)
}

// SignalFfiError* signal_donation_permit_response_get_expiration(uint64_t* out, SignalBorrowedBuffer response);
func Signal_donation_permit_response_get_expiration(
	out *C.uint64_t,
	response C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_response_get_expiration(out, response)
}

// SignalFfiError* signal_donation_permit_response_issue_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer request, SignalBorrowedBuffer key_pair, const SignalType_FixedArray32_uint8_t* seed);
func Signal_donation_permit_response_issue_deterministic(
	out *C.SignalOwnedBuffer,
	request C.SignalBorrowedBuffer,
	key_pair C.SignalBorrowedBuffer,
	seed *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_donation_permit_response_issue_deterministic(out, request, key_pair, seed)
}

// SignalFfiError* signal_donation_permit_spend_id(SignalOwnedBuffer* out, SignalBorrowedBuffer donation_permit);
func Signal_donation_permit_spend_id(
	out *C.SignalOwnedBuffer,
	donation_permit C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_spend_id(out, donation_permit)
}

// SignalFfiError* signal_donation_permit_verify(SignalBorrowedBuffer permit, uint64_t now, SignalBorrowedBuffer key_pair);
func Signal_donation_permit_verify(
	permit C.SignalBorrowedBuffer,
	now C.uint64_t,
	key_pair C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_donation_permit_verify(permit, now, key_pair)
}

// SignalFfiError* signal_encrypt_message(SignalMutPointerCiphertextMessage* out, SignalBorrowedBuffer ptext, SignalConstPointerProtocolAddress protocol_address, SignalConstPointerProtocolAddress local_address, SignalConstPointerFfiSessionStoreStruct session_store, SignalConstPointerFfiIdentityKeyStoreStruct identity_key_store, uint64_t now);
func Signal_encrypt_message(
	out *C.SignalMutPointerCiphertextMessage,
	ptext C.SignalBorrowedBuffer,
	protocol_address C.SignalConstPointerProtocolAddress,
	local_address C.SignalConstPointerProtocolAddress,
	session_store C.SignalConstPointerFfiSessionStoreStruct,
	identity_key_store C.SignalConstPointerFfiIdentityKeyStoreStruct,
	now C.uint64_t,
) *C.SignalFfiError {
	return C.signal_encrypt_message(out, ptext, protocol_address, local_address, session_store, identity_key_store, now)
}

// SignalFfiError* signal_error_get_address(SignalMutPointerProtocolAddress* out, const SignalFfiError* err);
func Signal_error_get_address(
	out *C.SignalMutPointerProtocolAddress,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_address(out, err)
}

// SignalFfiError* signal_error_get_invalid_protocol_address(SignalPairOfCStringPtru32* out, const SignalFfiError* err);
func Signal_error_get_invalid_protocol_address(
	out *C.SignalPairOfCStringPtru32,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_invalid_protocol_address(out, err)
}

// SignalFfiError* signal_error_get_message(SignalCStringPtr* out, const SignalFfiError* err);
func Signal_error_get_message(
	out *C.SignalCStringPtr,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_message(out, err)
}

// SignalFfiError* signal_error_get_mismatched_device_errors(SignalOwnedBufferOfFfiMismatchedDevicesError* out, const SignalFfiError* err);
func Signal_error_get_mismatched_device_errors(
	out *C.SignalOwnedBufferOfFfiMismatchedDevicesError,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_mismatched_device_errors(out, err)
}

// SignalFfiError* signal_error_get_our_fingerprint_version(uint32_t* out, const SignalFfiError* err);
func Signal_error_get_our_fingerprint_version(
	out *C.uint32_t,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_our_fingerprint_version(out, err)
}

// SignalFfiError* signal_error_get_rate_limit_challenge(SignalPairOfPairOfCStringPtrOwnedBufferi64* out, const SignalFfiError* err);
func Signal_error_get_rate_limit_challenge(
	out *C.SignalPairOfPairOfCStringPtrOwnedBufferi64,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_rate_limit_challenge(out, err)
}

// SignalFfiError* signal_error_get_registration_error_not_deliverable(SignalPairOfCStringPtrbool* out, const SignalFfiError* err);
func Signal_error_get_registration_error_not_deliverable(
	out *C.SignalPairOfCStringPtrbool,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_registration_error_not_deliverable(out, err)
}

// SignalFfiError* signal_error_get_registration_lock(uint64_t* out_time_remaining_seconds, SignalPairOfCStringPtrCStringPtr* out_svr2_credentials, const SignalFfiError* err);
func Signal_error_get_registration_lock(
	out_time_remaining_seconds *C.uint64_t,
	out_svr2_credentials *C.SignalPairOfCStringPtrCStringPtr,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_registration_lock(out_time_remaining_seconds, out_svr2_credentials, err)
}

// SignalFfiError* signal_error_get_retry_after_seconds(uint32_t* out, const SignalFfiError* err);
func Signal_error_get_retry_after_seconds(
	out *C.uint32_t,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_retry_after_seconds(out, err)
}

// SignalFfiError* signal_error_get_their_fingerprint_version(uint32_t* out, const SignalFfiError* err);
func Signal_error_get_their_fingerprint_version(
	out *C.uint32_t,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_their_fingerprint_version(out, err)
}

// SignalFfiError* signal_error_get_tries_remaining(uint32_t* out, const SignalFfiError* err);
func Signal_error_get_tries_remaining(
	out *C.uint32_t,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_tries_remaining(out, err)
}

// SignalFfiError* signal_error_get_unknown_fields(SignalBytestringArray* out, const SignalFfiError* err);
func Signal_error_get_unknown_fields(
	out *C.SignalBytestringArray,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_unknown_fields(out, err)
}

// SignalFfiError* signal_error_get_uuid(SignalUuid* out, const SignalFfiError* err);
func Signal_error_get_uuid(
	out *C.SignalUuid,
	err *C.SignalFfiError,
) *C.SignalFfiError {
	return C.signal_error_get_uuid(out, err)
}

// SignalFfiError* signal_expiring_profile_key_credential_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_expiring_profile_key_credential_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_expiring_profile_key_credential_check_valid_contents(buffer)
}

// SignalFfiError* signal_expiring_profile_key_credential_get_expiration_time(uint64_t* out, const SignalType_FixedArray153_uint8_t* credential);
func Signal_expiring_profile_key_credential_get_expiration_time(
	out *C.uint64_t,
	credential *C.SignalType_FixedArray153_uint8_t,
) *C.SignalFfiError {
	return C.signal_expiring_profile_key_credential_get_expiration_time(out, credential)
}

// SignalFfiError* signal_expiring_profile_key_credential_response_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_expiring_profile_key_credential_response_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_expiring_profile_key_credential_response_check_valid_contents(buffer)
}

// SignalFfiError* signal_fingerprint_clone(SignalMutPointerFingerprint* new_obj, SignalConstPointerFingerprint obj);
func Signal_fingerprint_clone(
	new_obj *C.SignalMutPointerFingerprint,
	obj C.SignalConstPointerFingerprint,
) *C.SignalFfiError {
	return C.signal_fingerprint_clone(new_obj, obj)
}

// SignalFfiError* signal_fingerprint_compare(bool* out, SignalBorrowedBuffer fprint1, SignalBorrowedBuffer fprint2);
func Signal_fingerprint_compare(
	out *C.bool,
	fprint1 C.SignalBorrowedBuffer,
	fprint2 C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_fingerprint_compare(out, fprint1, fprint2)
}

// SignalFfiError* signal_fingerprint_destroy(SignalMutPointerFingerprint p);
func Signal_fingerprint_destroy(
	p C.SignalMutPointerFingerprint,
) *C.SignalFfiError {
	return C.signal_fingerprint_destroy(p)
}

// SignalFfiError* signal_fingerprint_display_string(SignalCStringPtr* out, SignalConstPointerFingerprint obj);
func Signal_fingerprint_display_string(
	out *C.SignalCStringPtr,
	obj C.SignalConstPointerFingerprint,
) *C.SignalFfiError {
	return C.signal_fingerprint_display_string(out, obj)
}

// SignalFfiError* signal_fingerprint_new(SignalMutPointerFingerprint* out, uint32_t iterations, uint32_t version, SignalBorrowedBuffer local_identifier, SignalConstPointerPublicKey local_key, SignalBorrowedBuffer remote_identifier, SignalConstPointerPublicKey remote_key);
func Signal_fingerprint_new(
	out *C.SignalMutPointerFingerprint,
	iterations C.uint32_t,
	version C.uint32_t,
	local_identifier C.SignalBorrowedBuffer,
	local_key C.SignalConstPointerPublicKey,
	remote_identifier C.SignalBorrowedBuffer,
	remote_key C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_fingerprint_new(out, iterations, version, local_identifier, local_key, remote_identifier, remote_key)
}

// SignalFfiError* signal_fingerprint_scannable_encoding(SignalOwnedBuffer* out, SignalConstPointerFingerprint obj);
func Signal_fingerprint_scannable_encoding(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerFingerprint,
) *C.SignalFfiError {
	return C.signal_fingerprint_scannable_encoding(out, obj)
}

// SignalFfiError* signal_generic_server_public_params_check_valid_contents(SignalBorrowedBuffer params_bytes);
func Signal_generic_server_public_params_check_valid_contents(
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_generic_server_public_params_check_valid_contents(params_bytes)
}

// SignalFfiError* signal_generic_server_secret_params_check_valid_contents(SignalBorrowedBuffer params_bytes);
func Signal_generic_server_secret_params_check_valid_contents(
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_generic_server_secret_params_check_valid_contents(params_bytes)
}

// SignalFfiError* signal_generic_server_secret_params_generate_deterministic(SignalOwnedBuffer* out, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_generic_server_secret_params_generate_deterministic(
	out *C.SignalOwnedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_generic_server_secret_params_generate_deterministic(out, randomness)
}

// SignalFfiError* signal_generic_server_secret_params_get_public_params(SignalOwnedBuffer* out, SignalBorrowedBuffer params_bytes);
func Signal_generic_server_secret_params_get_public_params(
	out *C.SignalOwnedBuffer,
	params_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_generic_server_secret_params_get_public_params(out, params_bytes)
}

// SignalFfiError* signal_group_decrypt_message(SignalOwnedBuffer* out, SignalConstPointerProtocolAddress sender, SignalBorrowedBuffer message, SignalConstPointerFfiSenderKeyStoreStruct store);
func Signal_group_decrypt_message(
	out *C.SignalOwnedBuffer,
	sender C.SignalConstPointerProtocolAddress,
	message C.SignalBorrowedBuffer,
	store C.SignalConstPointerFfiSenderKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_group_decrypt_message(out, sender, message, store)
}

// SignalFfiError* signal_group_encrypt_message(SignalMutPointerCiphertextMessage* out, SignalConstPointerProtocolAddress sender, SignalUuid distribution_id, SignalBorrowedBuffer message, SignalConstPointerFfiSenderKeyStoreStruct store);
func Signal_group_encrypt_message(
	out *C.SignalMutPointerCiphertextMessage,
	sender C.SignalConstPointerProtocolAddress,
	distribution_id C.SignalUuid,
	message C.SignalBorrowedBuffer,
	store C.SignalConstPointerFfiSenderKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_group_encrypt_message(out, sender, distribution_id, message, store)
}

// SignalFfiError* signal_group_master_key_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_group_master_key_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_master_key_check_valid_contents(buffer)
}

// SignalFfiError* signal_group_public_params_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_group_public_params_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_public_params_check_valid_contents(buffer)
}

// SignalFfiError* signal_group_public_params_get_group_identifier(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray97_uint8_t* group_public_params);
func Signal_group_public_params_get_group_identifier(
	out *C.SignalType_FixedArray32_uint8_t,
	group_public_params *C.SignalType_FixedArray97_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_public_params_get_group_identifier(out, group_public_params)
}

// SignalFfiError* signal_group_secret_params_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_group_secret_params_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_secret_params_check_valid_contents(buffer)
}

// SignalFfiError* signal_group_secret_params_decrypt_blob_with_padding(SignalOwnedBuffer* out, const SignalType_FixedArray289_uint8_t* params, SignalBorrowedBuffer ciphertext);
func Signal_group_secret_params_decrypt_blob_with_padding(
	out *C.SignalOwnedBuffer,
	params *C.SignalType_FixedArray289_uint8_t,
	ciphertext C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_secret_params_decrypt_blob_with_padding(out, params, ciphertext)
}

// SignalFfiError* signal_group_secret_params_decrypt_profile_key(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray289_uint8_t* params, const SignalType_FixedArray65_uint8_t* profile_key, const SignalType_FixedArray17_uint8_t* user_id);
func Signal_group_secret_params_decrypt_profile_key(
	out *C.SignalType_FixedArray32_uint8_t,
	params *C.SignalType_FixedArray289_uint8_t,
	profile_key *C.SignalType_FixedArray65_uint8_t,
	user_id *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_decrypt_profile_key(out, params, profile_key, user_id)
}

// SignalFfiError* signal_group_secret_params_decrypt_service_id(SignalType_FixedArray17_uint8_t* out, const SignalType_FixedArray289_uint8_t* params, const SignalType_FixedArray65_uint8_t* ciphertext);
func Signal_group_secret_params_decrypt_service_id(
	out *C.SignalType_FixedArray17_uint8_t,
	params *C.SignalType_FixedArray289_uint8_t,
	ciphertext *C.SignalType_FixedArray65_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_decrypt_service_id(out, params, ciphertext)
}

// SignalFfiError* signal_group_secret_params_derive_from_master_key(SignalType_FixedArray289_uint8_t* out, const SignalType_FixedArray32_uint8_t* master_key);
func Signal_group_secret_params_derive_from_master_key(
	out *C.SignalType_FixedArray289_uint8_t,
	master_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_derive_from_master_key(out, master_key)
}

// SignalFfiError* signal_group_secret_params_encrypt_blob_with_padding_deterministic(SignalOwnedBuffer* out, const SignalType_FixedArray289_uint8_t* params, const SignalType_FixedArray32_uint8_t* randomness, SignalBorrowedBuffer plaintext, uint32_t padding_len);
func Signal_group_secret_params_encrypt_blob_with_padding_deterministic(
	out *C.SignalOwnedBuffer,
	params *C.SignalType_FixedArray289_uint8_t,
	randomness *C.SignalType_FixedArray32_uint8_t,
	plaintext C.SignalBorrowedBuffer,
	padding_len C.uint32_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_encrypt_blob_with_padding_deterministic(out, params, randomness, plaintext, padding_len)
}

// SignalFfiError* signal_group_secret_params_encrypt_profile_key(SignalType_FixedArray65_uint8_t* out, const SignalType_FixedArray289_uint8_t* params, const SignalType_FixedArray32_uint8_t* profile_key, const SignalType_FixedArray17_uint8_t* user_id);
func Signal_group_secret_params_encrypt_profile_key(
	out *C.SignalType_FixedArray65_uint8_t,
	params *C.SignalType_FixedArray289_uint8_t,
	profile_key *C.SignalType_FixedArray32_uint8_t,
	user_id *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_encrypt_profile_key(out, params, profile_key, user_id)
}

// SignalFfiError* signal_group_secret_params_encrypt_service_id(SignalType_FixedArray65_uint8_t* out, const SignalType_FixedArray289_uint8_t* params, const SignalType_FixedArray17_uint8_t* service_id);
func Signal_group_secret_params_encrypt_service_id(
	out *C.SignalType_FixedArray65_uint8_t,
	params *C.SignalType_FixedArray289_uint8_t,
	service_id *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_encrypt_service_id(out, params, service_id)
}

// SignalFfiError* signal_group_secret_params_generate_deterministic(SignalType_FixedArray289_uint8_t* out, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_group_secret_params_generate_deterministic(
	out *C.SignalType_FixedArray289_uint8_t,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_generate_deterministic(out, randomness)
}

// SignalFfiError* signal_group_secret_params_get_master_key(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray289_uint8_t* params);
func Signal_group_secret_params_get_master_key(
	out *C.SignalType_FixedArray32_uint8_t,
	params *C.SignalType_FixedArray289_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_get_master_key(out, params)
}

// SignalFfiError* signal_group_secret_params_get_public_params(SignalType_FixedArray97_uint8_t* out, const SignalType_FixedArray289_uint8_t* params);
func Signal_group_secret_params_get_public_params(
	out *C.SignalType_FixedArray97_uint8_t,
	params *C.SignalType_FixedArray289_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_secret_params_get_public_params(out, params)
}

// SignalFfiError* signal_group_send_derived_key_pair_check_valid_contents(SignalBorrowedBuffer bytes);
func Signal_group_send_derived_key_pair_check_valid_contents(
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_derived_key_pair_check_valid_contents(bytes)
}

// SignalFfiError* signal_group_send_derived_key_pair_for_expiration(SignalOwnedBuffer* out, uint64_t expiration, SignalConstPointerServerSecretParams server_params);
func Signal_group_send_derived_key_pair_for_expiration(
	out *C.SignalOwnedBuffer,
	expiration C.uint64_t,
	server_params C.SignalConstPointerServerSecretParams,
) *C.SignalFfiError {
	return C.signal_group_send_derived_key_pair_for_expiration(out, expiration, server_params)
}

// SignalFfiError* signal_group_send_endorsement_call_link_params_to_token(SignalOwnedBuffer* out, SignalBorrowedBuffer endorsement, SignalBorrowedBuffer call_link_secret_params_serialized);
func Signal_group_send_endorsement_call_link_params_to_token(
	out *C.SignalOwnedBuffer,
	endorsement C.SignalBorrowedBuffer,
	call_link_secret_params_serialized C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_endorsement_call_link_params_to_token(out, endorsement, call_link_secret_params_serialized)
}

// SignalFfiError* signal_group_send_endorsement_check_valid_contents(SignalBorrowedBuffer bytes);
func Signal_group_send_endorsement_check_valid_contents(
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_endorsement_check_valid_contents(bytes)
}

// SignalFfiError* signal_group_send_endorsement_combine(SignalOwnedBuffer* out, SignalBorrowedSliceOfBuffers endorsements);
func Signal_group_send_endorsement_combine(
	out *C.SignalOwnedBuffer,
	endorsements C.SignalBorrowedSliceOfBuffers,
) *C.SignalFfiError {
	return C.signal_group_send_endorsement_combine(out, endorsements)
}

// SignalFfiError* signal_group_send_endorsement_remove(SignalOwnedBuffer* out, SignalBorrowedBuffer endorsement, SignalBorrowedBuffer to_remove);
func Signal_group_send_endorsement_remove(
	out *C.SignalOwnedBuffer,
	endorsement C.SignalBorrowedBuffer,
	to_remove C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_endorsement_remove(out, endorsement, to_remove)
}

// SignalFfiError* signal_group_send_endorsement_to_token(SignalOwnedBuffer* out, SignalBorrowedBuffer endorsement, const SignalType_FixedArray289_uint8_t* group_params);
func Signal_group_send_endorsement_to_token(
	out *C.SignalOwnedBuffer,
	endorsement C.SignalBorrowedBuffer,
	group_params *C.SignalType_FixedArray289_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_send_endorsement_to_token(out, endorsement, group_params)
}

// SignalFfiError* signal_group_send_endorsements_response_check_valid_contents(SignalBorrowedBuffer bytes);
func Signal_group_send_endorsements_response_check_valid_contents(
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_endorsements_response_check_valid_contents(bytes)
}

// SignalFfiError* signal_group_send_endorsements_response_get_expiration(uint64_t* out, SignalBorrowedBuffer response_bytes);
func Signal_group_send_endorsements_response_get_expiration(
	out *C.uint64_t,
	response_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_endorsements_response_get_expiration(out, response_bytes)
}

// SignalFfiError* signal_group_send_endorsements_response_issue_deterministic(SignalOwnedBuffer* out, SignalBorrowedBuffer concatenated_group_member_ciphertexts, SignalBorrowedBuffer key_pair, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_group_send_endorsements_response_issue_deterministic(
	out *C.SignalOwnedBuffer,
	concatenated_group_member_ciphertexts C.SignalBorrowedBuffer,
	key_pair C.SignalBorrowedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_group_send_endorsements_response_issue_deterministic(out, concatenated_group_member_ciphertexts, key_pair, randomness)
}

// SignalFfiError* signal_group_send_endorsements_response_receive_and_combine_with_ciphertexts(SignalBytestringArray* out, SignalBorrowedBuffer response_bytes, SignalBorrowedBuffer concatenated_group_member_ciphertexts, SignalBorrowedBuffer local_user_ciphertext, uint64_t now, SignalConstPointerServerPublicParams server_params);
func Signal_group_send_endorsements_response_receive_and_combine_with_ciphertexts(
	out *C.SignalBytestringArray,
	response_bytes C.SignalBorrowedBuffer,
	concatenated_group_member_ciphertexts C.SignalBorrowedBuffer,
	local_user_ciphertext C.SignalBorrowedBuffer,
	now C.uint64_t,
	server_params C.SignalConstPointerServerPublicParams,
) *C.SignalFfiError {
	return C.signal_group_send_endorsements_response_receive_and_combine_with_ciphertexts(out, response_bytes, concatenated_group_member_ciphertexts, local_user_ciphertext, now, server_params)
}

// SignalFfiError* signal_group_send_endorsements_response_receive_and_combine_with_service_ids(SignalBytestringArray* out, SignalBorrowedBuffer response_bytes, SignalBorrowedBuffer group_members, const SignalType_FixedArray17_uint8_t* local_user, uint64_t now, const SignalType_FixedArray289_uint8_t* group_params, SignalConstPointerServerPublicParams server_params);
func Signal_group_send_endorsements_response_receive_and_combine_with_service_ids(
	out *C.SignalBytestringArray,
	response_bytes C.SignalBorrowedBuffer,
	group_members C.SignalBorrowedBuffer,
	local_user *C.SignalType_FixedArray17_uint8_t,
	now C.uint64_t,
	group_params *C.SignalType_FixedArray289_uint8_t,
	server_params C.SignalConstPointerServerPublicParams,
) *C.SignalFfiError {
	return C.signal_group_send_endorsements_response_receive_and_combine_with_service_ids(out, response_bytes, group_members, local_user, now, group_params, server_params)
}

// SignalFfiError* signal_group_send_full_token_check_valid_contents(SignalBorrowedBuffer bytes);
func Signal_group_send_full_token_check_valid_contents(
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_full_token_check_valid_contents(bytes)
}

// SignalFfiError* signal_group_send_full_token_get_expiration(uint64_t* out, SignalBorrowedBuffer token);
func Signal_group_send_full_token_get_expiration(
	out *C.uint64_t,
	token C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_full_token_get_expiration(out, token)
}

// SignalFfiError* signal_group_send_full_token_verify(SignalBorrowedBuffer token, SignalBorrowedBuffer user_ids, uint64_t now, SignalBorrowedBuffer key_pair);
func Signal_group_send_full_token_verify(
	token C.SignalBorrowedBuffer,
	user_ids C.SignalBorrowedBuffer,
	now C.uint64_t,
	key_pair C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_full_token_verify(token, user_ids, now, key_pair)
}

// SignalFfiError* signal_group_send_token_check_valid_contents(SignalBorrowedBuffer bytes);
func Signal_group_send_token_check_valid_contents(
	bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_group_send_token_check_valid_contents(bytes)
}

// SignalFfiError* signal_group_send_token_to_full_token(SignalOwnedBuffer* out, SignalBorrowedBuffer token, uint64_t expiration);
func Signal_group_send_token_to_full_token(
	out *C.SignalOwnedBuffer,
	token C.SignalBorrowedBuffer,
	expiration C.uint64_t,
) *C.SignalFfiError {
	return C.signal_group_send_token_to_full_token(out, token, expiration)
}

// SignalFfiError* signal_hex_encode(SignalBorrowedMutableBuffer output, SignalBorrowedBuffer input);
func Signal_hex_encode(
	output C.SignalBorrowedMutableBuffer,
	input C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_hex_encode(output, input)
}

// SignalFfiError* signal_hkdf_derive(SignalBorrowedMutableBuffer output, SignalBorrowedBuffer ikm, SignalBorrowedBuffer label, SignalBorrowedBuffer salt);
func Signal_hkdf_derive(
	output C.SignalBorrowedMutableBuffer,
	ikm C.SignalBorrowedBuffer,
	label C.SignalBorrowedBuffer,
	salt C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_hkdf_derive(output, ikm, label, salt)
}

// SignalFfiError* signal_hsm_enclave_client_complete_handshake(SignalMutPointerHsmEnclaveClient cli, SignalBorrowedBuffer handshake_received);
func Signal_hsm_enclave_client_complete_handshake(
	cli C.SignalMutPointerHsmEnclaveClient,
	handshake_received C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_hsm_enclave_client_complete_handshake(cli, handshake_received)
}

// SignalFfiError* signal_hsm_enclave_client_destroy(SignalMutPointerHsmEnclaveClient p);
func Signal_hsm_enclave_client_destroy(
	p C.SignalMutPointerHsmEnclaveClient,
) *C.SignalFfiError {
	return C.signal_hsm_enclave_client_destroy(p)
}

// SignalFfiError* signal_hsm_enclave_client_established_recv(SignalOwnedBuffer* out, SignalMutPointerHsmEnclaveClient cli, SignalBorrowedBuffer received_ciphertext);
func Signal_hsm_enclave_client_established_recv(
	out *C.SignalOwnedBuffer,
	cli C.SignalMutPointerHsmEnclaveClient,
	received_ciphertext C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_hsm_enclave_client_established_recv(out, cli, received_ciphertext)
}

// SignalFfiError* signal_hsm_enclave_client_established_send(SignalOwnedBuffer* out, SignalMutPointerHsmEnclaveClient cli, SignalBorrowedBuffer plaintext_to_send);
func Signal_hsm_enclave_client_established_send(
	out *C.SignalOwnedBuffer,
	cli C.SignalMutPointerHsmEnclaveClient,
	plaintext_to_send C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_hsm_enclave_client_established_send(out, cli, plaintext_to_send)
}

// SignalFfiError* signal_hsm_enclave_client_initial_request(SignalOwnedBuffer* out, SignalConstPointerHsmEnclaveClient obj);
func Signal_hsm_enclave_client_initial_request(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerHsmEnclaveClient,
) *C.SignalFfiError {
	return C.signal_hsm_enclave_client_initial_request(out, obj)
}

// SignalFfiError* signal_hsm_enclave_client_new(SignalMutPointerHsmEnclaveClient* out, SignalBorrowedBuffer trusted_public_key, SignalBorrowedBuffer trusted_code_hashes);
func Signal_hsm_enclave_client_new(
	out *C.SignalMutPointerHsmEnclaveClient,
	trusted_public_key C.SignalBorrowedBuffer,
	trusted_code_hashes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_hsm_enclave_client_new(out, trusted_public_key, trusted_code_hashes)
}

// SignalFfiError* signal_http_request_add_header(SignalConstPointerHttpRequest request, const int8_t* name, const int8_t* value);
func Signal_http_request_add_header(
	request C.SignalConstPointerHttpRequest,
	name *C.int8_t,
	value *C.int8_t,
) *C.SignalFfiError {
	return C.signal_http_request_add_header(request, name, value)
}

// SignalFfiError* signal_http_request_destroy(SignalMutPointerHttpRequest p);
func Signal_http_request_destroy(
	p C.SignalMutPointerHttpRequest,
) *C.SignalFfiError {
	return C.signal_http_request_destroy(p)
}

// SignalFfiError* signal_http_request_new_with_body(SignalMutPointerHttpRequest* out, const int8_t* method, const int8_t* path, SignalBorrowedBuffer body_as_slice);
func Signal_http_request_new_with_body(
	out *C.SignalMutPointerHttpRequest,
	method *C.int8_t,
	path *C.int8_t,
	body_as_slice C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_http_request_new_with_body(out, method, path, body_as_slice)
}

// SignalFfiError* signal_http_request_new_without_body(SignalMutPointerHttpRequest* out, const int8_t* method, const int8_t* path);
func Signal_http_request_new_without_body(
	out *C.SignalMutPointerHttpRequest,
	method *C.int8_t,
	path *C.int8_t,
) *C.SignalFfiError {
	return C.signal_http_request_new_without_body(out, method, path)
}

// SignalFfiError* signal_identitykey_verify_alternate_identity(bool* out, SignalConstPointerPublicKey public_key, SignalConstPointerPublicKey other_identity, SignalBorrowedBuffer signature);
func Signal_identitykey_verify_alternate_identity(
	out *C.bool,
	public_key C.SignalConstPointerPublicKey,
	other_identity C.SignalConstPointerPublicKey,
	signature C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_identitykey_verify_alternate_identity(out, public_key, other_identity, signature)
}

// SignalFfiError* signal_identitykeypair_deserialize(SignalPairOfMutPointerPublicKeyMutPointerPrivateKey* out, SignalBorrowedBuffer input);
func Signal_identitykeypair_deserialize(
	out *C.SignalPairOfMutPointerPublicKeyMutPointerPrivateKey,
	input C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_identitykeypair_deserialize(out, input)
}

// SignalFfiError* signal_identitykeypair_serialize(SignalOwnedBuffer* out, SignalConstPointerPublicKey public_key, SignalConstPointerPrivateKey private_key);
func Signal_identitykeypair_serialize(
	out *C.SignalOwnedBuffer,
	public_key C.SignalConstPointerPublicKey,
	private_key C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_identitykeypair_serialize(out, public_key, private_key)
}

// SignalFfiError* signal_identitykeypair_sign_alternate_identity(SignalOwnedBuffer* out, SignalConstPointerPublicKey public_key, SignalConstPointerPrivateKey private_key, SignalConstPointerPublicKey other_identity);
func Signal_identitykeypair_sign_alternate_identity(
	out *C.SignalOwnedBuffer,
	public_key C.SignalConstPointerPublicKey,
	private_key C.SignalConstPointerPrivateKey,
	other_identity C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_identitykeypair_sign_alternate_identity(out, public_key, private_key, other_identity)
}

// SignalFfiError* signal_incremental_mac_calculate_chunk_size(uint32_t* out, uint32_t data_size);
func Signal_incremental_mac_calculate_chunk_size(
	out *C.uint32_t,
	data_size C.uint32_t,
) *C.SignalFfiError {
	return C.signal_incremental_mac_calculate_chunk_size(out, data_size)
}

// SignalFfiError* signal_incremental_mac_destroy(SignalMutPointerIncrementalMac p);
func Signal_incremental_mac_destroy(
	p C.SignalMutPointerIncrementalMac,
) *C.SignalFfiError {
	return C.signal_incremental_mac_destroy(p)
}

// SignalFfiError* signal_incremental_mac_finalize(SignalOwnedBuffer* out, SignalMutPointerIncrementalMac mac);
func Signal_incremental_mac_finalize(
	out *C.SignalOwnedBuffer,
	mac C.SignalMutPointerIncrementalMac,
) *C.SignalFfiError {
	return C.signal_incremental_mac_finalize(out, mac)
}

// SignalFfiError* signal_incremental_mac_initialize(SignalMutPointerIncrementalMac* out, SignalBorrowedBuffer key, uint32_t chunk_size);
func Signal_incremental_mac_initialize(
	out *C.SignalMutPointerIncrementalMac,
	key C.SignalBorrowedBuffer,
	chunk_size C.uint32_t,
) *C.SignalFfiError {
	return C.signal_incremental_mac_initialize(out, key, chunk_size)
}

// SignalFfiError* signal_incremental_mac_update(SignalOwnedBuffer* out, SignalMutPointerIncrementalMac mac, SignalBorrowedBuffer bytes, uint32_t offset, uint32_t length);
func Signal_incremental_mac_update(
	out *C.SignalOwnedBuffer,
	mac C.SignalMutPointerIncrementalMac,
	bytes C.SignalBorrowedBuffer,
	offset C.uint32_t,
	length C.uint32_t,
) *C.SignalFfiError {
	return C.signal_incremental_mac_update(out, mac, bytes, offset, length)
}

// SignalFfiError* signal_key_transparency_aci_search_key(SignalOwnedBuffer* out, const SignalType_FixedArray17_uint8_t* aci);
func Signal_key_transparency_aci_search_key(
	out *C.SignalOwnedBuffer,
	aci *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_key_transparency_aci_search_key(out, aci)
}

// SignalFfiError* signal_key_transparency_check(SignalCPromisePairOfOwnedBufferOwnedBuffer* promise, SignalConstPointerTokioAsyncContext async_runtime, uint8_t environment, SignalConstPointerUnauthenticatedChatConnection chat_connection, const SignalType_FixedArray17_uint8_t* aci, SignalConstPointerPublicKey aci_identity_key, const int8_t* e164, SignalOptionalBorrowedSliceOfc_uchar unidentified_access_key, SignalOptionalBorrowedSliceOfc_uchar username_hash, SignalOptionalBorrowedSliceOfc_uchar account_data, SignalOptionalBorrowedSliceOfc_uchar last_distinguished_tree_head, bool is_self_check, bool is_e164_discoverable);
func Signal_key_transparency_check(
	promise *C.SignalCPromisePairOfOwnedBufferOwnedBuffer,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	environment C.uint8_t,
	chat_connection C.SignalConstPointerUnauthenticatedChatConnection,
	aci *C.SignalType_FixedArray17_uint8_t,
	aci_identity_key C.SignalConstPointerPublicKey,
	e164 *C.int8_t,
	unidentified_access_key C.SignalOptionalBorrowedSliceOfc_uchar,
	username_hash C.SignalOptionalBorrowedSliceOfc_uchar,
	account_data C.SignalOptionalBorrowedSliceOfc_uchar,
	last_distinguished_tree_head C.SignalOptionalBorrowedSliceOfc_uchar,
	is_self_check C.bool,
	is_e164_discoverable C.bool,
) *C.SignalFfiError {
	return C.signal_key_transparency_check(promise, async_runtime, environment, chat_connection, aci, aci_identity_key, e164, unidentified_access_key, username_hash, account_data, last_distinguished_tree_head, is_self_check, is_e164_discoverable)
}

// SignalFfiError* signal_key_transparency_e164_search_key(SignalOwnedBuffer* out, const int8_t* e164);
func Signal_key_transparency_e164_search_key(
	out *C.SignalOwnedBuffer,
	e164 *C.int8_t,
) *C.SignalFfiError {
	return C.signal_key_transparency_e164_search_key(out, e164)
}

// SignalFfiError* signal_key_transparency_reset_data_field(SignalOwnedBuffer* out, SignalBorrowedBuffer account_data, uint8_t field);
func Signal_key_transparency_reset_data_field(
	out *C.SignalOwnedBuffer,
	account_data C.SignalBorrowedBuffer,
	field C.uint8_t,
) *C.SignalFfiError {
	return C.signal_key_transparency_reset_data_field(out, account_data, field)
}

// SignalFfiError* signal_key_transparency_username_hash_search_key(SignalOwnedBuffer* out, SignalBorrowedBuffer hash);
func Signal_key_transparency_username_hash_search_key(
	out *C.SignalOwnedBuffer,
	hash C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_key_transparency_username_hash_search_key(out, hash)
}

// SignalFfiError* signal_kyber_key_pair_clone(SignalMutPointerKyberKeyPair* new_obj, SignalConstPointerKyberKeyPair obj);
func Signal_kyber_key_pair_clone(
	new_obj *C.SignalMutPointerKyberKeyPair,
	obj C.SignalConstPointerKyberKeyPair,
) *C.SignalFfiError {
	return C.signal_kyber_key_pair_clone(new_obj, obj)
}

// SignalFfiError* signal_kyber_key_pair_destroy(SignalMutPointerKyberKeyPair p);
func Signal_kyber_key_pair_destroy(
	p C.SignalMutPointerKyberKeyPair,
) *C.SignalFfiError {
	return C.signal_kyber_key_pair_destroy(p)
}

// SignalFfiError* signal_kyber_key_pair_generate(SignalMutPointerKyberKeyPair* out);
func Signal_kyber_key_pair_generate(
	out *C.SignalMutPointerKyberKeyPair,
) *C.SignalFfiError {
	return C.signal_kyber_key_pair_generate(out)
}

// SignalFfiError* signal_kyber_key_pair_get_public_key(SignalMutPointerKyberPublicKey* out, SignalConstPointerKyberKeyPair key_pair);
func Signal_kyber_key_pair_get_public_key(
	out *C.SignalMutPointerKyberPublicKey,
	key_pair C.SignalConstPointerKyberKeyPair,
) *C.SignalFfiError {
	return C.signal_kyber_key_pair_get_public_key(out, key_pair)
}

// SignalFfiError* signal_kyber_key_pair_get_secret_key(SignalMutPointerKyberSecretKey* out, SignalConstPointerKyberKeyPair key_pair);
func Signal_kyber_key_pair_get_secret_key(
	out *C.SignalMutPointerKyberSecretKey,
	key_pair C.SignalConstPointerKyberKeyPair,
) *C.SignalFfiError {
	return C.signal_kyber_key_pair_get_secret_key(out, key_pair)
}

// SignalFfiError* signal_kyber_pre_key_record_clone(SignalMutPointerKyberPreKeyRecord* new_obj, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_clone(
	new_obj *C.SignalMutPointerKyberPreKeyRecord,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_clone(new_obj, obj)
}

// SignalFfiError* signal_kyber_pre_key_record_deserialize(SignalMutPointerKyberPreKeyRecord* out, SignalBorrowedBuffer data);
func Signal_kyber_pre_key_record_deserialize(
	out *C.SignalMutPointerKyberPreKeyRecord,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_deserialize(out, data)
}

// SignalFfiError* signal_kyber_pre_key_record_destroy(SignalMutPointerKyberPreKeyRecord p);
func Signal_kyber_pre_key_record_destroy(
	p C.SignalMutPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_destroy(p)
}

// SignalFfiError* signal_kyber_pre_key_record_get_id(uint32_t* out, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_get_id(
	out *C.uint32_t,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_get_id(out, obj)
}

// SignalFfiError* signal_kyber_pre_key_record_get_key_pair(SignalMutPointerKyberKeyPair* out, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_get_key_pair(
	out *C.SignalMutPointerKyberKeyPair,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_get_key_pair(out, obj)
}

// SignalFfiError* signal_kyber_pre_key_record_get_public_key(SignalMutPointerKyberPublicKey* out, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_get_public_key(
	out *C.SignalMutPointerKyberPublicKey,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_get_public_key(out, obj)
}

// SignalFfiError* signal_kyber_pre_key_record_get_secret_key(SignalMutPointerKyberSecretKey* out, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_get_secret_key(
	out *C.SignalMutPointerKyberSecretKey,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_get_secret_key(out, obj)
}

// SignalFfiError* signal_kyber_pre_key_record_get_signature(SignalOwnedBuffer* out, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_get_signature(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_get_signature(out, obj)
}

// SignalFfiError* signal_kyber_pre_key_record_get_timestamp(uint64_t* out, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_get_timestamp(
	out *C.uint64_t,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_get_timestamp(out, obj)
}

// SignalFfiError* signal_kyber_pre_key_record_new(SignalMutPointerKyberPreKeyRecord* out, uint32_t id, uint64_t timestamp, SignalConstPointerKyberKeyPair key_pair, SignalBorrowedBuffer signature);
func Signal_kyber_pre_key_record_new(
	out *C.SignalMutPointerKyberPreKeyRecord,
	id C.uint32_t,
	timestamp C.uint64_t,
	key_pair C.SignalConstPointerKyberKeyPair,
	signature C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_new(out, id, timestamp, key_pair, signature)
}

// SignalFfiError* signal_kyber_pre_key_record_serialize(SignalOwnedBuffer* out, SignalConstPointerKyberPreKeyRecord obj);
func Signal_kyber_pre_key_record_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerKyberPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_kyber_pre_key_record_serialize(out, obj)
}

// SignalFfiError* signal_kyber_public_key_clone(SignalMutPointerKyberPublicKey* new_obj, SignalConstPointerKyberPublicKey obj);
func Signal_kyber_public_key_clone(
	new_obj *C.SignalMutPointerKyberPublicKey,
	obj C.SignalConstPointerKyberPublicKey,
) *C.SignalFfiError {
	return C.signal_kyber_public_key_clone(new_obj, obj)
}

// SignalFfiError* signal_kyber_public_key_deserialize(SignalMutPointerKyberPublicKey* out, SignalBorrowedBuffer data);
func Signal_kyber_public_key_deserialize(
	out *C.SignalMutPointerKyberPublicKey,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_kyber_public_key_deserialize(out, data)
}

// SignalFfiError* signal_kyber_public_key_destroy(SignalMutPointerKyberPublicKey p);
func Signal_kyber_public_key_destroy(
	p C.SignalMutPointerKyberPublicKey,
) *C.SignalFfiError {
	return C.signal_kyber_public_key_destroy(p)
}

// SignalFfiError* signal_kyber_public_key_equals(bool* out, SignalConstPointerKyberPublicKey lhs, SignalConstPointerKyberPublicKey rhs);
func Signal_kyber_public_key_equals(
	out *C.bool,
	lhs C.SignalConstPointerKyberPublicKey,
	rhs C.SignalConstPointerKyberPublicKey,
) *C.SignalFfiError {
	return C.signal_kyber_public_key_equals(out, lhs, rhs)
}

// SignalFfiError* signal_kyber_public_key_serialize(SignalOwnedBuffer* out, SignalConstPointerKyberPublicKey obj);
func Signal_kyber_public_key_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerKyberPublicKey,
) *C.SignalFfiError {
	return C.signal_kyber_public_key_serialize(out, obj)
}

// SignalFfiError* signal_kyber_secret_key_clone(SignalMutPointerKyberSecretKey* new_obj, SignalConstPointerKyberSecretKey obj);
func Signal_kyber_secret_key_clone(
	new_obj *C.SignalMutPointerKyberSecretKey,
	obj C.SignalConstPointerKyberSecretKey,
) *C.SignalFfiError {
	return C.signal_kyber_secret_key_clone(new_obj, obj)
}

// SignalFfiError* signal_kyber_secret_key_deserialize(SignalMutPointerKyberSecretKey* out, SignalBorrowedBuffer data);
func Signal_kyber_secret_key_deserialize(
	out *C.SignalMutPointerKyberSecretKey,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_kyber_secret_key_deserialize(out, data)
}

// SignalFfiError* signal_kyber_secret_key_destroy(SignalMutPointerKyberSecretKey p);
func Signal_kyber_secret_key_destroy(
	p C.SignalMutPointerKyberSecretKey,
) *C.SignalFfiError {
	return C.signal_kyber_secret_key_destroy(p)
}

// SignalFfiError* signal_kyber_secret_key_serialize(SignalOwnedBuffer* out, SignalConstPointerKyberSecretKey obj);
func Signal_kyber_secret_key_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerKyberSecretKey,
) *C.SignalFfiError {
	return C.signal_kyber_secret_key_serialize(out, obj)
}

// SignalFfiError* signal_lookup_request_add_aci_and_access_key(SignalConstPointerLookupRequest request, const SignalType_FixedArray17_uint8_t* aci, SignalBorrowedBuffer access_key);
func Signal_lookup_request_add_aci_and_access_key(
	request C.SignalConstPointerLookupRequest,
	aci *C.SignalType_FixedArray17_uint8_t,
	access_key C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_lookup_request_add_aci_and_access_key(request, aci, access_key)
}

// SignalFfiError* signal_lookup_request_add_e164(SignalConstPointerLookupRequest request, const int8_t* e164);
func Signal_lookup_request_add_e164(
	request C.SignalConstPointerLookupRequest,
	e164 *C.int8_t,
) *C.SignalFfiError {
	return C.signal_lookup_request_add_e164(request, e164)
}

// SignalFfiError* signal_lookup_request_add_previous_e164(SignalConstPointerLookupRequest request, const int8_t* e164);
func Signal_lookup_request_add_previous_e164(
	request C.SignalConstPointerLookupRequest,
	e164 *C.int8_t,
) *C.SignalFfiError {
	return C.signal_lookup_request_add_previous_e164(request, e164)
}

// SignalFfiError* signal_lookup_request_destroy(SignalMutPointerLookupRequest p);
func Signal_lookup_request_destroy(
	p C.SignalMutPointerLookupRequest,
) *C.SignalFfiError {
	return C.signal_lookup_request_destroy(p)
}

// SignalFfiError* signal_lookup_request_new(SignalMutPointerLookupRequest* out);
func Signal_lookup_request_new(
	out *C.SignalMutPointerLookupRequest,
) *C.SignalFfiError {
	return C.signal_lookup_request_new(out)
}

// SignalFfiError* signal_lookup_request_set_token(SignalConstPointerLookupRequest request, SignalBorrowedBuffer token);
func Signal_lookup_request_set_token(
	request C.SignalConstPointerLookupRequest,
	token C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_lookup_request_set_token(request, token)
}

// SignalFfiError* signal_message_backup_key_destroy(SignalMutPointerMessageBackupKey p);
func Signal_message_backup_key_destroy(
	p C.SignalMutPointerMessageBackupKey,
) *C.SignalFfiError {
	return C.signal_message_backup_key_destroy(p)
}

// SignalFfiError* signal_message_backup_key_from_account_entropy_pool(SignalMutPointerMessageBackupKey* out, const int8_t* account_entropy, const SignalType_FixedArray17_uint8_t* aci, const SignalType_FixedArray32_uint8_t* forward_secrecy_token);
func Signal_message_backup_key_from_account_entropy_pool(
	out *C.SignalMutPointerMessageBackupKey,
	account_entropy *C.int8_t,
	aci *C.SignalType_FixedArray17_uint8_t,
	forward_secrecy_token *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_message_backup_key_from_account_entropy_pool(out, account_entropy, aci, forward_secrecy_token)
}

// SignalFfiError* signal_message_backup_key_from_backup_key_and_backup_id(SignalMutPointerMessageBackupKey* out, const SignalType_FixedArray32_uint8_t* backup_key, const SignalType_FixedArray16_uint8_t* backup_id, const SignalType_FixedArray32_uint8_t* forward_secrecy_token);
func Signal_message_backup_key_from_backup_key_and_backup_id(
	out *C.SignalMutPointerMessageBackupKey,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	backup_id *C.SignalType_FixedArray16_uint8_t,
	forward_secrecy_token *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_message_backup_key_from_backup_key_and_backup_id(out, backup_key, backup_id, forward_secrecy_token)
}

// SignalFfiError* signal_message_backup_key_get_aes_key(SignalType_FixedArray32_uint8_t* out, SignalConstPointerMessageBackupKey key);
func Signal_message_backup_key_get_aes_key(
	out *C.SignalType_FixedArray32_uint8_t,
	key C.SignalConstPointerMessageBackupKey,
) *C.SignalFfiError {
	return C.signal_message_backup_key_get_aes_key(out, key)
}

// SignalFfiError* signal_message_backup_key_get_hmac_key(SignalType_FixedArray32_uint8_t* out, SignalConstPointerMessageBackupKey key);
func Signal_message_backup_key_get_hmac_key(
	out *C.SignalType_FixedArray32_uint8_t,
	key C.SignalConstPointerMessageBackupKey,
) *C.SignalFfiError {
	return C.signal_message_backup_key_get_hmac_key(out, key)
}

// SignalFfiError* signal_message_backup_validation_outcome_destroy(SignalMutPointerMessageBackupValidationOutcome p);
func Signal_message_backup_validation_outcome_destroy(
	p C.SignalMutPointerMessageBackupValidationOutcome,
) *C.SignalFfiError {
	return C.signal_message_backup_validation_outcome_destroy(p)
}

// SignalFfiError* signal_message_backup_validation_outcome_get_error_message(SignalCStringPtr* out, SignalConstPointerMessageBackupValidationOutcome outcome);
func Signal_message_backup_validation_outcome_get_error_message(
	out *C.SignalCStringPtr,
	outcome C.SignalConstPointerMessageBackupValidationOutcome,
) *C.SignalFfiError {
	return C.signal_message_backup_validation_outcome_get_error_message(out, outcome)
}

// SignalFfiError* signal_message_backup_validation_outcome_get_unknown_fields(SignalBytestringArray* out, SignalConstPointerMessageBackupValidationOutcome outcome);
func Signal_message_backup_validation_outcome_get_unknown_fields(
	out *C.SignalBytestringArray,
	outcome C.SignalConstPointerMessageBackupValidationOutcome,
) *C.SignalFfiError {
	return C.signal_message_backup_validation_outcome_get_unknown_fields(out, outcome)
}

// SignalFfiError* signal_message_backup_validator_validate(SignalMutPointerMessageBackupValidationOutcome* out, SignalConstPointerMessageBackupKey key, SignalConstPointerFfiSyncInputStreamStruct first_stream, SignalConstPointerFfiSyncInputStreamStruct second_stream, uint64_t len, uint8_t purpose);
func Signal_message_backup_validator_validate(
	out *C.SignalMutPointerMessageBackupValidationOutcome,
	key C.SignalConstPointerMessageBackupKey,
	first_stream C.SignalConstPointerFfiSyncInputStreamStruct,
	second_stream C.SignalConstPointerFfiSyncInputStreamStruct,
	len C.uint64_t,
	purpose C.uint8_t,
) *C.SignalFfiError {
	return C.signal_message_backup_validator_validate(out, key, first_stream, second_stream, len, purpose)
}

// SignalFfiError* signal_message_clone(SignalMutPointerSignalMessage* new_obj, SignalConstPointerSignalMessage obj);
func Signal_message_clone(
	new_obj *C.SignalMutPointerSignalMessage,
	obj C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_clone(new_obj, obj)
}

// SignalFfiError* signal_message_deserialize(SignalMutPointerSignalMessage* out, SignalBorrowedBuffer data);
func Signal_message_deserialize(
	out *C.SignalMutPointerSignalMessage,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_message_deserialize(out, data)
}

// SignalFfiError* signal_message_destroy(SignalMutPointerSignalMessage p);
func Signal_message_destroy(
	p C.SignalMutPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_destroy(p)
}

// SignalFfiError* signal_message_get_body(SignalOwnedBuffer* out, SignalConstPointerSignalMessage obj);
func Signal_message_get_body(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_get_body(out, obj)
}

// SignalFfiError* signal_message_get_counter(uint32_t* out, SignalConstPointerSignalMessage obj);
func Signal_message_get_counter(
	out *C.uint32_t,
	obj C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_get_counter(out, obj)
}

// SignalFfiError* signal_message_get_message_version(uint32_t* out, SignalConstPointerSignalMessage obj);
func Signal_message_get_message_version(
	out *C.uint32_t,
	obj C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_get_message_version(out, obj)
}

// SignalFfiError* signal_message_get_pq_ratchet(SignalOwnedBuffer* out, SignalConstPointerSignalMessage msg);
func Signal_message_get_pq_ratchet(
	out *C.SignalOwnedBuffer,
	msg C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_get_pq_ratchet(out, msg)
}

// SignalFfiError* signal_message_get_sender_ratchet_key(SignalMutPointerPublicKey* out, SignalConstPointerSignalMessage m);
func Signal_message_get_sender_ratchet_key(
	out *C.SignalMutPointerPublicKey,
	m C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_get_sender_ratchet_key(out, m)
}

// SignalFfiError* signal_message_get_serialized(SignalOwnedBuffer* out, SignalConstPointerSignalMessage obj);
func Signal_message_get_serialized(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_message_get_serialized(out, obj)
}

// SignalFfiError* signal_message_new(SignalMutPointerSignalMessage* out, uint8_t message_version, SignalBorrowedBuffer mac_key, SignalConstPointerPublicKey sender_ratchet_key, uint32_t counter, uint32_t previous_counter, SignalBorrowedBuffer ciphertext, SignalConstPointerPublicKey sender_identity_key, SignalConstPointerPublicKey receiver_identity_key, SignalBorrowedBuffer pq_ratchet);
func Signal_message_new(
	out *C.SignalMutPointerSignalMessage,
	message_version C.uint8_t,
	mac_key C.SignalBorrowedBuffer,
	sender_ratchet_key C.SignalConstPointerPublicKey,
	counter C.uint32_t,
	previous_counter C.uint32_t,
	ciphertext C.SignalBorrowedBuffer,
	sender_identity_key C.SignalConstPointerPublicKey,
	receiver_identity_key C.SignalConstPointerPublicKey,
	pq_ratchet C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_message_new(out, message_version, mac_key, sender_ratchet_key, counter, previous_counter, ciphertext, sender_identity_key, receiver_identity_key, pq_ratchet)
}

// SignalFfiError* signal_online_backup_validator_add_frame(SignalMutPointerOnlineBackupValidator backup, SignalBorrowedBuffer frame);
func Signal_online_backup_validator_add_frame(
	backup C.SignalMutPointerOnlineBackupValidator,
	frame C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_online_backup_validator_add_frame(backup, frame)
}

// SignalFfiError* signal_online_backup_validator_destroy(SignalMutPointerOnlineBackupValidator p);
func Signal_online_backup_validator_destroy(
	p C.SignalMutPointerOnlineBackupValidator,
) *C.SignalFfiError {
	return C.signal_online_backup_validator_destroy(p)
}

// SignalFfiError* signal_online_backup_validator_finalize(SignalMutPointerOnlineBackupValidator backup);
func Signal_online_backup_validator_finalize(
	backup C.SignalMutPointerOnlineBackupValidator,
) *C.SignalFfiError {
	return C.signal_online_backup_validator_finalize(backup)
}

// SignalFfiError* signal_online_backup_validator_new(SignalMutPointerOnlineBackupValidator* out, SignalBorrowedBuffer backup_info_frame, uint8_t purpose);
func Signal_online_backup_validator_new(
	out *C.SignalMutPointerOnlineBackupValidator,
	backup_info_frame C.SignalBorrowedBuffer,
	purpose C.uint8_t,
) *C.SignalFfiError {
	return C.signal_online_backup_validator_new(out, backup_info_frame, purpose)
}

// SignalFfiError* signal_pin_hash_access_key(SignalType_FixedArray32_uint8_t* out, SignalConstPointerPinHash ph);
func Signal_pin_hash_access_key(
	out *C.SignalType_FixedArray32_uint8_t,
	ph C.SignalConstPointerPinHash,
) *C.SignalFfiError {
	return C.signal_pin_hash_access_key(out, ph)
}

// SignalFfiError* signal_pin_hash_clone(SignalMutPointerPinHash* new_obj, SignalConstPointerPinHash obj);
func Signal_pin_hash_clone(
	new_obj *C.SignalMutPointerPinHash,
	obj C.SignalConstPointerPinHash,
) *C.SignalFfiError {
	return C.signal_pin_hash_clone(new_obj, obj)
}

// SignalFfiError* signal_pin_hash_destroy(SignalMutPointerPinHash p);
func Signal_pin_hash_destroy(
	p C.SignalMutPointerPinHash,
) *C.SignalFfiError {
	return C.signal_pin_hash_destroy(p)
}

// SignalFfiError* signal_pin_hash_encryption_key(SignalType_FixedArray32_uint8_t* out, SignalConstPointerPinHash ph);
func Signal_pin_hash_encryption_key(
	out *C.SignalType_FixedArray32_uint8_t,
	ph C.SignalConstPointerPinHash,
) *C.SignalFfiError {
	return C.signal_pin_hash_encryption_key(out, ph)
}

// SignalFfiError* signal_pin_hash_from_salt(SignalMutPointerPinHash* out, SignalBorrowedBuffer pin, const SignalType_FixedArray32_uint8_t* salt);
func Signal_pin_hash_from_salt(
	out *C.SignalMutPointerPinHash,
	pin C.SignalBorrowedBuffer,
	salt *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_pin_hash_from_salt(out, pin, salt)
}

// SignalFfiError* signal_pin_hash_from_username_mrenclave(SignalMutPointerPinHash* out, SignalBorrowedBuffer pin, const int8_t* username, SignalBorrowedBuffer mrenclave);
func Signal_pin_hash_from_username_mrenclave(
	out *C.SignalMutPointerPinHash,
	pin C.SignalBorrowedBuffer,
	username *C.int8_t,
	mrenclave C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_pin_hash_from_username_mrenclave(out, pin, username, mrenclave)
}

// SignalFfiError* signal_pin_local_hash(SignalCStringPtr* out, SignalBorrowedBuffer pin);
func Signal_pin_local_hash(
	out *C.SignalCStringPtr,
	pin C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_pin_local_hash(out, pin)
}

// SignalFfiError* signal_pin_verify_local_hash(bool* out, const int8_t* encoded_hash, SignalBorrowedBuffer pin);
func Signal_pin_verify_local_hash(
	out *C.bool,
	encoded_hash *C.int8_t,
	pin C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_pin_verify_local_hash(out, encoded_hash, pin)
}

// SignalFfiError* signal_plaintext_content_clone(SignalMutPointerPlaintextContent* new_obj, SignalConstPointerPlaintextContent obj);
func Signal_plaintext_content_clone(
	new_obj *C.SignalMutPointerPlaintextContent,
	obj C.SignalConstPointerPlaintextContent,
) *C.SignalFfiError {
	return C.signal_plaintext_content_clone(new_obj, obj)
}

// SignalFfiError* signal_plaintext_content_deserialize(SignalMutPointerPlaintextContent* out, SignalBorrowedBuffer data);
func Signal_plaintext_content_deserialize(
	out *C.SignalMutPointerPlaintextContent,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_plaintext_content_deserialize(out, data)
}

// SignalFfiError* signal_plaintext_content_destroy(SignalMutPointerPlaintextContent p);
func Signal_plaintext_content_destroy(
	p C.SignalMutPointerPlaintextContent,
) *C.SignalFfiError {
	return C.signal_plaintext_content_destroy(p)
}

// SignalFfiError* signal_plaintext_content_from_decryption_error_message(SignalMutPointerPlaintextContent* out, SignalConstPointerDecryptionErrorMessage m);
func Signal_plaintext_content_from_decryption_error_message(
	out *C.SignalMutPointerPlaintextContent,
	m C.SignalConstPointerDecryptionErrorMessage,
) *C.SignalFfiError {
	return C.signal_plaintext_content_from_decryption_error_message(out, m)
}

// SignalFfiError* signal_plaintext_content_get_body(SignalOwnedBuffer* out, SignalConstPointerPlaintextContent obj);
func Signal_plaintext_content_get_body(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPlaintextContent,
) *C.SignalFfiError {
	return C.signal_plaintext_content_get_body(out, obj)
}

// SignalFfiError* signal_plaintext_content_serialize(SignalOwnedBuffer* out, SignalConstPointerPlaintextContent obj);
func Signal_plaintext_content_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPlaintextContent,
) *C.SignalFfiError {
	return C.signal_plaintext_content_serialize(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_clone(SignalMutPointerPreKeyBundle* new_obj, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_clone(
	new_obj *C.SignalMutPointerPreKeyBundle,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_clone(new_obj, obj)
}

// SignalFfiError* signal_pre_key_bundle_destroy(SignalMutPointerPreKeyBundle p);
func Signal_pre_key_bundle_destroy(
	p C.SignalMutPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_destroy(p)
}

// SignalFfiError* signal_pre_key_bundle_get_device_id(uint32_t* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_device_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_device_id(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_identity_key(SignalMutPointerPublicKey* out, SignalConstPointerPreKeyBundle p);
func Signal_pre_key_bundle_get_identity_key(
	out *C.SignalMutPointerPublicKey,
	p C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_identity_key(out, p)
}

// SignalFfiError* signal_pre_key_bundle_get_kyber_pre_key_id(uint32_t* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_kyber_pre_key_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_kyber_pre_key_id(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_kyber_pre_key_public(SignalMutPointerKyberPublicKey* out, SignalConstPointerPreKeyBundle bundle);
func Signal_pre_key_bundle_get_kyber_pre_key_public(
	out *C.SignalMutPointerKyberPublicKey,
	bundle C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_kyber_pre_key_public(out, bundle)
}

// SignalFfiError* signal_pre_key_bundle_get_kyber_pre_key_signature(SignalOwnedBuffer* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_kyber_pre_key_signature(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_kyber_pre_key_signature(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_pre_key_id(uint32_t* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_pre_key_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_pre_key_id(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_pre_key_public(SignalMutPointerPublicKey* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_pre_key_public(
	out *C.SignalMutPointerPublicKey,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_pre_key_public(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_registration_id(uint32_t* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_registration_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_registration_id(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_signed_pre_key_id(uint32_t* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_signed_pre_key_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_signed_pre_key_id(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_signed_pre_key_public(SignalMutPointerPublicKey* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_signed_pre_key_public(
	out *C.SignalMutPointerPublicKey,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_signed_pre_key_public(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_get_signed_pre_key_signature(SignalOwnedBuffer* out, SignalConstPointerPreKeyBundle obj);
func Signal_pre_key_bundle_get_signed_pre_key_signature(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPreKeyBundle,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_get_signed_pre_key_signature(out, obj)
}

// SignalFfiError* signal_pre_key_bundle_new(SignalMutPointerPreKeyBundle* out, uint32_t registration_id, uint32_t device_id, uint32_t prekey_id, SignalConstPointerPublicKey prekey, uint32_t signed_prekey_id, SignalConstPointerPublicKey signed_prekey, SignalBorrowedBuffer signed_prekey_signature, SignalConstPointerPublicKey identity_key, uint32_t kyber_prekey_id, SignalConstPointerKyberPublicKey kyber_prekey, SignalBorrowedBuffer kyber_prekey_signature);
func Signal_pre_key_bundle_new(
	out *C.SignalMutPointerPreKeyBundle,
	registration_id C.uint32_t,
	device_id C.uint32_t,
	prekey_id C.uint32_t,
	prekey C.SignalConstPointerPublicKey,
	signed_prekey_id C.uint32_t,
	signed_prekey C.SignalConstPointerPublicKey,
	signed_prekey_signature C.SignalBorrowedBuffer,
	identity_key C.SignalConstPointerPublicKey,
	kyber_prekey_id C.uint32_t,
	kyber_prekey C.SignalConstPointerKyberPublicKey,
	kyber_prekey_signature C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_pre_key_bundle_new(out, registration_id, device_id, prekey_id, prekey, signed_prekey_id, signed_prekey, signed_prekey_signature, identity_key, kyber_prekey_id, kyber_prekey, kyber_prekey_signature)
}

// SignalFfiError* signal_pre_key_record_clone(SignalMutPointerPreKeyRecord* new_obj, SignalConstPointerPreKeyRecord obj);
func Signal_pre_key_record_clone(
	new_obj *C.SignalMutPointerPreKeyRecord,
	obj C.SignalConstPointerPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_pre_key_record_clone(new_obj, obj)
}

// SignalFfiError* signal_pre_key_record_deserialize(SignalMutPointerPreKeyRecord* out, SignalBorrowedBuffer data);
func Signal_pre_key_record_deserialize(
	out *C.SignalMutPointerPreKeyRecord,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_pre_key_record_deserialize(out, data)
}

// SignalFfiError* signal_pre_key_record_destroy(SignalMutPointerPreKeyRecord p);
func Signal_pre_key_record_destroy(
	p C.SignalMutPointerPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_pre_key_record_destroy(p)
}

// SignalFfiError* signal_pre_key_record_get_id(uint32_t* out, SignalConstPointerPreKeyRecord obj);
func Signal_pre_key_record_get_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_pre_key_record_get_id(out, obj)
}

// SignalFfiError* signal_pre_key_record_get_private_key(SignalMutPointerPrivateKey* out, SignalConstPointerPreKeyRecord obj);
func Signal_pre_key_record_get_private_key(
	out *C.SignalMutPointerPrivateKey,
	obj C.SignalConstPointerPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_pre_key_record_get_private_key(out, obj)
}

// SignalFfiError* signal_pre_key_record_get_public_key(SignalMutPointerPublicKey* out, SignalConstPointerPreKeyRecord obj);
func Signal_pre_key_record_get_public_key(
	out *C.SignalMutPointerPublicKey,
	obj C.SignalConstPointerPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_pre_key_record_get_public_key(out, obj)
}

// SignalFfiError* signal_pre_key_record_new(SignalMutPointerPreKeyRecord* out, uint32_t id, SignalConstPointerPublicKey pub_key, SignalConstPointerPrivateKey priv_key);
func Signal_pre_key_record_new(
	out *C.SignalMutPointerPreKeyRecord,
	id C.uint32_t,
	pub_key C.SignalConstPointerPublicKey,
	priv_key C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_pre_key_record_new(out, id, pub_key, priv_key)
}

// SignalFfiError* signal_pre_key_record_serialize(SignalOwnedBuffer* out, SignalConstPointerPreKeyRecord obj);
func Signal_pre_key_record_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_pre_key_record_serialize(out, obj)
}

// SignalFfiError* signal_pre_key_signal_message_clone(SignalMutPointerPreKeySignalMessage* new_obj, SignalConstPointerPreKeySignalMessage obj);
func Signal_pre_key_signal_message_clone(
	new_obj *C.SignalMutPointerPreKeySignalMessage,
	obj C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_clone(new_obj, obj)
}

// SignalFfiError* signal_pre_key_signal_message_deserialize(SignalMutPointerPreKeySignalMessage* out, SignalBorrowedBuffer data);
func Signal_pre_key_signal_message_deserialize(
	out *C.SignalMutPointerPreKeySignalMessage,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_deserialize(out, data)
}

// SignalFfiError* signal_pre_key_signal_message_destroy(SignalMutPointerPreKeySignalMessage p);
func Signal_pre_key_signal_message_destroy(
	p C.SignalMutPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_destroy(p)
}

// SignalFfiError* signal_pre_key_signal_message_get_base_key(SignalMutPointerPublicKey* out, SignalConstPointerPreKeySignalMessage m);
func Signal_pre_key_signal_message_get_base_key(
	out *C.SignalMutPointerPublicKey,
	m C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_get_base_key(out, m)
}

// SignalFfiError* signal_pre_key_signal_message_get_identity_key(SignalMutPointerPublicKey* out, SignalConstPointerPreKeySignalMessage m);
func Signal_pre_key_signal_message_get_identity_key(
	out *C.SignalMutPointerPublicKey,
	m C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_get_identity_key(out, m)
}

// SignalFfiError* signal_pre_key_signal_message_get_pre_key_id(uint32_t* out, SignalConstPointerPreKeySignalMessage obj);
func Signal_pre_key_signal_message_get_pre_key_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_get_pre_key_id(out, obj)
}

// SignalFfiError* signal_pre_key_signal_message_get_registration_id(uint32_t* out, SignalConstPointerPreKeySignalMessage obj);
func Signal_pre_key_signal_message_get_registration_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_get_registration_id(out, obj)
}

// SignalFfiError* signal_pre_key_signal_message_get_signal_message(SignalMutPointerSignalMessage* out, SignalConstPointerPreKeySignalMessage m);
func Signal_pre_key_signal_message_get_signal_message(
	out *C.SignalMutPointerSignalMessage,
	m C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_get_signal_message(out, m)
}

// SignalFfiError* signal_pre_key_signal_message_get_signed_pre_key_id(uint32_t* out, SignalConstPointerPreKeySignalMessage obj);
func Signal_pre_key_signal_message_get_signed_pre_key_id(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_get_signed_pre_key_id(out, obj)
}

// SignalFfiError* signal_pre_key_signal_message_get_version(uint32_t* out, SignalConstPointerPreKeySignalMessage obj);
func Signal_pre_key_signal_message_get_version(
	out *C.uint32_t,
	obj C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_get_version(out, obj)
}

// SignalFfiError* signal_pre_key_signal_message_new(SignalMutPointerPreKeySignalMessage* out, uint8_t message_version, uint32_t registration_id, uint32_t pre_key_id, uint32_t signed_pre_key_id, SignalConstPointerPublicKey base_key, SignalConstPointerPublicKey identity_key, SignalConstPointerSignalMessage signal_message);
func Signal_pre_key_signal_message_new(
	out *C.SignalMutPointerPreKeySignalMessage,
	message_version C.uint8_t,
	registration_id C.uint32_t,
	pre_key_id C.uint32_t,
	signed_pre_key_id C.uint32_t,
	base_key C.SignalConstPointerPublicKey,
	identity_key C.SignalConstPointerPublicKey,
	signal_message C.SignalConstPointerSignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_new(out, message_version, registration_id, pre_key_id, signed_pre_key_id, base_key, identity_key, signal_message)
}

// SignalFfiError* signal_pre_key_signal_message_serialize(SignalOwnedBuffer* out, SignalConstPointerPreKeySignalMessage obj);
func Signal_pre_key_signal_message_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPreKeySignalMessage,
) *C.SignalFfiError {
	return C.signal_pre_key_signal_message_serialize(out, obj)
}

// SignalFfiError* signal_privatekey_agree(SignalOwnedBuffer* out, SignalConstPointerPrivateKey private_key, SignalConstPointerPublicKey public_key);
func Signal_privatekey_agree(
	out *C.SignalOwnedBuffer,
	private_key C.SignalConstPointerPrivateKey,
	public_key C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_privatekey_agree(out, private_key, public_key)
}

// SignalFfiError* signal_privatekey_clone(SignalMutPointerPrivateKey* new_obj, SignalConstPointerPrivateKey obj);
func Signal_privatekey_clone(
	new_obj *C.SignalMutPointerPrivateKey,
	obj C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_privatekey_clone(new_obj, obj)
}

// SignalFfiError* signal_privatekey_deserialize(SignalMutPointerPrivateKey* out, SignalBorrowedBuffer data);
func Signal_privatekey_deserialize(
	out *C.SignalMutPointerPrivateKey,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_privatekey_deserialize(out, data)
}

// SignalFfiError* signal_privatekey_destroy(SignalMutPointerPrivateKey p);
func Signal_privatekey_destroy(
	p C.SignalMutPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_privatekey_destroy(p)
}

// SignalFfiError* signal_privatekey_generate(SignalMutPointerPrivateKey* out);
func Signal_privatekey_generate(
	out *C.SignalMutPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_privatekey_generate(out)
}

// SignalFfiError* signal_privatekey_get_public_key(SignalMutPointerPublicKey* out, SignalConstPointerPrivateKey k);
func Signal_privatekey_get_public_key(
	out *C.SignalMutPointerPublicKey,
	k C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_privatekey_get_public_key(out, k)
}

// SignalFfiError* signal_privatekey_hpke_open(SignalOwnedBuffer* out, SignalConstPointerPrivateKey sk, SignalBorrowedBuffer ciphertext, SignalBorrowedBuffer info, SignalBorrowedBuffer associated_data);
func Signal_privatekey_hpke_open(
	out *C.SignalOwnedBuffer,
	sk C.SignalConstPointerPrivateKey,
	ciphertext C.SignalBorrowedBuffer,
	info C.SignalBorrowedBuffer,
	associated_data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_privatekey_hpke_open(out, sk, ciphertext, info, associated_data)
}

// SignalFfiError* signal_privatekey_serialize(SignalOwnedBuffer* out, SignalConstPointerPrivateKey obj);
func Signal_privatekey_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_privatekey_serialize(out, obj)
}

// SignalFfiError* signal_privatekey_sign(SignalOwnedBuffer* out, SignalConstPointerPrivateKey key, SignalBorrowedBuffer message);
func Signal_privatekey_sign(
	out *C.SignalOwnedBuffer,
	key C.SignalConstPointerPrivateKey,
	message C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_privatekey_sign(out, key, message)
}

// SignalFfiError* signal_process_prekey_bundle(SignalConstPointerPreKeyBundle bundle, SignalConstPointerProtocolAddress protocol_address, SignalConstPointerProtocolAddress local_address, SignalConstPointerFfiSessionStoreStruct session_store, SignalConstPointerFfiIdentityKeyStoreStruct identity_key_store, uint64_t now);
func Signal_process_prekey_bundle(
	bundle C.SignalConstPointerPreKeyBundle,
	protocol_address C.SignalConstPointerProtocolAddress,
	local_address C.SignalConstPointerProtocolAddress,
	session_store C.SignalConstPointerFfiSessionStoreStruct,
	identity_key_store C.SignalConstPointerFfiIdentityKeyStoreStruct,
	now C.uint64_t,
) *C.SignalFfiError {
	return C.signal_process_prekey_bundle(bundle, protocol_address, local_address, session_store, identity_key_store, now)
}

// SignalFfiError* signal_process_sender_key_distribution_message(SignalConstPointerProtocolAddress sender, SignalConstPointerSenderKeyDistributionMessage sender_key_distribution_message, SignalConstPointerFfiSenderKeyStoreStruct store);
func Signal_process_sender_key_distribution_message(
	sender C.SignalConstPointerProtocolAddress,
	sender_key_distribution_message C.SignalConstPointerSenderKeyDistributionMessage,
	store C.SignalConstPointerFfiSenderKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_process_sender_key_distribution_message(sender, sender_key_distribution_message, store)
}

// SignalFfiError* signal_profile_key_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_profile_key_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_check_valid_contents(buffer)
}

// SignalFfiError* signal_profile_key_ciphertext_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_profile_key_ciphertext_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_ciphertext_check_valid_contents(buffer)
}

// SignalFfiError* signal_profile_key_commitment_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_profile_key_commitment_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_commitment_check_valid_contents(buffer)
}

// SignalFfiError* signal_profile_key_credential_presentation_check_valid_contents(SignalBorrowedBuffer presentation_bytes);
func Signal_profile_key_credential_presentation_check_valid_contents(
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_credential_presentation_check_valid_contents(presentation_bytes)
}

// SignalFfiError* signal_profile_key_credential_presentation_get_profile_key_ciphertext(SignalType_FixedArray65_uint8_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_profile_key_credential_presentation_get_profile_key_ciphertext(
	out *C.SignalType_FixedArray65_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_credential_presentation_get_profile_key_ciphertext(out, presentation_bytes)
}

// SignalFfiError* signal_profile_key_credential_presentation_get_uuid_ciphertext(SignalType_FixedArray65_uint8_t* out, SignalBorrowedBuffer presentation_bytes);
func Signal_profile_key_credential_presentation_get_uuid_ciphertext(
	out *C.SignalType_FixedArray65_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_credential_presentation_get_uuid_ciphertext(out, presentation_bytes)
}

// SignalFfiError* signal_profile_key_credential_request_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_profile_key_credential_request_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_credential_request_check_valid_contents(buffer)
}

// SignalFfiError* signal_profile_key_credential_request_context_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_profile_key_credential_request_context_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_profile_key_credential_request_context_check_valid_contents(buffer)
}

// SignalFfiError* signal_profile_key_credential_request_context_get_request(SignalType_FixedArray329_uint8_t* out, const SignalType_FixedArray473_uint8_t* context);
func Signal_profile_key_credential_request_context_get_request(
	out *C.SignalType_FixedArray329_uint8_t,
	context *C.SignalType_FixedArray473_uint8_t,
) *C.SignalFfiError {
	return C.signal_profile_key_credential_request_context_get_request(out, context)
}

// SignalFfiError* signal_profile_key_derive_access_key(SignalType_FixedArray16_uint8_t* out, const SignalType_FixedArray32_uint8_t* profile_key);
func Signal_profile_key_derive_access_key(
	out *C.SignalType_FixedArray16_uint8_t,
	profile_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_profile_key_derive_access_key(out, profile_key)
}

// SignalFfiError* signal_profile_key_get_commitment(SignalType_FixedArray97_uint8_t* out, const SignalType_FixedArray32_uint8_t* profile_key, const SignalType_FixedArray17_uint8_t* user_id);
func Signal_profile_key_get_commitment(
	out *C.SignalType_FixedArray97_uint8_t,
	profile_key *C.SignalType_FixedArray32_uint8_t,
	user_id *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_profile_key_get_commitment(out, profile_key, user_id)
}

// SignalFfiError* signal_profile_key_get_profile_key_version(SignalType_FixedArray64_uint8_t* out, const SignalType_FixedArray32_uint8_t* profile_key, const SignalType_FixedArray17_uint8_t* user_id);
func Signal_profile_key_get_profile_key_version(
	out *C.SignalType_FixedArray64_uint8_t,
	profile_key *C.SignalType_FixedArray32_uint8_t,
	user_id *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_profile_key_get_profile_key_version(out, profile_key, user_id)
}

// SignalFfiError* signal_provisioning_chat_connection_connect(SignalCPromiseMutPointerProvisioningChatConnection* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerConnectionManager connection_manager);
func Signal_provisioning_chat_connection_connect(
	promise *C.SignalCPromiseMutPointerProvisioningChatConnection,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	connection_manager C.SignalConstPointerConnectionManager,
) *C.SignalFfiError {
	return C.signal_provisioning_chat_connection_connect(promise, async_runtime, connection_manager)
}

// SignalFfiError* signal_provisioning_chat_connection_destroy(SignalMutPointerProvisioningChatConnection p);
func Signal_provisioning_chat_connection_destroy(
	p C.SignalMutPointerProvisioningChatConnection,
) *C.SignalFfiError {
	return C.signal_provisioning_chat_connection_destroy(p)
}

// SignalFfiError* signal_provisioning_chat_connection_disconnect(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerProvisioningChatConnection chat);
func Signal_provisioning_chat_connection_disconnect(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerProvisioningChatConnection,
) *C.SignalFfiError {
	return C.signal_provisioning_chat_connection_disconnect(promise, async_runtime, chat)
}

// SignalFfiError* signal_provisioning_chat_connection_info(SignalMutPointerChatConnectionInfo* out, SignalConstPointerProvisioningChatConnection chat);
func Signal_provisioning_chat_connection_info(
	out *C.SignalMutPointerChatConnectionInfo,
	chat C.SignalConstPointerProvisioningChatConnection,
) *C.SignalFfiError {
	return C.signal_provisioning_chat_connection_info(out, chat)
}

// SignalFfiError* signal_provisioning_chat_connection_init_listener(SignalConstPointerProvisioningChatConnection chat, SignalConstPointerFfiProvisioningListenerStruct listener);
func Signal_provisioning_chat_connection_init_listener(
	chat C.SignalConstPointerProvisioningChatConnection,
	listener C.SignalConstPointerFfiProvisioningListenerStruct,
) *C.SignalFfiError {
	return C.signal_provisioning_chat_connection_init_listener(chat, listener)
}

// SignalFfiError* signal_publickey_clone(SignalMutPointerPublicKey* new_obj, SignalConstPointerPublicKey obj);
func Signal_publickey_clone(
	new_obj *C.SignalMutPointerPublicKey,
	obj C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_publickey_clone(new_obj, obj)
}

// SignalFfiError* signal_publickey_deserialize(SignalMutPointerPublicKey* out, SignalBorrowedBuffer data);
func Signal_publickey_deserialize(
	out *C.SignalMutPointerPublicKey,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_publickey_deserialize(out, data)
}

// SignalFfiError* signal_publickey_destroy(SignalMutPointerPublicKey p);
func Signal_publickey_destroy(
	p C.SignalMutPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_publickey_destroy(p)
}

// SignalFfiError* signal_publickey_equals(bool* out, SignalConstPointerPublicKey lhs, SignalConstPointerPublicKey rhs);
func Signal_publickey_equals(
	out *C.bool,
	lhs C.SignalConstPointerPublicKey,
	rhs C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_publickey_equals(out, lhs, rhs)
}

// SignalFfiError* signal_publickey_get_public_key_bytes(SignalOwnedBuffer* out, SignalConstPointerPublicKey obj);
func Signal_publickey_get_public_key_bytes(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_publickey_get_public_key_bytes(out, obj)
}

// SignalFfiError* signal_publickey_hpke_seal(SignalOwnedBuffer* out, SignalConstPointerPublicKey pk, SignalBorrowedBuffer plaintext, SignalBorrowedBuffer info, SignalBorrowedBuffer associated_data);
func Signal_publickey_hpke_seal(
	out *C.SignalOwnedBuffer,
	pk C.SignalConstPointerPublicKey,
	plaintext C.SignalBorrowedBuffer,
	info C.SignalBorrowedBuffer,
	associated_data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_publickey_hpke_seal(out, pk, plaintext, info, associated_data)
}

// SignalFfiError* signal_publickey_serialize(SignalOwnedBuffer* out, SignalConstPointerPublicKey obj);
func Signal_publickey_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_publickey_serialize(out, obj)
}

// SignalFfiError* signal_publickey_verify(bool* out, SignalConstPointerPublicKey key, SignalBorrowedBuffer message, SignalBorrowedBuffer signature);
func Signal_publickey_verify(
	out *C.bool,
	key C.SignalConstPointerPublicKey,
	message C.SignalBorrowedBuffer,
	signature C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_publickey_verify(out, key, message, signature)
}

// SignalFfiError* signal_receipt_credential_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_receipt_credential_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_receipt_credential_check_valid_contents(buffer)
}

// SignalFfiError* signal_receipt_credential_get_receipt_expiration_time(uint64_t* out, const SignalType_FixedArray129_uint8_t* receipt_credential);
func Signal_receipt_credential_get_receipt_expiration_time(
	out *C.uint64_t,
	receipt_credential *C.SignalType_FixedArray129_uint8_t,
) *C.SignalFfiError {
	return C.signal_receipt_credential_get_receipt_expiration_time(out, receipt_credential)
}

// SignalFfiError* signal_receipt_credential_get_receipt_level(uint64_t* out, const SignalType_FixedArray129_uint8_t* receipt_credential);
func Signal_receipt_credential_get_receipt_level(
	out *C.uint64_t,
	receipt_credential *C.SignalType_FixedArray129_uint8_t,
) *C.SignalFfiError {
	return C.signal_receipt_credential_get_receipt_level(out, receipt_credential)
}

// SignalFfiError* signal_receipt_credential_presentation_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_receipt_credential_presentation_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_receipt_credential_presentation_check_valid_contents(buffer)
}

// SignalFfiError* signal_receipt_credential_presentation_get_receipt_expiration_time(uint64_t* out, const SignalType_FixedArray329_uint8_t* presentation);
func Signal_receipt_credential_presentation_get_receipt_expiration_time(
	out *C.uint64_t,
	presentation *C.SignalType_FixedArray329_uint8_t,
) *C.SignalFfiError {
	return C.signal_receipt_credential_presentation_get_receipt_expiration_time(out, presentation)
}

// SignalFfiError* signal_receipt_credential_presentation_get_receipt_level(uint64_t* out, const SignalType_FixedArray329_uint8_t* presentation);
func Signal_receipt_credential_presentation_get_receipt_level(
	out *C.uint64_t,
	presentation *C.SignalType_FixedArray329_uint8_t,
) *C.SignalFfiError {
	return C.signal_receipt_credential_presentation_get_receipt_level(out, presentation)
}

// SignalFfiError* signal_receipt_credential_presentation_get_receipt_serial(SignalType_FixedArray16_uint8_t* out, const SignalType_FixedArray329_uint8_t* presentation);
func Signal_receipt_credential_presentation_get_receipt_serial(
	out *C.SignalType_FixedArray16_uint8_t,
	presentation *C.SignalType_FixedArray329_uint8_t,
) *C.SignalFfiError {
	return C.signal_receipt_credential_presentation_get_receipt_serial(out, presentation)
}

// SignalFfiError* signal_receipt_credential_request_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_receipt_credential_request_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_receipt_credential_request_check_valid_contents(buffer)
}

// SignalFfiError* signal_receipt_credential_request_context_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_receipt_credential_request_context_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_receipt_credential_request_context_check_valid_contents(buffer)
}

// SignalFfiError* signal_receipt_credential_request_context_get_request(SignalType_FixedArray97_uint8_t* out, const SignalType_FixedArray177_uint8_t* request_context);
func Signal_receipt_credential_request_context_get_request(
	out *C.SignalType_FixedArray97_uint8_t,
	request_context *C.SignalType_FixedArray177_uint8_t,
) *C.SignalFfiError {
	return C.signal_receipt_credential_request_context_get_request(out, request_context)
}

// SignalFfiError* signal_receipt_credential_response_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_receipt_credential_response_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_receipt_credential_response_check_valid_contents(buffer)
}

// SignalFfiError* signal_register_account_request_create(SignalMutPointerRegisterAccountRequest* out);
func Signal_register_account_request_create(
	out *C.SignalMutPointerRegisterAccountRequest,
) *C.SignalFfiError {
	return C.signal_register_account_request_create(out)
}

// SignalFfiError* signal_register_account_request_destroy(SignalMutPointerRegisterAccountRequest p);
func Signal_register_account_request_destroy(
	p C.SignalMutPointerRegisterAccountRequest,
) *C.SignalFfiError {
	return C.signal_register_account_request_destroy(p)
}

// SignalFfiError* signal_register_account_request_set_account_password(SignalConstPointerRegisterAccountRequest register_account, const int8_t* account_password);
func Signal_register_account_request_set_account_password(
	register_account C.SignalConstPointerRegisterAccountRequest,
	account_password *C.int8_t,
) *C.SignalFfiError {
	return C.signal_register_account_request_set_account_password(register_account, account_password)
}

// SignalFfiError* signal_register_account_request_set_apn_push_token(SignalConstPointerRegisterAccountRequest register_account, const int8_t* apn_push_token);
func Signal_register_account_request_set_apn_push_token(
	register_account C.SignalConstPointerRegisterAccountRequest,
	apn_push_token *C.int8_t,
) *C.SignalFfiError {
	return C.signal_register_account_request_set_apn_push_token(register_account, apn_push_token)
}

// SignalFfiError* signal_register_account_request_set_identity_pq_last_resort_pre_key(SignalConstPointerRegisterAccountRequest register_account, uint8_t identity_type, SignalFfiSignedPublicPreKey pq_last_resort_pre_key);
func Signal_register_account_request_set_identity_pq_last_resort_pre_key(
	register_account C.SignalConstPointerRegisterAccountRequest,
	identity_type C.uint8_t,
	pq_last_resort_pre_key C.SignalFfiSignedPublicPreKey,
) *C.SignalFfiError {
	return C.signal_register_account_request_set_identity_pq_last_resort_pre_key(register_account, identity_type, pq_last_resort_pre_key)
}

// SignalFfiError* signal_register_account_request_set_identity_public_key(SignalConstPointerRegisterAccountRequest register_account, uint8_t identity_type, SignalConstPointerPublicKey identity_key);
func Signal_register_account_request_set_identity_public_key(
	register_account C.SignalConstPointerRegisterAccountRequest,
	identity_type C.uint8_t,
	identity_key C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_register_account_request_set_identity_public_key(register_account, identity_type, identity_key)
}

// SignalFfiError* signal_register_account_request_set_identity_signed_pre_key(SignalConstPointerRegisterAccountRequest register_account, uint8_t identity_type, SignalFfiSignedPublicPreKey signed_pre_key);
func Signal_register_account_request_set_identity_signed_pre_key(
	register_account C.SignalConstPointerRegisterAccountRequest,
	identity_type C.uint8_t,
	signed_pre_key C.SignalFfiSignedPublicPreKey,
) *C.SignalFfiError {
	return C.signal_register_account_request_set_identity_signed_pre_key(register_account, identity_type, signed_pre_key)
}

// SignalFfiError* signal_register_account_request_set_skip_device_transfer(SignalConstPointerRegisterAccountRequest register_account);
func Signal_register_account_request_set_skip_device_transfer(
	register_account C.SignalConstPointerRegisterAccountRequest,
) *C.SignalFfiError {
	return C.signal_register_account_request_set_skip_device_transfer(register_account)
}

// SignalFfiError* signal_register_account_response_destroy(SignalMutPointerRegisterAccountResponse p);
func Signal_register_account_response_destroy(
	p C.SignalMutPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_destroy(p)
}

// SignalFfiError* signal_register_account_response_get_entitlement_backup_expiration_seconds(uint64_t* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_entitlement_backup_expiration_seconds(
	out *C.uint64_t,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_entitlement_backup_expiration_seconds(out, response)
}

// SignalFfiError* signal_register_account_response_get_entitlement_backup_level(uint64_t* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_entitlement_backup_level(
	out *C.uint64_t,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_entitlement_backup_level(out, response)
}

// SignalFfiError* signal_register_account_response_get_entitlement_badges(SignalOwnedBufferOfFfiRegisterResponseBadge* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_entitlement_badges(
	out *C.SignalOwnedBufferOfFfiRegisterResponseBadge,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_entitlement_badges(out, response)
}

// SignalFfiError* signal_register_account_response_get_identity(SignalType_FixedArray17_uint8_t* out, SignalConstPointerRegisterAccountResponse response, uint8_t identity_type);
func Signal_register_account_response_get_identity(
	out *C.SignalType_FixedArray17_uint8_t,
	response C.SignalConstPointerRegisterAccountResponse,
	identity_type C.uint8_t,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_identity(out, response, identity_type)
}

// SignalFfiError* signal_register_account_response_get_number(SignalCStringPtr* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_number(
	out *C.SignalCStringPtr,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_number(out, response)
}

// SignalFfiError* signal_register_account_response_get_reregistration(bool* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_reregistration(
	out *C.bool,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_reregistration(out, response)
}

// SignalFfiError* signal_register_account_response_get_storage_capable(bool* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_storage_capable(
	out *C.bool,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_storage_capable(out, response)
}

// SignalFfiError* signal_register_account_response_get_username_hash(SignalOwnedBuffer* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_username_hash(
	out *C.SignalOwnedBuffer,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_username_hash(out, response)
}

// SignalFfiError* signal_register_account_response_get_username_link_handle(SignalOptionalUuid* out, SignalConstPointerRegisterAccountResponse response);
func Signal_register_account_response_get_username_link_handle(
	out *C.SignalOptionalUuid,
	response C.SignalConstPointerRegisterAccountResponse,
) *C.SignalFfiError {
	return C.signal_register_account_response_get_username_link_handle(out, response)
}

// SignalFfiError* signal_registration_account_attributes_create(SignalMutPointerRegistrationAccountAttributes* out, SignalBorrowedBuffer recovery_password, uint16_t aci_registration_id, uint16_t pni_registration_id, const int8_t* registration_lock, const SignalType_FixedArray16_uint8_t* unidentified_access_key, bool unrestricted_unidentified_access, SignalBorrowedBytestringArray capabilities, bool discoverable_by_phone_number);
func Signal_registration_account_attributes_create(
	out *C.SignalMutPointerRegistrationAccountAttributes,
	recovery_password C.SignalBorrowedBuffer,
	aci_registration_id C.uint16_t,
	pni_registration_id C.uint16_t,
	registration_lock *C.int8_t,
	unidentified_access_key *C.SignalType_FixedArray16_uint8_t,
	unrestricted_unidentified_access C.bool,
	capabilities C.SignalBorrowedBytestringArray,
	discoverable_by_phone_number C.bool,
) *C.SignalFfiError {
	return C.signal_registration_account_attributes_create(out, recovery_password, aci_registration_id, pni_registration_id, registration_lock, unidentified_access_key, unrestricted_unidentified_access, capabilities, discoverable_by_phone_number)
}

// SignalFfiError* signal_registration_account_attributes_destroy(SignalMutPointerRegistrationAccountAttributes p);
func Signal_registration_account_attributes_destroy(
	p C.SignalMutPointerRegistrationAccountAttributes,
) *C.SignalFfiError {
	return C.signal_registration_account_attributes_destroy(p)
}

// SignalFfiError* signal_registration_service_check_svr2_credentials(SignalCPromiseFfiCheckSvr2CredentialsResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerRegistrationService service, SignalBorrowedBytestringArray svr_tokens);
func Signal_registration_service_check_svr2_credentials(
	promise *C.SignalCPromiseFfiCheckSvr2CredentialsResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	service C.SignalConstPointerRegistrationService,
	svr_tokens C.SignalBorrowedBytestringArray,
) *C.SignalFfiError {
	return C.signal_registration_service_check_svr2_credentials(promise, async_runtime, service, svr_tokens)
}

// SignalFfiError* signal_registration_service_create_session(SignalCPromiseMutPointerRegistrationService* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalFfiRegistrationCreateSessionRequest create_session, SignalConstPointerFfiConnectChatBridgeStruct connect_chat);
func Signal_registration_service_create_session(
	promise *C.SignalCPromiseMutPointerRegistrationService,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	create_session C.SignalFfiRegistrationCreateSessionRequest,
	connect_chat C.SignalConstPointerFfiConnectChatBridgeStruct,
) *C.SignalFfiError {
	return C.signal_registration_service_create_session(promise, async_runtime, create_session, connect_chat)
}

// SignalFfiError* signal_registration_service_destroy(SignalMutPointerRegistrationService p);
func Signal_registration_service_destroy(
	p C.SignalMutPointerRegistrationService,
) *C.SignalFfiError {
	return C.signal_registration_service_destroy(p)
}

// SignalFfiError* signal_registration_service_register_account(SignalCPromiseMutPointerRegisterAccountResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerRegistrationService service, SignalConstPointerRegisterAccountRequest register_account, SignalConstPointerRegistrationAccountAttributes account_attributes);
func Signal_registration_service_register_account(
	promise *C.SignalCPromiseMutPointerRegisterAccountResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	service C.SignalConstPointerRegistrationService,
	register_account C.SignalConstPointerRegisterAccountRequest,
	account_attributes C.SignalConstPointerRegistrationAccountAttributes,
) *C.SignalFfiError {
	return C.signal_registration_service_register_account(promise, async_runtime, service, register_account, account_attributes)
}

// SignalFfiError* signal_registration_service_registration_session(SignalMutPointerRegistrationSession* out, SignalConstPointerRegistrationService service);
func Signal_registration_service_registration_session(
	out *C.SignalMutPointerRegistrationSession,
	service C.SignalConstPointerRegistrationService,
) *C.SignalFfiError {
	return C.signal_registration_service_registration_session(out, service)
}

// SignalFfiError* signal_registration_service_request_push_challenge(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerRegistrationService service, const int8_t* push_token);
func Signal_registration_service_request_push_challenge(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	service C.SignalConstPointerRegistrationService,
	push_token *C.int8_t,
) *C.SignalFfiError {
	return C.signal_registration_service_request_push_challenge(promise, async_runtime, service, push_token)
}

// SignalFfiError* signal_registration_service_request_verification_code(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerRegistrationService service, const int8_t* transport, const int8_t* client, SignalBorrowedBytestringArray languages);
func Signal_registration_service_request_verification_code(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	service C.SignalConstPointerRegistrationService,
	transport *C.int8_t,
	client *C.int8_t,
	languages C.SignalBorrowedBytestringArray,
) *C.SignalFfiError {
	return C.signal_registration_service_request_verification_code(promise, async_runtime, service, transport, client, languages)
}

// SignalFfiError* signal_registration_service_reregister_account(SignalCPromiseMutPointerRegisterAccountResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerFfiConnectChatBridgeStruct connect_chat, const int8_t* number, SignalConstPointerRegisterAccountRequest register_account, SignalConstPointerRegistrationAccountAttributes account_attributes);
func Signal_registration_service_reregister_account(
	promise *C.SignalCPromiseMutPointerRegisterAccountResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	connect_chat C.SignalConstPointerFfiConnectChatBridgeStruct,
	number *C.int8_t,
	register_account C.SignalConstPointerRegisterAccountRequest,
	account_attributes C.SignalConstPointerRegistrationAccountAttributes,
) *C.SignalFfiError {
	return C.signal_registration_service_reregister_account(promise, async_runtime, connect_chat, number, register_account, account_attributes)
}

// SignalFfiError* signal_registration_service_resume_session(SignalCPromiseMutPointerRegistrationService* promise, SignalConstPointerTokioAsyncContext async_runtime, const int8_t* session_id, const int8_t* number, SignalConstPointerFfiConnectChatBridgeStruct connect_chat);
func Signal_registration_service_resume_session(
	promise *C.SignalCPromiseMutPointerRegistrationService,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	session_id *C.int8_t,
	number *C.int8_t,
	connect_chat C.SignalConstPointerFfiConnectChatBridgeStruct,
) *C.SignalFfiError {
	return C.signal_registration_service_resume_session(promise, async_runtime, session_id, number, connect_chat)
}

// SignalFfiError* signal_registration_service_session_id(SignalCStringPtr* out, SignalConstPointerRegistrationService service);
func Signal_registration_service_session_id(
	out *C.SignalCStringPtr,
	service C.SignalConstPointerRegistrationService,
) *C.SignalFfiError {
	return C.signal_registration_service_session_id(out, service)
}

// SignalFfiError* signal_registration_service_submit_captcha(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerRegistrationService service, const int8_t* captcha_value);
func Signal_registration_service_submit_captcha(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	service C.SignalConstPointerRegistrationService,
	captcha_value *C.int8_t,
) *C.SignalFfiError {
	return C.signal_registration_service_submit_captcha(promise, async_runtime, service, captcha_value)
}

// SignalFfiError* signal_registration_service_submit_push_challenge(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerRegistrationService service, const int8_t* push_challenge);
func Signal_registration_service_submit_push_challenge(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	service C.SignalConstPointerRegistrationService,
	push_challenge *C.int8_t,
) *C.SignalFfiError {
	return C.signal_registration_service_submit_push_challenge(promise, async_runtime, service, push_challenge)
}

// SignalFfiError* signal_registration_service_submit_verification_code(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerRegistrationService service, const int8_t* code);
func Signal_registration_service_submit_verification_code(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	service C.SignalConstPointerRegistrationService,
	code *C.int8_t,
) *C.SignalFfiError {
	return C.signal_registration_service_submit_verification_code(promise, async_runtime, service, code)
}

// SignalFfiError* signal_registration_session_destroy(SignalMutPointerRegistrationSession p);
func Signal_registration_session_destroy(
	p C.SignalMutPointerRegistrationSession,
) *C.SignalFfiError {
	return C.signal_registration_session_destroy(p)
}

// SignalFfiError* signal_registration_session_get_allowed_to_request_code(bool* out, SignalConstPointerRegistrationSession session);
func Signal_registration_session_get_allowed_to_request_code(
	out *C.bool,
	session C.SignalConstPointerRegistrationSession,
) *C.SignalFfiError {
	return C.signal_registration_session_get_allowed_to_request_code(out, session)
}

// SignalFfiError* signal_registration_session_get_next_call_seconds(uint32_t* out, SignalConstPointerRegistrationSession session);
func Signal_registration_session_get_next_call_seconds(
	out *C.uint32_t,
	session C.SignalConstPointerRegistrationSession,
) *C.SignalFfiError {
	return C.signal_registration_session_get_next_call_seconds(out, session)
}

// SignalFfiError* signal_registration_session_get_next_sms_seconds(uint32_t* out, SignalConstPointerRegistrationSession session);
func Signal_registration_session_get_next_sms_seconds(
	out *C.uint32_t,
	session C.SignalConstPointerRegistrationSession,
) *C.SignalFfiError {
	return C.signal_registration_session_get_next_sms_seconds(out, session)
}

// SignalFfiError* signal_registration_session_get_next_verification_attempt_seconds(uint32_t* out, SignalConstPointerRegistrationSession session);
func Signal_registration_session_get_next_verification_attempt_seconds(
	out *C.uint32_t,
	session C.SignalConstPointerRegistrationSession,
) *C.SignalFfiError {
	return C.signal_registration_session_get_next_verification_attempt_seconds(out, session)
}

// SignalFfiError* signal_registration_session_get_requested_information(SignalOwnedBuffer* out, SignalConstPointerRegistrationSession session);
func Signal_registration_session_get_requested_information(
	out *C.SignalOwnedBuffer,
	session C.SignalConstPointerRegistrationSession,
) *C.SignalFfiError {
	return C.signal_registration_session_get_requested_information(out, session)
}

// SignalFfiError* signal_registration_session_get_verified(bool* out, SignalConstPointerRegistrationSession session);
func Signal_registration_session_get_verified(
	out *C.bool,
	session C.SignalConstPointerRegistrationSession,
) *C.SignalFfiError {
	return C.signal_registration_session_get_verified(out, session)
}

// SignalFfiError* signal_sealed_sender_multi_recipient_encrypt(SignalOwnedBuffer* out, SignalBorrowedSliceOfConstPointerProtocolAddress recipients, SignalBorrowedSliceOfConstPointerSessionRecord recipient_sessions, SignalBorrowedBuffer excluded_recipients, SignalConstPointerUnidentifiedSenderMessageContent content, SignalConstPointerFfiIdentityKeyStoreStruct identity_key_store);
func Signal_sealed_sender_multi_recipient_encrypt(
	out *C.SignalOwnedBuffer,
	recipients C.SignalBorrowedSliceOfConstPointerProtocolAddress,
	recipient_sessions C.SignalBorrowedSliceOfConstPointerSessionRecord,
	excluded_recipients C.SignalBorrowedBuffer,
	content C.SignalConstPointerUnidentifiedSenderMessageContent,
	identity_key_store C.SignalConstPointerFfiIdentityKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_sealed_sender_multi_recipient_encrypt(out, recipients, recipient_sessions, excluded_recipients, content, identity_key_store)
}

// SignalFfiError* signal_sealed_sender_multi_recipient_message_for_single_recipient(SignalOwnedBuffer* out, SignalBorrowedBuffer encoded_multi_recipient_message);
func Signal_sealed_sender_multi_recipient_message_for_single_recipient(
	out *C.SignalOwnedBuffer,
	encoded_multi_recipient_message C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sealed_sender_multi_recipient_message_for_single_recipient(out, encoded_multi_recipient_message)
}

// SignalFfiError* signal_sealed_session_cipher_decrypt_to_usmc(SignalMutPointerUnidentifiedSenderMessageContent* out, SignalBorrowedBuffer ctext, SignalConstPointerFfiIdentityKeyStoreStruct identity_store);
func Signal_sealed_session_cipher_decrypt_to_usmc(
	out *C.SignalMutPointerUnidentifiedSenderMessageContent,
	ctext C.SignalBorrowedBuffer,
	identity_store C.SignalConstPointerFfiIdentityKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_sealed_session_cipher_decrypt_to_usmc(out, ctext, identity_store)
}

// SignalFfiError* signal_sealed_session_cipher_encrypt(SignalOwnedBuffer* out, SignalConstPointerProtocolAddress destination, SignalConstPointerUnidentifiedSenderMessageContent content, SignalConstPointerFfiIdentityKeyStoreStruct identity_key_store);
func Signal_sealed_session_cipher_encrypt(
	out *C.SignalOwnedBuffer,
	destination C.SignalConstPointerProtocolAddress,
	content C.SignalConstPointerUnidentifiedSenderMessageContent,
	identity_key_store C.SignalConstPointerFfiIdentityKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_sealed_session_cipher_encrypt(out, destination, content, identity_key_store)
}

// SignalFfiError* signal_secure_value_recovery_for_backups_create_new_backup_chain(SignalOwnedBuffer* out, uint8_t environment, const SignalType_FixedArray32_uint8_t* backup_key);
func Signal_secure_value_recovery_for_backups_create_new_backup_chain(
	out *C.SignalOwnedBuffer,
	environment C.uint8_t,
	backup_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_secure_value_recovery_for_backups_create_new_backup_chain(out, environment, backup_key)
}

// SignalFfiError* signal_secure_value_recovery_for_backups_remove_backup(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerConnectionManager connection_manager, const int8_t* username, const int8_t* password);
func Signal_secure_value_recovery_for_backups_remove_backup(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	connection_manager C.SignalConstPointerConnectionManager,
	username *C.int8_t,
	password *C.int8_t,
) *C.SignalFfiError {
	return C.signal_secure_value_recovery_for_backups_remove_backup(promise, async_runtime, connection_manager, username, password)
}

// SignalFfiError* signal_secure_value_recovery_for_backups_restore_backup_from_server(SignalCPromiseMutPointerBackupRestoreResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, const SignalType_FixedArray32_uint8_t* backup_key, SignalBorrowedBuffer metadata, SignalConstPointerConnectionManager connection_manager, const int8_t* username, const int8_t* password);
func Signal_secure_value_recovery_for_backups_restore_backup_from_server(
	promise *C.SignalCPromiseMutPointerBackupRestoreResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	metadata C.SignalBorrowedBuffer,
	connection_manager C.SignalConstPointerConnectionManager,
	username *C.int8_t,
	password *C.int8_t,
) *C.SignalFfiError {
	return C.signal_secure_value_recovery_for_backups_restore_backup_from_server(promise, async_runtime, backup_key, metadata, connection_manager, username, password)
}

// SignalFfiError* signal_secure_value_recovery_for_backups_store_backup(SignalCPromiseMutPointerBackupStoreResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, const SignalType_FixedArray32_uint8_t* backup_key, SignalBorrowedBuffer previous_secret_data, SignalConstPointerConnectionManager connection_manager, const int8_t* username, const int8_t* password);
func Signal_secure_value_recovery_for_backups_store_backup(
	promise *C.SignalCPromiseMutPointerBackupStoreResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	backup_key *C.SignalType_FixedArray32_uint8_t,
	previous_secret_data C.SignalBorrowedBuffer,
	connection_manager C.SignalConstPointerConnectionManager,
	username *C.int8_t,
	password *C.int8_t,
) *C.SignalFfiError {
	return C.signal_secure_value_recovery_for_backups_store_backup(promise, async_runtime, backup_key, previous_secret_data, connection_manager, username, password)
}

// SignalFfiError* signal_sender_certificate_clone(SignalMutPointerSenderCertificate* new_obj, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_clone(
	new_obj *C.SignalMutPointerSenderCertificate,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_clone(new_obj, obj)
}

// SignalFfiError* signal_sender_certificate_deserialize(SignalMutPointerSenderCertificate* out, SignalBorrowedBuffer data);
func Signal_sender_certificate_deserialize(
	out *C.SignalMutPointerSenderCertificate,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sender_certificate_deserialize(out, data)
}

// SignalFfiError* signal_sender_certificate_destroy(SignalMutPointerSenderCertificate p);
func Signal_sender_certificate_destroy(
	p C.SignalMutPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_destroy(p)
}

// SignalFfiError* signal_sender_certificate_get_certificate(SignalOwnedBuffer* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_certificate(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_certificate(out, obj)
}

// SignalFfiError* signal_sender_certificate_get_device_id(uint32_t* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_device_id(
	out *C.uint32_t,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_device_id(out, obj)
}

// SignalFfiError* signal_sender_certificate_get_expiration(uint64_t* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_expiration(
	out *C.uint64_t,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_expiration(out, obj)
}

// SignalFfiError* signal_sender_certificate_get_key(SignalMutPointerPublicKey* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_key(
	out *C.SignalMutPointerPublicKey,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_key(out, obj)
}

// SignalFfiError* signal_sender_certificate_get_sender_e164(SignalCStringPtr* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_sender_e164(
	out *C.SignalCStringPtr,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_sender_e164(out, obj)
}

// SignalFfiError* signal_sender_certificate_get_sender_uuid(SignalCStringPtr* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_sender_uuid(
	out *C.SignalCStringPtr,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_sender_uuid(out, obj)
}

// SignalFfiError* signal_sender_certificate_get_serialized(SignalOwnedBuffer* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_serialized(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_serialized(out, obj)
}

// SignalFfiError* signal_sender_certificate_get_server_certificate(SignalMutPointerServerCertificate* out, SignalConstPointerSenderCertificate cert);
func Signal_sender_certificate_get_server_certificate(
	out *C.SignalMutPointerServerCertificate,
	cert C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_server_certificate(out, cert)
}

// SignalFfiError* signal_sender_certificate_get_signature(SignalOwnedBuffer* out, SignalConstPointerSenderCertificate obj);
func Signal_sender_certificate_get_signature(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderCertificate,
) *C.SignalFfiError {
	return C.signal_sender_certificate_get_signature(out, obj)
}

// SignalFfiError* signal_sender_certificate_new(SignalMutPointerSenderCertificate* out, const int8_t* sender_uuid, const int8_t* sender_e164, uint32_t sender_device_id, SignalConstPointerPublicKey sender_key, uint64_t expiration, SignalConstPointerServerCertificate signer_cert, SignalConstPointerPrivateKey signer_key);
func Signal_sender_certificate_new(
	out *C.SignalMutPointerSenderCertificate,
	sender_uuid *C.int8_t,
	sender_e164 *C.int8_t,
	sender_device_id C.uint32_t,
	sender_key C.SignalConstPointerPublicKey,
	expiration C.uint64_t,
	signer_cert C.SignalConstPointerServerCertificate,
	signer_key C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_sender_certificate_new(out, sender_uuid, sender_e164, sender_device_id, sender_key, expiration, signer_cert, signer_key)
}

// SignalFfiError* signal_sender_certificate_validate(bool* out, SignalConstPointerSenderCertificate cert, SignalBorrowedSliceOfConstPointerPublicKey trust_roots, uint64_t time);
func Signal_sender_certificate_validate(
	out *C.bool,
	cert C.SignalConstPointerSenderCertificate,
	trust_roots C.SignalBorrowedSliceOfConstPointerPublicKey,
	time C.uint64_t,
) *C.SignalFfiError {
	return C.signal_sender_certificate_validate(out, cert, trust_roots, time)
}

// SignalFfiError* signal_sender_key_distribution_message_clone(SignalMutPointerSenderKeyDistributionMessage* new_obj, SignalConstPointerSenderKeyDistributionMessage obj);
func Signal_sender_key_distribution_message_clone(
	new_obj *C.SignalMutPointerSenderKeyDistributionMessage,
	obj C.SignalConstPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_clone(new_obj, obj)
}

// SignalFfiError* signal_sender_key_distribution_message_create(SignalMutPointerSenderKeyDistributionMessage* out, SignalConstPointerProtocolAddress sender, SignalUuid distribution_id, SignalConstPointerFfiSenderKeyStoreStruct store);
func Signal_sender_key_distribution_message_create(
	out *C.SignalMutPointerSenderKeyDistributionMessage,
	sender C.SignalConstPointerProtocolAddress,
	distribution_id C.SignalUuid,
	store C.SignalConstPointerFfiSenderKeyStoreStruct,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_create(out, sender, distribution_id, store)
}

// SignalFfiError* signal_sender_key_distribution_message_deserialize(SignalMutPointerSenderKeyDistributionMessage* out, SignalBorrowedBuffer data);
func Signal_sender_key_distribution_message_deserialize(
	out *C.SignalMutPointerSenderKeyDistributionMessage,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_deserialize(out, data)
}

// SignalFfiError* signal_sender_key_distribution_message_destroy(SignalMutPointerSenderKeyDistributionMessage p);
func Signal_sender_key_distribution_message_destroy(
	p C.SignalMutPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_destroy(p)
}

// SignalFfiError* signal_sender_key_distribution_message_get_chain_id(uint32_t* out, SignalConstPointerSenderKeyDistributionMessage obj);
func Signal_sender_key_distribution_message_get_chain_id(
	out *C.uint32_t,
	obj C.SignalConstPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_get_chain_id(out, obj)
}

// SignalFfiError* signal_sender_key_distribution_message_get_chain_key(SignalOwnedBuffer* out, SignalConstPointerSenderKeyDistributionMessage obj);
func Signal_sender_key_distribution_message_get_chain_key(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_get_chain_key(out, obj)
}

// SignalFfiError* signal_sender_key_distribution_message_get_distribution_id(SignalUuid* out, SignalConstPointerSenderKeyDistributionMessage obj);
func Signal_sender_key_distribution_message_get_distribution_id(
	out *C.SignalUuid,
	obj C.SignalConstPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_get_distribution_id(out, obj)
}

// SignalFfiError* signal_sender_key_distribution_message_get_iteration(uint32_t* out, SignalConstPointerSenderKeyDistributionMessage obj);
func Signal_sender_key_distribution_message_get_iteration(
	out *C.uint32_t,
	obj C.SignalConstPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_get_iteration(out, obj)
}

// SignalFfiError* signal_sender_key_distribution_message_get_signature_key(SignalMutPointerPublicKey* out, SignalConstPointerSenderKeyDistributionMessage m);
func Signal_sender_key_distribution_message_get_signature_key(
	out *C.SignalMutPointerPublicKey,
	m C.SignalConstPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_get_signature_key(out, m)
}

// SignalFfiError* signal_sender_key_distribution_message_new(SignalMutPointerSenderKeyDistributionMessage* out, uint8_t message_version, SignalUuid distribution_id, uint32_t chain_id, uint32_t iteration, SignalBorrowedBuffer chainkey, SignalConstPointerPublicKey pk);
func Signal_sender_key_distribution_message_new(
	out *C.SignalMutPointerSenderKeyDistributionMessage,
	message_version C.uint8_t,
	distribution_id C.SignalUuid,
	chain_id C.uint32_t,
	iteration C.uint32_t,
	chainkey C.SignalBorrowedBuffer,
	pk C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_new(out, message_version, distribution_id, chain_id, iteration, chainkey, pk)
}

// SignalFfiError* signal_sender_key_distribution_message_serialize(SignalOwnedBuffer* out, SignalConstPointerSenderKeyDistributionMessage obj);
func Signal_sender_key_distribution_message_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderKeyDistributionMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_distribution_message_serialize(out, obj)
}

// SignalFfiError* signal_sender_key_message_clone(SignalMutPointerSenderKeyMessage* new_obj, SignalConstPointerSenderKeyMessage obj);
func Signal_sender_key_message_clone(
	new_obj *C.SignalMutPointerSenderKeyMessage,
	obj C.SignalConstPointerSenderKeyMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_message_clone(new_obj, obj)
}

// SignalFfiError* signal_sender_key_message_deserialize(SignalMutPointerSenderKeyMessage* out, SignalBorrowedBuffer data);
func Signal_sender_key_message_deserialize(
	out *C.SignalMutPointerSenderKeyMessage,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sender_key_message_deserialize(out, data)
}

// SignalFfiError* signal_sender_key_message_destroy(SignalMutPointerSenderKeyMessage p);
func Signal_sender_key_message_destroy(
	p C.SignalMutPointerSenderKeyMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_message_destroy(p)
}

// SignalFfiError* signal_sender_key_message_get_chain_id(uint32_t* out, SignalConstPointerSenderKeyMessage obj);
func Signal_sender_key_message_get_chain_id(
	out *C.uint32_t,
	obj C.SignalConstPointerSenderKeyMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_message_get_chain_id(out, obj)
}

// SignalFfiError* signal_sender_key_message_get_cipher_text(SignalOwnedBuffer* out, SignalConstPointerSenderKeyMessage obj);
func Signal_sender_key_message_get_cipher_text(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderKeyMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_message_get_cipher_text(out, obj)
}

// SignalFfiError* signal_sender_key_message_get_distribution_id(SignalUuid* out, SignalConstPointerSenderKeyMessage obj);
func Signal_sender_key_message_get_distribution_id(
	out *C.SignalUuid,
	obj C.SignalConstPointerSenderKeyMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_message_get_distribution_id(out, obj)
}

// SignalFfiError* signal_sender_key_message_get_iteration(uint32_t* out, SignalConstPointerSenderKeyMessage obj);
func Signal_sender_key_message_get_iteration(
	out *C.uint32_t,
	obj C.SignalConstPointerSenderKeyMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_message_get_iteration(out, obj)
}

// SignalFfiError* signal_sender_key_message_new(SignalMutPointerSenderKeyMessage* out, uint8_t message_version, SignalUuid distribution_id, uint32_t chain_id, uint32_t iteration, SignalBorrowedBuffer ciphertext, SignalConstPointerPrivateKey pk);
func Signal_sender_key_message_new(
	out *C.SignalMutPointerSenderKeyMessage,
	message_version C.uint8_t,
	distribution_id C.SignalUuid,
	chain_id C.uint32_t,
	iteration C.uint32_t,
	ciphertext C.SignalBorrowedBuffer,
	pk C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_sender_key_message_new(out, message_version, distribution_id, chain_id, iteration, ciphertext, pk)
}

// SignalFfiError* signal_sender_key_message_serialize(SignalOwnedBuffer* out, SignalConstPointerSenderKeyMessage obj);
func Signal_sender_key_message_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderKeyMessage,
) *C.SignalFfiError {
	return C.signal_sender_key_message_serialize(out, obj)
}

// SignalFfiError* signal_sender_key_message_verify_signature(bool* out, SignalConstPointerSenderKeyMessage skm, SignalConstPointerPublicKey pubkey);
func Signal_sender_key_message_verify_signature(
	out *C.bool,
	skm C.SignalConstPointerSenderKeyMessage,
	pubkey C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_sender_key_message_verify_signature(out, skm, pubkey)
}

// SignalFfiError* signal_sender_key_record_clone(SignalMutPointerSenderKeyRecord* new_obj, SignalConstPointerSenderKeyRecord obj);
func Signal_sender_key_record_clone(
	new_obj *C.SignalMutPointerSenderKeyRecord,
	obj C.SignalConstPointerSenderKeyRecord,
) *C.SignalFfiError {
	return C.signal_sender_key_record_clone(new_obj, obj)
}

// SignalFfiError* signal_sender_key_record_deserialize(SignalMutPointerSenderKeyRecord* out, SignalBorrowedBuffer data);
func Signal_sender_key_record_deserialize(
	out *C.SignalMutPointerSenderKeyRecord,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sender_key_record_deserialize(out, data)
}

// SignalFfiError* signal_sender_key_record_destroy(SignalMutPointerSenderKeyRecord p);
func Signal_sender_key_record_destroy(
	p C.SignalMutPointerSenderKeyRecord,
) *C.SignalFfiError {
	return C.signal_sender_key_record_destroy(p)
}

// SignalFfiError* signal_sender_key_record_serialize(SignalOwnedBuffer* out, SignalConstPointerSenderKeyRecord obj);
func Signal_sender_key_record_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSenderKeyRecord,
) *C.SignalFfiError {
	return C.signal_sender_key_record_serialize(out, obj)
}

// SignalFfiError* signal_server_certificate_clone(SignalMutPointerServerCertificate* new_obj, SignalConstPointerServerCertificate obj);
func Signal_server_certificate_clone(
	new_obj *C.SignalMutPointerServerCertificate,
	obj C.SignalConstPointerServerCertificate,
) *C.SignalFfiError {
	return C.signal_server_certificate_clone(new_obj, obj)
}

// SignalFfiError* signal_server_certificate_deserialize(SignalMutPointerServerCertificate* out, SignalBorrowedBuffer data);
func Signal_server_certificate_deserialize(
	out *C.SignalMutPointerServerCertificate,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_server_certificate_deserialize(out, data)
}

// SignalFfiError* signal_server_certificate_destroy(SignalMutPointerServerCertificate p);
func Signal_server_certificate_destroy(
	p C.SignalMutPointerServerCertificate,
) *C.SignalFfiError {
	return C.signal_server_certificate_destroy(p)
}

// SignalFfiError* signal_server_certificate_get_certificate(SignalOwnedBuffer* out, SignalConstPointerServerCertificate obj);
func Signal_server_certificate_get_certificate(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerServerCertificate,
) *C.SignalFfiError {
	return C.signal_server_certificate_get_certificate(out, obj)
}

// SignalFfiError* signal_server_certificate_get_key(SignalMutPointerPublicKey* out, SignalConstPointerServerCertificate obj);
func Signal_server_certificate_get_key(
	out *C.SignalMutPointerPublicKey,
	obj C.SignalConstPointerServerCertificate,
) *C.SignalFfiError {
	return C.signal_server_certificate_get_key(out, obj)
}

// SignalFfiError* signal_server_certificate_get_key_id(uint32_t* out, SignalConstPointerServerCertificate obj);
func Signal_server_certificate_get_key_id(
	out *C.uint32_t,
	obj C.SignalConstPointerServerCertificate,
) *C.SignalFfiError {
	return C.signal_server_certificate_get_key_id(out, obj)
}

// SignalFfiError* signal_server_certificate_get_serialized(SignalOwnedBuffer* out, SignalConstPointerServerCertificate obj);
func Signal_server_certificate_get_serialized(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerServerCertificate,
) *C.SignalFfiError {
	return C.signal_server_certificate_get_serialized(out, obj)
}

// SignalFfiError* signal_server_certificate_get_signature(SignalOwnedBuffer* out, SignalConstPointerServerCertificate obj);
func Signal_server_certificate_get_signature(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerServerCertificate,
) *C.SignalFfiError {
	return C.signal_server_certificate_get_signature(out, obj)
}

// SignalFfiError* signal_server_certificate_new(SignalMutPointerServerCertificate* out, uint32_t key_id, SignalConstPointerPublicKey server_key, SignalConstPointerPrivateKey trust_root);
func Signal_server_certificate_new(
	out *C.SignalMutPointerServerCertificate,
	key_id C.uint32_t,
	server_key C.SignalConstPointerPublicKey,
	trust_root C.SignalConstPointerPrivateKey,
) *C.SignalFfiError {
	return C.signal_server_certificate_new(out, key_id, server_key, trust_root)
}

// SignalFfiError* signal_server_message_ack_destroy(SignalMutPointerServerMessageAck p);
func Signal_server_message_ack_destroy(
	p C.SignalMutPointerServerMessageAck,
) *C.SignalFfiError {
	return C.signal_server_message_ack_destroy(p)
}

// SignalFfiError* signal_server_message_ack_send(SignalConstPointerServerMessageAck ack);
func Signal_server_message_ack_send(
	ack C.SignalConstPointerServerMessageAck,
) *C.SignalFfiError {
	return C.signal_server_message_ack_send(ack)
}

// SignalFfiError* signal_server_public_params_create_auth_credential_with_pni_presentation_deterministic(SignalOwnedBuffer* out, SignalConstPointerServerPublicParams server_public_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray289_uint8_t* group_secret_params, SignalBorrowedBuffer auth_credential_with_pni_bytes);
func Signal_server_public_params_create_auth_credential_with_pni_presentation_deterministic(
	out *C.SignalOwnedBuffer,
	server_public_params C.SignalConstPointerServerPublicParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	group_secret_params *C.SignalType_FixedArray289_uint8_t,
	auth_credential_with_pni_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_server_public_params_create_auth_credential_with_pni_presentation_deterministic(out, server_public_params, randomness, group_secret_params, auth_credential_with_pni_bytes)
}

// SignalFfiError* signal_server_public_params_create_expiring_profile_key_credential_presentation_deterministic(SignalOwnedBuffer* out, SignalConstPointerServerPublicParams server_public_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray289_uint8_t* group_secret_params, const SignalType_FixedArray153_uint8_t* profile_key_credential);
func Signal_server_public_params_create_expiring_profile_key_credential_presentation_deterministic(
	out *C.SignalOwnedBuffer,
	server_public_params C.SignalConstPointerServerPublicParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	group_secret_params *C.SignalType_FixedArray289_uint8_t,
	profile_key_credential *C.SignalType_FixedArray153_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_public_params_create_expiring_profile_key_credential_presentation_deterministic(out, server_public_params, randomness, group_secret_params, profile_key_credential)
}

// SignalFfiError* signal_server_public_params_create_profile_key_credential_request_context_deterministic(SignalType_FixedArray473_uint8_t* out, SignalConstPointerServerPublicParams server_public_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray17_uint8_t* user_id, const SignalType_FixedArray32_uint8_t* profile_key);
func Signal_server_public_params_create_profile_key_credential_request_context_deterministic(
	out *C.SignalType_FixedArray473_uint8_t,
	server_public_params C.SignalConstPointerServerPublicParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	user_id *C.SignalType_FixedArray17_uint8_t,
	profile_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_public_params_create_profile_key_credential_request_context_deterministic(out, server_public_params, randomness, user_id, profile_key)
}

// SignalFfiError* signal_server_public_params_create_receipt_credential_presentation_deterministic(SignalType_FixedArray329_uint8_t* out, SignalConstPointerServerPublicParams server_public_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray129_uint8_t* receipt_credential);
func Signal_server_public_params_create_receipt_credential_presentation_deterministic(
	out *C.SignalType_FixedArray329_uint8_t,
	server_public_params C.SignalConstPointerServerPublicParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	receipt_credential *C.SignalType_FixedArray129_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_public_params_create_receipt_credential_presentation_deterministic(out, server_public_params, randomness, receipt_credential)
}

// SignalFfiError* signal_server_public_params_create_receipt_credential_request_context_deterministic(SignalType_FixedArray177_uint8_t* out, SignalConstPointerServerPublicParams server_public_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray16_uint8_t* receipt_serial);
func Signal_server_public_params_create_receipt_credential_request_context_deterministic(
	out *C.SignalType_FixedArray177_uint8_t,
	server_public_params C.SignalConstPointerServerPublicParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	receipt_serial *C.SignalType_FixedArray16_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_public_params_create_receipt_credential_request_context_deterministic(out, server_public_params, randomness, receipt_serial)
}

// SignalFfiError* signal_server_public_params_deserialize(SignalMutPointerServerPublicParams* out, SignalBorrowedBuffer buffer);
func Signal_server_public_params_deserialize(
	out *C.SignalMutPointerServerPublicParams,
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_server_public_params_deserialize(out, buffer)
}

// SignalFfiError* signal_server_public_params_destroy(SignalMutPointerServerPublicParams p);
func Signal_server_public_params_destroy(
	p C.SignalMutPointerServerPublicParams,
) *C.SignalFfiError {
	return C.signal_server_public_params_destroy(p)
}

// SignalFfiError* signal_server_public_params_get_endorsement_public_key(SignalOwnedBuffer* out, SignalConstPointerServerPublicParams params);
func Signal_server_public_params_get_endorsement_public_key(
	out *C.SignalOwnedBuffer,
	params C.SignalConstPointerServerPublicParams,
) *C.SignalFfiError {
	return C.signal_server_public_params_get_endorsement_public_key(out, params)
}

// SignalFfiError* signal_server_public_params_receive_auth_credential_with_pni_as_service_id(SignalOwnedBuffer* out, SignalConstPointerServerPublicParams params, const SignalType_FixedArray17_uint8_t* aci, const SignalType_FixedArray17_uint8_t* pni, uint64_t redemption_time, SignalBorrowedBuffer auth_credential_with_pni_response_bytes);
func Signal_server_public_params_receive_auth_credential_with_pni_as_service_id(
	out *C.SignalOwnedBuffer,
	params C.SignalConstPointerServerPublicParams,
	aci *C.SignalType_FixedArray17_uint8_t,
	pni *C.SignalType_FixedArray17_uint8_t,
	redemption_time C.uint64_t,
	auth_credential_with_pni_response_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_server_public_params_receive_auth_credential_with_pni_as_service_id(out, params, aci, pni, redemption_time, auth_credential_with_pni_response_bytes)
}

// SignalFfiError* signal_server_public_params_receive_auth_credential_zkc_without_pni(SignalOwnedBuffer* out, SignalConstPointerServerPublicParams params, const SignalType_FixedArray17_uint8_t* aci, SignalBorrowedBuffer salt, uint64_t redemption_time, SignalBorrowedBuffer auth_credential_with_pni_response_bytes);
func Signal_server_public_params_receive_auth_credential_zkc_without_pni(
	out *C.SignalOwnedBuffer,
	params C.SignalConstPointerServerPublicParams,
	aci *C.SignalType_FixedArray17_uint8_t,
	salt C.SignalBorrowedBuffer,
	redemption_time C.uint64_t,
	auth_credential_with_pni_response_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_server_public_params_receive_auth_credential_zkc_without_pni(out, params, aci, salt, redemption_time, auth_credential_with_pni_response_bytes)
}

// SignalFfiError* signal_server_public_params_receive_expiring_profile_key_credential(SignalType_FixedArray153_uint8_t* out, SignalConstPointerServerPublicParams server_public_params, const SignalType_FixedArray473_uint8_t* request_context, const SignalType_FixedArray497_uint8_t* response, uint64_t current_time_in_seconds);
func Signal_server_public_params_receive_expiring_profile_key_credential(
	out *C.SignalType_FixedArray153_uint8_t,
	server_public_params C.SignalConstPointerServerPublicParams,
	request_context *C.SignalType_FixedArray473_uint8_t,
	response *C.SignalType_FixedArray497_uint8_t,
	current_time_in_seconds C.uint64_t,
) *C.SignalFfiError {
	return C.signal_server_public_params_receive_expiring_profile_key_credential(out, server_public_params, request_context, response, current_time_in_seconds)
}

// SignalFfiError* signal_server_public_params_receive_receipt_credential(SignalType_FixedArray129_uint8_t* out, SignalConstPointerServerPublicParams server_public_params, const SignalType_FixedArray177_uint8_t* request_context, const SignalType_FixedArray409_uint8_t* response);
func Signal_server_public_params_receive_receipt_credential(
	out *C.SignalType_FixedArray129_uint8_t,
	server_public_params C.SignalConstPointerServerPublicParams,
	request_context *C.SignalType_FixedArray177_uint8_t,
	response *C.SignalType_FixedArray409_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_public_params_receive_receipt_credential(out, server_public_params, request_context, response)
}

// SignalFfiError* signal_server_public_params_serialize(SignalOwnedBuffer* out, SignalConstPointerServerPublicParams handle);
func Signal_server_public_params_serialize(
	out *C.SignalOwnedBuffer,
	handle C.SignalConstPointerServerPublicParams,
) *C.SignalFfiError {
	return C.signal_server_public_params_serialize(out, handle)
}

// SignalFfiError* signal_server_public_params_verify_signature(SignalConstPointerServerPublicParams server_public_params, SignalBorrowedBuffer message, const SignalType_FixedArray64_uint8_t* notary_signature);
func Signal_server_public_params_verify_signature(
	server_public_params C.SignalConstPointerServerPublicParams,
	message C.SignalBorrowedBuffer,
	notary_signature *C.SignalType_FixedArray64_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_public_params_verify_signature(server_public_params, message, notary_signature)
}

// SignalFfiError* signal_server_secret_params_deserialize(SignalMutPointerServerSecretParams* out, SignalBorrowedBuffer buffer);
func Signal_server_secret_params_deserialize(
	out *C.SignalMutPointerServerSecretParams,
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_server_secret_params_deserialize(out, buffer)
}

// SignalFfiError* signal_server_secret_params_destroy(SignalMutPointerServerSecretParams p);
func Signal_server_secret_params_destroy(
	p C.SignalMutPointerServerSecretParams,
) *C.SignalFfiError {
	return C.signal_server_secret_params_destroy(p)
}

// SignalFfiError* signal_server_secret_params_generate_deterministic(SignalMutPointerServerSecretParams* out, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_server_secret_params_generate_deterministic(
	out *C.SignalMutPointerServerSecretParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_generate_deterministic(out, randomness)
}

// SignalFfiError* signal_server_secret_params_get_public_params(SignalMutPointerServerPublicParams* out, SignalConstPointerServerSecretParams params);
func Signal_server_secret_params_get_public_params(
	out *C.SignalMutPointerServerPublicParams,
	params C.SignalConstPointerServerSecretParams,
) *C.SignalFfiError {
	return C.signal_server_secret_params_get_public_params(out, params)
}

// SignalFfiError* signal_server_secret_params_issue_auth_credential_with_pni_zkc_deterministic(SignalOwnedBuffer* out, SignalConstPointerServerSecretParams server_secret_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray17_uint8_t* aci, const SignalType_FixedArray17_uint8_t* pni, uint64_t redemption_time);
func Signal_server_secret_params_issue_auth_credential_with_pni_zkc_deterministic(
	out *C.SignalOwnedBuffer,
	server_secret_params C.SignalConstPointerServerSecretParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	aci *C.SignalType_FixedArray17_uint8_t,
	pni *C.SignalType_FixedArray17_uint8_t,
	redemption_time C.uint64_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_issue_auth_credential_with_pni_zkc_deterministic(out, server_secret_params, randomness, aci, pni, redemption_time)
}

// SignalFfiError* signal_server_secret_params_issue_auth_credential_zkc_without_pni_deterministic(SignalOwnedBuffer* out, SignalConstPointerServerSecretParams server_secret_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray17_uint8_t* aci, SignalBorrowedBuffer salt, uint64_t redemption_time);
func Signal_server_secret_params_issue_auth_credential_zkc_without_pni_deterministic(
	out *C.SignalOwnedBuffer,
	server_secret_params C.SignalConstPointerServerSecretParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	aci *C.SignalType_FixedArray17_uint8_t,
	salt C.SignalBorrowedBuffer,
	redemption_time C.uint64_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_issue_auth_credential_zkc_without_pni_deterministic(out, server_secret_params, randomness, aci, salt, redemption_time)
}

// SignalFfiError* signal_server_secret_params_issue_expiring_profile_key_credential_deterministic(SignalType_FixedArray497_uint8_t* out, SignalConstPointerServerSecretParams server_secret_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray329_uint8_t* request, const SignalType_FixedArray17_uint8_t* user_id, const SignalType_FixedArray97_uint8_t* commitment, uint64_t expiration_in_seconds);
func Signal_server_secret_params_issue_expiring_profile_key_credential_deterministic(
	out *C.SignalType_FixedArray497_uint8_t,
	server_secret_params C.SignalConstPointerServerSecretParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	request *C.SignalType_FixedArray329_uint8_t,
	user_id *C.SignalType_FixedArray17_uint8_t,
	commitment *C.SignalType_FixedArray97_uint8_t,
	expiration_in_seconds C.uint64_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_issue_expiring_profile_key_credential_deterministic(out, server_secret_params, randomness, request, user_id, commitment, expiration_in_seconds)
}

// SignalFfiError* signal_server_secret_params_issue_receipt_credential_deterministic(SignalType_FixedArray409_uint8_t* out, SignalConstPointerServerSecretParams server_secret_params, const SignalType_FixedArray32_uint8_t* randomness, const SignalType_FixedArray97_uint8_t* request, uint64_t receipt_expiration_time, uint64_t receipt_level);
func Signal_server_secret_params_issue_receipt_credential_deterministic(
	out *C.SignalType_FixedArray409_uint8_t,
	server_secret_params C.SignalConstPointerServerSecretParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	request *C.SignalType_FixedArray97_uint8_t,
	receipt_expiration_time C.uint64_t,
	receipt_level C.uint64_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_issue_receipt_credential_deterministic(out, server_secret_params, randomness, request, receipt_expiration_time, receipt_level)
}

// SignalFfiError* signal_server_secret_params_serialize(SignalOwnedBuffer* out, SignalConstPointerServerSecretParams handle);
func Signal_server_secret_params_serialize(
	out *C.SignalOwnedBuffer,
	handle C.SignalConstPointerServerSecretParams,
) *C.SignalFfiError {
	return C.signal_server_secret_params_serialize(out, handle)
}

// SignalFfiError* signal_server_secret_params_sign_deterministic(SignalType_FixedArray64_uint8_t* out, SignalConstPointerServerSecretParams params, const SignalType_FixedArray32_uint8_t* randomness, SignalBorrowedBuffer message);
func Signal_server_secret_params_sign_deterministic(
	out *C.SignalType_FixedArray64_uint8_t,
	params C.SignalConstPointerServerSecretParams,
	randomness *C.SignalType_FixedArray32_uint8_t,
	message C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_server_secret_params_sign_deterministic(out, params, randomness, message)
}

// SignalFfiError* signal_server_secret_params_verify_auth_credential_presentation(SignalConstPointerServerSecretParams server_secret_params, const SignalType_FixedArray97_uint8_t* group_public_params, SignalBorrowedBuffer presentation_bytes, uint64_t current_time_in_seconds);
func Signal_server_secret_params_verify_auth_credential_presentation(
	server_secret_params C.SignalConstPointerServerSecretParams,
	group_public_params *C.SignalType_FixedArray97_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
	current_time_in_seconds C.uint64_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_verify_auth_credential_presentation(server_secret_params, group_public_params, presentation_bytes, current_time_in_seconds)
}

// SignalFfiError* signal_server_secret_params_verify_profile_key_credential_presentation(SignalConstPointerServerSecretParams server_secret_params, const SignalType_FixedArray97_uint8_t* group_public_params, SignalBorrowedBuffer presentation_bytes, uint64_t current_time_in_seconds);
func Signal_server_secret_params_verify_profile_key_credential_presentation(
	server_secret_params C.SignalConstPointerServerSecretParams,
	group_public_params *C.SignalType_FixedArray97_uint8_t,
	presentation_bytes C.SignalBorrowedBuffer,
	current_time_in_seconds C.uint64_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_verify_profile_key_credential_presentation(server_secret_params, group_public_params, presentation_bytes, current_time_in_seconds)
}

// SignalFfiError* signal_server_secret_params_verify_receipt_credential_presentation(SignalConstPointerServerSecretParams server_secret_params, const SignalType_FixedArray329_uint8_t* presentation);
func Signal_server_secret_params_verify_receipt_credential_presentation(
	server_secret_params C.SignalConstPointerServerSecretParams,
	presentation *C.SignalType_FixedArray329_uint8_t,
) *C.SignalFfiError {
	return C.signal_server_secret_params_verify_receipt_credential_presentation(server_secret_params, presentation)
}

// SignalFfiError* signal_service_id_parse_from_service_id_binary(SignalType_FixedArray17_uint8_t* out, SignalBorrowedBuffer input);
func Signal_service_id_parse_from_service_id_binary(
	out *C.SignalType_FixedArray17_uint8_t,
	input C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_service_id_parse_from_service_id_binary(out, input)
}

// SignalFfiError* signal_service_id_parse_from_service_id_string(SignalType_FixedArray17_uint8_t* out, const int8_t* input);
func Signal_service_id_parse_from_service_id_string(
	out *C.SignalType_FixedArray17_uint8_t,
	input *C.int8_t,
) *C.SignalFfiError {
	return C.signal_service_id_parse_from_service_id_string(out, input)
}

// SignalFfiError* signal_service_id_service_id_binary(SignalOwnedBuffer* out, const SignalType_FixedArray17_uint8_t* value);
func Signal_service_id_service_id_binary(
	out *C.SignalOwnedBuffer,
	value *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_service_id_service_id_binary(out, value)
}

// SignalFfiError* signal_service_id_service_id_log(SignalCStringPtr* out, const SignalType_FixedArray17_uint8_t* value);
func Signal_service_id_service_id_log(
	out *C.SignalCStringPtr,
	value *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_service_id_service_id_log(out, value)
}

// SignalFfiError* signal_service_id_service_id_string(SignalCStringPtr* out, const SignalType_FixedArray17_uint8_t* value);
func Signal_service_id_service_id_string(
	out *C.SignalCStringPtr,
	value *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_service_id_service_id_string(out, value)
}

// SignalFfiError* signal_session_record_archive_current_state(SignalMutPointerSessionRecord session_record);
func Signal_session_record_archive_current_state(
	session_record C.SignalMutPointerSessionRecord,
) *C.SignalFfiError {
	return C.signal_session_record_archive_current_state(session_record)
}

// SignalFfiError* signal_session_record_clone(SignalMutPointerSessionRecord* new_obj, SignalConstPointerSessionRecord obj);
func Signal_session_record_clone(
	new_obj *C.SignalMutPointerSessionRecord,
	obj C.SignalConstPointerSessionRecord,
) *C.SignalFfiError {
	return C.signal_session_record_clone(new_obj, obj)
}

// SignalFfiError* signal_session_record_current_ratchet_key_matches(bool* out, SignalConstPointerSessionRecord s, SignalConstPointerPublicKey key);
func Signal_session_record_current_ratchet_key_matches(
	out *C.bool,
	s C.SignalConstPointerSessionRecord,
	key C.SignalConstPointerPublicKey,
) *C.SignalFfiError {
	return C.signal_session_record_current_ratchet_key_matches(out, s, key)
}

// SignalFfiError* signal_session_record_deserialize(SignalMutPointerSessionRecord* out, SignalBorrowedBuffer data);
func Signal_session_record_deserialize(
	out *C.SignalMutPointerSessionRecord,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_session_record_deserialize(out, data)
}

// SignalFfiError* signal_session_record_destroy(SignalMutPointerSessionRecord p);
func Signal_session_record_destroy(
	p C.SignalMutPointerSessionRecord,
) *C.SignalFfiError {
	return C.signal_session_record_destroy(p)
}

// SignalFfiError* signal_session_record_get_local_registration_id(uint32_t* out, SignalConstPointerSessionRecord obj);
func Signal_session_record_get_local_registration_id(
	out *C.uint32_t,
	obj C.SignalConstPointerSessionRecord,
) *C.SignalFfiError {
	return C.signal_session_record_get_local_registration_id(out, obj)
}

// SignalFfiError* signal_session_record_get_remote_registration_id(uint32_t* out, SignalConstPointerSessionRecord obj);
func Signal_session_record_get_remote_registration_id(
	out *C.uint32_t,
	obj C.SignalConstPointerSessionRecord,
) *C.SignalFfiError {
	return C.signal_session_record_get_remote_registration_id(out, obj)
}

// SignalFfiError* signal_session_record_has_usable_sender_chain(bool* out, SignalConstPointerSessionRecord s, uint64_t now);
func Signal_session_record_has_usable_sender_chain(
	out *C.bool,
	s C.SignalConstPointerSessionRecord,
	now C.uint64_t,
) *C.SignalFfiError {
	return C.signal_session_record_has_usable_sender_chain(out, s, now)
}

// SignalFfiError* signal_session_record_serialize(SignalOwnedBuffer* out, SignalConstPointerSessionRecord obj);
func Signal_session_record_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSessionRecord,
) *C.SignalFfiError {
	return C.signal_session_record_serialize(out, obj)
}

// SignalFfiError* signal_sgx_client_state_complete_handshake(SignalMutPointerSgxClientState cli, SignalBorrowedBuffer handshake_received);
func Signal_sgx_client_state_complete_handshake(
	cli C.SignalMutPointerSgxClientState,
	handshake_received C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sgx_client_state_complete_handshake(cli, handshake_received)
}

// SignalFfiError* signal_sgx_client_state_destroy(SignalMutPointerSgxClientState p);
func Signal_sgx_client_state_destroy(
	p C.SignalMutPointerSgxClientState,
) *C.SignalFfiError {
	return C.signal_sgx_client_state_destroy(p)
}

// SignalFfiError* signal_sgx_client_state_established_recv(SignalOwnedBuffer* out, SignalMutPointerSgxClientState cli, SignalBorrowedBuffer received_ciphertext);
func Signal_sgx_client_state_established_recv(
	out *C.SignalOwnedBuffer,
	cli C.SignalMutPointerSgxClientState,
	received_ciphertext C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sgx_client_state_established_recv(out, cli, received_ciphertext)
}

// SignalFfiError* signal_sgx_client_state_established_send(SignalOwnedBuffer* out, SignalMutPointerSgxClientState cli, SignalBorrowedBuffer plaintext_to_send);
func Signal_sgx_client_state_established_send(
	out *C.SignalOwnedBuffer,
	cli C.SignalMutPointerSgxClientState,
	plaintext_to_send C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_sgx_client_state_established_send(out, cli, plaintext_to_send)
}

// SignalFfiError* signal_sgx_client_state_initial_request(SignalOwnedBuffer* out, SignalConstPointerSgxClientState obj);
func Signal_sgx_client_state_initial_request(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSgxClientState,
) *C.SignalFfiError {
	return C.signal_sgx_client_state_initial_request(out, obj)
}

// SignalFfiError* signal_signed_pre_key_record_clone(SignalMutPointerSignedPreKeyRecord* new_obj, SignalConstPointerSignedPreKeyRecord obj);
func Signal_signed_pre_key_record_clone(
	new_obj *C.SignalMutPointerSignedPreKeyRecord,
	obj C.SignalConstPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_clone(new_obj, obj)
}

// SignalFfiError* signal_signed_pre_key_record_deserialize(SignalMutPointerSignedPreKeyRecord* out, SignalBorrowedBuffer data);
func Signal_signed_pre_key_record_deserialize(
	out *C.SignalMutPointerSignedPreKeyRecord,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_deserialize(out, data)
}

// SignalFfiError* signal_signed_pre_key_record_destroy(SignalMutPointerSignedPreKeyRecord p);
func Signal_signed_pre_key_record_destroy(
	p C.SignalMutPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_destroy(p)
}

// SignalFfiError* signal_signed_pre_key_record_get_id(uint32_t* out, SignalConstPointerSignedPreKeyRecord obj);
func Signal_signed_pre_key_record_get_id(
	out *C.uint32_t,
	obj C.SignalConstPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_get_id(out, obj)
}

// SignalFfiError* signal_signed_pre_key_record_get_private_key(SignalMutPointerPrivateKey* out, SignalConstPointerSignedPreKeyRecord obj);
func Signal_signed_pre_key_record_get_private_key(
	out *C.SignalMutPointerPrivateKey,
	obj C.SignalConstPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_get_private_key(out, obj)
}

// SignalFfiError* signal_signed_pre_key_record_get_public_key(SignalMutPointerPublicKey* out, SignalConstPointerSignedPreKeyRecord obj);
func Signal_signed_pre_key_record_get_public_key(
	out *C.SignalMutPointerPublicKey,
	obj C.SignalConstPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_get_public_key(out, obj)
}

// SignalFfiError* signal_signed_pre_key_record_get_signature(SignalOwnedBuffer* out, SignalConstPointerSignedPreKeyRecord obj);
func Signal_signed_pre_key_record_get_signature(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_get_signature(out, obj)
}

// SignalFfiError* signal_signed_pre_key_record_get_timestamp(uint64_t* out, SignalConstPointerSignedPreKeyRecord obj);
func Signal_signed_pre_key_record_get_timestamp(
	out *C.uint64_t,
	obj C.SignalConstPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_get_timestamp(out, obj)
}

// SignalFfiError* signal_signed_pre_key_record_new(SignalMutPointerSignedPreKeyRecord* out, uint32_t id, uint64_t timestamp, SignalConstPointerPublicKey pub_key, SignalConstPointerPrivateKey priv_key, SignalBorrowedBuffer signature);
func Signal_signed_pre_key_record_new(
	out *C.SignalMutPointerSignedPreKeyRecord,
	id C.uint32_t,
	timestamp C.uint64_t,
	pub_key C.SignalConstPointerPublicKey,
	priv_key C.SignalConstPointerPrivateKey,
	signature C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_new(out, id, timestamp, pub_key, priv_key, signature)
}

// SignalFfiError* signal_signed_pre_key_record_serialize(SignalOwnedBuffer* out, SignalConstPointerSignedPreKeyRecord obj);
func Signal_signed_pre_key_record_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerSignedPreKeyRecord,
) *C.SignalFfiError {
	return C.signal_signed_pre_key_record_serialize(out, obj)
}

// SignalFfiError* signal_svr2_client_new(SignalMutPointerSgxClientState* out, SignalBorrowedBuffer mrenclave, SignalBorrowedBuffer attestation_msg, uint64_t current_timestamp);
func Signal_svr2_client_new(
	out *C.SignalMutPointerSgxClientState,
	mrenclave C.SignalBorrowedBuffer,
	attestation_msg C.SignalBorrowedBuffer,
	current_timestamp C.uint64_t,
) *C.SignalFfiError {
	return C.signal_svr2_client_new(out, mrenclave, attestation_msg, current_timestamp)
}

// SignalFfiError* signal_svr_key_derive_logging_key(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray32_uint8_t* svr_key);
func Signal_svr_key_derive_logging_key(
	out *C.SignalType_FixedArray32_uint8_t,
	svr_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_svr_key_derive_logging_key(out, svr_key)
}

// SignalFfiError* signal_svr_key_derive_registration_lock(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray32_uint8_t* svr_key);
func Signal_svr_key_derive_registration_lock(
	out *C.SignalType_FixedArray32_uint8_t,
	svr_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_svr_key_derive_registration_lock(out, svr_key)
}

// SignalFfiError* signal_svr_key_derive_registration_recovery_password(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray32_uint8_t* svr_key);
func Signal_svr_key_derive_registration_recovery_password(
	out *C.SignalType_FixedArray32_uint8_t,
	svr_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_svr_key_derive_registration_recovery_password(out, svr_key)
}

// SignalFfiError* signal_svr_key_derive_storage_service_key(SignalType_FixedArray32_uint8_t* out, const SignalType_FixedArray32_uint8_t* svr_key);
func Signal_svr_key_derive_storage_service_key(
	out *C.SignalType_FixedArray32_uint8_t,
	svr_key *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_svr_key_derive_storage_service_key(out, svr_key)
}

// SignalFfiError* signal_tokio_async_context_cancel(SignalConstPointerTokioAsyncContext context, uint64_t raw_cancellation_id);
func Signal_tokio_async_context_cancel(
	context C.SignalConstPointerTokioAsyncContext,
	raw_cancellation_id C.uint64_t,
) *C.SignalFfiError {
	return C.signal_tokio_async_context_cancel(context, raw_cancellation_id)
}

// SignalFfiError* signal_tokio_async_context_destroy(SignalMutPointerTokioAsyncContext p);
func Signal_tokio_async_context_destroy(
	p C.SignalMutPointerTokioAsyncContext,
) *C.SignalFfiError {
	return C.signal_tokio_async_context_destroy(p)
}

// SignalFfiError* signal_tokio_async_context_new(SignalMutPointerTokioAsyncContext* out);
func Signal_tokio_async_context_new(
	out *C.SignalMutPointerTokioAsyncContext,
) *C.SignalFfiError {
	return C.signal_tokio_async_context_new(out)
}

// SignalFfiError* signal_unauthenticated_chat_connection_account_exists(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, const SignalType_FixedArray17_uint8_t* account);
func Signal_unauthenticated_chat_connection_account_exists(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	account *C.SignalType_FixedArray17_uint8_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_account_exists(promise, async_runtime, chat, account)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_copy_media(SignalMutPointerCopyBackupMediaStream* out, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, SignalBorrowedSliceOfBridgeCopyBackupMediaItemFfiArg items, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_copy_media(
	out *C.SignalMutPointerCopyBackupMediaStream,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	items C.SignalBorrowedSliceOfBridgeCopyBackupMediaItemFfiArg,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_copy_media(out, chat, credential, server_keys, signing_key, items, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_delete_all(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_delete_all(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_delete_all(promise, async_runtime, chat, credential, server_keys, signing_key, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_delete_media(SignalMutPointerDeleteBackupMediaStream* out, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, SignalBorrowedSliceOfBridgeDeleteBackupMediaItemFfiArg items, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_delete_media(
	out *C.SignalMutPointerDeleteBackupMediaStream,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	items C.SignalBorrowedSliceOfBridgeDeleteBackupMediaItemFfiArg,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_delete_media(out, chat, credential, server_keys, signing_key, items, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_get_cdn_credentials(SignalCPromisePairOfOwnedBufferOfCStringPtrOwnedBufferOfCStringPtr* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, int32_t cdn, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_get_cdn_credentials(
	promise *C.SignalCPromisePairOfOwnedBufferOfCStringPtrOwnedBufferOfCStringPtr,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	cdn C.int32_t,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_get_cdn_credentials(promise, async_runtime, chat, credential, server_keys, signing_key, cdn, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_get_media_backup_info(SignalCPromiseBridgeMediaBackupInfoFfiResult* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_get_media_backup_info(
	promise *C.SignalCPromiseBridgeMediaBackupInfoFfiResult,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_get_media_backup_info(promise, async_runtime, chat, credential, server_keys, signing_key, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_get_media_upload_form(SignalCPromiseFfiUploadForm* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, uint64_t upload_size, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_get_media_upload_form(
	promise *C.SignalCPromiseFfiUploadForm,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	upload_size C.uint64_t,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_get_media_upload_form(promise, async_runtime, chat, credential, server_keys, signing_key, upload_size, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_get_message_backup_info(SignalCPromiseBridgeMessageBackupInfoFfiResult* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_get_message_backup_info(
	promise *C.SignalCPromiseBridgeMessageBackupInfoFfiResult,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_get_message_backup_info(promise, async_runtime, chat, credential, server_keys, signing_key, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_get_svrb_credentials(SignalCPromisePairOfCStringPtrCStringPtr* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_get_svrb_credentials(
	promise *C.SignalCPromisePairOfCStringPtrCStringPtr,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_get_svrb_credentials(promise, async_runtime, chat, credential, server_keys, signing_key, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_get_upload_form(SignalCPromiseFfiUploadForm* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, uint64_t upload_size, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_get_upload_form(
	promise *C.SignalCPromiseFfiUploadForm,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	upload_size C.uint64_t,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_get_upload_form(promise, async_runtime, chat, credential, server_keys, signing_key, upload_size, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_list_media(SignalCPromiseListMediaResponseFfiResult* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, const int8_t* cursor, int32_t limit, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_list_media(
	promise *C.SignalCPromiseListMediaResponseFfiResult,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	cursor *C.int8_t,
	limit C.int32_t,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_list_media(promise, async_runtime, chat, credential, server_keys, signing_key, cursor, limit, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_refresh(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_refresh(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_refresh(promise, async_runtime, chat, credential, server_keys, signing_key, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_backup_set_public_key(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer credential, SignalBorrowedBuffer server_keys, SignalConstPointerPrivateKey signing_key, int64_t rng);
func Signal_unauthenticated_chat_connection_backup_set_public_key(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	credential C.SignalBorrowedBuffer,
	server_keys C.SignalBorrowedBuffer,
	signing_key C.SignalConstPointerPrivateKey,
	rng C.int64_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_backup_set_public_key(promise, async_runtime, chat, credential, server_keys, signing_key, rng)
}

// SignalFfiError* signal_unauthenticated_chat_connection_connect(SignalCPromiseMutPointerUnauthenticatedChatConnection* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerConnectionManager connection_manager, SignalBorrowedBytestringArray languages);
func Signal_unauthenticated_chat_connection_connect(
	promise *C.SignalCPromiseMutPointerUnauthenticatedChatConnection,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	connection_manager C.SignalConstPointerConnectionManager,
	languages C.SignalBorrowedBytestringArray,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_connect(promise, async_runtime, connection_manager, languages)
}

// SignalFfiError* signal_unauthenticated_chat_connection_destroy(SignalMutPointerUnauthenticatedChatConnection p);
func Signal_unauthenticated_chat_connection_destroy(
	p C.SignalMutPointerUnauthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_destroy(p)
}

// SignalFfiError* signal_unauthenticated_chat_connection_disconnect(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat);
func Signal_unauthenticated_chat_connection_disconnect(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_disconnect(promise, async_runtime, chat)
}

// SignalFfiError* signal_unauthenticated_chat_connection_get_pre_keys_access_key_auth(SignalCPromiseFfiPreKeysResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, const SignalType_FixedArray16_uint8_t* auth, const SignalType_FixedArray17_uint8_t* target, int32_t device);
func Signal_unauthenticated_chat_connection_get_pre_keys_access_key_auth(
	promise *C.SignalCPromiseFfiPreKeysResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	auth *C.SignalType_FixedArray16_uint8_t,
	target *C.SignalType_FixedArray17_uint8_t,
	device C.int32_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_get_pre_keys_access_key_auth(promise, async_runtime, chat, auth, target, device)
}

// SignalFfiError* signal_unauthenticated_chat_connection_get_pre_keys_group_auth(SignalCPromiseFfiPreKeysResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer auth, const SignalType_FixedArray17_uint8_t* target, int32_t device);
func Signal_unauthenticated_chat_connection_get_pre_keys_group_auth(
	promise *C.SignalCPromiseFfiPreKeysResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	auth C.SignalBorrowedBuffer,
	target *C.SignalType_FixedArray17_uint8_t,
	device C.int32_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_get_pre_keys_group_auth(promise, async_runtime, chat, auth, target, device)
}

// SignalFfiError* signal_unauthenticated_chat_connection_get_pre_keys_unrestricted_auth(SignalCPromiseFfiPreKeysResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, const SignalType_FixedArray17_uint8_t* target, int32_t device);
func Signal_unauthenticated_chat_connection_get_pre_keys_unrestricted_auth(
	promise *C.SignalCPromiseFfiPreKeysResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	target *C.SignalType_FixedArray17_uint8_t,
	device C.int32_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_get_pre_keys_unrestricted_auth(promise, async_runtime, chat, target, device)
}

// SignalFfiError* signal_unauthenticated_chat_connection_info(SignalMutPointerChatConnectionInfo* out, SignalConstPointerUnauthenticatedChatConnection chat);
func Signal_unauthenticated_chat_connection_info(
	out *C.SignalMutPointerChatConnectionInfo,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_info(out, chat)
}

// SignalFfiError* signal_unauthenticated_chat_connection_init_listener(SignalConstPointerUnauthenticatedChatConnection chat, SignalConstPointerFfiChatListenerStruct listener);
func Signal_unauthenticated_chat_connection_init_listener(
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	listener C.SignalConstPointerFfiChatListenerStruct,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_init_listener(chat, listener)
}

// SignalFfiError* signal_unauthenticated_chat_connection_look_up_username_hash(SignalCPromiseOptionalUuid* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer hash);
func Signal_unauthenticated_chat_connection_look_up_username_hash(
	promise *C.SignalCPromiseOptionalUuid,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	hash C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_look_up_username_hash(promise, async_runtime, chat, hash)
}

// SignalFfiError* signal_unauthenticated_chat_connection_look_up_username_link(SignalCPromiseOptionalPairOfCStringPtrc_uchar32* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalUuid uuid, SignalBorrowedBuffer entropy);
func Signal_unauthenticated_chat_connection_look_up_username_link(
	promise *C.SignalCPromiseOptionalPairOfCStringPtrc_uchar32,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	uuid C.SignalUuid,
	entropy C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_look_up_username_link(promise, async_runtime, chat, uuid, entropy)
}

// SignalFfiError* signal_unauthenticated_chat_connection_send(SignalCPromiseFfiChatResponse* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalConstPointerHttpRequest http_request, uint32_t timeout_millis);
func Signal_unauthenticated_chat_connection_send(
	promise *C.SignalCPromiseFfiChatResponse,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	http_request C.SignalConstPointerHttpRequest,
	timeout_millis C.uint32_t,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_send(promise, async_runtime, chat, http_request, timeout_millis)
}

// SignalFfiError* signal_unauthenticated_chat_connection_send_message(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, const SignalType_FixedArray17_uint8_t* destination, uint64_t timestamp, SignalBorrowedSliceOfu32 device_ids, SignalBorrowedSliceOfu32 registration_ids, SignalBorrowedSliceOfBuffers contents, uint8_t auth_kind, SignalOptionalBorrowedSliceOfc_uchar auth_buffer, bool online_only, bool is_urgent);
func Signal_unauthenticated_chat_connection_send_message(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	destination *C.SignalType_FixedArray17_uint8_t,
	timestamp C.uint64_t,
	device_ids C.SignalBorrowedSliceOfu32,
	registration_ids C.SignalBorrowedSliceOfu32,
	contents C.SignalBorrowedSliceOfBuffers,
	auth_kind C.uint8_t,
	auth_buffer C.SignalOptionalBorrowedSliceOfc_uchar,
	online_only C.bool,
	is_urgent C.bool,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_send_message(promise, async_runtime, chat, destination, timestamp, device_ids, registration_ids, contents, auth_kind, auth_buffer, online_only, is_urgent)
}

// SignalFfiError* signal_unauthenticated_chat_connection_send_multi_recipient_message(SignalCPromiseOwnedBufferOfc_uchar17* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalBorrowedBuffer payload, uint64_t timestamp, SignalBorrowedBuffer auth, bool online_only, bool is_urgent);
func Signal_unauthenticated_chat_connection_send_multi_recipient_message(
	promise *C.SignalCPromiseOwnedBufferOfc_uchar17,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	payload C.SignalBorrowedBuffer,
	timestamp C.uint64_t,
	auth C.SignalBorrowedBuffer,
	online_only C.bool,
	is_urgent C.bool,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_send_multi_recipient_message(promise, async_runtime, chat, payload, timestamp, auth, online_only, is_urgent)
}

// SignalFfiError* signal_unauthenticated_chat_connection_send_raw_grpc(SignalCPromiseOwnedBuffer* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, const int8_t* service, const int8_t* method, SignalBorrowedBuffer payload);
func Signal_unauthenticated_chat_connection_send_raw_grpc(
	promise *C.SignalCPromiseOwnedBuffer,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	service *C.int8_t,
	method *C.int8_t,
	payload C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_send_raw_grpc(promise, async_runtime, chat, service, method, payload)
}

// SignalFfiError* signal_unauthenticated_chat_connection_submit_call_quality_survey(SignalCPromisebool* promise, SignalConstPointerTokioAsyncContext async_runtime, SignalConstPointerUnauthenticatedChatConnection chat, SignalCallQualitySurveyInternalFfiArg survey);
func Signal_unauthenticated_chat_connection_submit_call_quality_survey(
	promise *C.SignalCPromisebool,
	async_runtime C.SignalConstPointerTokioAsyncContext,
	chat C.SignalConstPointerUnauthenticatedChatConnection,
	survey C.SignalCallQualitySurveyInternalFfiArg,
) *C.SignalFfiError {
	return C.signal_unauthenticated_chat_connection_submit_call_quality_survey(promise, async_runtime, chat, survey)
}

// SignalFfiError* signal_unidentified_sender_message_content_deserialize(SignalMutPointerUnidentifiedSenderMessageContent* out, SignalBorrowedBuffer data);
func Signal_unidentified_sender_message_content_deserialize(
	out *C.SignalMutPointerUnidentifiedSenderMessageContent,
	data C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_deserialize(out, data)
}

// SignalFfiError* signal_unidentified_sender_message_content_destroy(SignalMutPointerUnidentifiedSenderMessageContent p);
func Signal_unidentified_sender_message_content_destroy(
	p C.SignalMutPointerUnidentifiedSenderMessageContent,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_destroy(p)
}

// SignalFfiError* signal_unidentified_sender_message_content_get_content_hint(uint32_t* out, SignalConstPointerUnidentifiedSenderMessageContent m);
func Signal_unidentified_sender_message_content_get_content_hint(
	out *C.uint32_t,
	m C.SignalConstPointerUnidentifiedSenderMessageContent,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_get_content_hint(out, m)
}

// SignalFfiError* signal_unidentified_sender_message_content_get_contents(SignalOwnedBuffer* out, SignalConstPointerUnidentifiedSenderMessageContent obj);
func Signal_unidentified_sender_message_content_get_contents(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerUnidentifiedSenderMessageContent,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_get_contents(out, obj)
}

// SignalFfiError* signal_unidentified_sender_message_content_get_group_id_or_empty(SignalOwnedBuffer* out, SignalConstPointerUnidentifiedSenderMessageContent m);
func Signal_unidentified_sender_message_content_get_group_id_or_empty(
	out *C.SignalOwnedBuffer,
	m C.SignalConstPointerUnidentifiedSenderMessageContent,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_get_group_id_or_empty(out, m)
}

// SignalFfiError* signal_unidentified_sender_message_content_get_msg_type(uint8_t* out, SignalConstPointerUnidentifiedSenderMessageContent m);
func Signal_unidentified_sender_message_content_get_msg_type(
	out *C.uint8_t,
	m C.SignalConstPointerUnidentifiedSenderMessageContent,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_get_msg_type(out, m)
}

// SignalFfiError* signal_unidentified_sender_message_content_get_sender_cert(SignalMutPointerSenderCertificate* out, SignalConstPointerUnidentifiedSenderMessageContent m);
func Signal_unidentified_sender_message_content_get_sender_cert(
	out *C.SignalMutPointerSenderCertificate,
	m C.SignalConstPointerUnidentifiedSenderMessageContent,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_get_sender_cert(out, m)
}

// SignalFfiError* signal_unidentified_sender_message_content_new(SignalMutPointerUnidentifiedSenderMessageContent* out, SignalConstPointerCiphertextMessage message, SignalConstPointerSenderCertificate sender, uint32_t content_hint, SignalBorrowedBuffer group_id);
func Signal_unidentified_sender_message_content_new(
	out *C.SignalMutPointerUnidentifiedSenderMessageContent,
	message C.SignalConstPointerCiphertextMessage,
	sender C.SignalConstPointerSenderCertificate,
	content_hint C.uint32_t,
	group_id C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_new(out, message, sender, content_hint, group_id)
}

// SignalFfiError* signal_unidentified_sender_message_content_new_from_content_and_type(SignalMutPointerUnidentifiedSenderMessageContent* out, SignalBorrowedBuffer message_content, uint8_t message_type, SignalConstPointerSenderCertificate sender, uint32_t content_hint, SignalBorrowedBuffer group_id);
func Signal_unidentified_sender_message_content_new_from_content_and_type(
	out *C.SignalMutPointerUnidentifiedSenderMessageContent,
	message_content C.SignalBorrowedBuffer,
	message_type C.uint8_t,
	sender C.SignalConstPointerSenderCertificate,
	content_hint C.uint32_t,
	group_id C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_new_from_content_and_type(out, message_content, message_type, sender, content_hint, group_id)
}

// SignalFfiError* signal_unidentified_sender_message_content_serialize(SignalOwnedBuffer* out, SignalConstPointerUnidentifiedSenderMessageContent obj);
func Signal_unidentified_sender_message_content_serialize(
	out *C.SignalOwnedBuffer,
	obj C.SignalConstPointerUnidentifiedSenderMessageContent,
) *C.SignalFfiError {
	return C.signal_unidentified_sender_message_content_serialize(out, obj)
}

// SignalFfiError* signal_username_candidates_from(SignalBytestringArray* out, const int8_t* nickname, uint32_t min_len, uint32_t max_len);
func Signal_username_candidates_from(
	out *C.SignalBytestringArray,
	nickname *C.int8_t,
	min_len C.uint32_t,
	max_len C.uint32_t,
) *C.SignalFfiError {
	return C.signal_username_candidates_from(out, nickname, min_len, max_len)
}

// SignalFfiError* signal_username_hash(SignalType_FixedArray32_uint8_t* out, const int8_t* username);
func Signal_username_hash(
	out *C.SignalType_FixedArray32_uint8_t,
	username *C.int8_t,
) *C.SignalFfiError {
	return C.signal_username_hash(out, username)
}

// SignalFfiError* signal_username_hash_from_parts(SignalType_FixedArray32_uint8_t* out, const int8_t* nickname, const int8_t* discriminator, uint32_t min_len, uint32_t max_len);
func Signal_username_hash_from_parts(
	out *C.SignalType_FixedArray32_uint8_t,
	nickname *C.int8_t,
	discriminator *C.int8_t,
	min_len C.uint32_t,
	max_len C.uint32_t,
) *C.SignalFfiError {
	return C.signal_username_hash_from_parts(out, nickname, discriminator, min_len, max_len)
}

// SignalFfiError* signal_username_link_create(SignalOwnedBuffer* out, const int8_t* username, SignalBorrowedBuffer entropy);
func Signal_username_link_create(
	out *C.SignalOwnedBuffer,
	username *C.int8_t,
	entropy C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_username_link_create(out, username, entropy)
}

// SignalFfiError* signal_username_link_decrypt_username(SignalCStringPtr* out, SignalBorrowedBuffer entropy, SignalBorrowedBuffer encrypted_username);
func Signal_username_link_decrypt_username(
	out *C.SignalCStringPtr,
	entropy C.SignalBorrowedBuffer,
	encrypted_username C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_username_link_decrypt_username(out, entropy, encrypted_username)
}

// SignalFfiError* signal_username_proof(SignalOwnedBuffer* out, const int8_t* username, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_username_proof(
	out *C.SignalOwnedBuffer,
	username *C.int8_t,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_username_proof(out, username, randomness)
}

// SignalFfiError* signal_username_verify(SignalBorrowedBuffer proof, SignalBorrowedBuffer hash);
func Signal_username_verify(
	proof C.SignalBorrowedBuffer,
	hash C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_username_verify(proof, hash)
}

// SignalFfiError* signal_uuid_ciphertext_check_valid_contents(SignalBorrowedBuffer buffer);
func Signal_uuid_ciphertext_check_valid_contents(
	buffer C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_uuid_ciphertext_check_valid_contents(buffer)
}

// SignalFfiError* signal_validating_mac_destroy(SignalMutPointerValidatingMac p);
func Signal_validating_mac_destroy(
	p C.SignalMutPointerValidatingMac,
) *C.SignalFfiError {
	return C.signal_validating_mac_destroy(p)
}

// SignalFfiError* signal_validating_mac_finalize(int32_t* out, SignalMutPointerValidatingMac mac);
func Signal_validating_mac_finalize(
	out *C.int32_t,
	mac C.SignalMutPointerValidatingMac,
) *C.SignalFfiError {
	return C.signal_validating_mac_finalize(out, mac)
}

// SignalFfiError* signal_validating_mac_initialize(SignalMutPointerValidatingMac* out, SignalBorrowedBuffer key, uint32_t chunk_size, SignalBorrowedBuffer digests);
func Signal_validating_mac_initialize(
	out *C.SignalMutPointerValidatingMac,
	key C.SignalBorrowedBuffer,
	chunk_size C.uint32_t,
	digests C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_validating_mac_initialize(out, key, chunk_size, digests)
}

// SignalFfiError* signal_validating_mac_update(int32_t* out, SignalMutPointerValidatingMac mac, SignalBorrowedBuffer bytes, uint32_t offset, uint32_t length);
func Signal_validating_mac_update(
	out *C.int32_t,
	mac C.SignalMutPointerValidatingMac,
	bytes C.SignalBorrowedBuffer,
	offset C.uint32_t,
	length C.uint32_t,
) *C.SignalFfiError {
	return C.signal_validating_mac_update(out, mac, bytes, offset, length)
}

// SignalFfiError* signal_zk_credential_key_pair_check_valid_contents(SignalBorrowedBuffer key_pair_bytes);
func Signal_zk_credential_key_pair_check_valid_contents(
	key_pair_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_zk_credential_key_pair_check_valid_contents(key_pair_bytes)
}

// SignalFfiError* signal_zk_credential_key_pair_generate_deterministic(SignalOwnedBuffer* out, const SignalType_FixedArray32_uint8_t* randomness);
func Signal_zk_credential_key_pair_generate_deterministic(
	out *C.SignalOwnedBuffer,
	randomness *C.SignalType_FixedArray32_uint8_t,
) *C.SignalFfiError {
	return C.signal_zk_credential_key_pair_generate_deterministic(out, randomness)
}

// SignalFfiError* signal_zk_credential_key_pair_get_public_key(SignalOwnedBuffer* out, SignalBorrowedBuffer key_pair_bytes);
func Signal_zk_credential_key_pair_get_public_key(
	out *C.SignalOwnedBuffer,
	key_pair_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_zk_credential_key_pair_get_public_key(out, key_pair_bytes)
}

// SignalFfiError* signal_zk_credential_public_key_check_valid_contents(SignalBorrowedBuffer public_key_bytes);
func Signal_zk_credential_public_key_check_valid_contents(
	public_key_bytes C.SignalBorrowedBuffer,
) *C.SignalFfiError {
	return C.signal_zk_credential_public_key_check_valid_contents(public_key_bytes)
}
