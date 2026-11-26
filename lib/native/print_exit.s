# print_exit.s — source of the 123-byte print+exit helper embedded in
# lib/codegen-native.oak (fn emitPrintExit). Assembled with gcc/ld -nostdlib.
# Contract: _start calls it with the top-level result in %rdi.
# Behavior: writes the signed decimal rendering of %rdi to stdout (fd 1),
# then exits the process with status 0. No callee-saved registers touched
# beyond the red zone; the caller's frame is untouched.
    .globl _print_exit
_print_exit:
    push %rbp
    mov  %rsp, %rbp
    sub  $64, %rsp              # digit scratch (also satisfies alignment)
    mov  %rdi, %rax
    xor  %r10d, %r10d           # length
    test %rax, %rax
    jne    .Lloop
    movb $'0', -1(%rbp)         # zero special case: digit at rbp-1,
    mov  $1, %r10d              # buffer lea must be rbp-2 (tail does inc %rsi)
    lea  -2(%rbp), %rsi
    jmp  .Lwrite
.Lloop:
    xor  %r11d, %r11d
    test %rax, %rax
    jns  .Ldiv
    mov  $1, %r11d              # negative flag
    neg  %rax
.Ldiv:
    lea  -1(%rbp), %rsi
    mov  $10, %ecx
.Ldigits:
    cqto
    idiv %rcx
    add  $'0', %dl
    mov  %dl, (%rsi)
    dec  %rsi
    inc  %r10d
    test %rax, %rax
    jne  .Ldigits
    test %r11d, %r11d
    je   .Lwrite
    movb $'-', (%rsi)
    dec  %rsi
    inc  %r10d
.Lwrite:
    inc  %rsi
    mov  $1, %eax               # sys_write
    mov  $1, %edi
    mov  %r10d, %edx
    syscall
    mov  $60, %eax              # sys_exit(0)
    xor  %edi, %edi
    syscall
    ret
