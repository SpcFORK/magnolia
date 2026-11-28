//go:build !windows && (!linux || !cgo)

package main

import (
	"fmt"
)

// lookupSysProc is the portable fallback for native shared-library procedure
// resolution when CGO is disabled or the target is not Linux/Windows.
// Native sysproc requires either Windows (LoadLibrary/GetProcAddress via
// golang.org/x/sys/windows) or Linux with cgo enabled (dlopen/dlsym).
func lookupSysProc(library, name string) (uintptr, error) {
	return 0, fmt.Errorf("native shared-library sysproc requires Windows or Linux with cgo enabled")
}
