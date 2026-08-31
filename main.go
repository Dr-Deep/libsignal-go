package main

/*
#cgo LDFLAGS: ./libsignal_ffi.a
*/
import "C"

// https://github.com/ihciah/rust2go

func main() {}

/*
[workspace]
# When adding new members, consider updating the log filter in rust/bridge/shared/src/logging.rs.
members = [
    "rust/attest",
    "rust/crypto",
    "rust/debug",
    "rust/device-transfer",
    "rust/keytrans",
    "rust/media",
    "rust/message-backup",
    "rust/net",
    "rust/net/chat",
    "rust/net/infra",
    "rust/account-keys",
    "rust/poksho",
    "rust/protocol",
    "rust/usernames",
    "rust/zkcredential",
    "rust/zkgroup",
    "rust/bridge/ffi",
    "rust/bridge/ffi/impl",
    "rust/bridge/ffi/native_swift",
    "rust/bridge/jni",
    "rust/bridge/jni/impl",
    "rust/bridge/jni/native_kt",
    "rust/bridge/jni/testing",
    "rust/bridge/node",
    "rust/bridge/node/native_ts",
]
default-members = [
    "rust/crypto",
    "rust/device-transfer",
    "rust/media",
    "rust/message-backup",
    "rust/account-keys",
    "rust/poksho",
    "rust/protocol",
    "rust/usernames",
    "rust/zkcredential",
    "rust/zkgroup",
]

[workspace.dependencies]
# Our own crates, so that we don't have to depend on them by inter-crate paths
attest = { path = "rust/attest" }
device-transfer = { path = "rust/device-transfer" }
libsignal-account-keys = { path = "rust/account-keys" }
libsignal-cli-utils = { path = "rust/cli-utils" }
libsignal-core = { path = "rust/core" }
libsignal-debug = { path = "rust/debug" }
libsignal-keytrans = { path = "rust/keytrans" }
libsignal-message-backup = { path = "rust/message-backup" }
libsignal-net = { path = "rust/net" }
libsignal-net-chat = { path = "rust/net/chat" }
libsignal-net-grpc = { path = "rust/net/grpc" }
libsignal-node = { path = "rust/bridge/node" }
libsignal-protocol = { path = "rust/protocol" }
libsignal-svr2 = { path = "rust/svr2" }
libsignal-svrb = { path = "rust/svrb" }
poksho = { path = "rust/poksho" }
signal-crypto = { path = "rust/crypto" }
signal-media = { path = "rust/media" }
usernames = { path = "rust/usernames" }
zkcredential = { path = "rust/zkcredential" }
zkgroup = { path = "rust/zkgroup" }

libsignal-bridge = { path = "rust/bridge/shared" }
libsignal-bridge-macros = { path = "rust/bridge/shared/macros" }
libsignal-bridge-testing = { path = "rust/bridge/shared/testing" }
libsignal-bridge-types = { path = "rust/bridge/shared/types" }
libsignal-jni-impl = { path = "rust/bridge/jni/impl" }
signal-neon-futures = { path = "rust/bridge/node/futures" }
*/
