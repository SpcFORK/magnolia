# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\windows-fonts.oak`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `wstr(s)`

### `gdi32(symbol, args...)`

### `user32(symbol, args...)`

- `FW_DONTCARE` · `0`
- `FW_THIN` · `100`
- `FW_EXTRALIGHT` · `200`
- `FW_LIGHT` · `300`
- `FW_NORMAL` · `400`
- `FW_MEDIUM` · `500`
- `FW_SEMIBOLD` · `600`
- `FW_BOLD` · `700`
- `FW_EXTRABOLD` · `800`
- `FW_HEAVY` · `900`
- `ANSI_CHARSET` · `0`
- `DEFAULT_CHARSET` · `1`
- `SYMBOL_CHARSET` · `2`
- `SHIFTJIS_CHARSET` · `128`
- `HANGUL_CHARSET` · `129`
- `GB2312_CHARSET` · `134`
- `CHINESEBIG5_CHARSET` · `136`
- `OEM_CHARSET` · `255`
- `JOHAB_CHARSET` · `130`
- `HEBREW_CHARSET` · `177`
- `ARABIC_CHARSET` · `178`
- `GREEK_CHARSET` · `161`
- `TURKISH_CHARSET` · `162`
- `VIETNAMESE_CHARSET` · `163`
- `THAI_CHARSET` · `222`
- `EASTEUROPE_CHARSET` · `238`
- `RUSSIAN_CHARSET` · `204`
- `BALTIC_CHARSET` · `186`
- `OUT_DEFAULT_PRECIS` · `0`
- `OUT_STRING_PRECIS` · `1`
- `OUT_STROKE_PRECIS` · `3`
- `OUT_TT_PRECIS` · `4`
- `OUT_DEVICE_PRECIS` · `5`
- `OUT_TT_ONLY_PRECIS` · `7`
- `OUT_OUTLINE_PRECIS` · `8`
- `CLIP_DEFAULT_PRECIS` · `0`
- `DEFAULT_QUALITY` · `0`
- `DRAFT_QUALITY` · `1`
- `PROOF_QUALITY` · `2`
- `NONANTIALIASED_QUALITY` · `3`
- `ANTIALIASED_QUALITY` · `4`
- `CLEARTYPE_QUALITY` · `5`
- `DEFAULT_PITCH` · `0`
- `FIXED_PITCH` · `1`
- `VARIABLE_PITCH` · `2`
- `FF_DONTCARE` · `0`
- `FF_ROMAN` · `16`
- `FF_SWISS` · `32`
- `FF_MODERN` · `48`
- `FF_SCRIPT` · `64`
- `FF_DECORATIVE` · `80`
- `OEM_FIXED_FONT` · `10`
- `ANSI_FIXED_FONT` · `11`
- `ANSI_VAR_FONT` · `12`
- `SYSTEM_FONT` · `13`
- `DEVICE_DEFAULT_FONT` · `14`
- `DEFAULT_GUI_FONT` · `17`
- `TA_NOUPDATECP` · `0`
- `TA_UPDATECP` · `1`
- `TA_LEFT` · `0`
- `TA_RIGHT` · `2`
- `TA_CENTER` · `6`
- `TA_TOP` · `0`
- `TA_BOTTOM` · `8`
- `TA_BASELINE` · `24`
- `TM_HEIGHT_OFF` · `0`
- `TM_ASCENT_OFF` · `4`
- `TM_DESCENT_OFF` · `8`
- `TM_INTERNAL_LEADING_OFF` · `12`
- `TM_EXTERNAL_LEADING_OFF` · `16`
- `TM_AVE_CHAR_WIDTH_OFF` · `20`
- `TM_MAX_CHAR_WIDTH_OFF` · `24`
- `TM_WEIGHT_OFF` · `28`
### `createFontEx(spec)`

> returns `:object`

### `deleteFont(handle)`

### `selectFont(hdc, fontHandle)`

### `getStockFont(stockIndex)`

> returns `:object`

### `getDefaultGuiFont()`

### `setTextAlign(hdc, flags)`

### `getTextMetrics(hdc)`

> returns `:object`

### `getTextExtent(hdc, text)`

> returns `:object`

- `_fontCacheMap` · `{}`
### `_fontCacheKey(spec)`

### `cachedFont(spec)`

### `releaseCachedFonts()`

> returns `:int`

### `removeCachedFont(spec)`

> returns `:int`

### `measureText(hwnd, spec, text)`

> returns `:object`

### `fontLineHeight(hdc)`

> returns `:object`

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `sys`

### `_isObject?(v)`

### `ok?(result)`

> returns `:bool`

### `error?(result)`

> returns `:bool`

### `resolve(library, symbol)`

> returns `:object`

### `call(target, args...)`

### `resolveAndCall(library, symbol, args...)`

### `valueOr(result, fallback)`

## Module: `windows-constants`

- `Kernel32` · `'kernel32.dll'`
- `Ntdll` · `'ntdll.dll'`
- `Psapi` · `'psapi.dll'`
- `User32` · `'user32.dll'`
- `Gdi32` · `'gdi32.dll'`
- `Advapi32` · `'advapi32.dll'`
- `Shell32` · `'shell32.dll'`
- `Ole32` · `'ole32.dll'`
- `Ws2_32` · `'ws2_32.dll'`
- `Comctl32` · `'comctl32.dll'`
- `Wininet` · `'wininet.dll'`
- `OpenGL32` · `'opengl32.dll'`
- `Vulkan1` · `'vulkan-1.dll'`
- `D3d9` · `'d3d9.dll'`
- `D3d11` · `'d3d11.dll'`
- `Dxgi` · `'dxgi.dll'`
- `Ddraw` · `'ddraw.dll'`
- `Msvcrt` · `'msvcrt.dll'`
- `Ucrtbase` · `'ucrtbase.dll'`
- `Vcruntime140` · `'vcruntime140.dll'`
- `ActionCenter` · `'actioncenter.dll'`
- `Aclui` · `'aclui.dll'`
- `Acledit` · `'acledit.dll'`
- `Acppage` · `'acppage.dll'`
- `Acproxy` · `'acproxy.dll'`
- `Adprovider` · `'adprovider.dll'`
- `Aeinv` · `'aeinv.dll'`
- `Aepic` · `'aepic.dll'`
- `Amstream` · `'amstream.dll'`
- `Adsldp` · `'adsldp.dll'`
- `Adsnt` · `'adsnt.dll'`
- `Adtschema` · `'adtschema.dll'`
- `Adsldpc` · `'adsldpc.dll'`
- `Adsmsext` · `'adsmsext.dll'`
- `Adhsvc` · `'adhsvc.dll'`
- `Advapi32res` · `'advapi32res.dll'`
- `Advpack` · `'advpack.dll'`
- `Aeevts` · `'aeevts.dll'`
- `Apds` · `'apds.dll'`
- `Winhttp` · `'winhttp.dll'`
- `Urlmon` · `'urlmon.dll'`
- `Crypt32` · `'crypt32.dll'`
- `Bcrypt` · `'bcrypt.dll'`
- `Secur32` · `'secur32.dll'`
- `Comdlg32` · `'comdlg32.dll'`
- `Oleaut32` · `'oleaut32.dll'`
- `Imm32` · `'imm32.dll'`
- `Shlwapi` · `'shlwapi.dll'`
- `Shcore` · `'shcore.dll'`
- `UxTheme` · `'uxtheme.dll'`
- `Dwmapi` · `'dwmapi.dll'`
- `Version` · `'version.dll'`
- `Setupapi` · `'setupapi.dll'`
- `Netapi32` · `'netapi32.dll'`
- `Winmm` · `'winmm.dll'`
- `Avrt` · `'avrt.dll'`
- `Mmdevapi` · `'mmdevapi.dll'`
- `Dsound` · `'dsound.dll'`
- `Mfplat` · `'mfplat.dll'`
- `Mfreadwrite` · `'mfreadwrite.dll'`
- `Mfuuid` · `'mfuuid.dll'`
- `Taskschd` · `'taskschd.dll'`
- `Wevtapi` · `'wevtapi.dll'`
- `Wlanapi` · `'wlanapi.dll'`
- `Mpr` · `'mpr.dll'`
- `Spoolss` · `'spoolss.dll'`
- `Wtsapi32` · `'wtsapi32.dll'`
- `Rasapi32` · `'rasapi32.dll'`
- `Msi` · `'msi.dll'`
- `Wimgapi` · `'wimgapi.dll'`
- `Cabinet` · `'cabinet.dll'`
- `Apphelp` · `'apphelp.dll'`
- `Wer` · `'wer.dll'`
- `Faultrep` · `'faultrep.dll'`
- `Dbghelp` · `'dbghelp.dll'`
- `Dbgeng` · `'dbgeng.dll'`
- `Pdh` · `'pdh.dll'`
- `Iphlpapi` · `'iphlpapi.dll'`
- `Wscapi` · `'wscapi.dll'`
- `Sensapi` · `'sensapi.dll'`
- `Ncrypt` · `'ncrypt.dll'`
- `Cryptui` · `'cryptui.dll'`
- `Wintrust` · `'wintrust.dll'`
- `Samlib` · `'samlib.dll'`
- `Netshell` · `'netshell.dll'`
- `Fwpuclnt` · `'fwpuclnt.dll'`
- `Dnsapi` · `'dnsapi.dll'`
- `Nlaapi` · `'nlaapi.dll'`
- `Httpapi` · `'httpapi.dll'`
- `Rpcrt4` · `'rpcrt4.dll'`
- `Srpapi` · `'srpapi.dll'`
- `Sxs` · `'sxs.dll'`
- `Msvcirt` · `'msvcirt.dll'`
- `ApiSetPrefix` · `'api-ms-win-'`
- `D3dx9Prefix` · `'d3dx9_'`
- `MsvcpPrefix` · `'msvcp'`
- `VcruntimePrefix` · `'vcruntime'`
- `AtlPrefix` · `'atl'`
- `MfcPrefix` · `'mfc'`
- `VcompPrefix` · `'vcomp'`
## Module: `writes`

### `_b0(v)`

### `_b1(v)`

### `_b2(v)`

### `_b3(v)`

### `_b4(v)`

### `_b5(v)`

### `_b6(v)`

### `_b7(v)`

### `readU32(address)`

### `writeU32(address, value)`

### `readU64(address)`

### `writeU64(address, value)`

