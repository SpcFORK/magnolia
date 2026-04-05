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

// --- Dead branch elimination ---

func TestFoldIfConstantConditionMatchesBranch(t *testing.T) {
	// if 1 { 1 -> 42, _ -> 0 } should fold to 42
	node := optimizedFirstNode(t, "if 1 { 1 -> 42, _ -> 0 }")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 42 {
		t.Errorf("Expected 42, got %d", n.payload)
	}
}

func TestFoldIfConstantConditionMatchesWildcard(t *testing.T) {
	// if 99 { 1 -> 10, _ -> 20 } should fold to 20 (wildcard)
	node := optimizedFirstNode(t, "if 99 { 1 -> 10, _ -> 20 }")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 20 {
		t.Errorf("Expected 20, got %d", n.payload)
	}
}

func TestFoldIfTrueCondition(t *testing.T) {
	expectOptimizedResult(t, "if true { true -> 'yes', _ -> 'no' }", MakeString("yes"))
}

func TestFoldIfFalseCondition(t *testing.T) {
	expectOptimizedResult(t, "if false { true -> 'yes', false -> 'no' }", MakeString("no"))
}

func TestFoldIfAtomCondition(t *testing.T) {
	node := optimizedFirstNode(t, "if :ok { :ok -> 1, :err -> 2 }")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 1 {
		t.Errorf("Expected 1, got %d", n.payload)
	}
}

func TestFoldIfStringCondition(t *testing.T) {
	node := optimizedFirstNode(t, "if 'a' { 'b' -> 1, 'a' -> 2, _ -> 3 }")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 2 {
		t.Errorf("Expected 2, got %d", n.payload)
	}
}

func TestFoldIfBranchPruningAfterWildcard(t *testing.T) {
	// Branches after _ are dead; the wildcard branch should be kept
	program := "if x { 1 -> 10, _ -> 20, 3 -> 30 }"
	node := optimizedFirstNode(t, program)
	ifNode, ok := node.(ifExprNode)
	if !ok {
		t.Fatalf("Expected ifExprNode, got %T", node)
	}
	if len(ifNode.branches) != 2 {
		t.Errorf("Expected 2 branches after pruning, got %d", len(ifNode.branches))
	}
}

// --- Identity / absorption algebraic simplifications ---

func TestSimplifyAddZero(t *testing.T) {
	expectOptimizedResult(t, "x := 5, x + 0", IntValue(5))
	expectOptimizedResult(t, "x := 5, 0 + x", IntValue(5))
}

func TestSimplifySubZero(t *testing.T) {
	expectOptimizedResult(t, "x := 7, x - 0", IntValue(7))
}

func TestSimplifyMulOne(t *testing.T) {
	expectOptimizedResult(t, "x := 9, x * 1", IntValue(9))
	expectOptimizedResult(t, "x := 9, 1 * x", IntValue(9))
}

func TestSimplifyMulZero(t *testing.T) {
	// x * 0 → 0 when x is pure (identifier)
	expectOptimizedResult(t, "x := 42, x * 0", IntValue(0))
	expectOptimizedResult(t, "x := 42, 0 * x", IntValue(0))
}

func TestSimplifyDivOne(t *testing.T) {
	expectOptimizedResult(t, "x := 15, x / 1", IntValue(15))
}

func TestSimplifyPowerOne(t *testing.T) {
	expectOptimizedResult(t, "x := 4, x ** 1", IntValue(4))
}

func TestSimplifyPowerZero(t *testing.T) {
	// x ** 0 → 1 when x is pure
	node := optimizedFirstNode(t, "x ** 0")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 1 {
		t.Errorf("Expected 1, got %d", n.payload)
	}
}

func TestSimplifyBitwiseOrZero(t *testing.T) {
	expectOptimizedResult(t, "x := 0xFF, x | 0", IntValue(0xFF))
	expectOptimizedResult(t, "x := 0xFF, 0 | x", IntValue(0xFF))
}

func TestSimplifyBitwiseAndZero(t *testing.T) {
	expectOptimizedResult(t, "x := 0xFF, x & 0", IntValue(0))
}

func TestSimplifyXorZero(t *testing.T) {
	expectOptimizedResult(t, "x := 7, x ^ 0", IntValue(7))
	expectOptimizedResult(t, "x := 7, 0 ^ x", IntValue(7))
}

func TestSimplifyShiftZero(t *testing.T) {
	expectOptimizedResult(t, "x := 8, x << 0", IntValue(8))
	expectOptimizedResult(t, "x := 8, x >> 0", IntValue(8))
}

// --- Strength reduction ---

func TestStrengthReducePowerTwo(t *testing.T) {
	// x ** 2 → x * x for simple identifiers
	expectOptimizedResult(t, "x := 7, x ** 2", IntValue(49))
}

// --- Boolean simplification ---

func TestSimplifyBoolOrTrue(t *testing.T) {
	expectOptimizedResult(t, "x := false, true | x", BoolValue(true))
}

func TestSimplifyBoolOrFalse(t *testing.T) {
	expectOptimizedResult(t, "x := true, false | x", BoolValue(true))
	expectOptimizedResult(t, "x := true, x | false", BoolValue(true))
}

func TestSimplifyBoolAndFalse(t *testing.T) {
	expectOptimizedResult(t, "x := true, false & x", BoolValue(false))
}

func TestSimplifyBoolAndTrue(t *testing.T) {
	expectOptimizedResult(t, "x := true, true & x", BoolValue(true))
	expectOptimizedResult(t, "x := true, x & true", BoolValue(true))
}

// --- Double negation elimination ---

func TestSimplifyDoubleNegation(t *testing.T) {
	expectOptimizedResult(t, "x := true, !!x", BoolValue(true))
	expectOptimizedResult(t, "x := false, !!x", BoolValue(false))
}

func TestSimplifyDoubleNumericNegation(t *testing.T) {
	node := optimizedFirstNode(t, "--x")
	// Should reduce to just identifierNode x
	if n, ok := node.(identifierNode); !ok {
		t.Errorf("Expected identifierNode, got %T", node)
	} else if n.payload != "x" {
		t.Errorf("Expected 'x', got '%s'", n.payload)
	}
}

func TestSimplifyDoubleBitwiseNot(t *testing.T) {
	node := optimizedFirstNode(t, "~~x")
	if n, ok := node.(identifierNode); !ok {
		t.Errorf("Expected identifierNode, got %T", node)
	} else if n.payload != "x" {
		t.Errorf("Expected 'x', got '%s'", n.payload)
	}
}

// --- End-to-end: optimization preserves correctness ---

func TestOptimizedComplexProgram(t *testing.T) {
	program := `
fn compute(x) {
	a := x ** 2
	b := a + 0
	c := b * 1
	d := c - 0
	d
}
compute(5)
`
	expectOptimizedResult(t, program, IntValue(25))
}

func TestOptimizedConstantIfInFunction(t *testing.T) {
	program := `
fn mode() if true {
	true -> :fast
	_ -> :slow
}
mode()
`
	expectOptimizedResult(t, program, AtomValue("fast"))
}

// --- Builtin constant folding ---

func TestFoldBuiltinLen(t *testing.T) {
	// len('hello') → 5
	node := optimizedFirstNode(t, "len('hello')")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 5 {
		t.Errorf("Expected 5, got %d", n.payload)
	}
}

func TestFoldBuiltinLenEmptyString(t *testing.T) {
	node := optimizedFirstNode(t, "len('')")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 0 {
		t.Errorf("Expected 0, got %d", n.payload)
	}
}

func TestFoldBuiltinLenList(t *testing.T) {
	node := optimizedFirstNode(t, "len([1, 2, 3])")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 3 {
		t.Errorf("Expected 3, got %d", n.payload)
	}
}

func TestFoldBuiltinLenNonConstant(t *testing.T) {
	// len(x) should NOT be folded
	node := optimizedFirstNode(t, "len(x)")
	if _, ok := node.(fnCallNode); !ok {
		t.Errorf("Expected fnCallNode, got %T", node)
	}
}

func TestFoldBuiltinType(t *testing.T) {
	node := optimizedFirstNode(t, "type(42)")
	if n, ok := node.(atomNode); !ok {
		t.Errorf("Expected atomNode, got %T", node)
	} else if n.payload != "int" {
		t.Errorf("Expected :int, got :%s", n.payload)
	}

	node = optimizedFirstNode(t, "type(3.14)")
	if n, ok := node.(atomNode); !ok {
		t.Errorf("Expected atomNode, got %T", node)
	} else if n.payload != "float" {
		t.Errorf("Expected :float, got :%s", n.payload)
	}

	node = optimizedFirstNode(t, "type(true)")
	if n, ok := node.(atomNode); !ok {
		t.Errorf("Expected atomNode, got %T", node)
	} else if n.payload != "bool" {
		t.Errorf("Expected :bool, got :%s", n.payload)
	}

	node = optimizedFirstNode(t, "type('hello')")
	if n, ok := node.(atomNode); !ok {
		t.Errorf("Expected atomNode, got %T", node)
	} else if n.payload != "string" {
		t.Errorf("Expected :string, got :%s", n.payload)
	}

	node = optimizedFirstNode(t, "type(:ok)")
	if n, ok := node.(atomNode); !ok {
		t.Errorf("Expected atomNode, got %T", node)
	} else if n.payload != "atom" {
		t.Errorf("Expected :atom, got :%s", n.payload)
	}

	node = optimizedFirstNode(t, "type(?)")
	if n, ok := node.(atomNode); !ok {
		t.Errorf("Expected atomNode, got %T", node)
	} else if n.payload != "null" {
		t.Errorf("Expected :null, got :%s", n.payload)
	}
}

func TestFoldBuiltinTypeNonConstant(t *testing.T) {
	node := optimizedFirstNode(t, "type(x)")
	if _, ok := node.(fnCallNode); !ok {
		t.Errorf("Expected fnCallNode, got %T", node)
	}
}

func TestFoldBuiltinInt(t *testing.T) {
	// int(3.7) → 3
	node := optimizedFirstNode(t, "int(3.7)")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 3 {
		t.Errorf("Expected 3, got %d", n.payload)
	}

	// int('42') → 42
	node = optimizedFirstNode(t, "int('42')")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 42 {
		t.Errorf("Expected 42, got %d", n.payload)
	}

	// int(10) → 10 (identity)
	node = optimizedFirstNode(t, "int(10)")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 10 {
		t.Errorf("Expected 10, got %d", n.payload)
	}
}

func TestFoldBuiltinIntBadString(t *testing.T) {
	// int('abc') should NOT be folded (returns ? at runtime)
	node := optimizedFirstNode(t, "int('abc')")
	if _, ok := node.(fnCallNode); !ok {
		t.Errorf("Expected fnCallNode, got %T", node)
	}
}

func TestFoldBuiltinFloat(t *testing.T) {
	// float(5) → 5.0
	node := optimizedFirstNode(t, "float(5)")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 5.0 {
		t.Errorf("Expected 5.0, got %f", n.payload)
	}

	// float('3.14') → 3.14
	node = optimizedFirstNode(t, "float('3.14')")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 3.14 {
		t.Errorf("Expected 3.14, got %f", n.payload)
	}
}

func TestFoldBuiltinString(t *testing.T) {
	// string(42) → '42'
	node := optimizedFirstNode(t, "string(42)")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "42" {
		t.Errorf("Expected '42', got '%s'", string(n.payload))
	}

	// string(3.14) → '3.14'
	node = optimizedFirstNode(t, "string(3.14)")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "3.14" {
		t.Errorf("Expected '3.14', got '%s'", string(n.payload))
	}

	// string(true) → 'true'
	node = optimizedFirstNode(t, "string(true)")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "true" {
		t.Errorf("Expected 'true', got '%s'", string(n.payload))
	}

	// string(:hello) → 'hello'
	node = optimizedFirstNode(t, "string(:hello)")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "hello" {
		t.Errorf("Expected 'hello', got '%s'", string(n.payload))
	}
}

// --- Builtin folding end-to-end ---

func TestFoldBuiltinLenEndToEnd(t *testing.T) {
	expectOptimizedResult(t, "len('hello world')", IntValue(11))
}

func TestFoldBuiltinChainedConversions(t *testing.T) {
	// string(int(3.7)) → '3'
	expectOptimizedResult(t, "string(int(3.7))", MakeString("3"))
}

func TestFoldBuiltinTypeComparison(t *testing.T) {
	// type(42) = :int should fold to true
	expectOptimizedResult(t, "type(42) = :int", BoolValue(true))
}

// --- String comparison folding ---

func TestFoldStringComparison(t *testing.T) {
	node := optimizedFirstNode(t, "'abc' > 'abb'")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}

	node = optimizedFirstNode(t, "'abc' < 'abd'")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}

	node = optimizedFirstNode(t, "'abc' >= 'abc'")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}

	node = optimizedFirstNode(t, "'abc' <= 'abc'")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload != true {
		t.Errorf("Expected true")
	}
}

// --- Dead assignment elimination ---

func TestDeadAssignmentEliminatedInBlock(t *testing.T) {
	// x := 1, x := 2, x → first assignment is dead
	expectOptimizedResult(t, "{ x := 1, x := 2, x }", IntValue(2))
}

func TestDeadAssignmentPreservedWhenImpure(t *testing.T) {
	// When the RHS has side effects, don't eliminate
	program := `
fn sideEffect() 42
{
	x := sideEffect()
	x := 10
	x
}
`
	expectOptimizedResult(t, program, IntValue(10))
}

// --- Single-branch if simplification ---

func TestSingleBranchWildcardSimplified(t *testing.T) {
	// if x { _ -> 42 } when x is pure → 42
	node := optimizedFirstNode(t, "if x { _ -> 42 }")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode from single-branch wildcard, got %T", node)
	} else if n.payload != 42 {
		t.Errorf("Expected 42, got %d", n.payload)
	}
}

func TestSingleBranchWildcardPreservedWhenImpure(t *testing.T) {
	// if f() { _ -> 42 } — f() may have side effects, keep the if
	node := optimizedFirstNode(t, "if f() { _ -> 42 }")
	if _, ok := node.(ifExprNode); !ok {
		t.Errorf("Expected ifExprNode preserved for impure cond, got %T", node)
	}
}

// --- Composite optimization: folding chains ---

func TestFoldChainedBuiltinsAndArithmetic(t *testing.T) {
	// len('hi') + len('bye') → 2 + 3 → 5
	node := optimizedFirstNode(t, "len('hi') + len('bye')")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 5 {
		t.Errorf("Expected 5, got %d", n.payload)
	}
}

func TestFoldTypeCheckInIf(t *testing.T) {
	// if type(42) { :int -> 'yes', _ -> 'no' } → 'yes'
	expectOptimizedResult(t, "if type(42) { :int -> 'yes', _ -> 'no' }", MakeString("yes"))
}

// --- Integer division optimization ---

func TestFoldIntDivExact(t *testing.T) {
	// 10 / 5 → int 2 (not float 2.0)
	node := optimizedFirstNode(t, "10 / 5")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode for exact division, got %T", node)
	} else if n.payload != 2 {
		t.Errorf("Expected 2, got %d", n.payload)
	}
}

func TestFoldIntDivNonExact(t *testing.T) {
	// 10 / 3 → float 3.333...
	node := optimizedFirstNode(t, "10 / 3")
	if _, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode for non-exact division, got %T", node)
	}
}

func TestIntDivExactRuntime(t *testing.T) {
	// Runtime: 10 / 2 should return int, not float
	expectOptimizedResult(t, "x := 10, x / 2", IntValue(5))
}

func TestIntDivNonExactRuntime(t *testing.T) {
	// Runtime: 10 / 4 should return 2.5
	expectOptimizedResult(t, "x := 10, x / 4", FloatValue(2.5))
}

// --- Constant property access folding ---

func TestFoldListIndex(t *testing.T) {
	// [10, 20, 30].1 → 20
	node := optimizedFirstNode(t, "[10, 20, 30].1")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 20 {
		t.Errorf("Expected 20, got %d", n.payload)
	}
}

func TestFoldListIndexFirst(t *testing.T) {
	node := optimizedFirstNode(t, "[1, 2, 3].0")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 1 {
		t.Errorf("Expected 1, got %d", n.payload)
	}
}

func TestFoldListIndexOutOfBounds(t *testing.T) {
	// Out of bounds should NOT be folded
	node := optimizedFirstNode(t, "[1, 2, 3].5")
	if _, ok := node.(propertyAccessNode); !ok {
		t.Errorf("Expected propertyAccessNode for OOB, got %T", node)
	}
}

func TestFoldStringIndex(t *testing.T) {
	// 'hello'.0 → 'h'
	node := optimizedFirstNode(t, "'hello'.0")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "h" {
		t.Errorf("Expected 'h', got '%s'", string(n.payload))
	}
}

func TestFoldStringIndexEndToEnd(t *testing.T) {
	expectOptimizedResult(t, "'abc'.2", MakeString("c"))
}

// --- Builtin codepoint/char folding ---

func TestFoldBuiltinCodepoint(t *testing.T) {
	// codepoint('A') → 65
	node := optimizedFirstNode(t, "codepoint('A')")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 65 {
		t.Errorf("Expected 65, got %d", n.payload)
	}
}

func TestFoldBuiltinCodepointMultibyte(t *testing.T) {
	// codepoint('AB') → not folded (multi-byte)
	node := optimizedFirstNode(t, "codepoint('AB')")
	if _, ok := node.(fnCallNode); !ok {
		t.Errorf("Expected fnCallNode for multi-byte, got %T", node)
	}
}

func TestFoldBuiltinChar(t *testing.T) {
	// char(65) → 'A'
	node := optimizedFirstNode(t, "char(65)")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "A" {
		t.Errorf("Expected 'A', got '%s'", string(n.payload))
	}
}

func TestFoldBuiltinCharClamped(t *testing.T) {
	// char(-1) → char(0) → '\0'
	node := optimizedFirstNode(t, "char(-1)")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if len(n.payload) != 1 || n.payload[0] != 0 {
		t.Errorf("Expected byte 0, got %v", n.payload)
	}
}

func TestFoldCodepointCharRoundTrip(t *testing.T) {
	// char(codepoint('Z')) → 'Z'
	expectOptimizedResult(t, "char(codepoint('Z'))", MakeString("Z"))
}

// --- Composite optimization: division + index ---

func TestFoldDivisionInContext(t *testing.T) {
	// 100 / 10 + [1, 2, 3].0 → 10 + 1 → 11
	node := optimizedFirstNode(t, "100 / 10 + [1, 2, 3].0")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 11 {
		t.Errorf("Expected 11, got %d", n.payload)
	}
}

// --- Object property access folding ---

func TestFoldObjectPropertyAccess(t *testing.T) {
	// {a: 1, b: 2}.a → 1
	node := optimizedFirstNode(t, "{a: 1, b: 2}.a")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 1 {
		t.Errorf("Expected 1, got %d", n.payload)
	}
}

func TestFoldObjectPropertyAccessString(t *testing.T) {
	// {name: 'alice'}.name → 'alice'
	node := optimizedFirstNode(t, "{name: 'alice'}.name")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "alice" {
		t.Errorf("Expected 'alice', got '%s'", string(n.payload))
	}
}

func TestFoldObjectPropertyAccessMiss(t *testing.T) {
	// {a: 1}.b → stays as propertyAccessNode (key not found)
	node := optimizedFirstNode(t, "{a: 1}.b")
	if _, ok := node.(propertyAccessNode); !ok {
		t.Errorf("Expected propertyAccessNode for missing key, got %T", node)
	}
}

// --- Negation of comparison ---

func TestNegateComparison(t *testing.T) {
	// !(x > 3) → x <= 3
	node := optimizedFirstNode(t, "!(x > 3)")
	if n, ok := node.(binaryNode); !ok {
		t.Errorf("Expected binaryNode, got %T", node)
	} else if n.op != leq {
		t.Errorf("Expected leq, got %v", n.op)
	}
}

func TestNegateEquality(t *testing.T) {
	// !(x = y) → x != y
	node := optimizedFirstNode(t, "!(x = y)")
	if n, ok := node.(binaryNode); !ok {
		t.Errorf("Expected binaryNode, got %T", node)
	} else if n.op != neq {
		t.Errorf("Expected neq, got %v", n.op)
	}
}

func TestNegateNeq(t *testing.T) {
	// !(x != y) → x = y
	node := optimizedFirstNode(t, "!(x != y)")
	if n, ok := node.(binaryNode); !ok {
		t.Errorf("Expected binaryNode, got %T", node)
	} else if n.op != eq {
		t.Errorf("Expected eq, got %v", n.op)
	}
}

func TestNegateLess(t *testing.T) {
	// !(a < b) → a >= b
	node := optimizedFirstNode(t, "!(a < b)")
	if n, ok := node.(binaryNode); !ok {
		t.Errorf("Expected binaryNode, got %T", node)
	} else if n.op != geq {
		t.Errorf("Expected geq, got %v", n.op)
	}
}

func TestNegateGeq(t *testing.T) {
	// !(a >= b) → a < b
	node := optimizedFirstNode(t, "!(a >= b)")
	if n, ok := node.(binaryNode); !ok {
		t.Errorf("Expected binaryNode, got %T", node)
	} else if n.op != less {
		t.Errorf("Expected less, got %v", n.op)
	}
}

// --- Self-equality folding ---

func TestSelfEqualityTrue(t *testing.T) {
	// x = x → true
	node := optimizedFirstNode(t, "x = x")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true")
	}
}

func TestSelfEqualityNeq(t *testing.T) {
	// x != x → false
	node := optimizedFirstNode(t, "x != x")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload {
		t.Errorf("Expected false")
	}
}

func TestSelfDeepEqTrue(t *testing.T) {
	// x == x → true
	node := optimizedFirstNode(t, "x == x")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true")
	}
}

// --- Block flattening ---

func TestNestedBlockFlattened(t *testing.T) {
	// ((42)) → 42
	node := optimizedFirstNode(t, "((42))")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode from block flattening, got %T", node)
	} else if n.payload != 42 {
		t.Errorf("Expected 42, got %d", n.payload)
	}
}

func TestSingleExprBlockUnwrap(t *testing.T) {
	// { 42 } → 42 (single constant expr blocks are unwrapped)
	node := optimizedFirstNode(t, "{ 42 }")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode from block unwrap, got %T", node)
	} else if n.payload != 42 {
		t.Errorf("Expected 42, got %d", n.payload)
	}
}

func TestSingleNonConstBlockPreserved(t *testing.T) {
	// { x + 1 } should remain a block (blocks mark computed property access)
	node := optimizedFirstNode(t, "{ x + 1 }")
	if _, ok := node.(blockNode); !ok {
		t.Errorf("Expected blockNode preserved for non-const single expr, got %T", node)
	}
}

// --- String concat runtime fast path ---

func TestStringConcatRuntime(t *testing.T) {
	expectOptimizedResult(t, "x := 'hello', y := ' world', x + y", MakeString("hello world"))
}

// --- Combined optimization chains ---

func TestNegateComparisonFoldsToConstant(t *testing.T) {
	// !(3 > 5) → 3 <= 5 → true
	node := optimizedFirstNode(t, "!(3 > 5)")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for !(3 > 5)")
	}
}

func TestObjectAccessEndToEnd(t *testing.T) {
	expectOptimizedResult(t, "{x: 10, y: 20}.y + 5", IntValue(25))
}

// --- Builtin pow folding ---

func TestFoldBuiltinPowIntInt(t *testing.T) {
	// pow(2, 10) → 1024
	node := optimizedFirstNode(t, "pow(2, 10)")
	if n, ok := node.(intNode); !ok {
		t.Errorf("Expected intNode, got %T", node)
	} else if n.payload != 1024 {
		t.Errorf("Expected 1024, got %d", n.payload)
	}
}

func TestFoldBuiltinPowNegExp(t *testing.T) {
	// pow(2, -1) → 0.5
	node := optimizedFirstNode(t, "pow(2, -1)")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 0.5 {
		t.Errorf("Expected 0.5, got %f", n.payload)
	}
}

func TestFoldBuiltinPowFloatBase(t *testing.T) {
	// pow(2.0, 3) → 8.0
	node := optimizedFirstNode(t, "pow(2.0, 3)")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 8.0 {
		t.Errorf("Expected 8.0, got %f", n.payload)
	}
}

func TestFoldBuiltinPowEndToEnd(t *testing.T) {
	expectOptimizedResult(t, "pow(3, 2) + 1", IntValue(10))
}

// --- VM fast path: modulus, neq, string concat (runtime) ---

func TestModulusRuntime(t *testing.T) {
	expectOptimizedResult(t, "x := 17, x % 5", IntValue(2))
}

func TestNeqRuntimeInt(t *testing.T) {
	expectOptimizedResult(t, "x := 5, x != 3", BoolValue(true))
	expectOptimizedResult(t, "x := 5, x != 5", BoolValue(false))
}

func TestStringConcatRuntimeVM(t *testing.T) {
	expectOptimizedResult(t, "'abc' + 'def'", MakeString("abcdef"))
}

// --- Block scope elision ---

func TestBlockHasLocalFlag(t *testing.T) {
	// Block with local assignment should have hasLocal set
	node := optimizedFirstNode(t, "{ x := 1, x + 2 }")
	if block, ok := node.(blockNode); !ok {
		t.Errorf("Expected blockNode, got %T", node)
	} else if !block.hasLocal {
		t.Error("Expected hasLocal=true for block with := assignment")
	}
}

func TestBlockScopeElisionCorrectness(t *testing.T) {
	// Local assignment in block should still work (scope is created)
	expectOptimizedResult(t, "{ x := 5, x * 2 }", IntValue(10))
	// Non-local blocks should see parent scope variables
	expectOptimizedResult(t, "x := 3, { x + 7 }", IntValue(10))
	// Nested blocks with mixed local/non-local
	expectOptimizedResult(t, "x := 1, { y := 2, { x + y } }", IntValue(3))
}

// --- If-expression fast paths ---

func TestIfExprIntFastPath(t *testing.T) {
	expectOptimizedResult(t, "x := 2, if x { 0 -> :zero, 1 -> :one, 2 -> :two, _ -> :other }", AtomValue("two"))
}

func TestIfExprAtomFastPath(t *testing.T) {
	expectOptimizedResult(t, "x := :hello, if x { :hello -> 1, :world -> 2, _ -> 0 }", IntValue(1))
}

func TestIfExprBoolFastPath(t *testing.T) {
	expectOptimizedResult(t, "x := true, if x { true -> :yes, false -> :no }", AtomValue("yes"))
	expectOptimizedResult(t, "x := false, if x { true -> :yes, false -> :no }", AtomValue("no"))
}

func TestIfExprWildcardMatch(t *testing.T) {
	// EmptyValue (_) wildcard should still work with fast paths
	expectOptimizedResult(t, "x := 99, if x { 0 -> :zero, _ -> :other }", AtomValue("other"))
}

// --- Property access identifier fast path ---

func TestPropertyAccessIdentFastPath(t *testing.T) {
	expectOptimizedResult(t, "obj := { a: 1, b: 2, c: 3 }, obj.a + obj.b + obj.c", IntValue(6))
}

func TestPropertyAccessIdentAssignFastPath(t *testing.T) {
	expectOptimizedResult(t, "obj := { x: 0 }, obj.x <- 42, obj.x", IntValue(42))
}

func TestPropertyAccessComputedStillWorks(t *testing.T) {
	// Computed property access (non-identifier) must still work
	expectOptimizedResult(t, "obj := { a: 1 }, key := 'a', obj.(key)", IntValue(1))
}

// --- Inline scope lookup for identifiers ---

func TestInlineScopeLookupLocal(t *testing.T) {
	// Variable in local scope should be found via inline fast path
	expectOptimizedResult(t, "fn f(x) x + 1, f(10)", IntValue(11))
}

func TestInlineScopeLookupParent(t *testing.T) {
	// Variable in parent scope (closure) should be found via parent chain
	expectOptimizedResult(t, "x := 5, fn f() x * 2, f()", IntValue(10))
}

func TestInlineScopeLookupBuiltin(t *testing.T) {
	// Builtins (in root scope) should still resolve
	expectOptimizedResult(t, "len('hello')", IntValue(5))
}

// --- EmptyValue wildcard early-exit in if-expr ---

func TestIfExprWildcardIntCond(t *testing.T) {
	expectOptimizedResult(t, "x := 42, if x { 0 -> :zero, _ -> :wildcard }", AtomValue("wildcard"))
}

func TestIfExprWildcardAtomCond(t *testing.T) {
	expectOptimizedResult(t, "x := :unknown, if x { :a -> 1, :b -> 2, _ -> 99 }", IntValue(99))
}

func TestIfExprWildcardBoolCond(t *testing.T) {
	expectOptimizedResult(t, "x := true, if x { false -> :no, _ -> :yes }", AtomValue("yes"))
}

// --- Inline fnCall identifier lookup ---

func TestFnCallInlineLookup(t *testing.T) {
	// Function call via identifier should use inline scope lookup
	expectOptimizedResult(t, "fn add(a, b) a + b, add(3, 4)", IntValue(7))
}

func TestFnCallRecursiveInlineLookup(t *testing.T) {
	// Recursive function should resolve itself via inline scope lookup
	expectOptimizedResult(t, `
fn fact(n) if n {
	0 -> 1
	_ -> n * fact(n - 1)
}
fact(5)
`, IntValue(120))
}

// --- Neq int fast path ---

func TestNeqIntFastPath(t *testing.T) {
	expectOptimizedResult(t, "x := 3, y := 5, x != y", BoolValue(true))
	expectOptimizedResult(t, "x := 7, y := 7, x != y", BoolValue(false))
}

// --- Inline put for local assignments ---

func TestInlinePutLocalAssignment(t *testing.T) {
	expectOptimizedResult(t, "{ x := 10, y := x + 5, y }", IntValue(15))
}

func TestInlinePutInLoop(t *testing.T) {
	// Loop with local assignments and mutations
	expectOptimizedResult(t, `
fn loop(n) {
	i := 0
	fn go {
		if i < n {
			true -> {
				i <- i + 1
				go()
			}
		}
	}
	go()
	i
}
loop(10)
`, IntValue(10))
}
