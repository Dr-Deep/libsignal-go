package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
SignalFfiError* signal_aes256_gcm_decryption_destroy(
SignalFfiError* signal_aes256_gcm_decryption_new(
SignalFfiError* signal_aes256_gcm_decryption_update(
SignalFfiError* signal_aes256_gcm_decryption_verify_tag(
SignalFfiError* signal_aes256_gcm_encryption_compute_tag(
SignalFfiError* signal_aes256_gcm_encryption_destroy(
SignalFfiError* signal_aes256_gcm_encryption_new(
SignalFfiError* signal_aes256_gcm_encryption_update(
SignalFfiError* signal_aes256_gcm_siv_decrypt(
SignalFfiError* signal_aes256_gcm_siv_destroy(
SignalFfiError* signal_aes256_gcm_siv_encrypt(
SignalFfiError* signal_aes256_gcm_siv_new(
*/
