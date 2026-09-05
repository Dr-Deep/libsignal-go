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
	var (
		connmgr = getConnMgr()
		ctx     = getTokioAsyncContext()
	)

	ffi.ConnectUnauthenticatedChat
	ffi.ConnectAuthenticatedChat
}

func getConnMgr() *ffi.ConnectionManager {
	connmgr, err := ffi.NewConnectionManager()
	if err != nil {
		panic(err)
	}

	return connmgr
}

func getTokioAsyncContext() *ffi.TokioAsyncContext {
	ctx, err := ffi.NewTokioAsyncContext()
	if err != nil {
		panic(err)
	}

	return ctx
}
