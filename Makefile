.POSIX:


# cargo, go, cbindgen

CARGO_BIN := cargo
CARGO_RUSTFLAGS := -Ctarget-feature=-crt-static

CBINDGEN_BIN := ~/.cargo/bin/cbindgen
CBINDGEN_CFG := cbindgen.toml

GO_BIN := go
GO_FLAGS := CGO_ENABLED=1

LIBSIGNAL_DIR := deps/libsignal
LIBSIGNAL_VERSION := v0.101.2
LIBSIGNAL_FFI_H := deps/libsignal_ffi.h
LIBSIGNAL_FFI_A := deps/libsignal_ffi.a


.PHONY: all
all: build

$(LIBSIGNAL_DIR):
	git submodule init && git submodule update
	cd $(LIBSIGNAL_DIR) && git checkout $(LIBSIGNAL_VERSION)

$(LIBSIGNAL_FFI_A): $(LIBSIGNAL_DIR)
	cd $(LIBSIGNAL_DIR) && RUSTFLAGS="$(CARGO_RUSTFLAGS)" $(CARGO_BIN) build -p libsignal-ffi --release
	@cp -v $(LIBSIGNAL_DIR)/target/release/libsignal_ffi.a $(LIBSIGNAL_FFI_A)

$(LIBSIGNAL_FFI_H): $(LIBSIGNAL_DIR)
	echo skipping cbindgen
	#$(CBINDGEN_BIN) --config $(CBINDGEN_CFG) --profile release --crate libsignal-ffi --output $(LIBSIGNAL_FFI_H) --lang c

build: $(LIBSIGNAL_FFI_A) $(LIBSIGNAL_FFI_H) test
	$(GO_FLAGS) $(GO_BIN) build -v .

test: $(LIBSIGNAL_FFI_A) $(LIBSIGNAL_FFI_H)
	$(GO_FLAGS) $(GO_BIN) test -v .

clean:
	@echo "clean?"
