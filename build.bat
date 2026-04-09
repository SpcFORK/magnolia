@echo off
setlocal enabledelayedexpansion

set "LOGO_1=./Magnolia.jpg"
set "LOGO_2=./Magnolia2.png"

set "ICON_1=./build/magnolia1.ico"
set "ICON_2=./build/magnolia2.ico"
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
set "_ORIG_GOOS=%GOOS%"
set "_ORIG_GOARCH=%GOARCH%"

set "GOOS=linux"
set "GOARCH=amd64"
go build -o ./build/magnolia
if errorlevel 1 exit /b 1

if defined _ORIG_GOOS (
    set "GOOS=%_ORIG_GOOS%"
) else (
    set "GOOS="
)

if defined _ORIG_GOARCH (
    set "GOARCH=%_ORIG_GOARCH%"
) else (
    set "GOARCH="
)

powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"Compress-Archive -Path './build/magnolia' -DestinationPath './build/magnolia.zip' -Force"
if errorlevel 1 exit /b 1

if exist "%LOGO_1%" (
    call :make_ico "%LOGO_1%" "%ICON_1%"
    if errorlevel 1 exit /b 1
)
if exist "%LOGO_2%" (
    call :make_ico "%LOGO_2%" "%ICON_2%"
    if errorlevel 1 exit /b 1
)

if not exist "%ICON_1%" if not exist "%ICON_2%" (
    echo No icon source images found.
    exit /b 1
)

if exist "%SYSO_FILE%" del "%SYSO_FILE%"

where rsrc >nul 2>&1
if errorlevel 1 (
    echo Installing rsrc tool...
    go install github.com/akavel/rsrc@latest
    if errorlevel 1 exit /b 1
)

echo Embedding icon resource...
if exist "%ICON_2%" (
    rsrc -ico "%ICON_2%" -o "%SYSO_FILE%"
) else (
    rsrc -ico "%ICON_1%" -o "%SYSO_FILE%"
)
if errorlevel 1 exit /b 1

echo Building EXE target...
set "GOOS=windows"
set "GOARCH=amd64"
go build -o ./build/magnolia.exe
if errorlevel 1 (
    if exist "%SYSO_FILE%" del "%SYSO_FILE%"
    exit /b 1
)

if defined _ORIG_GOOS (
    set "GOOS=%_ORIG_GOOS%"
) else (
    set "GOOS="
)

if defined _ORIG_GOARCH (
    set "GOARCH=%_ORIG_GOARCH%"
) else (
    set "GOARCH="
)

powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"Compress-Archive -Path './build/magnolia.exe' -DestinationPath './build/magnolia-exe.zip' -Force"
if errorlevel 1 exit /b 1

if exist "%SYSO_FILE%" del "%SYSO_FILE%"

echo Build complete!
echo   ELF: ./build/magnolia
echo   EXE: ./build/magnolia.exe
if exist "%ICON_1%" echo   ICO: %ICON_1%
if exist "%ICON_2%" echo   ICO: %ICON_2%

set "DO_WSL_INSTALL=0"
if "%1"=="--wsl" set "DO_WSL_INSTALL=1"
if "%DO_WSL_INSTALL%"=="1" (
    echo Installing to /usr/local/bin via WSL...
    wsl -u root -- sh -c "rm -f /usr/local/bin/magnolia; mv ./build/magnolia /usr/local/bin"
)

echo Generating spec docs for lib...
if not exist "docs\spec" mkdir "docs\spec"
set "_SPEC_OK=0"
set "_SPEC_FAIL=0"
for %%F in (lib\*.oak) do (
    set "_BASE=%%~nF"
    .\build\magnolia.exe build --entry "%%F" --doc --output "docs\spec\!_BASE!.md" >nul 2>&1
    if errorlevel 1 (
        echo   Failed to generate spec for %%F
        set /a _SPEC_FAIL+=1
    ) else (
        set /a _SPEC_OK+=1
    )
)
echo   Spec docs: !_SPEC_OK! generated, !_SPEC_FAIL! skipped

goto :eof

:make_ico
set "_SRC=%~1"
set "_DST=%~2"
echo Creating icon from %_SRC% -^> %_DST%...
powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"Add-Type -AssemblyName System.Drawing; ^
$pngPath = '%_SRC%'; ^
$icoPath = '%_DST%'; ^
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
goto :eof