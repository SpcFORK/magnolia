# lib/beam-vm

`lib/beam-vm.oak` is the **BEAM (Bogdan's/Björn's Abstract Machine) virtual
machine** written in Oak, on top of the Oak VM. It is modeled on the JVM
runtime in `lib/applet-runtime.oak`: the same state-record shape, the same
`execute`/`_execute` pair with a nested `tick(pc)` dispatch loop over a flat
`[op, arg, ref]` instruction stream, the same `fail`/depth/steps budget, and
the same `run(program, options)` entry point.

## What the BEAM is

The BEAM is a **register machine** (descended from JAM/WAM), **not** a stack
machine. Its defining properties, all modeled here:

- **1024 X registers** (X0..X1023). X0 is the return value; call arguments
  are passed in X0...
- **Y registers** = stack-frame slots, caller-saved across calls. A stack
  frame holds the saved continuation (CP), the frame pointer (FP), and the
  Y slots; the stack top is E.
- **Special registers**: `I` (instruction pointer), `CP` (continuation /
  return address), `FP` (frame pointer), `Htop` (top of heap), `E` (top of
  stack), `fcalls` (reduction counter).
- **Direct-threaded dispatch** — a single `tick` switch over the opcode set,
  the analogue of the JVM runtime's instruction dispatch.
- **`allocate` / `test_heap`** frame and heap management.
- **A process mailbox** for actor message passing (`send`, `self`,
  `receive`, `remove_message`).
- **Reduction-count preemption** — each instruction consumes one reduction;
  `fcalls` is the budget's sibling to the step counter.
- **A shared atom table** — atoms are interned strings, never GC'd.

## Instruction set

The generic (external) instruction tier: a compact opcode stream that the
loader expands. Each instruction is a `[op, arg, ref]` array; `op` selects,
`arg` carries an immediate, `ref` a pool index or target. X0 is the working
and return register.

| op | name | meaning |
|----|------|---------|
| 0  | nop | no-op |
| 1  | get_constant | load a literal from the pool into X0 |
| 2  | get_small | small integer literal (arg) |
| 3  | get_integer | integer literal via pool ref |
| 4  | get_float | float literal via pool ref |
| 5  | move | X0 = X(arg) |
| 6  | movex | X(ref) = X(arg) |
| 7  | movey | X0 = Y(arg) |
| 8  | sety | Y(arg) = X0 |
| 9  | get_tuple_element | X0 = tuple(arg).(ref) |
| 10 | set_tuple_element | tuple(arg).(ref) = X0 |
| 11 | put_tuple_element | fresh tuple with X(arg) at ref |
| 12-16 | add/sub/mul/div/rem | X0 = X(arg) OP X0 |
| 17-19 | neg/not/bnot | unary |
| 20-25 | cmp_lt/gt/eq/neq/leq/geq | X0 = bool |
| 26-33 | type tests | X0 = is_integer/float/atom/list/tuple/pid/ref/fun |
| 34 | make_tuple | X0 = tuple of arg slots |
| 35 | make_list | X0 = list of arg elements from Y at ref |
| 36 | cons_list | X0 = [X(arg) \| X0] |
| 37 | cons_tuple | X0 = {X(arg)} |
| 38 | jump | next = ref |
| 39 | jump_if | if X(arg) truthy, next = ref |
| 40 | jump_if_not | if X(arg) falsy, next = ref |
| 41 | jump_if_eq | if X(arg) = X0, next = ref |
| 42 | jump_if_neq | if X(arg) != X0, next = ref |
| 43 | fail | raise with X0 as the reason |
| 44 | call | call module function arg |
| 45 | return | pop frame, restore CP, X0 = result |
| 46 | call_ext | host call via the bounded host adapter |
| 47 | return_ext | return from an external call |
| 48 | apply | call the fun in X(arg) |
| 49 | call_fun | invoke a fun value directly (ref = index) |
| 50 | allocate | reserve arg heap words |
| 51 | test_heap | ensure arg words are live |
| 52 | deallocate | release arg words |
| 53 | self | X0 = current PID |
| 54 | send | send X0 to PID(arg) |
| 55 | receive | X0 = next mailbox payload |
| 56 | remove_message | discard oldest message |
| 57 | link | link to PID(arg) |
| 58 | unlink | remove link |
| 59 | make_pid | X0 = fresh PID |
| 60 | make_reference | X0 = fresh reference |
| 61 | try | set a trap |
| 62 | try_case | trap keyed on X(arg) |
| 63 | try_end | clear the trap |
| 64 | halt | stop the process; X0 = exit value |

## Program model

A compiled BEAM program is a `BeamModule`: a list of `BeamFunction`s (each
with a flat `[op, arg, ref]` code array, an arity, and a shared constant
pool), the module's exported function indices, and the shared pool. This is
the Oak-VM analogue of the JVM program's class table.

## Term representation

BEAM terms are tagged values. To stay inside Oak's type system while keeping
the machine faithful: atoms are plain Oak strings, integers/floats/booleans
are native, PIDs/references/ports are small tagged objects, lists are Oak
lists, and tuples are tagged objects. This mirrors the tagged-term model of
the BEAM without introducing a parallel GC heap.

## Compiler

`compileProgram(ast)` lowers an Oak-like AST (the same node shapes the
`wasm-vm` compiler handles — `:int`, `:float`, `:string`, `:atom`,
`:identifier`, `:unary`, `:binary`, `:list`, `:object`, `:block`,
`:ifExpr`, `:function`, `:fnCall`) into the generic BEAM instruction stream.
This is the loader's job in a real BEAM: expand the compact external tier
into the specific tier.

## Extending the VM

The VM is deliberately a single, self-contained module. To extend it:

1. Add a checked instruction to the opcode table above.
2. Implement it in the `step` dispatch.
3. Add parity tests against native Oak / the other VM forms.

The bounded host adapter (`options.host`) is the interop boundary, exactly
as the JVM runtime's bounded host adapter is — external effects pass through
a checked contract, keeping the VM self-contained and bounded.
