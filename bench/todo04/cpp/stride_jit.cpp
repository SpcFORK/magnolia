// stride_jit: runtime-parameterized stride/divisor/threshold kernel; params derive
// from loaded data, so -O3 cannot fold them (static AOT's weak spot, plan §5.5). REPS=200
#include <cstdio>
static long long kernel(long long x, long long s, long long d, long long t) {
    return ((x * d + ((x >> s) & 1) * t) & 65535) % 251;
}
int main() {
    const int n = 20000;
    static long long xs[20000];
    for (int i = 0; i < n; i++)
        xs[i] = ((i + 1) * 7919LL) % 1000003LL;  // multiplier fits int32 (see §12 note)
    long long s = (xs[1] % 5) + 2;   // 2..6 — data-derived at runtime
    long long d = (xs[2] % 7) + 3;   // 3..9
    long long t = (xs[3] % 997) + 1;
    long long acc = 0;
    for (int r = 0; r < 50; r++)
        for (int i = 0; i < n; i++)
            acc += kernel(xs[i], s, d, t);
    printf("stride=%lld\n", acc);
}
