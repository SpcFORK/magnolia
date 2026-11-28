//go:build windows

package main

// oakJITCallSupported reports whether oakSyscallN can invoke an arbitrary
// native function pointer (required by the ___runtime_jit_call builtin).
const oakJITCallSupported = true

import "syscall"

func oakSyscallN(trap uintptr, args ...uintptr) (uintptr, uintptr, syscall.Errno) {
	return syscall.SyscallN(trap, args...)
}
