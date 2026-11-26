package main

import (
	"runtime"
	"testing"
)

// jitTestCodeXPlusOne is x86-64 machine code implementing fn(x) = x + 1:
//   48 89 f8     mov %rdi, %rax
//   48 83 c0 01  add $1, %rax
//   c3           ret
var jitTestCodeXPlusOne = []byte{0x48, 0x89, 0xF8, 0x48, 0x83, 0xC0, 0x01, 0xC3}

// jitTestCodeConst42 is fn() = 42: mov $42, %eax; ret
var jitTestCodeConst42 = []byte{0xB8, 0x2A, 0x00, 0x00, 0x00, 0xC3}

// jitTestCodeMul7 is fn(x) = x * 7 using imul: mov %rdi, %rax; imul $7, %rax; ret
//   48 89 f8        mov %rdi, %rax
//   48 6b c0 07     imul $7, %rax, %rax
//   c3              ret
var jitTestCodeMul7 = []byte{0x48, 0x89, 0xF8, 0x48, 0x6B, 0xC0, 0x07, 0xC3}

func jitCompileForTest(t *testing.T, ctx *Context, code []byte) int64 {
	t.Helper()
	list := make([]Value, len(code))
	for i, b := range code {
		list[i] = IntValue(int64(b))
	}
	lv := ListValue(list)
	v, rerr := ctx.oakRuntimeJITCompile([]Value{&lv})
	if rerr != nil {
		t.Fatalf("compile runtime error: %v", rerr)
	}
	if obj, ok := v.(ObjectValue); ok {
		if _, isErr := obj["error"]; isErr {
			t.Fatalf("compile returned error object: %v", obj["error"])
		}
	}
	addr, ok := v.(IntValue)
	if !ok {
		t.Fatalf("compile returned non-int: %T", v)
	}
	return int64(addr)
}

func TestRuntimeJITBridge(t *testing.T) {
	if runtime.GOOS != "linux" && runtime.GOOS != "windows" {
		t.Skipf("no executable-memory bridge on %s", runtime.GOOS)
	}
	if !oakJITCallSupported {
		t.Skip("function-pointer calls unsupported in this build configuration")
	}
	ctx := NewContextWithCwd()
	ctxPtr := &ctx

	addr := jitCompileForTest(t, ctxPtr, jitTestCodeXPlusOne)

	// call with one argument (SysV: x in %rdi)
	v, rerr := ctxPtr.oakRuntimeJITCall([]Value{IntValue(addr), IntValue(41)})
	if rerr != nil {
		t.Fatalf("call runtime error: %v", rerr)
	}
	if obj, ok := v.(ObjectValue); ok {
		if _, isErr := obj["error"]; isErr {
			t.Fatalf("call returned error object: %v", obj["error"])
		}
	}
	if got, ok := v.(IntValue); !ok || int64(got) != 42 {
		t.Fatalf("x+1(41) = %v, want 42", v)
	}

	// const fn with no args
	addr2 := jitCompileForTest(t, ctxPtr, jitTestCodeConst42)
	v2, _ := ctxPtr.oakRuntimeJITCall([]Value{IntValue(addr2)})
	if got, ok := v2.(IntValue); !ok || int64(got) != 42 {
		t.Fatalf("const42() = %v, want 42", v2)
	}

	// multiplication via imul immediate
	addr3 := jitCompileForTest(t, ctxPtr, jitTestCodeMul7)
	v3, _ := ctxPtr.oakRuntimeJITCall([]Value{IntValue(addr3), IntValue(6)})
	if got, ok := v3.(IntValue); !ok || int64(got) != 42 {
		t.Fatalf("x*7(6) = %v, want 42", v3)
	}

	// free both mappings
	for _, a := range []int64{addr, addr2, addr3} {
		fv, rerr := ctxPtr.oakRuntimeJITFree([]Value{IntValue(a)})
		if rerr != nil {
			t.Fatalf("free runtime error: %v", rerr)
		}
		if fv != BoolValue(true) {
			t.Fatalf("free returned %v, want true", fv)
		}
	}

	// double free must fail cleanly
	if fv, _ := ctxPtr.oakRuntimeJITFree([]Value{IntValue(addr)}); fv != BoolValue(true) {
		if _, ok := fv.(ObjectValue); !ok {
			t.Fatalf("double free must return error object, got %T", fv)
		}
	}

	// calling a freed address must fail cleanly (registry check)
	if fv, _ := ctxPtr.oakRuntimeJITCall([]Value{IntValue(addr), IntValue(1)}); fv == BoolValue(true) {
		t.Fatalf("call after free must not succeed")
	}
}

func TestRuntimeJITCallUnsupportedBuild(t *testing.T) {
	if oakJITCallSupported {
		t.Skip("supported build")
	}
	ctx := NewContextWithCwd()
	ctxPtr := &ctx
	v, _ := ctxPtr.oakRuntimeJITCall([]Value{IntValue(0x1000), IntValue(1)})
	if obj, ok := v.(ObjectValue); !ok {
		t.Fatalf("unsupported build must return error object, got %T", v)
	} else if _, isErr := obj["error"]; !isErr {
		t.Fatalf("unsupported build must return error object, got %v", v)
	}
}
