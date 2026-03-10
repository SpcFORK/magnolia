@echo off
setlocal enabledelayedexpansion

set "LOGO_PNG=./SushaLogo.png"
set "ICON_PATH=./build/magnolia.ico"
set "SYSO_FILE=./magnolia.syso"

if not exist build (
    mkdir build
)

echo Building ELF target...
go build -o ./build/magnolia
if errorlevel 1 exit /b 1

if not exist "%LOGO_PNG%" (
    echo Icon source not found: %LOGO_PNG%
    exit /b 1
)

echo Creating icon from %LOGO_PNG%...
powershell -NoProfile -ExecutionPolicy Bypass -Command "Add-Type -AssemblyName PresentationCore; $pngPath = (Resolve-Path '%LOGO_PNG%').Path; $icoPath = (Resolve-Path './build').Path + '\\magnolia.ico'; $uri = New-Object System.Uri($pngPath); $bmp = New-Object System.Windows.Media.Imaging.BitmapImage; $bmp.BeginInit(); $bmp.UriSource = $uri; $bmp.DecodePixelWidth = 256; $bmp.DecodePixelHeight = 256; $bmp.EndInit(); $frame = [System.Windows.Media.Imaging.BitmapFrame]::Create($bmp); $encoder = New-Object System.Windows.Media.Imaging.IconBitmapEncoder; $encoder.Frames.Add($frame); $fs = [System.IO.File]::Open($icoPath, [System.IO.FileMode]::Create); try { $encoder.Save($fs) } finally { $fs.Dispose() }"
if errorlevel 1 exit /b 1

if exist "%SYSO_FILE%" del "%SYSO_FILE%"

where rsrc >nul 2>&1
if errorlevel 1 (
    echo Installing rsrc tool...
    go install github.com/akavel/rsrc@latest
    if errorlevel 1 exit /b 1
)

echo Embedding icon resource...
rsrc -ico "%ICON_PATH%" -o "%SYSO_FILE%"
if errorlevel 1 exit /b 1

echo Building EXE target...
go build -o ./build/magnolia.exe
if errorlevel 1 (
    if exist "%SYSO_FILE%" del "%SYSO_FILE%"
    exit /b 1
)

if exist "%SYSO_FILE%" del "%SYSO_FILE%"

echo Build complete!
echo   ELF: ./build/magnolia
echo   EXE: ./build/magnolia.exe
echo   ICO: ./build/magnolia.ico