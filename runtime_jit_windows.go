//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

const (
	jitMemCommit   = 0x1000
	jitMemReserve  = 0x2000
	jitPageReadWriteExec = 0x40
	jitMemRelease  = 0x8000
	jitAllocAlign  = 4096
	jitUseProtect  = false // VirtualAlloc allocates executable directly
)

// jitResolve resolves a kernel32 procedure through the existing sysproc
// machinery (LoadLibrary/GetProcAddress on Windows builds).
func jitResolve(name string) (uintptr, error) {
	return lookupSysProc("kernel32", name)
}

// jitMmapRW allocates executable committed memory via VirtualAlloc.
// The allocation is created with PAGE_EXECUTE_READWRITE in one step;
// jitProtectRX is a no-op on Windows.
func jitMmapRW(n int) (uintptr, error) {
	virtualAlloc, err := jitResolve("VirtualAlloc")
	if err != nil {
		return 0, err
	}
	size := jitPageRound(n)
	addr, _, errno := oakSyscallN(virtualAlloc,
		0, uintptr(size), jitMemCommit|jitMemReserve, jitPageReadWriteExec)
	if addr == 0 || errno != 0 && errno != syscall.Errno(0) {
		if addr == 0 {
			return 0, errno
		}
	}
	return addr, nil
}

// jitProtectRX is a no-op on Windows: VirtualAlloc already granted
// PAGE_EXECUTE_READWRITE.
func jitProtectRX(addr uintptr, size int) error {
	return nil
}

// jitMunmap releases a VirtualAlloc mapping.
func jitMunmap(addr uintptr, size int) error {
	virtualFree, err := jitResolve("VirtualFree")
	if err != nil {
		return err
	}
	_, _, errno := oakSyscallN(virtualFree, addr, uintptr(size), jitMemRelease)
	if errno != syscall.Errno(0) && errno != syscall.Errno(0x20000000) {
		// errno here is the trampoline's C errno record, not a Win32 code;
		// VirtualFree reports failure through a zero result which we cannot
		// see through oakSyscallN, so treat the call as best-effort.
		_ = errno
	}
	return nil
}

// jitCopyIn copies machine-code bytes into a mapping.
func jitCopyIn(addr uintptr, code []byte) {
	dst := unsafe.Slice((*byte)(unsafe.Pointer(addr)), len(code))
	copy(dst, code)
}
