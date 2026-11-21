//go:build linux && cgo

package main

import (
	"os"
	"runtime"
	"syscall"
	"testing"
	"unsafe"
)

func TestLinuxSysprocCallsSharedLibraryNotKernelSyscall(t *testing.T) {
	pid, err := lookupSysProc("libc.so.6", "getpid")
	if err != nil || pid == 0 {
		t.Fatalf("resolve libc.getpid: addr=%#x err=%v", pid, err)
	}
	result, _, callErr := oakSyscallN(pid)
	if callErr != 0 || int(result) != os.Getpid() {
		t.Fatalf("libc.getpid: got %d errno %v want %d", result,
			callErr, os.Getpid())
	}
	strlen, err := lookupSysProc("libc.so.6", "strlen")
	if err != nil || strlen == 0 {
		t.Fatalf("resolve libc.strlen: addr=%#x err=%v", strlen, err)
	}
	raw := []byte("Magnolia\x00")
	length, _, callErr := oakSyscallN(strlen,
		uintptr(unsafe.Pointer(&raw[0])))
	runtime.KeepAlive(raw)
	if callErr != 0 || length != 8 {
		t.Fatalf("libc.strlen: got %d errno %v", length, callErr)
	}
	if _, err := lookupSysProc("libc.so.6", "not_an_oak_symbol"); err == nil {
		t.Fatal("missing symbol unexpectedly resolved")
	}
	_, _, callErr = oakSyscallN(pid, make([]uintptr, 13)...)
	if callErr != syscall.E2BIG {
		t.Fatalf("overlong native call: got %v want E2BIG", callErr)
	}
}