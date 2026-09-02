package ffi

/*
#include "../deps/signal_ffi.h"
#include "../deps/signal_ffi_testing.h"
#cgo LDFLAGS: -L${SRCDIR}/../deps -lsignal_ffi -ldl -lpthread -lm -lstdc++
*/
import "C"

/*
* SignalAuthenticatedChatConnection
 */
type AuthenticatedChatConnection struct{}

// C.signal_authenticated_chat_connection_init_listener()
func (conn *AuthenticatedChatConnection) InitListener() SignalError
func (conn *AuthenticatedChatConnection) ClearPushToken() SignalError
func (conn *AuthenticatedChatConnection) ClearRegistrationLock() SignalError
func (conn *AuthenticatedChatConnection) ConfirmUsername() SignalError
func (conn *AuthenticatedChatConnection) Connect() SignalError
func (conn *AuthenticatedChatConnection) DeleteUsernameHash() SignalError
func (conn *AuthenticatedChatConnection) DeleteUsernameLink() SignalError
func (conn *AuthenticatedChatConnection) Destroy() SignalError
func (conn *AuthenticatedChatConnection) Disconnect() SignalError
func (conn *AuthenticatedChatConnection) GetDevices() SignalError
func (conn *AuthenticatedChatConnection) GetUploadForm() SignalError
func (conn *AuthenticatedChatConnection) Info() SignalError
func (conn *AuthenticatedChatConnection) Preconnect() SignalError
func (conn *AuthenticatedChatConnection) RedeemBackupReceipt() SignalError
func (conn *AuthenticatedChatConnection) RemoveDevice() SignalError
func (conn *AuthenticatedChatConnection) ReserveUsernameHash() SignalError
func (conn *AuthenticatedChatConnection) Send() SignalError
func (conn *AuthenticatedChatConnection) SendMessage() SignalError
func (conn *AuthenticatedChatConnection) SendRawGRPC() SignalError
func (conn *AuthenticatedChatConnection) SendSyncMessage() SignalError
func (conn *AuthenticatedChatConnection) SetDeviceName() SignalError
func (conn *AuthenticatedChatConnection) SetDiscoverableByPhoneNumber() SignalError
func (conn *AuthenticatedChatConnection) SetPushToken_apns() SignalError
func (conn *AuthenticatedChatConnection) SetRegistrationLock() SignalError
func (conn *AuthenticatedChatConnection) SetRegistrationRecoveryPassword() SignalError
func (conn *AuthenticatedChatConnection) SetUsernameLink() SignalError

/*
* SignalUnauthenticatedChatConnection
 */
type UnauthenticatedChatConnection struct{}

// C.signal_unauthenticated_chat_connection_init_listener()
func (conn *UnauthenticatedChatConnection) InitListener() SignalError
func (conn *UnauthenticatedChatConnection) AccountExists() SignalError
func (conn *UnauthenticatedChatConnection) Connect() SignalError
func (conn *UnauthenticatedChatConnection) Destroy() SignalError
func (conn *UnauthenticatedChatConnection) Disconnect() SignalError
func (conn *UnauthenticatedChatConnection) Info() SignalError
func (conn *UnauthenticatedChatConnection) Send() SignalError
func (conn *UnauthenticatedChatConnection) SendMessage() SignalError
func (conn *UnauthenticatedChatConnection) SendMultiRecipientMessage() SignalError
func (conn *UnauthenticatedChatConnection) SendRawGRPC() SignalError
func (conn *UnauthenticatedChatConnection) SubmitCallQualitySurvey() SignalError
func (conn *UnauthenticatedChatConnection) GetPreKeysAccessKeyAuth() SignalError
func (conn *UnauthenticatedChatConnection) GetPreKeysGroupAuth() SignalError
func (conn *UnauthenticatedChatConnection) GetPreKeysUnrestrictedAuth() SignalError
func (conn *UnauthenticatedChatConnection) LookUpUsernameHash() SignalError
func (conn *UnauthenticatedChatConnection) LookUpUsernameLink() SignalError
func (conn *UnauthenticatedChatConnection) BackupCopyMedia() SignalError
func (conn *UnauthenticatedChatConnection) BackupDeleteAll() SignalError
func (conn *UnauthenticatedChatConnection) BackupDeleteMedia() SignalError
func (conn *UnauthenticatedChatConnection) BackupGetCDNCredentials() SignalError
func (conn *UnauthenticatedChatConnection) BackupGetMediaBackupInfo() SignalError
func (conn *UnauthenticatedChatConnection) BackupGetUploadForm() SignalError
func (conn *UnauthenticatedChatConnection) BackupGetMessageBackupInfo() SignalError
func (conn *UnauthenticatedChatConnection) BackupGetSVRBCredentials() SignalError
func (conn *UnauthenticatedChatConnection) BackupListMedia() SignalError
func (conn *UnauthenticatedChatConnection) BackupRefresh() SignalError
func (conn *UnauthenticatedChatConnection) BackupSetPublicKey() SignalError

/*
* SignalProvisioningChatConnection
 */
type ProvisioningChatConnection struct{}

// C.signal_provisioning_chat_connection_init_listener()
func (conn *ProvisioningChatConnection) InitListener() SignalError

func (conn *ProvisioningChatConnection) Connect() SignalError
func (conn *ProvisioningChatConnection) Destroy() SignalError
func (conn *ProvisioningChatConnection) Disconnect() SignalError
func (conn *ProvisioningChatConnection) Info() SignalError

/*
* SignalChatConnectionInfo
 */
type ChatConnectionInfo struct{}

func (*ChatConnectionInfo) Description() SignalError
func (*ChatConnectionInfo) Destroy() SignalError
func (*ChatConnectionInfo) IPVersion() SignalError
func (*ChatConnectionInfo) LocalPort() SignalError

/*
* SignalConnectionManager
 */
type ConnectionManager struct{}

func (*ConnectionManager) ClearProxy() SignalError
func (*ConnectionManager) Destroy() SignalError
func (*ConnectionManager) New() SignalError
func (*ConnectionManager) OnNetworkChange() SignalError
func (*ConnectionManager) SetCensorshipCircumventionEnabled() SignalError
func (*ConnectionManager) SetInvalidProxy() SignalError
func (*ConnectionManager) SetProxy() SignalError
func (*ConnectionManager) SetRemoteConfig() SignalError

/*
* SignalConnectionProxyConfig
 */
type ConnectionProxyConfig struct{}

func (*ConnectionProxyConfig) New() SignalError
func (*ConnectionProxyConfig) Clone() SignalError
func (*ConnectionProxyConfig) Destroy() SignalError
