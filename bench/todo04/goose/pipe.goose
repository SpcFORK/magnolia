// pipe: map / filter / fold chain — stream pipeline workload.
import std;

// REPS=5
fn main() {
    var xs: i64[>..] = [];
    for i in 20000 { xs.push(i); }
    var r: i64 = 0;
    for rep in 5 {
        r = xs.map() { it * 2 + 1 }.filter() { it % 3 == 0 }.fold(0) { a, b => a + b };
    }
    print("pipe=", r);
}
