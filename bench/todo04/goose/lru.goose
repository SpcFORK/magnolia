// lru: hash map + doubly-linked recency list over a grow-only node arena.
// The key universe is fixed (4096), so the map is a direct-address table of
// node indices — the vector-arena-with-i64-indices shape Goose's own lru
// bench calls the expert C++/Rust version. All relink traffic is plain
// array stores.
struct Node { k: i64, v: i64, p: i64, n: i64 }

// REPS=3
fn main() {
    let OPS = 60000;
    let CAP = 512;
    let KMAX = 4096;
    var table: i64[>..] = [];     // key -> node slot, -1 = empty
    for i in KMAX { table.push(-1); }
    var slots: Node[>..] = [];
    var head: i64 = -1;
    var tail: i64 = -1;
    var size: i64 = 0;
    var hits: i64 = 0;
    var rep: i64 = 0;
    var i: i64 = 0;
    for rep_i in 3 {
        rep = rep_i;
        i = 0;
        size = 0;
        head = -1;
        tail = -1;
        hits = 0;
        for t in KMAX { table[t] = -1; }
        slots.clear();
        while i < OPS {
        var k: i64 = (i * 2481) % KMAX;
        if i % 4 == 0 { k = i % 128; }
        let probe = table[k];
        i++;
        if probe >= 0 {
            let nd = probe;
            // hit: unlink + move to front
            let p = slots[nd].p;
            let nx = slots[nd].n;
            if p >= 0 { slots[p].n = nx; }
            if nx >= 0 { slots[nx].p = p; }
            if head == nd { head = nx; }
            if tail == nd { tail = p; }
            slots[nd].p = -1;
            slots[nd].n = head;
            if head >= 0 { slots[head].p = nd; }
            head = nd;
            if tail < 0 { tail = nd; }
            hits++;
        } else {
            // miss: evict tail at capacity, then insert at front
            if size == CAP {
                let old = tail;
                let op = slots[old].p;
                let onx = slots[old].n;
                if op >= 0 { slots[op].n = onx; }
                if onx >= 0 { slots[onx].p = op; }
                tail = op;
                table[slots[old].k] = -1;
                size--;
            }
            slots.push(Node { k: k, v: i, p: -1, n: head });
            let nd = slots.len - 1;
            if head >= 0 { slots[head].p = nd; }
            head = nd;
            if tail < 0 { tail = nd; }
            table[k] = nd;
            size++;
        }
        }
    }
    print("lru=", hits);
}
