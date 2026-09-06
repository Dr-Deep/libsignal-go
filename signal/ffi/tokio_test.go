package ffi

import (
	"context"
	"testing"
)

const (
	count = 32
)

// race conditions?
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
