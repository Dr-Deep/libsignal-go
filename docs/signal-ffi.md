# signal_ffi FFI Übersicht

## Übersicht

Diese Übersicht kann den eigentlichen `signal_ffi`-Vertrag in diesem
Arbeitsbaum derzeit nicht auflisten: `deps/libsignal` ist als Git-Submodule
bei Commit `eb7864c4d15435ee33681ce828930d9a4296f155` eingetragen, enthält
aber keine ausgecheckten Dateien. Das Submodule ist für
`https://github.com/Dr-Deep/libsignal.git` konfiguriert. Auch `signal_ffi.h`
liegt weder unter `deps/` noch unter `deps/libsignal/` vor. Daher werden hier
bewusst keine C-Definitionen, Rust-Typen oder Funktionszuordnungen ergänzt,
die nicht aus dem vorhandenen Repository ableitbar sind.

Die Build-Regeln beschreiben die vorgesehene Architektur: Das Rust-Paket
`libsignal-ffi` soll im Submodule gebaut werden; die Header sollen aus
`deps/libsignal/swift/Sources/SignalFfi/` nach `deps/` kopiert werden. Die
Go-Quelle bindet anschließend `./deps/signal_ffi.h` und
`./deps/signal_ffi_testing.h` ein.

## Typen und Objekte

Es sind keine FFI-Typen oder Objekte dokumentierbar, weil die maßgebliche
Header-Datei im Arbeitsbaum fehlt. Insbesondere können keine C-Definitionen,
zugehörigen Rust-Typen, Rust-Dateien oder Crate-/Modulzuordnungen verifiziert
werden.

### Funktionen

| C-Funktion | Rust-Funktion | Rust-Datei | Beschreibung |
|---|---|---|---|
| Nicht eindeutig aus dem Quellcode ableitbar. | Nicht eindeutig aus dem Quellcode ableitbar. | Nicht eindeutig aus dem Quellcode ableitbar. | `signal_ffi.h` und die Rust-Implementierung sind nicht ausgecheckt. |

### Callbacks

| Callback | Parameter | Verwendung | Rust-Seite |
|---|---|---|---|
| Nicht eindeutig aus dem Quellcode ableitbar. | Nicht eindeutig aus dem Quellcode ableitbar. | Nicht eindeutig aus dem Quellcode ableitbar. | `signal_ffi.h` und die Rust-Implementierung sind nicht ausgecheckt. |

## Fehlerbehandlung

Aus dem vorhandenen FFI-Code ist kein Fehlermechanismus ableitbar, weil weder
der FFI-Header noch dessen Rust-Quellen vorhanden sind. Aussagen zu
Fehlerobjekten, Rückgabecodes, NULL-Werten oder Callback-Fehlern würden eine
nicht belegte Annahme darstellen.

## Ownership und Lifetime

Ownership- und Lifetime-Regeln sind nicht dokumentierbar. Der vorhandene
Go-Quelltext zeigt nur die Cgo-Includes; er enthält keine verwendeten
`signal_ffi`-Signaturen oder Freigabeaufrufe.

## Unklare Zuordnungen

* **Alle C-zu-Rust-Zuordnungen:** Nicht eindeutig aus dem Quellcode
  ableitbar. `signal_ffi.h` fehlt und `deps/libsignal` enthält keine
  ausgecheckte Rust-Implementierung.
* **Vorgesehene, aber nicht verifizierbare Pfade:** Die Build-Regeln nennen
  `deps/libsignal/swift/Sources/SignalFfi/signal_ffi.h` sowie
  `libsignal/rust/bridge/ffi/cbindgen.toml`; diese Dateien konnten im
  Arbeitsbaum nicht gelesen werden.
