#!/usr/bin/env python3
"""TODO-04 cross-language benchmark runner (plan §6.3, §11.1).

Compares, per workload:
  - magnolia -n   (baseline tree-walker engine)
  - magnolia -b   (Magnolia-Hyper bytecode engine)
  - goose         (aardappel/goose compiled to native C via gcc)
  - c++           (g++ -O3 -march=native)

Protocol: build once, run each variant N times, keep the MINIMUM wall time
(min is the stable statistic on shared machines), and verify that every
engine prints the same checksum line (differential oracle). Results are
written as CSV to bench/todo04/results.csv and printed as a table.

Usage: python3 bench/todo04/run_bench.py [--runs 3] [--filter sieve]
"""

import argparse
import os
import re
import subprocess
import sys
import tempfile
import time

HERE = os.path.dirname(os.path.abspath(__file__))
REPO = os.path.dirname(os.path.dirname(HERE))
MAGNOLIA = os.environ.get("MAGNOLIA_BIN", "/tmp/ref/magnolia")
GOOSE = os.environ.get("GOOSE_BIN", "/tmp/ref/goose/build/goose")

WORKLOADS = ["fib", "arith", "sieve", "lru", "records", "pipe", "computefuse"]


def sh(cmd, **kw):
    return subprocess.run(cmd, shell=True, capture_output=True, text=True, **kw)


def build_all():
    # Magnolia binary (both engines are the same binary: -n / -b)
    if not os.path.exists(MAGNOLIA):
        env = dict(os.environ)
        env.setdefault("GOROOT", "/tmp/ref/go_bin/go")
        env.setdefault("GOPROXY", "direct")
        env.setdefault("GOSUMDB", "off")
        env["PATH"] = "/tmp/ref/go_bin/go/bin:" + env["PATH"]
        r = sh(f"go build -o {MAGNOLIA} .", cwd=REPO, env=env)
        if r.returncode != 0:
            sys.exit(f"magnolia build failed:\n{r.stderr}")

    # Goose: compile each program to a native binary (skipped gracefully when
    # the goose compiler is unavailable — column reported as '-')
    goose_ok = os.path.exists(GOOSE)
    if not goose_ok:
        print(f"goose compiler not found at {GOOSE}: skipping goose column")
    for w in (WORKLOADS if goose_ok else []):
        src = os.path.join(HERE, "goose", f"{w}.goose")
        dst = os.path.join(HERE, "goose", w)
        if not os.path.exists(dst):
            r = sh(f"{GOOSE} build {src}")
            if r.returncode != 0:
                sys.exit(f"goose build {w} failed:\n{r.stdout}\n{r.stderr}")
            # goose writes the binary next to the source as `<name>` (no ext)
            built = os.path.join(HERE, "goose", w)
            if not os.path.exists(built):
                # some versions drop it beside the compiler cwd; move it
                cand = os.path.join("/tmp/ref", w)
                if os.path.exists(cand):
                    os.rename(cand, built)
                else:
                    sys.exit(f"goose build {w}: binary not found")

    # C++: g++ -O3 -march=native
    for w in WORKLOADS:
        src = os.path.join(HERE, "cpp", f"{w}.cpp")
        dst = os.path.join(HERE, "cpp", w)
        if not os.path.exists(dst):
            r = sh(f"g++ -O3 -march=native -std=c++20 -o {dst} {src}")
            if r.returncode != 0:
                sys.exit(f"g++ build {w} failed:\n{r.stderr}")


def reps_of(kind, w):
    if kind in ("magnolia-n", "magnolia-b"):
        f = os.path.join(HERE, "oak", f"{w}.oak")
    elif kind == "goose":
        f = os.path.join(HERE, "goose", f"{w}.goose")
    else:
        f = os.path.join(HERE, "cpp", f"{w}.cpp")
    m = re.search(r"REPS=(\d+)", open(f).read())
    return int(m.group(1)) if m else 1


def run_variant(kind, w):
    """Run one variant once; return (checksum_line, per-rep seconds)."""
    if kind in ("magnolia-n", "magnolia-b"):
        flag = "-n" if kind == "magnolia-n" else "-b"
        cmd = [MAGNOLIA, flag, os.path.join(HERE, "oak", f"{w}.oak")]
    elif kind == "goose":
        cmd = [os.path.join(HERE, "goose", w)]
    else:
        cmd = [os.path.join(HERE, "cpp", w)]
    t0 = time.perf_counter()
    r = subprocess.run(cmd, capture_output=True, text=True)
    dt = (time.perf_counter() - t0) / reps_of(kind, w)
    if r.returncode != 0:
        raise RuntimeError(f"{kind}/{w} rc={r.returncode}: {r.stderr[:400]}")
    m = re.search(r"(\w+)=(\-?\d+)", r.stdout)
    if not m:
        raise RuntimeError(f"{kind}/{w}: no checksum in output {r.stdout!r}")
    return m.group(1) + "=" + m.group(2), dt


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--runs", type=int, default=3)
    ap.add_argument("--filter", default="")
    args = ap.parse_args()

    build_all()
    todo04 = sys.argv[0] if False else os.path.join(REPO, "docs", "magnolia-hyper-optimization-plan.md")
    workloads = [w for w in WORKLOADS if args.filter in w]
    kinds = ["magnolia-n", "magnolia-b", "goose", "cpp"]
    if not os.path.exists(GOOSE):
        kinds.remove("goose")

    rows = []
    for w in workloads:
        entry = {"workload": w}
        checksums = set()
        for kind in kinds:
            best, csum = None, None
            for _ in range(args.runs):
                csum, dt = run_variant(kind, w)
                if best is None or dt < best:
                    best = dt
            checksums.add(csum)
            entry[kind] = best
        for kind in ("goose",):
            if kind not in kinds:
                entry[kind] = None
        entry["oracle_ok"] = len(checksums) == 1
        entry["checksum"] = sorted(checksums)[0] if len(checksums) else "?"
        rows.append(entry)

    # report
    print()
    hdr = f"{'workload':<10} {'checksum':<16} {'magnolia -n':>12} {'magnolia -b':>12} {'goose':>12} {'c++ -O3':>12}  oracle"
    print(hdr)
    print("-" * len(hdr))
    print("(values are per-rep wall time; each engine's REPS is normalized out)")
    csv_lines = ["workload,checksum,magnolia_n,magnolia_b,goose,cpp,oracle_ok"]
    for e in rows:
        goose_s = f"{e['goose']*1e6:>11.0f}us" if e['goose'] is not None else f"{'skip':>12}"
        goose_v = f"{e['goose']:.4f}" if e['goose'] is not None else "-"
        print(f"{e['workload']:<10} {e['checksum']:<16} "
              f"{e['magnolia-n']*1e6:>11.0f}us {e['magnolia-b']*1e6:>11.0f}us "
              f"{goose_s} {e['cpp']*1e6:>11.0f}us  "
              f"{'OK' if e['oracle_ok'] else 'MISMATCH'}")
        csv_lines.append(f"{e['workload']},{e['checksum']},{e['magnolia-n']:.4f},"
                         f"{e['magnolia-b']:.4f},{goose_v},{e['cpp']:.4f},{e['oracle_ok']}")

    out = os.path.join(HERE, "results.csv")
    with open(out, "w") as f:
        f.write("\n".join(csv_lines) + "\n")
    print(f"\nresults written to {out}")


if __name__ == "__main__":
    main()
