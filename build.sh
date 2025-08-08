#!/bin/bash

# Exit on error
set -e

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o bin/go-socks5-proxy-linux-amd64 main.go

# Build for Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o bin/go-socks5-proxy-windows-amd64.exe main.go

# Build for macOS
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o bin/go-socks5-proxy-darwin-amd64 main.go

echo "Build complete!"
