@echo off
:menu
echo Choose the target platform (GOOS) for building:
echo 1. Windows (amd64)
echo 2. Linux (arm)
set /p choice=Enter your choice (1 or 2): 

if %choice%==1 (
    echo Building for Windows (amd64)...
    set GOARCH=amd64
    set GOOS=windows
    go build
    goto end
) else if %choice%==2 (
    echo Building for Linux (arm)...
    set GOARCH=arm
    set GOOS=linux
    go build
    goto end
) else (
    echo Invalid choice, please select 1 or 2.
    goto menu
)

:end
