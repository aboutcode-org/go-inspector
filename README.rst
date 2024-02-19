Go-Inspector
================================

- To enable the GoReSym plugin, first you need to install goresym from https://github.com/mandiant/GoReSym/releases/download/v2.6.4/GoReSym.zip
- Unzip the GoReSym.zip, extract goresym for linux and add it in src/go_inspector/bin.
- then change it to executable  ```chmod u+x src/go_inspector/bin/GoReSym_lin```
- Install requirements and dependencies using ```make dev```
- Use ```scancode --json-pp - --go-symbol <PATH> --verbose``` to get debug symbols.


How to generate test binaries
============================

- Run `go tool dist list` to get all possible pairs of OSes and arches to compile the binary.
- Then use a OS/arch pair like this ``GOOS=<OS> GOARCH=<arch> go build -o ./tests/data/app_exe ./tests/data/main.go``
  to get compiled binary.
