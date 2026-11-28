// sieve: list mutation sweep (idiomatic C++; vector<char>). REPS=50
#include <cstdio>
#include <vector>
#include <algorithm>
int main() {
    const long long n = 100000;
    std::vector<char> xs(n, 1);
    long long count = 0;
    for (int rep = 0; rep < 50; rep++) {
        std::fill(xs.begin(), xs.end(), 1);
        for (long long i = 2; i * i < n; i++)
            if (xs[i])
                for (long long j = i * i; j < n; j += i) xs[j] = 0;
        count = 0;
        for (long long k = 2; k < n; k++) count += xs[k];
    }
    printf("sieve=%lld\n", count);
}
