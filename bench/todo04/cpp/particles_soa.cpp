// particles_soa: homogeneous struct list update, AoS layout (idiomatic C++; -O3). REPS=200
#include <cstdio>
#include <vector>
struct Particle { double x, y, z, vx, vy, vz; };
int main() {
    volatile int vn = 20000;
    const int n = vn;
    std::vector<Particle> ps(n);
    for (int i = 0; i < n; i++) {
        ps[i].x = i * 0.5;
        ps[i].y = (i % 11) * 0.25;
        ps[i].z = (i % 13) * 2.0;
        ps[i].vx = (i % 7) * 1.0;
        ps[i].vy = (i % 5) * 0.5;
        ps[i].vz = (i % 3) * 1.5;
    }
    volatile double vdt = 0.5;
    const double dt = vdt;
    long long acc = 0;
    for (int r = 0; r < 10; r++) {
        for (int i = 0; i < n; i++) {
            ps[i].x += ps[i].vx * dt;
            ps[i].y += ps[i].vy * dt;
            ps[i].z += ps[i].vz * dt;
        }
        if (r == 0 || r == 9) {
            long long s = 0;
            for (int i = 0; i < n; i++)
                s += (long long)ps[i].x + (long long)ps[i].y + (long long)ps[i].z;
            acc += s;
        }
    }
    printf("particles=%lld\n", acc);
}
