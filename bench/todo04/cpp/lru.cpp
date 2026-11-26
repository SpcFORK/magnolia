// lru: hash map + doubly-linked recency list over a grow-only node arena
// (the shape Goose's lru bench calls the idiomatic C++ expert version). REPS=20
#include <cstdio>
#include <unordered_map>
#include <vector>
struct Node { long long k, v, p, n; };
int main() {
    const long long OPS = 60000, CAP = 512;
    std::unordered_map<long long, long long> cache;
    std::vector<Node> slots;
    slots.reserve(CAP);
    long long head = -1, tail = -1, size = 0, hits = 0;
    auto detach = [&](long long nd) {
        long long p = slots[nd].p, nx = slots[nd].n;
        if (p >= 0) slots[p].n = nx;
        if (nx >= 0) slots[nx].p = p;
        if (head == nd) head = nx;
        if (tail == nd) tail = p;
    };
    for (int rep = 0; rep < 20; rep++) {
        cache.clear();
        slots.clear();
        head = tail = size = hits = 0; head--; tail--;
        for (long long i = 0; i < OPS; i++) {
            long long k = (i * 2481LL) % 4096;
            if (i % 4 == 0) k = i % 128;
            auto it = cache.find(k);
            if (it != cache.end()) {
                long long nd = it->second;
                detach(nd);
                slots[nd].p = -1;
                slots[nd].n = head;
                if (head >= 0) slots[head].p = nd;
                head = nd;
                if (tail < 0) tail = nd;
                hits++;
            } else {
                if (size == CAP) {
                    long long old = tail;
                    detach(old);
                    cache.erase(slots[old].k);
                    size--;
                }
                slots.push_back(Node{k, i, -1, head});
                long long nd = (long long)slots.size() - 1;
                if (head >= 0) slots[head].p = nd;
                head = nd;
                if (tail < 0) tail = nd;
                cache.emplace(k, nd);
                size++;
            }
        }
    }
    printf("lru=%lld\n", hits);
}
