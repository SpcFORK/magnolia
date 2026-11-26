//go:build !windows

package main

import (
	"syscall"
	"unsafe"
)

const (
	jitProtRW      = 0x3 // PROT_READ | PROT_WRITE
	jitProtRE      = 0x5 // PROT_READ | PROT_EXEC
	jitMapPriv     = 0x2 // MAP_PRIVATE
	jitMapAnon     = 0x20 // x86-64/ARM64 Linux MAP_ANONYMOUS
	jitAllocAlign  = 4096
	jitUseProtect  = true // two-step RW -> RX (W^X)
)

// jitMmapRW allocates a private anonymous RW mapping of page-rounded size.
func jitMmapRW(n int) (uintptr, error) {
	size := jitPageRound(n)
	addr, _, errno := syscall.Syscall6(
		syscall.SYS_MMAP, 0, uintptr(size), jitProtRW, jitMapPriv|jitMapAnon,
		^uintptr(0), 0)
	if errno != 0 {
		return 0, errno
	}
	return addr, nil
}

// jitProtectRX flips an existing RW mapping to read-execute (W^X).
func jitProtectRX(addr uintptr, size int) error {
	_, _, errno := syscall.Syscall(syscall.SYS_MPROTECT, addr, uintptr(size), jitProtRE)
	if errno != 0 {
		return errno
	}
	return nil
}

// jitMunmap releases a mapping.
func jitMunmap(addr uintptr, size int) error {
	_, _, errno := syscall.Syscall(syscall.SYS_MUNMAP, addr, uintptr(size), 0)
	if errno != 0 {
		return errno
	}
	return nil
}

// jitCopyIn copies machine-code bytes into a mapping.
func jitCopyIn(addr uintptr, code []byte) {
	dst := unsafe.Slice((*byte)(unsafe.Pointer(addr)), len(code))
	copy(dst, code)
}
