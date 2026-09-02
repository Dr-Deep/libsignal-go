# libsignal-go

A deliberately small Go wrapper around Signal's `libsignal` C FFI. The current
milestone provides ownership-safe error handling, a Tokio runtime handle, and
Signal service-ID conversion. It aims to implement account provisioning,
message transport, or encryption stores.

## FFI harness

Requirements are Linux, CGO, a C toolchain, and the checked-in
`deps/libsignal_ffi.a` built for the current architecture.

If the archive was produced on another operating system, rebuild it locally:

```sh
make -B deps/libsignal_ffi.a
```

```sh
CGO_ENABLED=1 go test -race ./...
CGO_ENABLED=1 go run ./cmd/ffi-harness
```

