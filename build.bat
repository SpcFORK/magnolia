@echo off
setlocal enabledelayedexpansion

set "LOGO_PNG=./SushaLogo.png"
set "ICON_PATH=./build/magnolia.ico"
set "SYSO_FILE=./magnolia.syso"

set "OAK_VSCODE_DIR=.\tools\oak-vscode"
set "OAK_VSCODE_ZIP=.\tools\oak-vscode.zip"

if not exist "%OAK_VSCODE_DIR%" (
    echo VS Code extension folder not found: %OAK_VSCODE_DIR%
    exit /b 1
)

if exist "%OAK_VSCODE_ZIP%" del "%OAK_VSCODE_ZIP%"

echo Zipping %OAK_VSCODE_DIR%...
powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"Compress-Archive -Path '%OAK_VSCODE_DIR%' -DestinationPath '%OAK_VSCODE_ZIP%' -Force"
if errorlevel 1 exit /b 1

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

powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"Compress-Archive -Path './build/magnolia' -DestinationPath './build/magnolia.zip' -Force"
if errorlevel 1 exit /b 1

echo Creating icon from %LOGO_PNG%...
powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"Add-Type -AssemblyName System.Drawing; ^
$pngPath = '%LOGO_PNG%'; ^
$icoPath = '%ICON_PATH%'; ^
$src = [System.Drawing.Image]::FromFile($pngPath); ^
try { ^
    $bmp = [System.Drawing.Bitmap]::new(256, 256, [System.Drawing.Imaging.PixelFormat]::Format32bppArgb); ^
    $gfx = [System.Drawing.Graphics]::FromImage($bmp); ^
    try { ^
        $gfx.Clear([System.Drawing.Color]::Transparent); ^
        $gfx.CompositingQuality = [System.Drawing.Drawing2D.CompositingQuality]::HighQuality; ^
        $gfx.InterpolationMode = [System.Drawing.Drawing2D.InterpolationMode]::HighQualityBicubic; ^
        $gfx.SmoothingMode = [System.Drawing.Drawing2D.SmoothingMode]::HighQuality; ^
        $gfx.DrawImage($src, 0, 0, 256, 256); ^
    } finally { ^
        $gfx.Dispose(); ^
    }; ^
    $ms = [System.IO.MemoryStream]::new(); ^
    try { ^
        $bmp.Save($ms, [System.Drawing.Imaging.ImageFormat]::Bmp); ^
        $bmpBytes = $ms.ToArray(); ^
        $dib = [byte[]]::new($bmpBytes.Length - 14); ^
        [Array]::Copy($bmpBytes, 14, $dib, 0, $dib.Length); ^
        $height2 = [BitConverter]::GetBytes([int]($bmp.Height * 2)); ^
        [Array]::Copy($height2, 0, $dib, 8, 4); ^
        $maskRowBytes = [int]([Math]::Ceiling($bmp.Width / 32.0) * 4); ^
        $maskSize = $maskRowBytes * $bmp.Height; ^
        $mask = [byte[]]::new($maskSize); ^
        $imageSize = $dib.Length + $mask.Length; ^
        $fs = [System.IO.File]::Open($icoPath, [System.IO.FileMode]::Create); ^
        $bw = [System.IO.BinaryWriter]::new($fs); ^
        try { ^
            $bw.Write([UInt16]0); ^
            $bw.Write([UInt16]1); ^
            $bw.Write([UInt16]1); ^
            $bw.Write([byte]0); ^
            $bw.Write([byte]0); ^
            $bw.Write([byte]0); ^
            $bw.Write([byte]0); ^
            $bw.Write([UInt16]1); ^
            $bw.Write([UInt16]32); ^
            $bw.Write([UInt32]$imageSize); ^
            $bw.Write([UInt32]22); ^
            $bw.Write($dib); ^
            $bw.Write($mask); ^
        } finally { ^
            $bw.Dispose(); ^
            $fs.Dispose(); ^
        } ^
    } finally { ^
        $ms.Dispose(); ^
    }; ^
    $bmp.Dispose(); ^
} finally { ^
    $src.Dispose(); ^
}"
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

powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"Compress-Archive -Path './build/magnolia.exe' -DestinationPath './build/magnolia-exe.zip' -Force"
if errorlevel 1 exit /b 1

if exist "%SYSO_FILE%" del "%SYSO_FILE%"

echo Build complete!
echo   ELF: ./build/magnolia
echo   EXE: ./build/magnolia.exe
echo   ICO: ./build/magnolia.ico

set "DO_WSL_INSTALL=0"
if "%1"=="--wsl" set "DO_WSL_INSTALL=1"
if "%DO_WSL_INSTALL%"=="1" (
    echo Installing to /usr/local/bin via WSL...
    wsl -u root -- sh -c "rm -f /usr/local/bin/magnolia;"
    wsl -u root -- sh -c "mv ./build/magnolia /usr/local/bin"
    wsl -u root -- sh -c "exec bash -l"
)