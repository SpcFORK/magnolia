//go:build !windows && (!linux || !cgo)
	return 0, fmt.Errorf("native shared-library sysproc requires Windows or Linux with cgo enabled")
