# Windows Flags Library (windows-flags)

## Overview

`windows-flags` stores shared Win32 numeric constants and flags.

## Import

```oak
wf := import('windows-flags')
```

## Export groups

- Process access flags (`PROCESS_*`)
- Virtual memory flags (`MEM_*`)
- Page protection flags (`PAGE_*`)
- FormatMessage flags (`FORMAT_MESSAGE_*`, `ERROR_SUCCESS`)
- Winsock constants (`AF_INET`, `SOCK_*`, `IPPROTO_*`, etc.)
- WinINet constants (`INTERNET_*`)
- Registry constants (`HKEY_*`, `KEY_*`, `REG_*`)
- Window class/style/message flags (`CS_*`, `WS_*`, `WM_*`, `SW_*`, `PM_*`)
- Resource/message box constants (`IDC_ARROW`, `IDI_APPLICATION`, `MB_*`)

## Notes

- For an exhaustive constant list, see `windows.md`.
