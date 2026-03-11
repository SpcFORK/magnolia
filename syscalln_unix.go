//go:build !windows

package main

import "syscall"

func oakSyscallN(trap uintptr, args ...uintptr) (uintptr, uintptr, syscall.Errno) {
	var a [6]uintptr
	if len(args) > len(a) {
		return 0, 0, syscall.E2BIG
	}
	copy(a[:], args)

	if len(args) <= 3 {
		return syscall.Syscall(trap, a[0], a[1], a[2])
	}
	return syscall.Syscall6(trap, a[0], a[1], a[2], a[3], a[4], a[5])
}
