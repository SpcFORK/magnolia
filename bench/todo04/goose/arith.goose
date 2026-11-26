// arith: integer tail-recursive loop.
// REPS=10
recursive fn go(i: i64, acc: i64) -> i64 {
    if i == 0 { return acc; }
    return go(i - 1, acc + ((i * 3) % 1000007));
}
fn main() {
    var acc: i64 = 0;
    for r in 10 { acc += go(1000000, 0); }
    print("arith=", acc / 10);
}
