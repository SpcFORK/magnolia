package main

import (
	"os"
	"testing"
)

func TestIsStdinReadableWithRegularFile(t *testing.T) {
	original := os.Stdin
	t.Cleanup(func() {
		os.Stdin = original
	})

	tmp, err := os.CreateTemp(t.TempDir(), "stdin-*")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer tmp.Close()

	os.Stdin = tmp
	if !isStdinReadable() {
		t.Fatalf("expected stdin to be readable when using regular file")
	}
}

func TestPerformCommandIfExistsUnknownCommand(t *testing.T) {
	if performCommandIfExists("definitely-not-a-real-command") {
		t.Fatalf("expected unknown command to return false")
	}
}
