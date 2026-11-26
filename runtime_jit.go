// runtime_jit.go — executable-memory bridge for lib/jit.oak (plan §9.2).
//
// Exposes three builtins to Oak:
//
//	___runtime_jit_compile(codeBytes) -> address | error
//	    Allocates executable memory (mmap PROT_READ|PROT_WRITE then
//	    mprotect PROT_READ|PROT_EXEC on Unix; VirtualAlloc with
//	    PAGE_EXECUTE_READWRITE on Windows), copies the machine-code bytes
//	    in, and returns the entry address as an integer.
//	___runtime_jit_call(addr, args...) -> int64 result | error
//	    Invokes the native function at addr with the integer machine-word
//	    arguments using the platform function-pointer trampoline
//	    (oakSyscallN — C ABI trampoline on linux/cgo, syscall.SyscallN on
//	    Windows). Arguments follow the SysV / Windows calling conventions,
//	    so up to six integer arguments land in the conventional registers.
//	___runtime_jit_free(addr) -> true | error
//	    Releases the executable mapping. The mapping length is remembered
//	    at compile time, so the same address must be passed back.
//
// Non-cgo Unix builds cannot invoke function pointers safely
// (oakSyscallN's portable fallback interprets the target as a kernel
// syscall number), so ___runtime_jit_call reports an error there while
// compile/free still work (useful for tests and inspection).
package main

import "sync"

var (
	jitRegionsMu sync.Mutex
	jitRegions   = map[uintptr]uintptr{} // addr -> mapped length
)

func jitErr(message string) ObjectValue {
	return syscallErrObj(message)
}

// jitPageRound rounds a byte count up to the allocation granularity.
func jitPageRound(n int) int {
	return (n + jitAllocAlign - 1) & ^(jitAllocAlign - 1)
}

func (c *Context) oakRuntimeJITCompile(args []Value) (Value, *runtimeError) {
	if len(args) != 1 {
		return jitErr("___runtime_jit_compile requires 1 argument (list of byte values)"), nil
	}
	list, ok := args[0].(*ListValue)
	if !ok {
		return jitErr("___runtime_jit_compile argument must be a list of bytes"), nil
	}
	code := make([]byte, len(*list))
	for i, b := range *list {
		iv, ok := b.(IntValue)
		if !ok {
			return jitErr("___runtime_jit_compile bytes must be integers"), nil
		}
		code[i] = byte(int64(iv) & 0xff)
	}
	if len(code) == 0 {
		return jitErr("___runtime_jit_compile requires non-empty code"), nil
	}

	addr, err := jitMmapRW(len(code))
	if err != nil {
		return jitErr("jit mmap failed: " + err.Error()), nil
	}
	jitCopyIn(addr, code)
	if err := jitProtectRX(addr, len(code)); err != nil {
		_ = jitMunmap(addr, len(code))
		return jitErr("jit mprotect failed: " + err.Error()), nil
	}
	jitRegionsMu.Lock()
	jitRegions[addr] = uintptr(len(code))
	jitRegionsMu.Unlock()
	return IntValue(int64(addr)), nil
}

func (c *Context) oakRuntimeJITCall(args []Value) (Value, *runtimeError) {
	if !oakJITCallSupported {
		return jitErr("___runtime_jit_call requires a cgo Linux or Windows build"), nil
	}
	if len(args) < 1 {
		return jitErr("___runtime_jit_call requires at least the function address"), nil
	}
	addr, ok := args[0].(IntValue)
	if !ok || int64(addr) <= 0 {
		return jitErr("___runtime_jit_call address must be a positive integer"), nil
	}
	jitRegionsMu.Lock()
	_, known := jitRegions[uintptr(addr)]
	jitRegionsMu.Unlock()
	if !known {
		return jitErr("___runtime_jit_call address was not produced by ___runtime_jit_compile"), nil
	}
	words := make([]uintptr, 0, len(args)-1)
	for _, arg := range args[1:] {
		switch a := arg.(type) {
		case IntValue:
			words = append(words, uintptr(int64(a)))
		case BoolValue:
			if a {
				words = append(words, 1)
			} else {
				words = append(words, 0)
			}
		case PointerValue:
			words = append(words, uintptr(a))
		default:
			return jitErr("___runtime_jit_call arguments must be integers (machine words)"), nil
		}
	}
	result, _, _ := oakSyscallN(uintptr(addr), words...)
	return IntValue(int64(result)), nil
}

func (c *Context) oakRuntimeJITFree(args []Value) (Value, *runtimeError) {
	if len(args) != 1 {
		return jitErr("___runtime_jit_free requires 1 argument (address)"), nil
	}
	addr, ok := args[0].(IntValue)
	if !ok || int64(addr) <= 0 {
		return jitErr("___runtime_jit_free address must be a positive integer"), nil
	}
	jitRegionsMu.Lock()
	size, known := jitRegions[uintptr(addr)]
	if known {
		delete(jitRegions, uintptr(addr))
	}
	jitRegionsMu.Unlock()
	if !known {
		return jitErr("___runtime_jit_free address was not produced by ___runtime_jit_compile"), nil
	}
	if err := jitMunmap(uintptr(addr), int(size)); err != nil {
		return jitErr("jit munmap failed: " + err.Error()), nil
	}
	return BoolValue(true), nil
}
