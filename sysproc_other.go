//go:build !windows

package main

import "fmt"

func lookupSysProc(library, name string) (uintptr, error) {
	return 0, fmt.Errorf("sysproc is only supported on Windows")
}
