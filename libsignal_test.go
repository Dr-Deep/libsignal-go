package libsignal

import (
	"testing"

	"github.com/Dr-Deep/libsignal-go/signal/ffi"
	_ "github.com/Dr-Deep/libsignal-go/signal/ffi"
)

/*
	ProvisioningURL = "wss://chat.signal.org/v1/websocket/provisioning/"
	APIURL          = "https://chat.signal.org"
	WSURL           = "wss://chat.signal.org"
    Username?
    Password?
*/

func Test(t *testing.T) {

	ffi.NewConnectionManager()
	ffi.NewTokioAsyncContext()

	ffi.ConnectUnauthenticatedChat
	ffi.ConnectAuthenticatedChat
}
