//go:build windows

package main

import (
	"fmt"
	"sync"

	"golang.org/x/sys/windows"
)

var (
	sysProcMu    sync.Mutex
	sysProcCache = map[string]uintptr{}
	dllCache     = map[string]*windows.LazyDLL{}
)

func lookupSysProc(library, name string) (uintptr, error) {
	cacheKey := library + "\x00" + name

	sysProcMu.Lock()
	if addr, ok := sysProcCache[cacheKey]; ok {
		sysProcMu.Unlock()
		return addr, nil
	}

	dll, ok := dllCache[library]
	if !ok {
		dll = windows.NewLazyDLL(library)
		dllCache[library] = dll
	}
	sysProcMu.Unlock()

	proc := dll.NewProc(name)
	if err := proc.Find(); err != nil {
		return 0, fmt.Errorf("GetProcAddress(%s, %s): %w", library, name, err)
	}
	addr := proc.Addr()

	sysProcMu.Lock()
	sysProcCache[cacheKey] = addr
	sysProcMu.Unlock()
	return addr, nil
}
