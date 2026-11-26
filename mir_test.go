package main

import (
	"strings"
	"testing"
)

// MIR (M2 TODO-13) differential: the same analysis driver must produce an
// identical result string on the tree-walker (-n) and the bytecode VM (-b).
// The driver parses production-style Oak source (qualified HOF pipe chain,
// recursion, capture/escape) with the self-hosted parser and runs the four
// mir.oak passes over the resulting AST.
const mirDriverProg = `
pp := import('syntax-parse')
mir := import('mir')

src := '
std := import(\'std\')
fn main_work() std.range(0, 20000, 1) |> std.map(fn(x) x * 2 + 1) |> std.filter(fn(x) x % 3 = 0) |> with std.reduce(0) fn(a, b) a + b

fn fib(n) if n < 2 {
	true -> n
	_ -> fib(n - 1) + fib(n - 2)
}

fn caps() {
	local := []
	out := []
	f := fn() local <- local + [1]
	out
}

fn trappy() {
	obj := {a: 1}
	if obj {
		true -> 1
		_ -> 0
	}
}
'

nodes := pp.parse(src)
b := mir.analyze(nodes)
w := b.functions.main_work
c := w.fused.0
out := ''
out <- out + 'n=' + string(len(b.functions))
out <- out + ' ops=' + string(c.ops.0) + ',' + string(c.ops.1) + ',' + string(c.ops.2)
out <- out + ' lit=' + string(c.stages.0.literal?) + string(c.stages.1.literal?) + string(c.stages.2.literal?)
out <- out + ' fusedCount=' + string(w.fusedCount)
out <- out + ' returns=' + string(w.returns)
fb := b.functions.fib
out <- out + ' fib.n.mutable=' + string(fb.bindings.n.mutable?)
out <- out + ' fib.returns=' + string(fb.returns)
cp := b.functions.caps
out <- out + ' caps.local.region=' + string(cp.bindings.local.region)
out <- out + ' caps.out.region=' + string(cp.bindings.out.region)
dg := b.diagnostics
out <- out + ' diag=' + string(len(dg))
if len(dg) > 0 {
	true -> {
		d0 := dg.0
		out <- out + ' d0=' + string(d0.func) + '/' + string(d0.kind) + '/' + string(d0.type)
	}
	_ -> ?
}
out
`

const mirDriverWant = "n=5 ops=map,filter,reduce lit=truetruetrue fusedCount=1 returns=int " +
	"fib.n.mutable=false fib.returns=unknown caps.local.region=shared caps.out.region=shared " +
	"diag=1 d0=trappy/ifTrueOnNonBool/object"

func runMirDriver(t *testing.T, bytecode bool) string {
	t.Helper()
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	var (
		v   Value
		err error
	)
	if bytecode {
		v, err = ctx.EvalBytecode(strings.NewReader(mirDriverProg))
	} else {
		v, err = ctx.Eval(strings.NewReader(mirDriverProg))
	}
	if err != nil {
		t.Fatalf("bytecode=%v error: %s", bytecode, err.Error())
	}
	s, ok := v.(*StringValue)
	if !ok {
		t.Fatalf("bytecode=%v: driver returned %s, want string", bytecode, v.String())
	}
	return string(*s)
}

func TestMIRDifferential(t *testing.T) {
	want := mirDriverWant
	gotN := runMirDriver(t, false)
	if gotN != want {
		t.Fatalf("tree-walker MIR result:\n got: %s\nwant: %s", gotN, want)
	}
	gotB := runMirDriver(t, true)
	if gotB != want {
		t.Fatalf("bytecode MIR result:\n got: %s\nwant: %s", gotB, want)
	}
}
