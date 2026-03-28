package main

import (
	"strings"
	"testing"
)

// helper: parse program and run optimizeAST, return the first top-level node
func optimizedFirstNode(t *testing.T, program string) astNode {
	t.Helper()
	tokenizer := newTokenizer(program, "(test)")
	tokens := tokenizer.tokenize()
	parser := newParser(tokens)
	nodes, err := parser.parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	optimized := optimizeAST(nodes)
	if len(optimized) == 0 {
		t.Fatalf("Expected at least one node")
	}
	return optimized[0]
}

// helper: check that a program produces the expected value after optimization
func expectOptimizedResult(t *testing.T, program string, expected Value) {
	t.Helper()
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	val, err := ctx.Eval(strings.NewReader(program))
	if err != nil {
		t.Errorf("Did not expect error: %s", err.Error())
	}
	if val == nil {
		t.Errorf("Return value should not be nil")
	} else if !val.Eq(expected) {
		t.Errorf("Expected %s, got %s", expected.String(), val.String())
	}
}

// --- Constant folding: unary ---

func TestFoldUnaryNegInt(t *testing.T) {
	node := optimizedFirstNode(t, "-42")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != -42 {
		t.Errorf("Expected -42, got %d", n.payload)
	}
}

func TestFoldUnaryNegFloat(t *testing.T) {
	node := optimizedFirstNode(t, "-3.14")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != -3.14 {
		t.Errorf("Expected -3.14, got %f", n.payload)
	}
}

func TestFoldUnaryNotBool(t *testing.T) {
	node := optimizedFirstNode(t, "!true")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != false {
		t.Errorf("Expected false, got %v", n.payload)
	}
}

func TestFoldUnaryBitwiseNot(t *testing.T) {
	node := optimizedFirstNode(t, "~0")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != -1 {
		t.Errorf("Expected -1, got %d", n.payload)
	}
}

// --- Constant folding: binary int ---

func TestFoldIntAdd(t *testing.T) {
	node := optimizedFirstNode(t, "2 + 3")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 5 {
		t.Errorf("Expected 5, got %d", n.payload)
	}
}

func TestFoldIntSub(t *testing.T) {
	node := optimizedFirstNode(t, "10 - 7")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 3 {
		t.Errorf("Expected 3, got %d", n.payload)
	}
}

func TestFoldIntMul(t *testing.T) {
	node := optimizedFirstNode(t, "6 * 7")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 42 {
		t.Errorf("Expected 42, got %d", n.payload)
	}
}

func TestFoldIntDiv(t *testing.T) {
	node := optimizedFirstNode(t, "10 / 4")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 2.5 {
		t.Errorf("Expected 2.5, got %f", n.payload)
	}
}

func TestFoldIntDivByZeroNotFolded(t *testing.T) {
	node := optimizedFirstNode(t, "10 / 0")
	if _, ok := node.(binaryNode); !ok {
		t.Errorf("Division by zero should not be folded, got %T", node)
	}
}

func TestFoldIntMod(t *testing.T) {
	node := optimizedFirstNode(t, "17 % 5")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 2 {
		t.Errorf("Expected 2, got %d", n.payload)
	}
}

func TestFoldIntPower(t *testing.T) {
	node := optimizedFirstNode(t, "2 ** 8")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 256 {
		t.Errorf("Expected 256, got %d", n.payload)
	}
}

func TestFoldIntBitwiseOps(t *testing.T) {
	// AND
	node := optimizedFirstNode(t, "0xFF & 0x0F")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode for &, got %T", node)
	} else if n.payload != 15 {
		t.Errorf("Expected 15, got %d", n.payload)
	}

	// OR
	node = optimizedFirstNode(t, "0x01 | 0x02")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode for |, got %T", node)
	} else if n.payload != 3 {
		t.Errorf("Expected 3, got %d", n.payload)
	}

	// XOR
	node = optimizedFirstNode(t, "5 ^ 3")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode for ^, got %T", node)
	} else if n.payload != 6 {
		t.Errorf("Expected 6, got %d", n.payload)
	}
}

func TestFoldIntShift(t *testing.T) {
	expectOptimizedResult(t, "4 << 2", IntValue(16))
	expectOptimizedResult(t, "16 >> 2", IntValue(4))
}

func TestFoldIntComparison(t *testing.T) {
	node := optimizedFirstNode(t, "3 > 2")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}

	node = optimizedFirstNode(t, "2 < 1")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != false {
		t.Errorf("Expected false")
	}
}

// --- Constant folding: binary float ---

func TestFoldFloatAdd(t *testing.T) {
	node := optimizedFirstNode(t, "1.5 + 2.5")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 4.0 {
		t.Errorf("Expected 4.0, got %f", n.payload)
	}
}

func TestFoldMixedIntFloat(t *testing.T) {
	node := optimizedFirstNode(t, "2 + 3.5")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 5.5 {
		t.Errorf("Expected 5.5, got %f", n.payload)
	}

	node = optimizedFirstNode(t, "3.5 + 2")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 5.5 {
		t.Errorf("Expected 5.5, got %f", n.payload)
	}
}

// --- Constant folding: string concatenation ---

func TestFoldStringConcat(t *testing.T) {
	node := optimizedFirstNode(t, "'hello ' + 'world'")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", string(n.payload))
	}
}

// --- Constant folding: boolean ---

func TestFoldBoolOr(t *testing.T) {
	node := optimizedFirstNode(t, "false | true")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}
}

func TestFoldBoolAnd(t *testing.T) {
	node := optimizedFirstNode(t, "true & false")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != false {
		t.Errorf("Expected false")
	}
}

// --- Constant folding: equality ---

func TestFoldEquality(t *testing.T) {
	node := optimizedFirstNode(t, "3 = 3")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}

	node = optimizedFirstNode(t, "3 != 4")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}
}

// --- Nested folding ---

func TestFoldNestedExpr(t *testing.T) {
	// (2 + 3) * (4 + 1) should fold to 25
	node := optimizedFirstNode(t, "(2 + 3) * (4 + 1)")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 25 {
		t.Errorf("Expected 25, got %d", n.payload)
	}
}

func TestFoldChainedArithmetic(t *testing.T) {
	expectOptimizedResult(t, "2 + 3 + 4", IntValue(9))
	expectOptimizedResult(t, "10 * 2 - 5", IntValue(15))
}

// --- Folding inside composite structures ---

func TestFoldInsideFnBody(t *testing.T) {
	expectOptimizedResult(t, "fn f() 2 + 3, f()", IntValue(5))
}

func TestFoldInsideIfBranch(t *testing.T) {
	expectOptimizedResult(t, "if true { true -> 2 + 3, _ -> 0 }", IntValue(5))
}

func TestFoldInsideList(t *testing.T) {
	expectOptimizedResult(t, "[1 + 2, 3 + 4].0", IntValue(3))
}

func TestFoldInsideAssignment(t *testing.T) {
	expectOptimizedResult(t, "x := 2 + 3, x", IntValue(5))
}

// --- Non-folded cases ---

func TestNoFoldVariableExpr(t *testing.T) {
	// x + 1 should NOT be folded since x is a variable
	node := optimizedFirstNode(t, "x + 1")
	if _, ok := node.(binaryNode); !ok {
		t.Errorf("Expected binaryNode to be preserved, got %T", node)
	}
}

// --- End-to-end correctness after optimization ---

func TestOptimizedFibonacci(t *testing.T) {
	program := `
fn fib(n) if n {
	0 -> 0
	1 -> 1
	_ -> fib(n - 1) + fib(n - 2)
}
fib(10)
`
	expectOptimizedResult(t, program, IntValue(55))
}

func TestOptimizedStringOps(t *testing.T) {
	expectOptimizedResult(t, "'hello' + ' ' + 'world'", MakeString("hello world"))
}

func TestOptimizedBooleanLogic(t *testing.T) {
	expectOptimizedResult(t, "!false & true", BoolValue(true))
	expectOptimizedResult(t, "!(true & false)", BoolValue(true))
}
