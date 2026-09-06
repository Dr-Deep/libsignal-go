# libsignal-go

A deliberately small Go wrapper around Signal's `libsignal` C FFI. The current
milestone provides ownership-safe error handling, a Tokio runtime handle, and
Signal service-ID conversion. It aims to implement account provisioning,
message transport, or encryption stores.


passing Pointers is Okay
storing Pointers only Go; eleminate all C Pointers


$(GO_FLAGS) $(GO_BIN) test -c -o /tmp/libsignal-go.test ./signal/ffi
	@echo "[Valgrind] running..."
	valgrind \
		--leak-check=full \
		--show-leak-kinds=all \
		--track-origins=yes \
		--error-exitcode=1 \
		--suppressions=valgrind.supp \
		/tmp/libsignal-go.test -test.v -test.run TestTokio
