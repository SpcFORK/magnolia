// computefuse: fused map-square + reduce-sum over 20k ints (REPS=50).
// Mirrors oak/computefuse.oak using lib/compute.oak's fused pipeline.
#include <cstdio>
int main() {
    volatile int vn = 20000;
    const int n = vn;
    long long total = 0;
    for (int rep = 0; rep < 50; rep++) {
        long long sse = 0;
        for (int i = 1; i <= n; i++) sse += (long long)i * i; // fused: one pass
        total += sse % 7;
    }
    printf("computefuse=%lld\n", total);
}
