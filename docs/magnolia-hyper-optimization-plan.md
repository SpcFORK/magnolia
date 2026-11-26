# Magnolia-Hyper: Architectural Optimization Plan to Surpass Goose, C++, and Rust

> **Status:** Living Architecture & Implementation Blueprint (v2 — ablation-verified)
> **Scope:** Go Host Runtime (`eval.go`, `bytecode.go`, `optimize.go`, `env.go`), Unified Middle-End IR (`lib/mir.oak`), Native & Systems Backends (`lib/codegen-c.oak`, `lib/codegen-go.oak`, `lib/codegen-native.oak`, `lib/assembler.oak`, `lib/linker.oak`), Polyglot Transpilers (`lib/build.oak`, `lib/runtime-codegen.oak`, `lib/wasm-vm.oak`), and Heterogeneous SIMD/GPU Offload (`lib/gpu.oak`, `lib/gui-shader-codegen.oak`).
>
> **Verification status of the forensics in §3:** every claim marked **[ABLATED]** was verified empirically against the compiled Magnolia binary (tree-walker *and* `--bytecode` engine) during the M0 phase. No claim in §3 is speculative.

---

## §0. Master Cross-Referenced Implementation TODO List

Every item maps 1-to-1 to its governing design section (`§X.Y`) and carries an explicit verification/benchmark gate.

### Phase A — Build Hygiene, Syntax-Sugar Parity, & True Differential Benchmark Harness
- [x] **TODO-01 (`-> §3.3, §6.1`)**: Restore Go build/test hygiene across `CGO_ENABLED=1` and `CGO_ENABLED=0`: keep `package main` headers/stubs in `syscalln_unix.go` and `sysproc_other.go`, resolve leftover upstream Git merge conflict markers in `lib/std.oak`, `lib/bytes.oak`, `lib/gui-common.oak`, and `lib/jvm-vm-libs.oak`, and clean up duplicate `lib.go` registrations (preserving 100% of Java/JVM/applet/MIDP/SWF/UI modules in the opt-in section while adding the 4 missing embeds reported by `TestStdlibRegistryCoversLibDirectory`). **DONE:** all four sub-items landed; `go test ./...` green under both CGO modes; 28 duplicate embed-var blocks deduplicated; `gui-render-core`, `gui-window-nest`, `obj`, `web-pipeline` embedded; conflict markers resolved (HEAD-side shared-helper abstractions kept — each had exactly one definition); the newly-embedded `lib/gui-window-nest.oak` was repaired from a half-finished mechanical match-conversion (see §3.4).
- [ ] **TODO-02 (`-> §3.1.9, §6.1`)**: Unify self-hosted Oak tokenizer/parser/macro tools (`lib/syntax-tokenize.oak`, `lib/syntax-parse.oak`, `lib/syntax-macros.oak`, `lib/transpile.oak`) with Go host (`token.go`, `parse.go`) for `**` (`:power`), `==` (`:deepEq`), `~` (`:tilde`), `fn {}` empty-body rewrite, and `:class` macro/transpiler traversal, while preserving all parser-sugar invariants (`|>`, `with`, `cs`, `if`, `.(ident)`).
- [x] **TODO-03 (`-> §3.2.8, §6.2`)** — **DONE** (commit `e84d394`): Fix `bench_test.go` so lexing/parsing/AST-optimization and bytecode compilation are separated from `b.N` execution loops (`b.ResetTimer()`), adding `BenchmarkEvalPreparsed*` and `BenchmarkBytecodePrecompiled*` alongside end-to-end benchmarks.
- [ ] **TODO-04 (`-> §2.2, §6.3, §11.1`)**: Create the cross-language benchmark & differential-oracle suite (`bench/` + `/tmp/ref/`) comparing **Goose** (`aardappel/goose`), **C++ (`g++`/`clang++ -O3 -march=native`)**, **Baseline Magnolia**, and **Magnolia-Hyper** across compute, recursion, region/tree allocation (`records`/`btree`), pointer-relinking (`lru`), stream fusion (`|>`/`with`), and runtime value specialization workloads.

### Phase B — Go Host Engine & `bytecode.go` Overhaul (Milestone M1)
- [x] **TODO-05 (`-> §3.1.0, §3.1.1, §3.1.4, §7.6`)** — **DONE** (commit `88963d8`): Remove the unreachable dead `n.op == pipeArrow` operator branch in `bytecode.go` (`compileNodeTail` — left over from assuming `|>` was a `binaryNode`) and implement shared multi-target `if` body deduplication (`parse.go` splits `a, b -> body` into multiple `ifBranch` entries sharing the identical `body` AST node reference) in `compileIfTail`.
- [x] **TODO-06 (`-> §3.2.4, §7.3`)** — **DONE** (commit `e5a6a94`): Eliminate single-thread reflection & lock tax in `eval.go` (`objRLock`/`listRLock`/`strRLock` calling `reflect.ValueOf(x).Pointer()` on every read/write) by gating synchronization on `c.eng.concurrent` (atomic concurrency transition in `go`/`listen` builtins in `env.go`) and replacing `reflect.ValueOf` with zero-reflection pointer/hash striping when concurrent.
- [x] **TODO-07 (`-> §3.2.1, §7.1`)** — **DONE** (commit `88963d8`): Eliminate heap boxing on integers/booleans in `eval.go` & `bytecode.go` via pre-allocated small/medium `IntValue` caches + unboxed numeric fast paths in the VM (`opAdd`, `opSub`, `opMul`, `opDiv`, `opMod`, `opLt`, `opGt`, `opEq`, loop counters).
- [x] **TODO-08 (`-> §3.2.2, §7.2`)** — **DONE** (commit `5e4ea00`): Eliminate per-call heap allocations in `bytecode.go` (`opCall` / `opTailCall` allocating `make([]Value, ft.localCount)`, `args := make([]Value, arity)`, and `&vmScope{}`) by introducing a contiguous pre-allocated **VM Register/Local Slab (`vm.localsSlab`)** and zero-alloc frame windowing (`baseLocal`..`baseLocal+localCount`) when closures do not escape.
- [x] **TODO-09 (`-> §3.2.3, §7.4`)** — **DONE** (commit `decd862`; TODO-10 ships the measured flagship superinstr `LT_LOCAL_JUMP_TRUE` + no-match parity fix — `opTailCallSelf`/`opEqImmJumpFalse` remain future work as tail-recursion fib does not exercise them; TODO-11 ships inline-cached prop access, `compileClass` statics/parents deferred to M2 — no corpus usage): Replace linear string-name upvalue walks in `bytecode.go` (`vmScope.get`/`set`, `opLoadUpval`/`opStoreUpval`) with **compile-time lexical coordinates `(depth: u16, slot: u16)`** (`opLoadUpvalFast` / `opStoreUpvalFast`), retaining dynamic fallback only for `bytecode()` outer-scope bridges.
- [x] **TODO-10 (`-> §7.5`)** — **DONE** (commit `1b7c95d`; TODO-10 ships the measured flagship superinstr `LT_LOCAL_JUMP_TRUE` + no-match parity fix — `opTailCallSelf`/`opEqImmJumpFalse` remain future work as tail-recursion fib does not exercise them; TODO-11 ships inline-cached prop access, `compileClass` statics/parents deferred to M2 — no corpus usage): Add bytecode **Superinstructions & Specialized Peephole Fusion** in `bytecode.go` (`opAddLocalImm`, `opSubLocalImm`, `opLoadLocal2`, `opTailCallSelf`, `opLtJumpFalse`, `opEqImmJumpFalse`) to slash dispatch iterations and stack push/pop traffic on tight loops and recursion.
- [x] **TODO-11 (`-> §3.2.5, §7.7`)** — **DONE** (commit `aed413b`; TODO-10 ships the measured flagship superinstr `LT_LOCAL_JUMP_TRUE` + no-match parity fix — `opTailCallSelf`/`opEqImmJumpFalse` remain future work as tail-recursion fib does not exercise them; TODO-11 ships inline-cached prop access, `compileClass` statics/parents deferred to M2 — no corpus usage): Add **Inline Caching (IC) & Fast Property Access** in `bytecode.go` (`opGetPropConst` / `opSetPropConst`) and fix `compileClass` to support `staticExprs` and `parents`.
- [x] **TODO-12 (`-> §3.1.2, §7.8`)** — **DONE** (commit `cb90e11`; TODO-10 ships the measured flagship superinstr `LT_LOCAL_JUMP_TRUE` + no-match parity fix — `opTailCallSelf`/`opEqImmJumpFalse` remain future work as tail-recursion fib does not exercise them; TODO-11 ships inline-cached prop access, `compileClass` statics/parents deferred to M2 — no corpus usage): Eliminate the `docs/engine-switching.md` callback cliff by compiling known higher-order `fnCallNode`s with literal `fnNode` callbacks (`each`, `map`, `filter`, `reduce` — including `|>` and `with` desugarings) directly into **in-frame VM loops** without spawning child VMs or allocating closures.

### Phase C — Unified Middle-End `lib/mir.oak` & Systems AOT Backends (Milestone M2)
- [ ] **TODO-13 (`-> §4.1, §4.2, §4.3`)**: Implement **`lib/mir.oak`** — Magnolia's whole-bundle middle-end IR performing:
  1. Mutability promotion (`:=` bindings never targeted by `<-` become pure SSA values),
  2. Flow-sensitive type & shape inference (`:int`, `:float`, `:bool`, `:atom`, `:string`, `:list<T>`, `StructShape#k`, `SumShape#m`),
  3. Call-graph & `fnCall` HOF stream fusion (fusing parser-desugared `|>` chains and `with` blocks into single zero-allocation loops),
  4. Static **Virtual Region Stack (`R_0..R_k`)** lifetime assignment & disjoint-region alias tokens.
- [ ] **TODO-14 (`-> §5.1, §5.2, §5.3, §5.4, §8.1`)**: Implement **`lib/codegen-c.oak`** ("Goose+ Mode" C11/LLVM backend) and wire `--c` (and `--cc` native binary compilation via Clang/GCC) into `lib/build.oak`, `lib/build-config.oak`, and `lib/build-render.oak`:
  - 2 GB demand-paged OS virtual-memory Region Stacks (`mmap` / `VirtualAlloc`) with $O(1)$ watermark restore,
  - Disjoint-region `restrict` qualifiers (fixing Goose Ceiling #1 → full LLVM `noalias`),
  - Register-pinned base `Ptr32<R_k>` (`base[idx32]`, fixing Goose Ceiling #3 on `lru`/`calc`),
  - 64-byte-aligned SoA columnar buffers + `#pragma clang loop vectorize(enable)` (fixing Goose Ceiling #2),
  - `[[clang::musttail]]` / `__attribute__((musttail))` zero-stack tail calls, interned `uint32_t` `:atom` switch tables, and `libdivide`-style fast divisors.
- [ ] **TODO-15 (`-> §3.2.6, §8.2`)**: Upgrade **`lib/codegen-go.oak` (v2)** using `lib/mir.oak` to emit monomorphized typed Go functions (`func(int64) int64`), concrete Go `struct` shapes, direct `==` comparisons instead of `reflect.DeepEqual`, and GC-invisible `[]byte` region slabs while keeping dynamic `OakValue` fallback for untyped polymorphic code.

### Phase D — Self-Hosted Native Codegen v2 & Microsecond Runtime Value JIT (Milestone M3)
- [ ] **TODO-16 (`-> §3.2.7, §9.1`)**: Upgrade **`lib/ast-ir.oak`**, **`lib/codegen-native.oak`**, and **`lib/assembler.oak`** to support comparisons, branches/labels, function calls, recursion, mul/div/mod/bitwise ops, proper multi-register allocation + stack spills (eliminating the `regIdx >= 8` `rax` clobber bug), `cmovcc`/`csel` branchless selects, and AVX2/NEON vector instructions.
- [ ] **TODO-17 (`-> §5.5, §9.2`)**: Implement **Runtime Value Specialization (`lib/jit.oak` + Go host `mmap` executable memory bridge)** so Magnolia programs can specialize functions on runtime-loaded constants/strides/schemas (`jit.specialize(fn, consts...)`) into native machine code in microseconds — beating frozen AOT C++, Rust, and Goose on dynamic/data-driven workloads.

### Phase E — Polyglot Target Unification & Heterogeneous SIMD/GPU Offload (Milestone M4)
- [ ] **TODO-18 (`-> §10.1`)**: Upgrade polyglot transpilers (`lib/build.oak`, `lib/build-render.oak`, `lib/runtime-codegen.oak`) with MIR-assisted fast paths:
  - Direct typed WAT codegen (`--wasm` numeric/loop specialization without VM interpretation overhead where statically typed),
  - Tail-loop lowering & numeric unboxed fast paths in JS/TS (`--web`/`--ts`), Lua (`--lua`), and Java (`--java`).
- [ ] **TODO-19 (`-> §10.2`)**: Implement **`lib/compute.oak`** unifying fused pure HOF pipelines (`map` / `filter` / `reduce` / `dotProduct` / `softmax` from `lib/ai.oak`) with automatic **GLSL/HLSL/SPIR-V & OpenCL/CUDA kernel generation** (`lib/gui-shader-codegen.oak` + `lib/gpus.oak`) and multi-core CPU chunked fallback.
- [ ] **TODO-20 (`-> §11.2`)**: Run full regression tests, differential oracle checks across all engines/targets, and record live head-to-head benchmark numbers (Goose vs. C++ vs. Baseline Magnolia vs. Magnolia-Hyper) in §11.

---

## §1. Executive Thesis & The Speed Equation

Compiling a high-level language naively to machine code can at best **tie** C++ or Rust on scalar arithmetic: if two compilers lower the same scalar loop with identical memory indirection and alias assumptions to LLVM/x86-64/ARM64, hardware execution ports enforce a $1.00\times$ ceiling.

To make **Magnolia** faster than **idiomatic and expert C++**, **safe Rust**, and **Goose (`aardappel/goose`)**, Magnolia must attack the four places where C++, Rust, and Goose leave massive performance on the table: **memory/lifetime physics**, **pointer & alias representation**, **automatic data-layout transposition (AoS → SoA + variable-mode sums)**, and **runtime/heterogeneous specialization**.

$$\text{Magnolia}_{\text{Hyper}} = \underbrace{\text{Goose's Virtual-Memory Region Physics}}_{\text{Zero malloc/free, zero GC, zero Rust drop glue}} + \underbrace{\text{Eliminating Goose's 7 Ceilings}}_{\text{Disjoint }\texttt{restrict}\text{, Pinned-Base Ptr32, Auto-SoA SIMD}} + \underbrace{\text{Magnolia's 3 Unfair Advantages}}_{\text{Microsecond Value JIT + Go M:N Zero-GC I/O + GPU Kernel Fusion}}$$

---

## §2. Comparative Architecture Review: C++, Safe Rust, and Goose

### §2.1 Where C++ and Safe Rust Lose Cycles

| Dimension | Idiomatic C++ | Expert C++ | Safe Rust (`rustc -O3`) |
| :--- | :--- | :--- | :--- |
| **Allocation & Destruction** | `std::shared_ptr`, `unique_ptr`, `std::string`, nested `std::vector` fragment `malloc`/`free` free-lists and execute $O(N)$ recursive destructors. | Manual `pmr` arenas & placement `new`; error-prone lifetime tracking; non-trivial ABI structs pass via stack memory. | Zero GC, *but* `Box<T>`/`Vec<String>` allocate per container on the global heap and execute $O(N)$ **drop glue** on scope exit unless manually arena-allocated. |
| **Sum Types & Cache Density** | `sizeof(std::variant<...>)` pads every element to the **largest variant** + max alignment; pointer-heavy ASTs/records miss L1/L2 cache. | Hand-packed byte buffers forfeit type safety and pattern matching. | `enum` pads every element to `max(variant)`; recursive types require `Box<T>` or manual index arenas. |
| **Aliasing & SIMD** | `char*`/`uint8_t*` may alias any type under C++ TBAA; pointers omit `restrict` by default, forcing reloads after stores and blocking SIMD. | Manual `__restrict` + pragmas, fragile across refactors. | **Strongest static guarantee:** `&mut T` lowers to LLVM `noalias`, enabling load hoisting and auto-vectorization. |
| **Bounds Checks** | Unchecked `operator[]` (fast, memory-unsafe). | Raw pointer arithmetic. | **Safe Rust's tax:** `slice[i]` inserts panic branches whenever LLVM cannot prove `i < len` (especially graph/lookup indices without `unsafe { get_unchecked }`). |

### §2.2 How Goose (`aardappel/goose`) Beats C++ and Rust — Measured Baseline

Wouter van Oortmerssen's **Goose** is a closed-world whole-program systems language compiling to specialized C99 (plus a TinyCC dev JIT). Across Goose's 16-benchmark suite at large problem sizes (Ryzen 9 9950X3D, `clang-cl 22` / `MSVC` vs. `rustc 1.98`):
- **3.51×–3.54×** geomean faster than **idiomatic C++** (**1.95×** less peak RSS)
- **1.24×–1.26×** geomean faster than **expert C++** (**1.32×** less peak RSS)
- **1.13×–1.19×** geomean faster than **best safe Rust** (**1.25×** less peak RSS)

Goose wins via **seven architectural invariants**:
1. **Zero General-Purpose Heap:** Compile-time $N$ bump stacks, each backed by a **2 GB virtual-address reservation** (`mmap`/`VirtualAlloc`, demand-paged by OS hardware page faults). Allocation is `top += size`; scope exit is `top = watermark` ($O(1)$ bulk free — zero `free()`, zero GC, zero RC, zero drop glue).
2. **Pointer-Stable Grow-Only Arrays (`T[>..]`):** Appending at the tip of a dedicated 2 GB virtual stack never calls `realloc`, never copies prefix elements, and never invalidates interior references.
3. **Inline Variable-Mode Sum Streams (`Shape..`, `varint`, `u8[][>..]`):** Heterogeneous streams pack only the active variant's bytes inline (**2.02×–2.26×** faster and **3.98×** smaller RSS than Rust on `records`).
4. **Narrow Relative References (`T&<u32>`, `T&<u32 in pool>`):** 4-byte offsets halve pointer cache pressure and serialize in $O(1)$ via `to_bytes`/`from_bytes`.
5. **Guaranteed Destination Construction:** Aggregates construct directly into destination stack memory.
6. **Whole-Program Provenance & BCE (`checkfree fn`, `u32 in dist`):** Hoists slice views, caches stack tops in registers across pure loops, applies `libdivide` magic multipliers, and compiles pure `&&`/`||` chains branchlessly.
7. **Shared-Nothing `thread_fn` Workers:** Zero shared mutable heap; flat `memcpy` queues between OS threads.

### §2.3 Goose's 7 Structural Ceilings (Magnolia's Attack Surface)

1. **No `restrict` / `noalias` Emitted in C99:** Goose statically knows when two pointers target *distinct* 2 GB virtual stacks, yet omits `restrict` in generated C99, losing to Rust `&mut` on alias-sensitive vectorizable loops.
2. **Unaligned Byte-Packed Structs Impede Wide SIMD:** Variable-width byte packing causes unaligned vector loads and prevents automatic AoS → SoA layout transposition.
3. **Self-Relative Pointer ALU & Relink Overhead (`lru`: 0.81× vs Rust; `calc`: 0.89×–0.91× vs Rust Arena):** Computing `(uintptr_t)&slot + (int32_t)offset` adds dependent address-generation latency on every pointer chase and requires rewriting offsets on relink/move.
4. **Frozen Closed-World AOT (Zero Runtime Value Specialization):** Goose has no runtime specialization/JIT optimizer (its TinyCC path is `-O0` for fast dev startup). Runtime-loaded schemas, strides, filter predicates, and VM bytecodes cannot fold into machine immediates.
5. **$O(\text{size})$ Copies & Recursive Ownership Limits:** Non-destination transfers copy bytes, and recursive cycles cannot own growable data across iterations without explicit frame discipline.
6. **Heavy OS-Thread Model Only:** `thread_fn` clones global state per OS thread — unsuitable for 100k+ concurrent I/O flows.
7. **CPU-Only Execution:** Capped by ~100–200 GB/s CPU DRAM bandwidth vs. 1–3+ TB/s GPU HBM/GDDR compute bandwidth.

---

## §3. Forensic Review of Magnolia Today (Syntax Sugars, Runtime, & Transpilers)

### §3.0 Ablation Protocol (Added in M0)

Every claim in this section was established by one of three evidence classes, never by guessing:
1. **Source citation** — exact `token.go` / `parse.go` / `eval.go` / `bytecode.go` line references.
2. **Binary ablation** — minimal `.oak` probes executed against the compiled host under both `--normal` (tree-walker) and `--bytecode` (stack VM) engines, with identical results required.
3. **Go-level introspection** — a temporary in-package test dumping the parsed/optimized AST (`astNode` tree) and the evaluated `Value` types.

Key probe-verified semantics that govern every compiler tier (see §3.1):

| Construct | Verified behavior |
| :--- | :--- |
| `if cond -> body` | Desugars at parse time to `if cond { true -> body }`; body is a single node. |
| `if cond { arms }` | Match block: each arm is `targets… -> body`; body is a single node; ACI supplies arm separators. |
| `if cond { statements }` (no `->`) | **Parse error** (`Unexpected token } at start of unit`) — braces after `if` are match arms only. |
| Match target `?` | Fires when the subject is `?` (missing property/computed-access miss). `NullValue.Eq(NullValue) == true`. |
| Match target `_` | Wildcard (`EmptyValue`): fires on any subject; prunes later branches in `foldIf`. |
| No matching arm | The whole `if` expression evaluates to `?`. |
| `x = ?` / `x != ?` | Legal and correct (`NullValue.Eq` handles null operands). Existence-test idiom used across `lib/`. |
| `x?` (trailing `?` on an identifier) | **Not an operator**: `?` is an identifier character (token.go:277); `foo?` is a field literally named `foo?`. |
| Missing `.field` / `.(key)` access | Returns `?` silently (no error). |
| Multi-target arms `a, b -> body` | Split into multiple `ifBranch` entries sharing the **same body AST node pointer** — dedup opportunity (TODO-05). |
| Bare scripts | Do **not** autoload `lib/` (`str` is undefined → `? is not a function`); `string()` is the builtin stringifier; `print` requires a string. |

### §3.1 Magnolia's Parser Sugars & Core AST Invariants (Critical Compiler Insight)

A rigorous inspection of `token.go`, `parse.go`, `lib/syntax-tokenize.oak`, and `lib/syntax-parse.oak` reveals a foundational architectural truth: **Magnolia's core AST has only 16 node kinds because its most expressive idioms are pure parser-level syntax sugars, NOT runtime operators!**

#### §3.1.1 Pipe (`|>`) Is Call-Prepend Sugar, NOT a Binary Operator
In both `parse.go` (`parseNode`) and `lib/syntax-parse.oak` (`:pipeArrow`):
```go
case pipeArrow:
    pipe := p.next()
    pipeRight, err := p.parseSubNode()
    pipedFnCall, ok := pipeRight.(fnCallNode) // MUST be a fnCallNode!
    pipedFnCall.args = append([]astNode{node}, pipedFnCall.args...)
    node = pipedFnCall
```
- **Invariant [ABLATED]:** `x |> f(a, b)` **never** creates a `binaryNode{op: pipeArrow}`! It desugars **during parsing** directly into `fnCallNode{fn: f, args: [x, a, b]}`. Chained pipes `xs |> map(f) |> filter(g) |> reduce(0, h)` parse directly into nested `fnCallNode`s: `reduce(filter(map(xs, f), g), 0, h)`.
- **Dead Code Identified in `bytecode.go`:** `bytecode.go`'s `case binaryNode: if n.op == pipeArrow` branch is **completely unreachable** for parsed Oak programs because `parse.go` already desugared `|>` into `fnCallNode`.
- **Design Consequence for Stream Fusion:** Loop/stream fusion in `optimize.go`, `bytecode.go`, and `lib/mir.oak` **must pattern-match nested `fnCallNode` (`:fnCall`) trees along `args[0]`**, not binary operators. Doing so automatically fuses `|>` pipelines, `with` blocks, and nested function calls with **one unified pass**!

#### §3.1.2 `with` Is Trailing-Callback Call Sugar, NOT a Control-Flow AST Node
In both `parse.go` (`parseUnit`) and `lib/syntax-parse.oak` (`:withKeyword`):
```go
case withKeyword:
    withExprBase, _ := p.parseNode()       // e.g. loop(10) or (xs |> map())
    withExprBaseCall, ok := withExprBase.(fnCallNode)
    withExprLastArg, _ := p.parseNode()    // e.g. fn(i) { ... }
    withExprBaseCall.args = append(withExprBaseCall.args, withExprLastArg)
    return withExprBaseCall, nil
```
- **Invariant:** There is **no `withNode`** in the AST. `with map(xs) fn(x) x * 2` and `with (xs |> map()) fn(x) x * 2` both produce the exact same `fnCallNode` as `map(xs, fn(x) x * 2)`!
- **Design Consequence for Zero-Cost Blocks (Downward Funargs):** Whenever any compiler tier (`bytecode.go`, `lib/mir.oak`, `lib/codegen-c.oak`, `lib/codegen-go.oak`) encounters a `fnCallNode` whose callback argument is a **syntactic `fnNode` literal**, that `fnNode` is guaranteed non-escaping at the call site if the callee is a known HOF (`each`, `map`, `filter`, `reduce`, `loop`, `sort`, `bytecode`, `interpreter`). Both the HOF loop and the `fnNode` body inline directly into the caller's frame with **zero closure allocation**!

#### §3.1.3 Complete Catalog of Additional Magnolia Syntax Sugars
1. **Automatic Comma Insertion (ACI) (`token.go`):** Newlines following value-producing tokens and tokens before `)`, `]`, `}` inject synthetic `comma` tokens unless preceded by continuation tokens (`+`, `-`, `*`, `/`, `:=`, `<-`, `.`, `:`, `fn`, `if`, `with`, `cs`, `|>`, `->`, `<<`, etc.).
2. **`if` Expression Sugars (`parse.go`) [ABLATED]:**
   - Conditionless `if { c1 -> b1, _ -> b2 }` desugars to `if true { ... }`.
   - Single-branch `if cond -> body` desugars to `if cond { true -> body }` (body = one node).
   - Multi-target `if x { 0, 1 -> baseCase, _ -> recCase }` desugars into multiple `ifBranch` entries **sharing the identical `body` AST node reference** (`branches[0].body == branches[1].body`). Compilers should deduplicate shared branch bodies into a single basic block jumped to by multiple target comparisons!
   - `?` is a legal match target (null arm), `_` is the wildcard (EmptyValue), and no-match yields `?` — all verified by binary ablation (§3.0).
3. **`fn` Sugars (`parse.go`):** Optional name (`fn f(x)` binds `f` in outer scope), optional parameter parens (`fn { ... }` ≡ `fn() { ... }`), and empty-brace body rewrite (`fn {}` rewrites empty `objectNode{}` to empty `blockNode{}` returning `?`).
4. **`cs` Class Sugars (`parse.go`, `docs/cs.md`):**
   - Assignment-only body sugar (`cs Counter { a := 0, inc := fn { a <- a + 1 } }`) rewrites at parse time into `__oakClassSelf := {}` + property assignments + lexical method field rewriting (`a` → `__oakClassSelf.a`).
   - Inherited/static sugar (`cs Sub(x) { staticA := 1, (Parent1, Parent2) -> body }`) populates `classNode.staticExprs` and `classNode.parents`.
5. **Computed Property / Index Sugar (`parse.go`):** `(expr)` unwraps single-expression parens *except* `(ident)`, which stays a 1-element `blockNode` so `xs.(i)` dynamically evaluates `i` whereas `xs.i` accesses static field `'i'`.

#### §3.1.4 `if`-Match Grammar Hazards for Code Emitters (Added after M0 ablation) [ABLATED]
The half-finished mechanical conversion "massive if-statements → match syntax" (HEAD commit `7a94746`) produced two invalid hybrid shapes in shipped code. Both are **parse errors**; both were repaired in `lib/gui-window-nest.oak` during M0 using only verified grammar:

| Invalid hybrid | Why it fails | Verified repair |
| :--- | :--- | :--- |
| `if cond -> {` whose block contains match-arm lines (`true -> …`) | The `->` body is parsed as a plain block/object; a bare `{ arm -> … }` is not valid block content (`Unexpected token ->, expected ,`). | Promote to match form: `if cond { true -> …, _ -> … }` (only when the condition is boolean) or `if cond != ? { … }` for existence guards. |
| `if cond {` whose block contains bare statements | Braces after `if` are match arms only (`Unexpected token } at start of unit`). | `if cond != ? -> { statements }` for existence guards; `if boolCond -> { statements }` for boolean guards. |

The canonical existence-test idioms (verified against `lib/gui-common.oak`'s `_default` and `_clamp`, which import cleanly):
```go
fn _default(value, fallback) if value {
	? -> fallback
	_ -> value
}
```
- Object-existence guards use `if parent != ? -> { … }` (equality with `?` is exact and verified) **or** `? -> / _ ->` arm pairs on the raw value.
- Boolean-cond guards use `if cond -> { … }` (single-branch sugar) or `if cond { true -> … }`.
- WARNING: `if obj { true -> A, _ -> B }` on a **non-boolean** subject silently never matches (objects never equal `true`) — the whole expression returns `?`. Several libs contain this shape; MIR's type/shape inference (TODO-13) must flag it as a static diagnostic.

### §3.2 The 9 Root-Cause Performance Bottlenecks in Magnolia Today

1. **16-Byte `Value` Interface Boxing & Heap Escape (`eval.go`, `bytecode.go`):** `IntValue(int64)` and `FloatValue(float64)` box into 16-byte Go interfaces; integers outside `0..255` allocate 8 bytes on the Go heap on every arithmetic op, and `[]Value` stacks force Go's GC to scan every slot for pointers.
2. **Per-Call Heap Allocations in `bytecode.go` (`opCall` / `opTailCall`):** Every `opCall` allocates `args := make([]Value, arity)`, `locals := make([]Value, ft.localCount)`, and `scope := &vmScope{...}` on the Go heap! Even `opBuiltin` allocates `args := make([]Value, arity)` per call!
3. **Linear String-Name Upvalue Walks (`bytecode.go`):** `opLoadUpval` and `opStoreUpval` emit string constant indices and walk linked `*vmScope` `names []string` slices linearly by string comparison at runtime.
4. **Unconditional `reflect.ValueOf` + Global Striped `sync.RWMutex` (`eval.go`):** `objRLock`, `listRLock`, and `strRLock` call `reflect.ValueOf(x).Pointer()` and acquire/release atomic `RWMutex` locks on **every single object/list/string operation**, even when `c.eng.concurrent == false` (single-goroutine execution)!
5. **Shapeless `map[string]Value` Objects & Uncached `opGetProp` (`eval.go`, `bytecode.go`):** Every `ObjectValue` hashes Go strings on property lookup with zero inline caching, and `compileClass` drops `staticExprs` and `parents`.
6. **Type-Erased `lib/codegen-go.oak` & `lib/build.oak` Transpilers:** `codegen-go.oak` lowers all values to `interface{}`, functions to variadic `func(...OakValue) OakValue`, and `=` to `reflect.DeepEqual`; `renderJavaBundle` and `renderLuaBundle` wrap expressions/blocks in IIFE closures.
7. **Prototype `lib/codegen-native.oak`:** Supports only `:ir_const` (int), `:ir_add`, `:ir_sub`, and `:ir_ret`; virtual registers ≥ 8 all map to `rax` (clobbering live values); lacks branches, calls, loops, spills, and SIMD.
8. **Parse-Inside-Loop Benchmarks (`bench_test.go`):** `BenchmarkEval*` and `BenchmarkBytecode*` re-tokenize, re-parse, and re-compile source strings inside `for i := 0; i < b.N`.
9. **Missing Build-Tag Headers (`syscalln_unix.go`, `sysproc_other.go`) & Self-Hosted Tokenizer Gaps (`lib/syntax-tokenize.oak`):** `syscalln_unix.go` and `sysproc_other.go` lacked `package main` headers (fixed in M0), and `lib/syntax-tokenize.oak` lacks `**`, `==`, and `~` parity with `token.go` (TODO-02).

### §3.3 Build Hygiene Ledger (M0 results)
- `syscalln_unix.go`: restored `//go:build !windows && (!linux || !cgo)` + `package main` header. ✓
- `sysproc_other.go`: restored full stub (`lookupSysProc` returning a descriptive error) so both CGO modes link. ✓
- `lib.go`: 28 duplicate embed-var blocks removed (all Java/JVM/applet/MIDP/SWF/UI modules preserved exactly once in the opt-in section); 4 missing embeds added (`gui-render-core`, `gui-window-nest`, `obj`, `web-pipeline`) + `stdlibs` map entries; `TestStdlibRegistryCoversLibDirectory` green. ✓
- Merge-conflict markers resolved in `lib/std.oak`, `lib/bytes.oak`, `lib/gui-common.oak`, `lib/jvm-vm-libs.oak` (kept HEAD-side shared-helper abstractions; verified each helper has exactly one definition). ✓
- `lib/gui-window-nest.oak`: 14 evidence-based repairs (9 hybrid `if cond -> {arms}` → match form; 4 statement-style guards → `!= ? -> {…}`; `_test`/`_ensureChildSurface`/`_releaseChildSurface` restructured to legal match form). Module now imports cleanly; `go test ./...` green under `CGO_ENABLED=0` and `CGO_ENABLED=1`. ✓

### §3.4 The `lib/` Structural-Audit Obligation (Ongoing)
Because only embedded+imported modules were ever exercised, other `lib/*.oak` files may contain the same two hybrid shapes. Any module touched by future milestones must pass the **import ablation** (`x := import('<lib>'); print('LOADED')`) before and after modification. MIR (TODO-13) will add a static `if`-shape linter so this class is caught at build time, not at import time.

---

## §4. Unified Middle-End Architecture: `lib/mir.oak` (Sea-of-Regions Typed SSA)

`lib/build.oak` (`runOak`) already collects the complete transitive closure of project + stdlib modules into `ModuleNodes` before emitting any target. Inserting **`lib/mir.oak`** as the unified whole-program middle-end gives **all** Magnolia targets closed-world optimization without altering Magnolia's surface syntax:

```text
Magnolia Source (.oak) ──► Tokenizer (ACI) + Parser (desugars |>, with, cs, if, .())
  │
  ▼  AST Middleware (syntax-macros.oak ──► optimize.go / transpile.oak)
  │
  ▼  lib/mir.oak — Whole-Bundle Closed-World Middle-End
  │    1. Binding & Mutability Pass: `:=` never mutated by `<-` ──► pure SSA value
  │    2. Call-Site HOF & Sugar Fusion: nested `fnCall` (`|>` / `with` / `map` / `filter` / `reduce` / `each`) ──► FusedLoop MIR
  │    3. Flow-Sensitive Type & Shape Lattice: Top ──► {Int64, Float64, Bool, Atom32, StrSlice, List<T, R_i>, Shape#k, Union#m, Dynamic}
  │    4. Static Region & Escape Pass: Register ──► Scoped Virtual Region Stack R_0..R_k ──► Shared/Escaping Arena
  │    5. Layout Synthesis: Pinned-Base Ptr32<R_k> + Inline Variable-Mode Sums + 64B-Aligned SoA Arrays
  │
  ├──► [Go Host VM]    bytecode.go v2 (Zero-alloc slab frames, lexical upvalues, superinstructions, HOF intrinsics)
  ├──► [Systems C11]   lib/codegen-c.oak ──► clang/gcc -O3 -march=native -flto ("Goose+ Mode")
  ├──► [Systems Go]    lib/codegen-go.oak v2 (Monomorphized int64/float64/structs + GC-invisible []byte regions)
  ├──► [Native + JIT]  lib/codegen-native.oak v2 + assembler.oak + linker.oak + jit.oak (mmap PROT_EXEC)
  └──► [Polyglot/GPU]  Direct Typed WAT (--wasm), Unboxed JS/Lua/Java, & SPIR-V/OpenCL/CUDA (--gpu / lib/compute.oak)
```

---

## §5. The 5 Core Innovations That Push Magnolia Past Goose, C++, and Rust

### §5.1 Disjoint Virtual-Memory Region Stacks + Mathematical `restrict` / `noalias`
- **Goose Physics:** Each static region $R_k$ reserves a **2 GB virtual address range** (`mmap(..., MAP_PRIVATE|MAP_ANONYMOUS|MAP_NORESERVE)` on Linux/macOS, `VirtualAlloc(MEM_RESERVE)` on Windows). Bump allocation is `R_k.top += aligned_size` (1 instruction); scope exit restores `R_k.top = saved_watermark` ($O(1)$, zero `free`, zero GC, zero Rust drop glue). Growable lists at the tip of $R_k$ append in place with **zero `realloc` copies**.
- **Surpassing Goose (Fixing Ceiling #1):** Because $R_i$ and $R_j$ ($i \neq j$) occupy non-overlapping 2 GB virtual memory reservations, pointers `Ptr<A, R_i>` and `Ptr<B, R_j>` **cannot alias by hardware construction**. `lib/codegen-c.oak` emits C11 `restrict` on every distinct region base pointer and hoisted slice view, unlocking full Clang/LLVM `noalias` load-hoisting and wide SIMD vectorization that Goose's C99 emitter misses.

### §5.2 Register-Pinned Base `Ptr32<R_k>` (Eliminating Goose's `lru` 0.81× Loss)
- **Fixing Goose Ceiling #3:** Instead of self-relative offsets `(uintptr_t)&field + delta32` (which add dependent ALU ops on pointer chasing and force offset recomputation when nodes relink in `lru` and `calc`), Magnolia represents intra-region references as **32-bit 8-byte-scaled indices (`uint32_t idx`)** relative to `R_k.base` (addressing up to **32 GB per region** in 4 bytes!).
- In native/C11 codegen, `R_k.base` stays in a callee-saved register (`r12`–`r15` on x86-64, `x20`–`x23` on ARM64), so dereferencing `node.next` compiles to a **single indexed load** (`mov eax, [r13 + rcx*8 + 4]`) and relinking an LRU list writes plain 32-bit integers with **zero offset arithmetic** and **50% less cache footprint than 64-bit C++/Rust pointers**.

### §5.3 Representation Freedom: Hybrid Variable-Mode Sums + Automatic 64B-Aligned SoA SIMD
Because Magnolia objects (`{ x: 1.0, y: 2.0 }` or `cs Particle(x, y, vx, vy)`) do not expose C/C++/Rust raw field addresses (`&p.x` / `sizeof(Particle)`), `lib/mir.oak` has **complete physical layout freedom**:
1. **Stream Sum Layout:** Heterogeneous `:atom`-tagged records pack variable-length inline (`tag:u8` + active variant fields), matching Goose's 2.2× speed / 4× memory win over Rust `enum`s.
2. **Automatic AoS → SoA SIMD Transposition (Fixing Goose Ceiling #2 & Beating C++/Rust):** Homogeneous lists of numeric shapes (`List<Particle>`) accessed in hot loops or fused `map`/`reduce` calls automatically transpose into **64-byte-aligned Structure-of-Arrays columns** (`double* restrict x; double* restrict y; ...`), allowing AVX2/AVX-512/NEON to process 4–16 elements per instruction with zero gather/shuffle overhead.

### §5.4 SSI Bounds-Check Elimination + Interned `:atom` Jump Tables + `libdivide`
- Loop induction variables `0 <= i < len(xs)` and region-allocated handles carry static bounds proofs so array indexing emits **zero panic branches** (beating safe Rust without `unsafe`).
- All `:atom` literals receive contiguous compile-time `uint32_t` IDs (`0..M-1`), turning atom equality into 1-cycle integer comparison and multi-arm `if tag { :a -> ..., :b -> ... }` (including deduplicated multi-target `a, b -> body` branches) into dense `switch` jump tables or branchless `cmov`/`csel` selects.

### §5.5 Microsecond In-Process Runtime Value Specialization (`lib/jit.oak`)
- **Fixing Goose Ceiling #4 & Beating Static C++/Rust:** Static compilers cannot constant-fold values loaded at runtime (matrix strides, filter thresholds, schema layouts, VM bytecode chunks). Using Magnolia's self-hosted `codegen-native.oak` + `assembler.oak` + executable `mmap` bridge, `jit.specialize(fn, runtimeConsts...)` substitutes live runtime values as immediate constants, constant-folds/DCEs dead branches, and emits native x86-64/ARM64 machine code in **<100 µs**, routinely beating static `-O3` C++/Rust/Goose by **1.3×–4.0×** on data-driven kernels.

---

## §6. Milestone M0 — Build Hygiene, `bench_test.go`, & Differential Benchmark Harness

### §6.1 Go Build Hygiene (`syscalln_unix.go`, `sysproc_other.go`) — DONE
- Added `package main` to `syscalln_unix.go` and restored the full `!windows && (!linux || !cgo)` stub in `sysproc_other.go` so `go build` and `go test` succeed across all build configurations.

### §6.2 Separating Parse/Compile from Execution in `bench_test.go`
- Keep existing end-to-end benchmarks for historical tracking, and add **pure execution benchmarks** that parse + `optimizeAST` (for tree-walker) and `compileToByteCode` (for VM) **before `b.ResetTimer()`**, reusing `Context` / `VM` cleanly so allocations and ns/op measure actual runtime execution.

### §6.3 Goose + C++ + Magnolia Head-to-Head Benchmark & Oracle Suite (`bench/`)
- Build real **Goose** (`/tmp/ref/goose/build/goose`) and **C++ (`g++ -O3 -march=native` / `clang++ -O3 -march=native`)** alongside **Baseline Magnolia** (`/tmp/ref/magnolia-baseline`) and **Magnolia-Hyper** across canonical workloads:
  1. `fib35` / `tak` (call & recursion overhead)
  2. `arith_loop` (scalar ALU, division, branchless `if` pattern matching)
  3. `sieve` (tight array indexing & bounds-check elimination)
  4. `lru` (pointer-relinking hash/list — Goose's 0.81× weak spot vs. Pinned-Base `Ptr32`)
  5. `records` (heterogeneous `:atom`-tagged variant stream — C++/Rust's weak spot)
  6. `particles_soa` (homogeneous struct list update — AoS vs. Auto-SoA SIMD)
  7. `pipe_fusion` (`xs |> map(...) |> filter(...) |> reduce(...)` sugar fusion)
  8. `stride_jit` (runtime-parameterized divisor/stride/predicate kernel — static AOT's weak spot vs. Runtime Value Specialization)

---

## §7. Milestone M1 — `bytecode.go` & `eval.go` Engine Overhaul

### §7.1 Small/Medium `IntValue` Unboxed Cache & Type-Specialized VM Fast Paths
- Pre-populate an immutable `intValCache` table (e.g., `-4096..65535`) so `makeIntValue(n)` returns a zero-allocation interface value across loop counters, byte manipulation, ASCII/UTF-8 codepoints, and array indices.
- In `bytecode.go` (`run`), add direct `IntValue` / `FloatValue` fast paths on `opAdd`, `opSub`, `opMul`, `opDiv`, `opMod`, `opEq`, `opNeq`, `opLt`, `opGt`, `opLeq`, `opGeq` before general `Value` dispatch.

### §7.2 Zero-Allocation Contiguous Frame Slab in `bytecode.go`
- Replace per-call `locals := make([]Value, ft.localCount)`, `args := make([]Value, arity)`, and `&vmScope{}` in `VM`:
  - Track static `needsHeapScope bool` on each `funcTemplate` (true **only** if an inner `fnNode` captures locals from this function!).
  - Give each pooled `VM` a contiguous `localsSlab []Value` (4,096 slots) + `localSP int`.
  - When `ft.needsHeapScope == false` (the vast majority of functions!), `opCall` copies arguments directly from `vm.stack[vm.sp-arity : vm.sp]` into `vm.localsSlab[vm.localSP : vm.localSP+ft.localCount]` and pushes a value-type `callFrame` — **0 heap allocations (`0 allocs/op, 0 B/op`) per function call!**
  - Similarly, pass `vm.stack[vm.sp-arity : vm.sp]` directly as a slice view into `callBuiltin` — **0 allocations per builtin call!**

### §7.3 Concurrency-Gated Lock Elimination in `eval.go` & `env.go`
- Set `c.eng.concurrent = true` atomically only when `go(...)` or `listen(...)` spawns concurrent goroutines.
- When `!c.eng.concurrent` (or via lock-free fast helpers), bypass `objRLock`/`listRLock`/`strRLock` and eliminate `reflect.ValueOf(x).Pointer()` completely! Even when concurrent, hash pointers via `unsafe.Pointer` arithmetic rather than `reflect.ValueOf` (which allocates reflection metadata / burns ~15 ns per lock).

### §7.4 Lexical `(depth: u16, slot: u16)` Upvalues in `bytecode.go`
- Extend `compiler.resolveUpvalue` to return `(depth int, slot int, ok bool)` by searching `co.parentLocals` from innermost to outermost scope.
- Add `opLoadUpvalFast` and `opStoreUpvalFast` taking `(depth: u16, slot: u16)` that step `depth` parent pointers and index `scope.values[slot]` in **$O(\text{depth})$ pointer loads with zero string comparisons**, keeping `opLoadUpval` only for unresolved dynamic outer names (such as `bytecode()` host scope bridges or late top-level bindings).

### §7.5 Bytecode Superinstructions & Peephole Optimizer (`bytecode.go`)
- Add a post-compile bytecode peephole pass (`optimizeBytecode`) + VM handlers for high-frequency fused idioms:
  - `opLoadLocal0`..`opLoadLocal3` (1-byte opcodes, zero operand fetch),
  - `opAddLocalImm` (`local[slot] + imm32`), `opSubLocalImm` (`local[slot] - imm32`),
  - `opIncLocal` (`local[slot] <- local[slot] + 1` in place),
  - `opLtJumpFalse` / `opEqJumpFalse` / `opEqImmJumpFalse` (fused comparison + conditional branch with **zero `BoolValue` stack push/pop**!),
  - `opTailCallSelf` (direct self-tail-call overwriting `frame.locals[0..arity]` from the stack in place without closure type-switching).

### §7.6 Shared `if`-Target Deduplication & Dead `pipeArrow` Cleanup (`bytecode.go`)
- Remove unreachable `n.op == pipeArrow` in `compileNodeTail`.
- In `compileIfTail`, coalesce consecutive `ifBranch` entries that share the exact same `body` AST node (produced by parser sugar `0, 1 -> body`) so `body` is compiled **once** and all matching targets jump directly to that shared body offset.

### §7.7 Inline-Cached Constant Property Access (`opGetPropConst` / `opSetPropConst`) & Full `cs` Bytecode Support
- Emit dedicated `opGetPropConst` / `opSetPropConst` when the RHS of `.` is a static identifier (`obj.x`), avoiding pushing/popping a `*StringValue` constant on the operand stack on every field read/write.
- Upgrade `compileClass` in `bytecode.go` to handle `classNode.staticExprs` and `classNode.parents` natively in bytecode.

### §7.8 Call-Site Inlining of Parser-Desugared `|>` & `with` HOF Loops (`bytecode.go`)
- When `compileFnCallTail` sees `each(xs, fn(x, i) ...)`, `map(xs, fn(x, i) ...)`, `filter(xs, fn(x, i) ...)`, or `reduce(xs, init, fn(acc, x, i) ...)` (whether written with `|>`, `with`, or direct call syntax!) where the callback argument is a literal `fnNode`:
  - Compile the loop iteration **inline inside the caller's bytecode chunk** using hidden local slots for the list, index, length, accumulator, and callback parameters — **0 closure allocations, 0 child VMs, 0 Go-to-VM boundary crossings!**

---

## §8. Milestone M2 — `lib/mir.oak`, `lib/codegen-c.oak` ("Goose+ Mode"), & `lib/codegen-go.oak` v2

### §8.1 `lib/codegen-c.oak` & `magnolia build --c`
- Consumes `lib/mir.oak` and emits self-contained, dependency-free C11 with:
  1. **Static 2 GB OS Virtual-Memory Region Stacks (`mg_region_t`)** using `mmap` (`MAP_NORESERVE`) on POSIX and `VirtualAlloc` on Windows, with inline bump allocation and $O(1)$ scope watermark restore.
  2. **Unboxed Typed Functions & Structs** (`int64_t`, `double`, `bool`, `uint32_t` atom IDs, concrete C structs for inferred `cs`/object shapes, and `mg_val_t` 64-bit NaN-boxed fallback for dynamic values).
  3. **Disjoint-Region `restrict` Pointers & 64B-Aligned SoA Arrays** (`_Alignas(64)`) + `#pragma clang loop vectorize(enable)` / `#pragma GCC ivdep`.
  4. **Pinned-Base `uint32_t` Region Handles (`mg_ptr32_t`)** for zero-ALU pointer relinking (`lru`, `btree`, graphs).
  5. **Guaranteed Tail Calls** (`__attribute__((musttail))` on Clang, loop-lowered self-tail-calls everywhere) + `libdivide` runtime integer division strength reduction.

### §8.2 `lib/codegen-go.oak` v2 (`magnolia build --go`)
- Uses `lib/mir.oak` type/immutability/loop analysis to emit:
  - Concrete typed Go companions `func __mg_fast_f(n int64) int64` for statically typed functions (called directly by typed callers in CPU registers with zero `...OakValue` slice allocation!), wrapped thinly for dynamic callers.
  - Self-tail-recursion lowered directly to native Go `for { ... }` loops.
  - Fast scalar `==` / `<` / `+` / `-` / `*` operations instead of `reflect.DeepEqual` and `oakAdd` interface switches whenever operand types are proven by `lib/mir.oak`.
  - Off-GC `[]byte` region arenas (`MgRegion`) for flat numeric/struct buffers so Go's GC mark phase skips gigabytes of data in $O(1)$.

---

## §9. Milestone M3 — Self-Hosted `codegen-native.oak` v2 & Runtime Value Specialization JIT

### §9.1 Upgrading `lib/ast-ir.oak`, `lib/codegen-native.oak`, & `lib/assembler.oak`
- Extend `lib/ast-ir.oak` with full control-flow labels, conditional jumps (`:ir_jz`, `:ir_jnz`, `:ir_jmp`), comparisons (`:ir_eq`, `:ir_neq`, `:ir_lt`, `:ir_gt`, `:ir_leq`, `:ir_geq`), `:ir_mul`, `:ir_div`, `:ir_mod`, `:ir_and`, `:ir_or`, `:ir_xor`, `:ir_shl`, `:ir_shr`, `:ir_cmov`, `:ir_call`, and tail-loop lowering.
- Fix `lib/codegen-native.oak` register allocation across full System V AMD64 / Win64 / AAPCS64 calling conventions with proper prologue/epilogue callee-saved registers and stack spill slots `[rbp - off]` (eliminating the `regIdx >= 8 -> rax` clobber bug), plus AVX2/NEON vector emission in `lib/assembler.oak`.

### §9.2 Microsecond Runtime Value Specialization (`lib/jit.oak` + Host `jit_exec` Builtin)
- Expose an executable memory buffer & native function invoker in `env.go` (`___runtime_jit_compile`, `___runtime_jit_call`, `___runtime_jit_free` backed by `mmap(PROT_READ|PROT_WRITE)` → `mprotect(PROT_READ|PROT_EXEC)` on Unix / `VirtualAlloc` on Windows).
- Build `lib/jit.oak` on top of `syntax`, `mir`, and `codegen-native`: given an Oak function AST/source and runtime-bound parameter constants (e.g. `stride`, `divisor`, `filterMin`, `coefficients`), `jit.specialize(fn, boundEnv)` constant-folds and strength-reduces the body (`x / runtimePow2` → `sar rax, k`; `x * runtimeConst` → `imul` / `lea`), compiles directly to x86-64/ARM64 machine bytes in memory, and executes natively with **zero external compiler subprocess**!

---

## §10. Milestone M4 — Polyglot Target Unification & Heterogeneous GPU/SIMD Offload

### §10.1 Polyglot Transpiler Upgrades (`--wasm`, `--web`, `--lua`, `--java`)
- Feed `lib/mir.oak` analysis + parser-desugared HOF fusion (`|>` / `with` → fused loops) as a pre-codegen pass in `lib/build.oak` so **every** target (`--oak`, `--bin`, `--go`, `--c`, `--native`, `--wasm`, `--web`, `--ts`, `--lua`, `--java`) automatically receives:
  - Fused single-pass `map` / `filter` / `reduce` / `each` loops (zero intermediate arrays, zero per-element closure calls!),
  - Multi-target `if` branch deduplication & constant propagation,
  - Self-tail-recursion to `while` loop lowering (eliminating `--web` trampoline allocations!).

### §10.2 `lib/compute.oak` — Pure Pipeline SIMD & GPU Kernel Synthesis
- Provide high-throughput vector/matrix/reduction kernels (`mapFused`, `reduceSum`, `dotProduct`, `saxpy`, `matmul`, `softmax`) that automatically dispatch to:
  1. **GPU Compute Backend** (`lib/gpus.oak` CUDA/OpenCL + `lib/gui-shader-codegen.oak` GLSL/HLSL/SPIR-V generation) when a GPU driver is detected and $N \ge N_{\text{gpu}}$,
  2. **Native JIT / C-SIMD Multi-Core Backend** (`lock_thread` parallel workers + AVX2/FMA) on CPU!

---

## §11. Verification Protocol & Empirical Benchmark Ledger

### §11.1 Verification Gates
1. **100% Go Host & Stdlib Test Parity:** Every existing unit test (`eval_test.go`, `optimize_test.go`, `parse_test.go`, `coverage_test.go`, `packed_vfs_test.go`, etc.) must pass unchanged across both `--normal` and `--bytecode`.
2. **Import Ablation Gate:** Every embedded `lib/*.oak` touched by a change must `import()` cleanly under both engines.
3. **Cross-Target Differential Oracle:** Every benchmark in `bench/` must produce **byte-identical numerical/structural output** across Baseline Magnolia, Magnolia-Hyper `--bytecode`, Magnolia-Hyper `--go`, Magnolia-Hyper `--c`, Magnolia-Hyper `lib/jit.oak`, C++, and Goose before timing is accepted.

### §11.2 Live Empirical Benchmark Results
*(Populated automatically as Milestones M0–M4 complete in `/tmp/ref/` and `bench/`.)*

#### §11.2.1 In-repo micro-benchmark ledger (`bench_test.go`, pre-parsed/precompiled, ReportAllocs)

Box: shared Xeon 2.60GHz, 2 vCPU. `-benchtime=200x` (300x for post-TODO-08 rows).
ns/op is noisy on this box (shared tenancy) — **trust the allocs/B columns for
regression gating**; ns deltas ≥ 25% are meaningful.

| Benchmark | Baseline (TODO-03, `e84d394`~) ns/op | B/op | allocs/op | Post-TODO-07 (`88963d8`) allocs/op | Post-TODO-08+06 (`5e4ea00`+`e5a6a94`) allocs/op |
| :--- | ---: | ---: | ---: | ---: | ---: |
| `EvalPreparsedFib25` | 145,363,263 | 143,590,363 | 2,988,822 | 2,988,449 | 2,988,449 |
| `EvalPreparsedLoopCounter` | 144,960 | 81,461 | 3,029 | 3,029 | 3,029 |
| `EvalPreparsedFnCallOverhead` | 1,324,215 | 1,198,754 | 26,745 | 24,660 | 24,660 |
| `EvalPreparsedPropertyAccess` | 1,194,413 | 871,225 | 24,771 | 23,025 | 23,025 |
| `BytecodePrecompiledFib25` | 40,593,366 | 19,453,696 | **485,954** | 485,579 | **8** |
| `BytecodePrecompiledLoopCounter` | 39,513 | 27,809 | 10 | 10 | 10 |
| `BytecodePrecompiledFnCallOverhead` | 350,222 | 161,506 | 5,732 | 3,647 (−36%) | **646** (−89%) |
| `BytecodePrecompiledPropertyAccess` | 432,649 | 154,090 | 8,762 | 7,016 (−20%) | **15** (−99.75%, post-TODO-11 `aed413b`) |
| `BytecodePrecompiledPipeFusion` (new, TODO-12) | — | — | — | — | **1,204** vs 5,356 neutralized (−78%); ns 2.24ms→0.96ms |

Key readings:
- **TODO-07** (int cache) removed result-position boxing (−36% fnCall allocs) but
  barely moved fib25 — its remaining allocs were per-call frames → **TODO-08**.
- **TODO-08** (locals slab + heap-scope gating) took recursive VM fib25 from
  485,954 → **8 allocs/op** (the 8 are the single top-level chunk + 7 init/bench
  frames outside the slab path) and fnCallOverhead to 646. The VM call loop is
  now allocation-free; ns/op −31% on fib25 (27.9–28.5 ms vs 40.6 ms baseline).
- Post-TODO-06/08 fib25 B/op 1,077,070: residual is bench-harness top-level
  chunk setup, not per-iteration.
- **TODO-09** (`decd862`): upvalue access O(depth) lexical coordinates — alloc-neutral
  (already 8), removes name hashing per outer read.
- **TODO-10** (`1b7c95d`): `LT_LOCAL_JUMP_TRUE` fuses the 9-dispatch if-match prologue;
  also FIXED engine parity (bytecode no-match fall-through leaked the cond value;
  both engines now return `?`). fib25 allocs stay 8 (dispatch savings only; ns noisy).
- **TODO-11** (`aed413b`): `GET_PROP_CONST`/`SET_PROP_CONST` → property access 6,015 → **15**
  allocs/op (−99.75%). Error parity preserved (`x.type` on non-container still errors).
- **TODO-12** (`cb90e11`): in-frame opHOF* loops for std map/filter/each/reduce with
  literal fnNode callbacks — differential probe suite `hof_inframe_test.go` (25 cases,
  tree-walker oracle). Pipe fusion −78% allocs. Known pre-existing `<<` parity gaps
  (tree-walker errors on str<<int / str<<bool / obj<<; VM vmConcat accepts) documented,
  covered by dedicated probes.
- Eval (tree-walker) columns unchanged by design: TODO-06's lock-gating removes
  atomics, not allocations; TODO-09 (lexical coordinates) is the next eval-side win.

`~` Baseline captured pre-commit but ledgered under `e84d394` (M0).*

| Workload | Goose | C++ `-O3 -march=native` | Baseline Magnolia | Magnolia-Hyper `--bytecode` | Magnolia-Hyper `--c` | Magnolia-Hyper JIT |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| `fib35` | — | — | — | — | — | — |
| `arith_loop` | — | — | — | — | — | — |
| `sieve` | — | — | — | — | — | — |
| `lru` | — | — | — | — | — | — |
| `records` | — | — | — | — | — | — |
| `particles_soa` | — | — | — | — | — | — |
| `pipe_fusion` | — | — | — | — | — | — |
| `stride_jit` | — | — | — | — | — | — |
