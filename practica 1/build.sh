#!/bin/bash
go mod tidy
go build -o build/server cmd/server/main.go
