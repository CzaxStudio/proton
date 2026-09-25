#!/bin/bash

# Proton Build Script

# Creates self-contained binaries for Windows, macOS, and Linux

APP_NAME="ProtonCyberTool
"
SRC_PATH="./examples/cybertool/main.go"

OUT_DIR="./build"

mkdir -p $OUT_DIR

echo " Starting build process..."

echo " Welcome to Proton Cyber App"

echo " This app is made using the Proton cross platform Go library"

echo " This is just a basic example app, you can make even better versions!!!!"

# 1. Linux (64-bit)

echo " Building for Linux..."

GOOS=linux GOARCH=amd64 go build -o $OUT_DIR/${APP_NAME}_linux $SRC_PATH

# 2. Windows (64-bit)

echo " Building for Windows..."

GOOS=windows GOARCH=amd64 go build -o $OUT_DIR/${APP_NAME}.exe $SRC_PATH

# 3. macOS (Intel)

echo " Building for macOS (Intel)..."

GOOS=darwin GOARCH=amd64 go build -o $OUT_DIR/${APP_NAME}_macos_intel $SRC_PATH

# 4. macOS (Apple Silicon/M1/M2)
echo " Building for macOS (Apple Silicon)..."

GOOS=darwin GOARCH=arm64 go build -o $OUT_DIR/${APP_NAME}_macos_arm $SRC_PATH

echo " Build complete! Binaries are in the $OUT_DIR folder."

echo " Thanks for choosing Proton"
