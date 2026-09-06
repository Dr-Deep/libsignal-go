package ffi

import (
	"context"
	"testing"
)

const (
	count = 32
)

/*
race conditions?
? CGO_ENABLED=1 go test -race ./...

`valgrind --leak-check=full ./libsignal-go`
`CGO_ENABLED=1 go test -race -v ./signal/ffi -run TestTokio`
*/
func TestTokio(t *testing.T) {
	var (
		ctx = context.Background()
	)

	tokio, err := NewTokioAsyncContext()
	if err != nil {
		panic(err)
	}

	out, err := TokioTesting_FutureSuccessBytes(
		ctx, tokio, count,
	)
	if err != nil {
		panic(err)
	}

	t.Logf("OUT: '%#v'\n", out)
}
