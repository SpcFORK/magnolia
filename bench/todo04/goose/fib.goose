// fib: naive double recursion (recursion is opt-in via `recursive fn`).
// REPS=1 (8 distinct fib inputs; goose collapses repeated pure workloads)
recursive fn fib(n: i64) -> i64 {
    if n < 2 { return n; }
    return fib(n - 1) + fib(n - 2);
}
fn main() {
    var acc: i64 = 0;
    for n in 8 { acc += fib(20 + n); }
    print("fib=", acc);
}
