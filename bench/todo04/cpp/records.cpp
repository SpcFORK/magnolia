// records: build/check/throw-away binary trees (idiomatic C++:
// one allocation per node via make_unique, freed by scope exit). REPS=50
#include <cstdio>
#include <memory>
struct T { std::unique_ptr<T> l, r; };
using P = std::unique_ptr<T>;
static P make(int d) {
    if (d == 0) return nullptr;
    auto t = std::make_unique<T>();
    t->l = make(d - 1);
    t->r = make(d - 1);
    return t;
}
static long long check(const P& t) {
    if (!t) return 1;
    return 1 + check(t->l) + check(t->r);
}
int main() {
    auto long_lived = make(12);
    long long total = check(long_lived);
    for (int rep = 0; rep < 50; rep++) {
        long long acc = 0;
        for (int d = 4; d <= 12; d += 2) {
            auto a = make(d);
            auto b = make(d);
            acc += check(a) + check(b);
        }
    }
    long long one = 0;
    for (int d = 4; d <= 12; d += 2) {
        auto a = make(d);
        auto b = make(d);
        one += check(a) + check(b);
    }
    printf("records=%lld\n", total + one);
}
