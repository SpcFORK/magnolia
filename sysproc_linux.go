//go:build linux && cgo

package main

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

// Resolve only explicit caller-selected shared libraries. The existing Oak
// sysproc/syscall API already exposes native code; this adds the Linux
// implementation rather than misusing syscall.Syscall (a KERNEL syscall).
// Handles remain open for the process lifetime so cached procedure addresses
// cannot dangle between GUI frames.
var linuxProcMu sync.Mutex
var linuxLibraries = map[string]unsafe.Pointer{}
var linuxProcedures = map[string]uintptr{}

func lookupSysProc(library, name string) (uintptr, error) {
	if library == "" || name == "" {
		return 0, fmt.Errorf("empty shared library or symbol")
	}
	key := library + "\x00" + name
	linuxProcMu.Lock()
	defer linuxProcMu.Unlock()
	if address := linuxProcedures[key]; address != 0 {
		return address, nil
	}
	handle := linuxLibraries[library]
	if handle == nil {
		path := C.CString(library)
		defer C.free(unsafe.Pointer(path))
		handle = C.dlopen(path, C.RTLD_NOW|C.RTLD_LOCAL)
		if handle == nil {
			detail := C.dlerror()
			if detail != nil {
				return 0, fmt.Errorf("dlopen(%s): %s", library, C.GoString(detail))
			}
			return 0, fmt.Errorf("dlopen(%s) failed", library)
		}
		linuxLibraries[library] = handle
	}
	symbol := C.CString(name)
	defer C.free(unsafe.Pointer(symbol))
	C.dlerror() // clear this thread's previous error
	address := C.dlsym(handle, symbol)
	if address == nil {
		detail := C.dlerror()
		if detail != nil {
			return 0, fmt.Errorf("dlsym(%s, %s): %s", library, name, C.GoString(detail))
		}
		return 0, fmt.Errorf("dlsym(%s, %s) failed", library, name)
	}
	resolved := uintptr(address)
	linuxProcedures[key] = resolved
	return resolved, nil
}