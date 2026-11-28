// pipe: map / filter / fold chain (idiomatic C++; ranges-free classic form). REPS=200
#include <cstdio>
#include <vector>
int main() {
    std::vector<long long> xs;
    for (long long i = 0; i < 20000; i++) xs.push_back(i);
    long long acc = 0;
    for (int rep = 0; rep < 200; rep++) {
        acc = 0;
        for (auto x : xs) {
            long long v = x * 2 + 1;
            if (v % 3 == 0) acc += v;
        }
    }
    printf("pipe=%lld\n", acc);
}
