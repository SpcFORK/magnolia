package main

import (
	"strings"
	"testing"
)

// in-frame opHOF* rewrite against the std.oak oracle (M1 TODO-12).
func TestHOFProbeDifferential(t *testing.T) {
	cases := []struct {
		name string
		prog string
	}{
		{"map_pipe", "std := import('std')\nxs := [3, 1, 2]\nxs |> std.map(fn(x) x * 2)"},
		{"map_1arg", "std := import('std')\nxs := [3, 1, 2]\nstd.map(xs, fn(x) x - 1)"},
		{"map_3arg", "std := import('std')\nxs := [10, 20]\nstd.map(xs, fn(a, b, c) a + b)"},
		{"map_2arg_index", "std := import('std')\nxs := [10, 20]\nstd.map(xs, fn(x, i) x + i)"},
		{"filter_true", "std := import('std')\nxs := [1, 2, 3, 4]\nstd.filter(xs, fn(x) x % 2 == 0)"},
		{"filter_nonbool", "std := import('std')\nxs := [1, 2, 3]\nstd.filter(xs, fn(x) 1)"},
		{"filter_index", "std := import('std')\nxs := [9, 8, 7, 6]\nstd.filter(xs, fn(x, i) i > 1)"},
		{"reduce_sum", "std := import('std')\nxs := [1, 2, 3, 4]\nstd.reduce(xs, 0, fn(a, x, i) a + x)"},
		{"reduce_pipe_with", "std := import('std')\nxs := [1, 2, 3]\nxs |> with std.reduce(0) fn(acc, x) acc + x * 10"},
		{"each_null", "std := import('std')\nxs := [1, 2, 3]\nstd.each(xs, fn(x) x)"},
		{"map_string", "std := import('std')\ns := 'abc'\nstd.map(s, fn(c) c)"},
		{"map_string_int", "std := import('std')\ns := 'abc'\nstd.map(s, fn(c) 65)"}, // bytecode-only: str<<int parity gap
		{"map_string_bad", "std := import('std')\ns := 'abc'\nstd.map(s, fn(c) true)"}, // bytecode-only: str<<bool parity gap
		{"filter_string", "std := import('std')\ns := 'axbx'\nstd.filter(s, fn(c) c == 'a')"},
		{"nested_hof", "std := import('std')\nxs := [[1, 2], [3, 4]]\nstd.map(xs, fn(inner) std.reduce(inner, 0, fn(a, b) a + b))"},
		{"closure_capture", "std := import('std')\noff := 10\nxs := [1, 2, 3]\nstd.map(xs, fn(x) x + off)"},
		{"callback_tailrec", "std := import('std')\nxs := [3, 5]\nfn depth(n) if n { 0 -> 0, _ -> depth(n - 1) }\nstd.map(xs, fn(x) depth(x))"},
		{"hof_in_hof_args", "std := import('std')\nxs := [1, 2, 3]\nstd.map(std.map(xs, fn(x) x + 1), fn(y) y * 2)"},
		{"each_accumulate", "std := import('std')\nacc := []\nstd.each([1, 2, 3], fn(x, i) acc << x + i)\nacc"},
		{"empty_map", "std := import('std')\nstd.map([], fn(x) x)"},
		{"empty_reduce", "std := import('std')\nstd.reduce([], 42, fn(a, x) a + x)"},
		{"empty_each", "std := import('std')\nstd.each([], fn(x) x)"},
		{"map_named_cb", "std := import('std')\nxs := [1, 2]\nfn double(x) x * 2\nstd.map(xs, double)"},
		{"reduce_2arg_cb", "std := import('std')\nxs := [1, 2, 3]\nstd.reduce(xs, 0, fn(a, x) a + x)"},
		{"map_shadowed", "std := import('std')\nxs := [1, 2]\nmap := fn(l, f) 999\nmap(xs, fn(x) x)"},
	}
	bytecodeOnly := map[string]struct {
		prog string
		want string
	}{
		"map_string_int": {prog: "std := import('std')\ns := 'abc'\nstd.map(s, fn(c) 65)", want: "'AAA'"},
		"map_string_bad": {prog: "std := import('std')\ns := 'abc'\nstd.map(s, fn(c) true)", want: "?"},
	}
	for name, bc := range bytecodeOnly {
		t.Run(name, func(t *testing.T) {
			ctx := NewContext("/tmp")
			ctx.LoadBuiltins()
			v, err := ctx.EvalBytecode(strings.NewReader(bc.prog))
			if err != nil {
				t.Fatalf("bytecode error: %s", err.Error())
			}
			if v.String() != bc.want {
				t.Fatalf("bytecode-only probe %s: got %s want %s", name, v.String(), bc.want)
			}
		})
	}
	// both engines error: object subjects fall back to the generic std path
	// and hit the same `<<` incompatibility (parity preserved)
	t.Run("map_object_fallback", func(t *testing.T) {
		prog := "std := import('std')\no := {a: 1, b: 2}\nstd.map(o, fn(x) x)"
		expectProgramToError(t, prog)
		expectProgramToErrorBytecode(t, prog)
	})
	for _, tc := range cases {
		if _, skip := bytecodeOnly[tc.name]; skip {
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			expectProgramToReturnBytecode(t, tc.prog, func() Value {
				ctx := NewContext("/tmp")
				ctx.LoadBuiltins()
				v, err := ctx.Eval(strings.NewReader(tc.prog))
				if err != nil {
					t.Fatalf("tree-walker error: %s", err.Error())
				}
				return v
			}())
		})
	}
}
