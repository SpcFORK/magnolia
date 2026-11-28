# todo04 Oak workloads (faithful re-construction)

The original TODO-04 Oak sources were never committed; TODO-20 reconstructed
them from `cpp/` and TODO-20/M2 re-verified them against the interpreter and
the C backend. These are those faithful ports, named after the cpp workload
files (`arith_loop`, `sieve`, `lru`, `records`, `pipe_fusion`, `fib35`,
`computefuse`).

File names match `run_bench.py` WORKLOADS (plus `fib35`, matching
cpp/fib35.cpp). Checksum parity (interpreter = cpp = ledger = `--c`):

| workload    | checksum      |
| ----------- | ------------- |
| fib         | 503283        |
| fib35       | 9227465       |
| arith       | 500000500042  |
| sieve       | 9592          |
| lru         | 14968         |
| records     | 30005         |
| pipe        | 133346667     |
| computefuse | 50            |

Run through any engine:

    magnolia build --entry oak/arith.oak --output /tmp/a.c --c
    gcc -O2 -o /tmp/a /tmp/a.c -lm && /tmp/a         # C backend

`computefuse` imports `lib/compute.oak`; the bundler resolves it from the
repo root. The other six are self-contained.
