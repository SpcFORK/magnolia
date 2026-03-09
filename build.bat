@echo off
setlocal enabledelayedexpansion

if not exist build (
    mkdir build
)

echo Building ELF target...
go build -o ./build/magnolia
if errorlevel 1 exit /b 1

echo Building EXE target...
go build -o ./build/magnolia.exe
if errorlevel 1 exit /b 1

echo Build complete!
echo   ELF: ./build/magnolia
echo   EXE: ./build/magnolia.exe