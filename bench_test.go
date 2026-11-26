package main

import (
	"strings"
	"testing"
)

func benchEval(b *testing.B, program string) {
	b.Helper()
	// Pre-parse once to isolate eval overhead
	ctx := NewContext(".")
	ctx.LoadBuiltins()
	for i := 0; i < b.N; i++ {
		_, err := ctx.Eval(strings.NewReader(program))
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ---------------------------------------------------------------------------
// Milestone M0 (TODO-03, §6.2): pure-execution benchmarks.
//
// The end-to-end benchmarks above measure tokenize + parse + optimize +
// compile + execute on every iteration. The harnesses below hoist all
// front-end work before b.ResetTimer() so ns/op and allocs/op isolate the
// actual runtime engine (tree-walker or bytecode VM).
// ---------------------------------------------------------------------------

// benchParseOpt lexes, parses, and constant-folds a program once.
func benchParseOpt(b *testing.B, program string) []astNode {
	b.Helper()
	tk := newTokenizer(program, "bench.oak")
	tokens := tk.tokenize()
	p := newParser(tokens)
	nodes, err := p.parse()
	if err != nil {
		b.Fatal(err)
	}
	return optimizeAST(nodes)
}

// benchEvalPreparsed measures pure tree-walking evaluation of pre-parsed AST.
func benchEvalPreparsed(b *testing.B, program string) {
	b.Helper()
	ctx := NewContext(".")
	ctx.LoadBuiltins()
	nodes := benchParseOpt(b, program)
	if _, rerr := ctx.evalNodes(nodes); rerr != nil {
		b.Fatal(rerr)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, rerr := ctx.evalNodes(nodes); rerr != nil {
			b.Fatal(rerr)
		}
	}
}

// benchBytecodePrecompiled measures pure VM execution over a chunk compiled
// once outside the timer; each iteration runs a freshly pooled VM.
func benchBytecodePrecompiled(b *testing.B, program string) {
	b.Helper()
	ctx := NewContext(".")
	ctx.LoadBuiltins()
	nodes := benchParseOpt(b, program)
	chunk := compileToByteCode(nodes)
	if _, rerr := acquireVM(chunk, &ctx).run(); rerr != nil {
		b.Fatal(rerr)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, rerr := acquireVM(chunk, &ctx).run(); rerr != nil {
			b.Fatal(rerr)
		}
	}
}

// Pure tree-walker execution benchmarks (pre-parsed)

func BenchmarkEvalPreparsedFib25(b *testing.B) {
	benchEvalPreparsed(b, `
fn fib(n) if n {
	0 -> 0
	1 -> 1
	_ -> fib(n - 1) + fib(n - 2)
}
fib(25)
`)
}

func BenchmarkEvalPreparsedLoopCounter(b *testing.B) {
	benchEvalPreparsed(b, `
fn loop(n) {
	i := 0
	fn go {
		if i < n {
			true -> {
				i <- i + 1
				go()
			}
		}
	}
	go()
	i
}
loop(200)
`)
}

func BenchmarkEvalPreparsedFnCallOverhead(b *testing.B) {
	benchEvalPreparsed(b, `
fn add(a, b) a + b
total := 0
fn acc(i) if i < 1000 {
	true -> {
		total <- add(total, i)
		acc(i + 1)
	}
	_ -> total
}
acc(0)
`)
}

func BenchmarkEvalPreparsedPropertyAccess(b *testing.B) {
	benchEvalPreparsed(b, `
obj := { x: 1, y: 2, z: 3 }
sum := 0
fn go(i) if i < 1000 {
	true -> {
		sum <- obj.x + obj.y + obj.z
		go(i + 1)
	}
	_ -> sum
}
go(0)
`)
}

// Pure bytecode-VM execution benchmarks (pre-compiled)

func BenchmarkBytecodePrecompiledFib25(b *testing.B) {
	benchBytecodePrecompiled(b, `
fn fib(n) if n {
	0 -> 0
	1 -> 1
	_ -> fib(n - 1) + fib(n - 2)
}
fib(25)
`)
}

func BenchmarkBytecodePrecompiledLoopCounter(b *testing.B) {
	benchBytecodePrecompiled(b, `
fn loop(n) {
	i := 0
	fn go {
		if i < n {
			true -> {
				i <- i + 1
				go()
			}
		}
	}
	go()
	i
}
loop(200)
`)
}

func BenchmarkBytecodePrecompiledFnCallOverhead(b *testing.B) {
	benchBytecodePrecompiled(b, `
fn add(a, b) a + b
total := 0
fn acc(i) if i < 1000 {
	true -> {
		total <- add(total, i)
		acc(i + 1)
	}
	_ -> total
}
acc(0)
`)
}

func BenchmarkBytecodePrecompiledPropertyAccess(b *testing.B) {
	benchBytecodePrecompiled(b, `
obj := { x: 1, y: 2, z: 3 }
sum := 0
fn go(i) if i < 1000 {
	true -> {
		sum <- obj.x + obj.y + obj.z
		go(i + 1)
	}
	_ -> sum
}
go(0)
`)
}

// Core allocation benchmarks

func BenchmarkFib20(b *testing.B) {
	benchEval(b, `
fn fib(n) if n {
	0 -> 0
	1 -> 1
	_ -> fib(n - 1) + fib(n - 2)
}
fib(20)
`)
}

func BenchmarkFib25(b *testing.B) {
	benchEval(b, `
fn fib(n) if n {
	0 -> 0
	1 -> 1
	_ -> fib(n - 1) + fib(n - 2)
}
fib(25)
`)
}

func BenchmarkLoopCounter(b *testing.B) {
	benchEval(b, `
fn loop(n) {
	i := 0
	fn go {
		if i < n {
			true -> {
				i <- i + 1
				go()
			}
		}
	}
	go()
	i
}
loop(200)
`)
}

func BenchmarkListBuild(b *testing.B) {
	benchEval(b, `
fn buildList(n) {
	result := []
	fn go(i) if i < n {
		true -> {
			result << i
			go(i + 1)
		}
	}
	go(0)
	result
}
buildList(500)
`)
}

func BenchmarkObjectCreation(b *testing.B) {
	benchEval(b, `
fn makeObjects(n) {
	fn go(i, acc) if i < n {
		true -> go(i + 1, acc << { x: i, y: i * 2, name: 'item' })
		_ -> acc
	}
	go(0, [])
}
makeObjects(200)
`)
}

func BenchmarkStringConcat(b *testing.B) {
	benchEval(b, `
fn buildStr(n) {
	s := ''
	fn go(i) if i < n {
		true -> {
			s <- s + 'ab'
			go(i + 1)
		}
	}
	go(0)
	s
}
buildStr(200)
`)
}

func BenchmarkScopeCreation(b *testing.B) {
	benchEval(b, `
fn nested(n) if n {
	0 -> 0
	_ -> {
		x := n
		nested(n - 1) + x
	}
}
nested(500)
`)
}

func BenchmarkFnCallOverhead(b *testing.B) {
	benchEval(b, `
fn add(a, b) a + b
fn callMany(n) {
	result := 0
	fn go(i) if i < n {
		true -> {
			result <- add(result, 1)
			go(i + 1)
		}
	}
	go(0)
	result
}
callMany(200)
`)
}

func BenchmarkIfMatch(b *testing.B) {
	benchEval(b, `
fn classify(n) if n % 3 {
	0 -> :fizz
	1 -> :one
	2 -> :two
}
fn run(n) {
	fn go(i) if i < n {
		true -> {
			classify(i)
			go(i + 1)
		}
	}
	go(0)
}
run(200)
`)
}

func BenchmarkPropertyAccess(b *testing.B) {
	benchEval(b, `
obj := { a: 1, b: 2, c: 3, d: 4, e: 5 }
fn access(n) {
	result := 0
	fn go(i) if i < n {
		true -> {
			result <- obj.a + obj.b + obj.c + obj.d + obj.e
			go(i + 1)
		}
	}
	go(0)
	result
}
access(500)
`)
}

// Memory allocation benchmarks

func BenchmarkAllocFib20(b *testing.B) {
	b.ReportAllocs()
	benchEval(b, `
fn fib(n) if n {
	0 -> 0
	1 -> 1
	_ -> fib(n - 1) + fib(n - 2)
}
fib(20)
`)
}

func BenchmarkAllocLoopCounter(b *testing.B) {
	b.ReportAllocs()
	benchEval(b, `
fn loop(n) {
	i := 0
	fn go {
		if i < n {
			true -> {
				i <- i + 1
				go()
			}
		}
	}
	go()
	i
}
loop(200)
`)
}

func BenchmarkAllocFnCalls(b *testing.B) {
	b.ReportAllocs()
	benchEval(b, `
fn add(a, b) a + b
fn callMany(n) {
	result := 0
	fn go(i) if i < n {
		true -> {
			result <- add(result, 1)
			go(i + 1)
		}
	}
	go(0)
	result
}
callMany(200)
`)
}

func BenchmarkAllocScopeNested(b *testing.B) {
	b.ReportAllocs()
	benchEval(b, `
fn nested(n) if n {
	0 -> 0
	_ -> {
		x := n
		nested(n - 1) + x
	}
}
nested(500)
`)
}

// Pipe-fusion workloads: std.map/filter chains with literal fnNode callbacks.
// The bytecode engine compiles these into in-frame opHOF* loops (TODO-12);
// the tree-walker runs the std.oak FnValue loops with per-element closure
// calls (§7.8, §11.2 pipe_fusion row).
func BenchmarkEvalPreparsedPipeFusion(b *testing.B) {
	benchEvalPreparsed(b, `
std := import('std')
xs := std.range(0, 64, 1)
xs |> std.map(fn(x) x + 1) |> std.filter(fn(x) x % 2 == 0)
`)
}

func BenchmarkBytecodePrecompiledPipeFusion(b *testing.B) {
	benchBytecodePrecompiled(b, `
std := import('std')
xs := std.range(0, 64, 1)
xs |> std.map(fn(x) x + 1) |> std.filter(fn(x) x % 2 == 0)
`)
}
