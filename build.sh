#!/bin/bash

echo "Generate static"
go run ./cmd/gen/gen.go
echo "Build..."
GOOS=linux GOARCH=amd64 go build -v -ldflags='-s -w' -o ./dist/gonote ./cmd/main
