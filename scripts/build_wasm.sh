#!/bin/sh

GOOS=js GOARCH=wasm go build -o ../dist/galaxy.wasm ../cmd/galaxy/main.go

cp $(go env GOROOT)/misc/wasm/wasm_exec.js ../dist/
cp ../cmd/wasm/index.html ../dist/index.html
