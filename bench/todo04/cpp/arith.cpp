// arith: integer tail-recursive loop (idiomatic C++; -O3 tail call). REPS=50
#include <cstdio>
long long go(long long i, long long acc) {
    return i == 0 ? acc : go(i - 1, acc + (i * 3) % 1000007);
}
int main() {
    long long acc = 0;
    for (int r = 0; r < 50; r++) acc += go(1000000, 0);
    printf("arith=%lld\n", acc / 50);
}
