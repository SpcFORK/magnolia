package main

import (
	"strings"
	"testing"
)

func TestClassProbe(t *testing.T) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	progs := []string{
		"cs C {\n  x := 1\n}\nc := C()\nc.x\n",
		"cs C {\n  base := 10\n  (p) -> {\n    y := p.x + 1\n  }\n}\nc := C()\nc.y\n",
	}
	for i, prog := range progs {
		tv, err := ctx.Eval(strings.NewReader(prog))
		t.Logf("prog%d tree: %v err=%v", i, tv, err)
		bv, err2 := ctx.EvalBytecode(strings.NewReader(prog))
		t.Logf("prog%d byte: %v err=%v", i, bv, err2)
	}
}
