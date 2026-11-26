//go:build linux && cgo

package main


/*
#include <stdint.h>
#include <errno.h>

// The existing Oak syscall builtin accepts a resolved function pointer.
// On Unix that is NOT a kernel syscall number. Use a C ABI trampoline, with
// an exact prototype for each supported arity; all current GUI/libc inputs
// are integer, pointer, or boolean machine words (never floating-point ABI).
static uintptr_t oak_call_function(uintptr_t address, uintptr_t *a,
                                   int n, int *call_errno) {
    uintptr_t result = 0;
    errno = 0;
    switch (n) {
    case 0: result = ((uintptr_t (*)(void))address)(); break;
    case 1: result = ((uintptr_t (*)(uintptr_t))address)(a[0]); break;
    case 2: result = ((uintptr_t (*)(uintptr_t,uintptr_t))address)(a[0],a[1]); break;
    case 3: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2]); break;
    case 4: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3]); break;
    case 5: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4]); break;
    case 6: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4],a[5]); break;
    case 7: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4],a[5],a[6]); break;
    case 8: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4],a[5],a[6],a[7]); break;
    case 9: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4],a[5],a[6],a[7],a[8]); break;
    case 10: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4],a[5],a[6],a[7],a[8],a[9]); break;
    case 11: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4],a[5],a[6],a[7],a[8],a[9],a[10]); break;
    case 12: result = ((uintptr_t (*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))address)(a[0],a[1],a[2],a[3],a[4],a[5],a[6],a[7],a[8],a[9],a[10],a[11]); break;
    }
    *call_errno = errno;
    return result;
}
*/
import "C"

import "syscall"

// oakJITCallSupported reports whether oakSyscallN can invoke an arbitrary
// native function pointer (required by the ___runtime_jit_call builtin).
const oakJITCallSupported = true

func oakSyscallN(trap uintptr, args ...uintptr) (uintptr, uintptr, syscall.Errno) {
	if trap == 0 {
		return 0, 0, syscall.EINVAL
	}
	if len(args) > 12 {
		return 0, 0, syscall.E2BIG
	}
	var words [12]C.uintptr_t
	for i, arg := range args {
		words[i] = C.uintptr_t(arg)
	}
	var errno C.int
	result := C.oak_call_function(C.uintptr_t(trap), &words[0],
		C.int(len(args)), &errno)
	return uintptr(result), 0, syscall.Errno(errno)
}