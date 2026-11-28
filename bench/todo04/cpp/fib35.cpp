// fib35: single naive fib(35) (idiomatic C++; -O3 does NOT fold: n is volatile-derived). REPS=1
#include <cstdio>
long long fib(long long n) { return n < 2 ? n : fib(n - 1) + fib(n - 2); }
int main() {
    volatile long long vn = 35;
    printf("fib35=%lld\n", fib(vn));
}
