// records: build/check/throw-away binary trees over a reset-per-tree node
// arena (the Benchmarks Game bintrees shape; pool cleared per tree).
// REPS=5
recursive fn make(pool: Node[>..]&, d: i64) -> i64 {
    if d == 0 { return -1; }
    let l = make(pool, d - 1);
    let r = make(pool, d - 1);
    pool.push(Node { l: l, r: r });
    return pool.len - 1;
}
recursive fn check(pool: Node[:], idx: i64) -> i64 {
    if idx < 0 { return 1; }
    return 1 + check(pool, pool[idx].l) + check(pool, pool[idx].r);
}
struct Node { l: i64, r: i64 }

fn main() {
    var long: Node[>..] = [];
    let li = make(long, 12);
    var total = check(long, li);
    for rep in 5 {
        var acc: i64 = 0;
        var d: i64 = 4;
        while d <= 12 {
            var pool: Node[>..] = [];
            let a = make(pool, d);
            let b = make(pool, d);
            acc += check(pool, a) + check(pool, b);
            pool.clear();
            d += 2;
        }
    }
    var one: i64 = 0;
    var d: i64 = 4;
    while d <= 12 {
        var pool: Node[>..] = [];
        let a = make(pool, d);
        let b = make(pool, d);
        one += check(pool, a) + check(pool, b);
        pool.clear();
        d += 2;
    }
    print("records=", total + one);
}
