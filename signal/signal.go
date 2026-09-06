package signal

import _ "github.com/Dr-Deep/libsignal-go/signal/ffi"

type Client struct{}

/*
event call handler
signal.AddHandler(
func(signal *signal.Client, event signal.EventTypeMessage)
)
*/
func New(credentials any) (*Client, error) {
	return nil, nil
}
