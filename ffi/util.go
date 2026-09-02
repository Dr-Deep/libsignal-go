package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
void signal_free_buffer(
void signal_free_bytestring_array(
void signal_free_list_of_mismatched_device_errors(
void signal_free_list_of_register_response_badges(
void signal_free_list_of_service_ids(
void signal_free_list_of_strings(
void signal_free_lookup_response_entry_list(
void signal_free_outer_buffer_list_of_prekey_bundles(
void signal_free_owned_buffer_of_max_aligned(
void signal_free_string(
*/

/*
SignalFfiError* signal_bridged_string_map_clone(
SignalFfiError* signal_bridged_string_map_destroy(
SignalFfiError* signal_bridged_string_map_insert(
SignalFfiError* signal_bridged_string_map_new(
*/
