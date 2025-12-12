#!/bin/bash

echo "Build..."
GOOS=linux GOARCH=amd64 go build -v -ldflags='-s -w' -o ./dist/gonote ./cmd/main
