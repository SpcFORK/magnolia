//go:build windows

package main

import "syscall"

func oakSyscallN(trap uintptr, args ...uintptr) (uintptr, uintptr, syscall.Errno) {
	return syscall.SyscallN(trap, args...)
}
