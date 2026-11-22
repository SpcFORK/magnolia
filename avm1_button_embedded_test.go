package main

import (
	"strings"
	"testing"
)

// This is an owned synthetic action sequence. It checks that button-only
// _root.play is opt-in and does not add it to the old preloader host.
func TestEmbeddedAVM1ButtonPlayIsOptIn(t *testing.T) {
	ctx := NewContext(t.TempDir())
	defer ctx.Wait()
	ctx.LoadBuiltins()
	val, err := ctx.Eval(strings.NewReader(`
		std := import('std')
		vm := import('avm1-preloader')
		// Push int 0, Push string _root, GetVariable, Push string play,
		// CallMethod, Pop, End. No original media bytes are loaded.
		code := [150,12,0,7,0,0,0,0,0,95,114,111,111,116,0,
			28,150,6,0,0,112,108,97,121,0,82,23,0] |>
			std.reduce('', fn(acc, v) acc + char(v))
		old := vm.run(code, 1, 1)
		button := vm.runButton(code)
		old.type = :error & button.type = :ok & button.stackDepth = 0 &
			len(button.events) = 1 & button.events.(0).target = '_root' &
			button.events.(0).name = 'play'
	`))
	if err != nil {
		t.Fatalf("embedded AVM1 button action failed: %v", err)
	}
	if !val.Eq(BoolValue(true)) {
		t.Fatalf("button/preloader capability separation did not hold: %v", val)
	}
}