#!/bin/sh

GOOS=windows GOARCH=amd64 go build -o ../bin/galaxy.exe ../cmd/galaxy/main.go
