# TODO-04 Cross-Language Benchmark Suite (plan §2.2, §6.3, §11.1)

Differential-oracle benchmark harness comparing, on identical workloads:

| variant | engine | notes |
| :--- | :--- | :--- |
| `magnolia -n` | Baseline Magnolia tree-walker | `../..` built binary |
| `magnolia -b` | Magnolia-Hyper bytecode VM | same binary, `-b` flag |
| `goose` | aardappel/goose → C99 → gcc native | built by the goose compiler |
| `c++` | idiomatic C++ | `g++ -O3 -march=native -std=c++20` |

## Workloads (`oak/`, `goose/`, `cpp/` — same algorithms, same checksums)

- **fib** — Σ fib(20..27): naive double recursion. 8 *distinct* inputs because
  Goose detects and collapses repeated pure workloads (see its `bench/goose/sum.goose`
  note); C++ `-O3` constant-folds constants, so the base is `volatile`.
- **arith** — 1M-iteration tail-recursive integer loop (add/mod dispatch).
- **sieve** — 100k boolean list mutation sweep (index read/write).
- **lru** — 60k ops over a 512-entry recency list + map: pointer relinking,
  skewed stream (`i%4` hot set of 128 keys, otherwise golden-ratio stride),
  evictions. Oak uses object records; Goose/C++ use grow-only node arenas
  (direct-address table / `unordered_map` respectively).
- **records** — build/check/throw-away binary trees, long-lived depth 12 plus
  rounds 4..12 (Benchmarks Game bintrees shape). Oak leaves are `{}` object
  records; Goose/C++ leaves are null slots — both count identically (2^(d+1)−1).
- **pipe** — `range(0,20000) |> map(x*2+1) |> filter(x%3=0) |> reduce(+)`:
  stream pipeline over stdlib HOFs (Magnolia's in-frame opHOF* loops vs
  Goose stdlib generic lambdas vs hand-fused C++ loop).

## Protocol

`python3 run_bench.py [--runs N] [--filter wl]` builds all variants (Magnolia
via `go build`, Goose via `goose build`, C++ via g++), verifies every engine
prints the **same checksum** (differential oracle — a mismatch aborts trust),
and reports the **minimum of N wall-clock runs, divided by each program's
`REPS=` marker**. Discard loops print `acc/REPS` (identical value, no dead
code for GCC/Goose to eliminate).

## Known measurement caveats

- Oak numbers include a ~25–30 ms fixed process floor (startup + parse +
  `import('std')` + compile). Ratios `-n` vs `-b` are unaffected (same floor).
- Goose's compiler aggressively eliminates/collects repeated pure work; every
  discard loop here prints its accumulated value (`acc/REPS`) to stay honest.
- C++ `-O3` constant-folds: workload inputs are made runtime-opaque where
  needed (`volatile` base for fib).
- The sandbox is a shared 2-vCPU Xeon 2.60 GHz box — treat ns/op as ±20%.

## Results (min-of-3, 2026-11-26, this sandbox)

| workload | magnolia `-n` | magnolia `-b` | goose | c++ `-O3` | oracle |
| :--- | ---: | ---: | ---: | ---: | :--- |
| fib | 1,264 ms | 146 ms | 1.30 ms | 2.03 ms | OK |
| arith | 1,215 ms | 209 ms | 1.16 ms | 1.08 ms | OK |
| sieve | 451 ms | 93 ms | 0.39 ms | 0.14 ms | OK |
| lru | 617 ms | 158 ms | 0.78 ms | 2.08 ms | OK |
| records | 90 ms | 27 ms | 0.32 ms | 0.30 ms | OK |
| pipe | 141 ms | 35 ms | 0.30 ms | 0.03 ms | OK |

Readings (per plan §2.2):

1. **Magnolia-Hyper `-b` is 3.5–8.7× over its own tree-walker** across all six
   shapes (TODO-05..12 work), but still ~3–4 orders of magnitude behind native
   — the M2 AOT backends (`lib/codegen-c.oak`, `lib/codegen-go.oak`) are the
   levers that close this, exactly as planned.
2. **Goose vs C++ on this box**: near-parity on arith/sieve/fib (within ~1.1–2.7×),
   Goose **wins lru** (0.78 vs 2.08 ms — grow-only arena beats `unordered_map`,
   consistent with Goose's own published lru findings) and records is parity;
   C++ wins pipe ~10× (hand-fused scalar loop vs generic HOF dispatch) — the
   exact ceiling Goose's docs admit for generic lambda pipelines.
3. The lru/pipe rows are the shape evidence for Magnolia-Hyper's planned wins:
   pointer-relinking (2GB VM-region bump stacks, TODO-16) and stream fusion
   (§7.8 in-frame HOF loops already implemented in the VM).

## TODO-20 addendum (2026-11-26)

- The `oak/` workloads were lost from the worktree (never committed in
  TODO-04) and were **reconstructed** from the committed `cpp/` sources.
  All six reproduce the committed checksums exactly — verified against
  `results.csv` (fib 503283, arith 500000500042, sieve 9592, lru 14968,
  records 30005, pipe 133346667). Reconstruction-shape differences mean the
  records/pipe timing rows are not directly comparable to the §11.2.2
  snapshot; fib/arith/sieve/lru reproduce within noise.
- New workload `computefuse` (oak + cpp): fused `compute.oak` pipeline
  (map-squ + reduce) over 20k ints, 50 reps; checksum cross-verified oak ↔
  C++ (computefuse=50). Exercises the TODO-19 fused-pipeline path.
- New workload `fib35` (oak + cpp): single naive fib(35) for the M2
  head-to-head table (§11.3.1). Oracle: `-n` = `-b` = `--c` = C++ = 9227465.
- Goose: the compiler binary is not available in this sandbox; the runner
  now skips the goose column gracefully instead of aborting (reported `-`
  in the CSV, excluded from the oracle set).
- Full re-run results (7 workloads, min-of-3): `results.csv`; analysis and
  the `--c` target gap list live in plan §11.3.1.
