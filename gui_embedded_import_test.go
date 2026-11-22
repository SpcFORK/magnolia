package main

import (
	"strings"
	"testing"
)

// A library imported by name is backed by embedded source, not a filesystem
// directory. Its own GUI dependencies must therefore also resolve by name,
// even when the caller has no Magnolia checkout in its working directory.
func TestEmbeddedGUIImportsWithoutCheckout(t *testing.T) {
	ctx := NewContext(t.TempDir())
	defer ctx.Wait()
	ctx.LoadBuiltins()
	value, err := ctx.Eval(strings.NewReader(`
		gui := import('GUI')
		type(gui.createWindow) = :function & type(gui.backend) = :function
	`))
	if err != nil {
		t.Fatalf("embedded GUI import failed without checkout: %v", err)
	}
	if !value.Eq(BoolValue(true)) {
		t.Fatalf("embedded GUI facade did not expose its expected functions: %v", value)
	}
}