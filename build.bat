@echo off
echo This is a test file.....
set APP_NAME=ProtonCyberTool
set SRC_PATH=.\examples\cybertool\main.go
set OUT_DIR=.\build

if not exist %OUT_DIR% mkdir %OUT_DIR%

echo Starting build process...

echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o %OUT_DIR%\%APP_NAME%.exe %SRC_PATH%

echo Build complete! Binary is in the %OUT_DIR% folder.
pause
