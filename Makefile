.POSIX:

# rust version = nightly-2026-07-16
CARGO_BIN := /usr/local/bin/cargo
CARGO_RUSTFLAGS := -Ctarget-feature=-crt-static

RUSTC_BIN := /usr/local/bin/rustc

GO_BIN := go
GO_FLAGS := CGO_ENABLED=1

LIBSIGNAL_DIR := deps/libsignal
LIBSIGNAL_VERSION := v0.101.2
LIBSIGNAL_FEATURES := libsignal-bridge-testing
LIBSIGNAL_FFI_CBINDGEN := libsignal/rust/bridge/ffi/cbindgen.toml
LIBSIGNAL_FFI_H := deps/signal_ffi.h
LIBSIGNAL_FFI_A := deps/libsignal_ffi.a

# protoc?

.PHONY: all
all: clean test

$(LIBSIGNAL_DIR):
	@echo "[$(LIBSIGNAL_DIR)] ..."
	#git submodule init && git submodule update
	#cd $(LIBSIGNAL_DIR) && git checkout $(LIBSIGNAL_VERSION)
	# checkout submodules recursive

#cargo +${{ matrix.toolchain }} build --workspace --features libsignal-ffi/signal-media --verbose --keep-going
$(LIBSIGNAL_FFI_A): $(LIBSIGNAL_DIR)
	@echo "[$(LIBSIGNAL_FFI_A)] ..."
	cd $(LIBSIGNAL_DIR) && RUSTFLAGS="$(CARGO_RUSTFLAGS)" $(CARGO_BIN) build -p libsignal-ffi --release --features $(LIBSIGNAL_FEATURES)
	@cp -v $(LIBSIGNAL_DIR)/target/release/libsignal_ffi.a $(LIBSIGNAL_FFI_A)

#deps/libsignal/swift/Sources/SignalFfi/signal_ffi.h
# doas sysctl hardening.harden_rtld=0
# swift/build_ffi.sh --release --generate-ffi
$(LIBSIGNAL_FFI_H): $(LIBSIGNAL_DIR)
	@echo "[$(LIBSIGNAL_FFI_H)] ..."
	cp -nv deps/libsignal/swift/Sources/SignalFfi/*.h deps/

build: $(LIBSIGNAL_FFI_A) $(LIBSIGNAL_FFI_H) test
	@echo "[Building] ..."
	$(GO_FLAGS) $(GO_BIN) build -v .

test: $(LIBSIGNAL_FFI_A) $(LIBSIGNAL_FFI_H)
	@echo "[Testing] ..."
	$(GO_FLAGS) $(GO_BIN) test -v ./...

clean:
	@echo "[Cleaning] ..."
	@rm -fv libsignal-go
