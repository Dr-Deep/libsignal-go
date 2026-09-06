package ffi

/*
#include "../../deps/signal_ffi.h"
#include "../../deps/signal_ffi_testing.h"
#include <stdlib.h>
#cgo LDFLAGS: -L${SRCDIR}/../../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/*
void signal_error_free(
SignalFfiError* signal_error_get_address(
SignalFfiError* signal_error_get_invalid_protocol_address(
SignalFfiError* signal_error_get_message(
SignalFfiError* signal_error_get_mismatched_device_errors(
SignalFfiError* signal_error_get_our_fingerprint_version(
SignalFfiError* signal_error_get_rate_limit_challenge(
SignalFfiError* signal_error_get_registration_error_not_deliverable(
SignalFfiError* signal_error_get_registration_lock(
SignalFfiError* signal_error_get_retry_after_seconds(
SignalFfiError* signal_error_get_their_fingerprint_version(
SignalFfiError* signal_error_get_tries_remaining(
uint32_t signal_error_get_type(
SignalFfiError* signal_error_get_unknown_fields(
SignalFfiError* signal_error_get_uuid(
*/

type SignalErrorCode int

const (
	SignalErrorCodeUnknownError                                SignalErrorCode = 1
	SignalErrorCodeInvalidState                                SignalErrorCode = 2
	SignalErrorCodeInternalError                               SignalErrorCode = 3
	SignalErrorCodeNullParameter                               SignalErrorCode = 4
	SignalErrorCodeInvalidArgument                             SignalErrorCode = 5
	SignalErrorCodeInvalidType                                 SignalErrorCode = 6
	SignalErrorCodeInvalidUtf8String                           SignalErrorCode = 7
	SignalErrorCodeCancelled                                   SignalErrorCode = 8
	SignalErrorCodeProtobufError                               SignalErrorCode = 10
	SignalErrorCodeLegacyCiphertextVersion                     SignalErrorCode = 21
	SignalErrorCodeUnknownCiphertextVersion                    SignalErrorCode = 22
	SignalErrorCodeUnrecognizedMessageVersion                  SignalErrorCode = 23
	SignalErrorCodeInvalidMessage                              SignalErrorCode = 30
	SignalErrorCodeSealedSenderSelfSend                        SignalErrorCode = 31
	SignalErrorCodeInvalidKey                                  SignalErrorCode = 40
	SignalErrorCodeInvalidSignature                            SignalErrorCode = 41
	SignalErrorCodeInvalidAttestationData                      SignalErrorCode = 42
	SignalErrorCodeFingerprintVersionMismatch                  SignalErrorCode = 51
	SignalErrorCodeFingerprintParsingError                     SignalErrorCode = 52
	SignalErrorCodeUntrustedIdentity                           SignalErrorCode = 60
	SignalErrorCodeInvalidKeyIdentifier                        SignalErrorCode = 70
	SignalErrorCodeSessionNotFound                             SignalErrorCode = 80
	SignalErrorCodeInvalidRegistrationId                       SignalErrorCode = 81
	SignalErrorCodeInvalidSession                              SignalErrorCode = 82
	SignalErrorCodeInvalidSenderKeySession                     SignalErrorCode = 83
	SignalErrorCodeInvalidProtocolAddress                      SignalErrorCode = 84
	SignalErrorCodeDuplicatedMessage                           SignalErrorCode = 90
	SignalErrorCodeCallbackError                               SignalErrorCode = 100
	SignalErrorCodeVerificationFailure                         SignalErrorCode = 110
	SignalErrorCodeUsernameCannotBeEmpty                       SignalErrorCode = 120
	SignalErrorCodeUsernameCannotStartWithDigit                SignalErrorCode = 121
	SignalErrorCodeUsernameMissingSeparator                    SignalErrorCode = 122
	SignalErrorCodeUsernameBadDiscriminatorCharacter           SignalErrorCode = 123
	SignalErrorCodeUsernameBadNicknameCharacter                SignalErrorCode = 124
	SignalErrorCodeUsernameTooShort                            SignalErrorCode = 125
	SignalErrorCodeUsernameTooLong                             SignalErrorCode = 126
	SignalErrorCodeUsernameLinkInvalidEntropyDataLength        SignalErrorCode = 127
	SignalErrorCodeUsernameLinkInvalid                         SignalErrorCode = 128
	SignalErrorCodeUsernameDiscriminatorCannotBeEmpty          SignalErrorCode = 130
	SignalErrorCodeUsernameDiscriminatorCannotBeZero           SignalErrorCode = 131
	SignalErrorCodeUsernameDiscriminatorCannotBeSingleDigit    SignalErrorCode = 132
	SignalErrorCodeUsernameDiscriminatorCannotHaveLeadingZeros SignalErrorCode = 133
	SignalErrorCodeUsernameDiscriminatorTooLarge               SignalErrorCode = 134
	SignalErrorCodeIoError                                     SignalErrorCode = 140
	SignalErrorCodeInvalidMediaInput                           SignalErrorCode = 141
	SignalErrorCodeUnsupportedMediaInput                       SignalErrorCode = 142
	SignalErrorCodeConnectionTimedOut                          SignalErrorCode = 143
	SignalErrorCodeNetworkProtocol                             SignalErrorCode = 144
	SignalErrorCodeRateLimited                                 SignalErrorCode = 145
	SignalErrorCodeWebSocket                                   SignalErrorCode = 146
	SignalErrorCodeCdsiInvalidToken                            SignalErrorCode = 147
	SignalErrorCodeConnectionFailed                            SignalErrorCode = 148
	SignalErrorCodeChatServiceInactive                         SignalErrorCode = 149
	SignalErrorCodeRequestTimedOut                             SignalErrorCode = 150
	SignalErrorCodeRateLimitChallenge                          SignalErrorCode = 151
	SignalErrorCodePossibleCaptiveNetwork                      SignalErrorCode = 152
	SignalErrorCodeSvrDataMissing                              SignalErrorCode = 160
	SignalErrorCodeSvrRestoreFailed                            SignalErrorCode = 161
	SignalErrorCodeSvrRotationMachineTooManySteps              SignalErrorCode = 162
	SignalErrorCodeSvrRequestFailed                            SignalErrorCode = 163
	SignalErrorCodeAppExpired                                  SignalErrorCode = 170
	SignalErrorCodeDeviceDeregistered                          SignalErrorCode = 171
	SignalErrorCodeConnectionInvalidated                       SignalErrorCode = 172
	SignalErrorCodeConnectedElsewhere                          SignalErrorCode = 173
	SignalErrorCodeBackupValidation                            SignalErrorCode = 180
	SignalErrorCodeRegistrationInvalidSessionId                SignalErrorCode = 190
	SignalErrorCodeRegistrationUnknown                         SignalErrorCode = 192
	SignalErrorCodeRegistrationSessionNotFound                 SignalErrorCode = 193
	SignalErrorCodeRegistrationNotReadyForVerification         SignalErrorCode = 194
	SignalErrorCodeRegistrationSendVerificationCodeFailed      SignalErrorCode = 195
	SignalErrorCodeRegistrationCodeNotDeliverable              SignalErrorCode = 196
	SignalErrorCodeRegistrationSessionUpdateRejected           SignalErrorCode = 197
	SignalErrorCodeRegistrationCredentialsCouldNotBeParsed     SignalErrorCode = 198
	SignalErrorCodeRegistrationDeviceTransferPossible          SignalErrorCode = 199
	SignalErrorCodeRegistrationRecoveryVerificationFailed      SignalErrorCode = 200
	SignalErrorCodeRegistrationLock                            SignalErrorCode = 201
	SignalErrorCodeKeyTransparencyError                        SignalErrorCode = 210
	SignalErrorCodeKeyTransparencyVerificationFailed           SignalErrorCode = 211
	SignalErrorCodeRequestUnauthorized                         SignalErrorCode = 220
	SignalErrorCodeMismatchedDevices                           SignalErrorCode = 221
	SignalErrorCodeServiceIdNotFound                           SignalErrorCode = 222
	SignalErrorCodeUploadTooLarge                              SignalErrorCode = 223
	SignalErrorCodeDeviceIdNotFound                            SignalErrorCode = 224
	SignalErrorCodeUsernameNotAvailable                        SignalErrorCode = 225
	SignalErrorCodeUsernameNotSet                              SignalErrorCode = 226
	SignalErrorCodeUsernameReservationNotFound                 SignalErrorCode = 227
	SignalErrorCodeInvalidReceipt                              SignalErrorCode = 228
	SignalErrorCodeMissingBackupId                             SignalErrorCode = 229
)

type SignalError struct {
	Code    SignalErrorCode
	Message string
}

func (e *SignalError) Error() string {
	if e.Message != "" {
		//! print error message
		return fmt.Sprintf("SignalError: (%d): %s", e.Code, e.Message)
	}

	return fmt.Sprintf("SignalError: (%d)", e.Code)
}

func convertError(signal_err *C.SignalFfiError) error {
	if signal_err == nil {
		return nil
	}
	defer C.signal_error_free(signal_err)

	var (
		err_code = SignalErrorCode(
			C.signal_error_get_type(signal_err),
		)
		err_message C.SignalCStringPtr
	)

	// signal_error_get_message()
	if err := C.signal_error_get_message(&err_message, signal_err); err != nil {
		C.signal_error_free(err)
	}

	return &SignalError{
		Code:    err_code,
		Message: SignalCStringPtrToString(err_message),
	}
}
