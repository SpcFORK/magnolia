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
