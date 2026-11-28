// sieve: list mutation sweep.
// REPS=5
fn main() {
    let n = 100000;
    var xs: bool[>..] = [];
    for i in n { xs.push(true); }
    var count: i64 = 0;
    for rep in 5 {
        for q in n { xs[q] = true; }
        var i: i64 = 2;
        while i * i < n {
            if xs[i] {
                var j = i * i;
                while j < n { xs[j] = false; j += i; }
            }
            i++;
        }
        count = 0;
        var k: i64 = 2;
        while k < n {
            if xs[k] { count++; }
            k++;
        }
    }
    print("sieve=", count);
}
