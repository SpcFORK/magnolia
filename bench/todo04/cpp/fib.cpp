// fib: naive double recursion (idiomatic C++; -O3). REPS=1
#include <cstdio>
long long fib(long long n) { return n < 2 ? n : fib(n - 1) + fib(n - 2); }
int main() {
    volatile int vbase = 20;
    int base = vbase;
    long long acc = 0;
    for (int n = 0; n < 8; n++) acc += fib(base + n);
    printf("fib=%lld\n", acc);
}
