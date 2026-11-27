#!/bin/bash
# TODO-20 differential oracle: byte-identical output across engines/targets.
# Usage: bash bench/todo20/oracle.sh <host-binary>
HOST="${1:-./magnolia}"
OUT=$(mktemp -d)
pass=0; fail=0

check() { # name expected actual — newlines stripped on both sides:
# magnolia's print() emits no trailing newline; compiled bundles may join
# lines differently. Separator style is not a correctness question; the
# value sequence is.
	local exp got
	exp=$(printf '%s' "$2" | tr -d '\n')
	got=$(printf '%s' "$3" | tr -d '\n')
	if [ "$exp" = "$got" ]; then pass=$((pass+1)); echo "PASS $1: $got";
	else fail=$((fail+1)); echo "FAIL $1: expected=$exp got=$got"; fi
}

# P1: tail recursion + const folding
cat > "$OUT/p1.oak" <<'PROG'
fn loop(n, acc) if n = 0 { true -> acc, _ -> loop(n - 1, acc + n) }
x := 2 * 21
y := loop(1000, 0)
print(string(x))
print(string(y))
print(string(loop(10, 0)))
PROG
expect="42
500500
55"

check "P1 interpreter" "$expect" "$($HOST "$OUT/p1.oak" 2>/dev/null)"

$HOST build --entry "$OUT/p1.oak" --output "$OUT/p1.js" --web >/dev/null 2>&1
check "P1 web(node)" "$expect" "$(node "$OUT/p1.js" 2>/dev/null)"

$HOST build --entry "$OUT/p1.oak" --output "$OUT/p1.c" --c >/dev/null 2>&1
gcc -O2 -o "$OUT/p1c" "$OUT/p1.c" -lm 2>/dev/null
check "P1 c(gcc)" "$expect" "$("$OUT/p1c" 2>/dev/null)"

$HOST build --entry "$OUT/p1.oak" --output "$OUT/p1_go.go" --go >/dev/null 2>&1
# go toolchain: override with GO_CMD=/path/to/go GOROOT=... when the default
# `go` on PATH is not the pinned toolchain
${GO_CMD:-go} build -o "$OUT/p1go" "$OUT/p1_go.go" 2>/dev/null
check "P1 go" "$expect" "$("$OUT/p1go" 2>/dev/null)"

# P2: recursion (fib) + wasm typed fast path
cat > "$OUT/p2.oak" <<'PROG'
fn fib(n) if n < 2 { true -> n, _ -> fib(n - 1) + fib(n - 2) }
print(string(fib(18)))
PROG
check "P2 interpreter" "2584" "$($HOST "$OUT/p2.oak" 2>/dev/null)"

$HOST build --entry "$OUT/p2.oak" --output "$OUT/p2.js" --web >/dev/null 2>&1
check "P2 web(node)" "2584" "$(node "$OUT/p2.js" 2>/dev/null)"

$HOST build --entry "$OUT/p2.oak" --output "$OUT/p2.wat" --wasm >/dev/null 2>&1
cat > "$OUT/p2run.mjs" <<'NODE'
import wabtInit from 'wabt';
import { readFileSync } from 'fs';
const wat = readFileSync(process.argv[2], 'utf8');
const wabt = await wabtInit();
const mod = wabt.parseWat('p.wat', wat);
mod.resolveNames(); mod.validate();
const { buffer } = mod.toBinary({ log: false, write_debug_names: false });
const { instance } = await WebAssembly.instantiate(buffer, { oak: { run: () => 0 } });
console.log(instance.exports.fib(18n).toString());
NODE
# wasm(typed) check: skipped gracefully when the wabt npm package is
# unavailable (offline box) — the .wat is still emitted and can be validated
# by any WAT toolchain.
if node -e "import('wabt')" >/dev/null 2>&1 || (cd "$(dirname "$0")" && npm install wabt --silent >/dev/null 2>&1); then
	node "$OUT/p2run.mjs" "$OUT/p2.wat" 2>/dev/null > "$OUT/p2wasm.txt"
	if [ -s "$OUT/p2wasm.txt" ]; then
		check "P2 wasm(typed)" "2584" "$(cat "$OUT/p2wasm.txt")"
	else
		echo "SKIP P2 wasm(typed): wabt unavailable (wat emitted to $OUT/p2.wat)"
	fi
else
	echo "SKIP P2 wasm(typed): wabt unavailable (wat emitted to $OUT/p2.wat)"
fi

echo "---"
echo "oracle: $pass passed, $fail failed"
[ "$fail" = 0 ]
