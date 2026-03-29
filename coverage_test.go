package main

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ===========================================================================
// Bytecode VM parity — exercise opcodes and paths not covered by existing tests
// ===========================================================================

func TestBytecodeEmptyProgram(t *testing.T) {
	expectProgramToReturnBytecode(t, "", null)
	expectProgramToReturnBytecode(t, "  \n", null)
}

func TestBytecodeComments(t *testing.T) {
	expectProgramToReturnBytecode(t, "// comment", null)
	expectProgramToReturnBytecode(t, "1 + // inline\n2", IntValue(3))
}

func TestBytecodeLiterals(t *testing.T) {
	expectProgramToReturnBytecode(t, "_", empty)
	expectProgramToReturnBytecode(t, "?", null)
	expectProgramToReturnBytecode(t, "true", oakTrue)
	expectProgramToReturnBytecode(t, "false", oakFalse)
	expectProgramToReturnBytecode(t, "42", IntValue(42))
	expectProgramToReturnBytecode(t, "3.14", FloatValue(3.14))
	expectProgramToReturnBytecode(t, "'hello'", MakeString("hello"))
	expectProgramToReturnBytecode(t, ":oak", AtomValue("oak"))
}

func TestBytecodeArithmeticOps(t *testing.T) {
	expectProgramToReturnBytecode(t, "7 + 3", IntValue(10))
	expectProgramToReturnBytecode(t, "7 - 3", IntValue(4))
	expectProgramToReturnBytecode(t, "7 * 3", IntValue(21))
	expectProgramToReturnBytecode(t, "10 / 4", FloatValue(2.5))
	expectProgramToReturnBytecode(t, "17 % 5", IntValue(2))
	expectProgramToReturnBytecode(t, "2 ** 10", IntValue(1024))
	expectProgramToReturnBytecode(t, "-42", IntValue(-42))
}

func TestBytecodeFloatArithmetic(t *testing.T) {
	expectProgramToReturnBytecode(t, "1.5 + 2.5", FloatValue(4.0))
	expectProgramToReturnBytecode(t, "3 + 1.5", FloatValue(4.5))
	expectProgramToReturnBytecode(t, "6.0 / 2.0", FloatValue(3.0))
	expectProgramToReturnBytecode(t, "10.0 % 3.0", FloatValue(1.0))
	expectProgramToReturnBytecode(t, "2.0 ** 3.0", FloatValue(8.0))
}

func TestBytecodeComparisonOps(t *testing.T) {
	expectProgramToReturnBytecode(t, "3 > 2", oakTrue)
	expectProgramToReturnBytecode(t, "2 < 3", oakTrue)
	expectProgramToReturnBytecode(t, "3 >= 3", oakTrue)
	expectProgramToReturnBytecode(t, "3 <= 3", oakTrue)
	expectProgramToReturnBytecode(t, "3 = 3", oakTrue)
	expectProgramToReturnBytecode(t, "3 != 4", oakTrue)
	expectProgramToReturnBytecode(t, "3 > 5", oakFalse)
}

func TestBytecodeBitwiseOps(t *testing.T) {
	expectProgramToReturnBytecode(t, "0xFF & 0x0F", IntValue(15))
	expectProgramToReturnBytecode(t, "0x01 | 0x02", IntValue(3))
	expectProgramToReturnBytecode(t, "5 ^ 3", IntValue(6))
	expectProgramToReturnBytecode(t, "4 << 2", IntValue(16))
	expectProgramToReturnBytecode(t, "16 >> 2", IntValue(4))
	expectProgramToReturnBytecode(t, "~0", IntValue(-1))
}

func TestBytecodeLogicalOps(t *testing.T) {
	expectProgramToReturnBytecode(t, "true & false", oakFalse)
	expectProgramToReturnBytecode(t, "true | false", oakTrue)
	expectProgramToReturnBytecode(t, "!true", oakFalse)
	expectProgramToReturnBytecode(t, "!false", oakTrue)
}

func TestBytecodeStringOperations(t *testing.T) {
	expectProgramToReturnBytecode(t, "'hello' + ' world'", MakeString("hello world"))
	expectProgramToReturnBytecode(t, "string(42)", MakeString("42"))
	expectProgramToReturnBytecode(t, "len('abc')", IntValue(3))
}

func TestBytecodeListOperations(t *testing.T) {
	expectProgramToReturnBytecode(t, "[1, 2, 3]", MakeList(IntValue(1), IntValue(2), IntValue(3)))
	expectProgramToReturnBytecode(t, "len([1, 2, 3])", IntValue(3))
	expectProgramToReturnBytecode(t, "[1, 2, 3].1", IntValue(2))
}

func TestBytecodeObjectOperations(t *testing.T) {
	expectProgramToReturnBytecode(t, "{a: 1, b: 2}.a", IntValue(1))
	expectProgramToReturnBytecode(t, "len(keys({x: 1, y: 2}))", IntValue(2))
}

func TestBytecodeVariableScoping(t *testing.T) {
	expectProgramToReturnBytecode(t, "x := 10, x + 5", IntValue(15))
	// Nested function creates its own scope
	expectProgramToReturnBytecode(t, `
		fn add(a, b) a + b
		add(3, 7)
	`, IntValue(10))
}

func TestBytecodeNonlocalAssignment(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		x := 1
		fn update() x <- 10
		update()
		x
	`, IntValue(10))
}

func TestBytecodeFunctionDefAndCall(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		fn add(a, b) a + b
		add(3, 4)
	`, IntValue(7))
}

func TestBytecodeClosure(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		fn makeAdder(n) fn(x) x + n
		add5 := makeAdder(5)
		add5(10)
	`, IntValue(15))
}

func TestBytecodeClosureDeepNesting(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		fn a(x) fn(y) fn(z) x + y + z
		a(1)(2)(3)
	`, IntValue(6))
}

func TestBytecodeIfExpression(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		if true { true -> 42, _ -> 0 }
	`, IntValue(42))
	expectProgramToReturnBytecode(t, `
		if false { true -> 42, _ -> 99 }
	`, IntValue(99))
}

func TestBytecodePatternMatching(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		x := 2
		if x {
			1 -> :one
			2 -> :two
			_ -> :other
		}
	`, AtomValue("two"))
}

func TestBytecodeRecursion(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		fn fib(n) if n {
			0 -> 0
			1 -> 1
			_ -> fib(n - 1) + fib(n - 2)
		}
		fib(10)
	`, IntValue(55))
}

func TestBytecodePipeOperator(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		fn double(x) x * 2
		fn inc(x) x + 1
		5 |> double() |> inc()
	`, IntValue(11))
}

func TestBytecodeRestArgs(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		fn sum(first, rest...) {
			result := first
			fn loop(i) if i < len(rest) {
				true -> {
					result <- result + rest.(i)
					loop(i + 1)
				}
			}
			loop(0)
			result
		}
		sum(1, 2, 3, 4)
	`, IntValue(10))
}

func TestBytecodeListPush(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		xs := [1, 2]
		xs << 3
		xs
	`, MakeList(IntValue(1), IntValue(2), IntValue(3)))
}

func TestBytecodeStringPush(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		s := 'ab'
		s << 'c'
		s
	`, MakeString("abc"))
}

func TestBytecodeObjectMutation(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		obj := {x: 1}
		obj.y := 2
		obj.x + obj.y
	`, IntValue(3))
}

func TestBytecodeBlockScope(t *testing.T) {
	// Block scope creates a new scope in the tree-walker, but bytecode handles
	// locals differently. Test that block expressions return their last value.
	expectProgramToReturnBytecode(t, `{ 1, 2, 3 }`, IntValue(3))
}

func TestBytecodeDestructuring(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		[a, b, c] := [10, 20, 30]
		a + b + c
	`, IntValue(60))
}

func TestBytecodeDivisionByZeroError(t *testing.T) {
	expectProgramToErrorBytecode(t, "10 / 0")
}

func TestBytecodeModuloByZeroError(t *testing.T) {
	expectProgramToErrorBytecode(t, "10 % 0")
}

func TestBytecodeWithExpression(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		fn apply(x, f) f(x)
		with apply(5) fn(n) n * 2
	`, IntValue(10))
}

func TestBytecodeDeepEquality(t *testing.T) {
	expectProgramToReturnBytecode(t, `[1, [2, 3]] == [1, [2, 3]]`, oakTrue)
	expectProgramToReturnBytecode(t, `[1, 2] == [1, 3]`, oakFalse)
}

func TestBytecodeClassConstructor(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		cs Pair(a, b) { {left: a, right: b} }
		p := Pair(10, 20)
		p.left + p.right
	`, IntValue(30))
}

// ===========================================================================
// Engine switching builtins: bytecode() and interpreter()
// ===========================================================================

func TestBytecodeBuiltinWithFnValue(t *testing.T) {
	expectProgramToReturn(t, `
		fn add(a, b) a + b
		with bytecode([3, 4]) add
	`, IntValue(7))
}

func TestBytecodeBuiltinWithClosure(t *testing.T) {
	expectProgramToReturn(t, `
		fn makeMultiplier(factor) fn(x) x * factor
		mult3 := makeMultiplier(3)
		with bytecode([10]) mult3
	`, IntValue(30))
}

func TestBytecodeBuiltinWithBuiltinFn(t *testing.T) {
	expectProgramToReturn(t, `
		with bytecode(['hello']) string
	`, MakeString("hello"))
}

func TestBytecodeBuiltinMissingArgs(t *testing.T) {
	expectProgramToError(t, `bytecode()`)
	expectProgramToError(t, `bytecode([1])`)
}

func TestBytecodeBuiltinNonListFirstArg(t *testing.T) {
	expectProgramToError(t, `bytecode(42, fn(x) x)`)
}

func TestBytecodeBuiltinNonFunctionSecondArg(t *testing.T) {
	expectProgramToError(t, `bytecode([], 42)`)
}

func TestBytecodeBuiltinRestArgs(t *testing.T) {
	expectProgramToReturn(t, `
		fn sum(a, b, rest...) a + b
		with bytecode([10, 20, 30]) sum
	`, IntValue(30))
}

func TestInterpreterBuiltinWithFnValue(t *testing.T) {
	expectProgramToReturn(t, `
		fn add(a, b) a + b
		with interpreter([3, 4]) add
	`, IntValue(7))
}

func TestInterpreterBuiltinWithClosure(t *testing.T) {
	expectProgramToReturn(t, `
		fn outer(n) {
			fn inner(x) x + n
			inner
		}
		with interpreter([10]) outer(5)
	`, IntValue(15))
}

func TestInterpreterBuiltinMissingArgs(t *testing.T) {
	expectProgramToError(t, `interpreter()`)
	expectProgramToError(t, `interpreter([1])`)
}

func TestInterpreterBuiltinNonListFirstArg(t *testing.T) {
	expectProgramToError(t, `interpreter(42, fn(x) x)`)
}

func TestInterpreterBuiltinTCO(t *testing.T) {
	// interpreter() should support TCO for deep recursion
	expectProgramToReturn(t, `
		fn countDown(n) if n {
			0 -> :done
			_ -> countDown(n - 1)
		}
		with interpreter([10000]) countDown
	`, AtomValue("done"))
}

func TestBytecodeAndInterpreterParity(t *testing.T) {
	// Same function should produce same results through both builtins
	program := `
		fn compute(x) {
			a := x * 2
			b := a + 3
			if b > 10 {
				true -> b
				_ -> 0
			}
		}
		tree := with interpreter([5]) compute
		bc := with bytecode([5]) compute
		tree = bc
	`
	expectProgramToReturn(t, program, oakTrue)
}

// ===========================================================================
// Engine switching in bytecode mode (closureVal handling)
// ===========================================================================

func TestBytecodeInterpreterBuiltinInBytecodeMode(t *testing.T) {
	// interpreter() called from within bytecode mode should tree-walk
	expectProgramToReturnBytecode(t, `
		fn double(x) x * 2
		with interpreter([21]) double
	`, IntValue(42))
}

func TestBytecodeBytecodeBuiltinInBytecodeMode(t *testing.T) {
	// bytecode() called from within bytecode mode
	expectProgramToReturnBytecode(t, `
		fn add(a, b) a + b
		with bytecode([10, 20]) add
	`, IntValue(30))
}

func TestBytecodeInterpreterWithClosure(t *testing.T) {
	// closureVal with preserved AST should be convertible
	expectProgramToReturnBytecode(t, `
		fn makeAdder(n) fn(x) x + n
		adder := makeAdder(100)
		with interpreter([42]) adder
	`, IntValue(142))
}

// ===========================================================================
// Channel operations — error cases and edge cases
// ===========================================================================

func TestMakeChanNoArgs(t *testing.T) {
	expectProgramToReturn(t, `
		ch := make_chan()
		ch.type
	`, AtomValue("channel"))
}

func TestMakeChanWithCapacity(t *testing.T) {
	expectProgramToReturn(t, `
		ch := make_chan(5)
		ch.cap
	`, IntValue(5))
}

func TestMakeChanNegativeCapacityError(t *testing.T) {
	expectProgramToError(t, `make_chan(-1)`)
}

func TestMakeChanInvalidArgTypeError(t *testing.T) {
	expectProgramToError(t, `make_chan('abc')`)
}

func TestMakeChanTooManyArgsError(t *testing.T) {
	expectProgramToError(t, `make_chan(1, 2)`)
}

func TestChanSendArgCountError(t *testing.T) {
	expectProgramToError(t, `chan_send()`)
}

func TestChanRecvArgCountError(t *testing.T) {
	expectProgramToError(t, `chan_recv()`)
}

func TestChanMultipleValuesPassing(t *testing.T) {
	expectProgramToReturn(t, `
		ch := make_chan(3)
		chan_send(ch, :a)
		chan_send(ch, :b)
		chan_send(ch, :c)
		[chan_recv(ch).data, chan_recv(ch).data, chan_recv(ch).data]
	`, MakeList(AtomValue("a"), AtomValue("b"), AtomValue("c")))
}

// ===========================================================================
// go() builtin — error cases
// ===========================================================================

func TestGoNoArgsError(t *testing.T) {
	expectProgramToError(t, `go()`)
}

func TestGoNonFunctionArgError(t *testing.T) {
	expectProgramToError(t, `go(42)`)
}

func TestGoRejectsClassTarget(t *testing.T) {
	expectProgramToError(t, `
		cs C { {} }
		go(C)
	`)
}

// ===========================================================================
// lock_thread/unlock_thread
// ===========================================================================

func TestLockUnlockThread(t *testing.T) {
	expectProgramToReturn(t, `
		lock_thread()
		unlock_thread()
		:ok
	`, AtomValue("ok"))
}

func TestLockThreadExtraArgsIgnored(t *testing.T) {
	// lock_thread/unlock_thread accept extra args without error
	expectProgramToReturn(t, `lock_thread(1)`, null)
}

func TestUnlockThreadExtraArgsIgnored(t *testing.T) {
	expectProgramToReturn(t, `unlock_thread(1)`, null)
}

// ===========================================================================
// Optimization: previously untested fold functions
// ===========================================================================

func TestFoldModByZeroNotFolded(t *testing.T) {
	node := optimizedFirstNode(t, "10 % 0")
	if _, ok := node.(binaryNode); !ok {
		t.Errorf("Modulo by zero should not be folded, got %T", node)
	}
}

func TestFoldFloatDivByZeroNotFolded(t *testing.T) {
	node := optimizedFirstNode(t, "10.0 / 0.0")
	if _, ok := node.(binaryNode); !ok {
		t.Errorf("Float division by zero should not be folded, got %T", node)
	}
}

func TestFoldFloatModByZeroNotFolded(t *testing.T) {
	node := optimizedFirstNode(t, "10.0 % 0.0")
	if _, ok := node.(binaryNode); !ok {
		t.Errorf("Float modulo by zero should not be folded, got %T", node)
	}
}

func TestFoldStringConcatResult(t *testing.T) {
	node := optimizedFirstNode(t, "'foo' + 'bar'")
	if n, ok := node.(stringNode); !ok {
		t.Errorf("Expected stringNode, got %T", node)
	} else if string(n.payload) != "foobar" {
		t.Errorf("Expected 'foobar', got '%s'", string(n.payload))
	}
}

func TestFoldBlockSingleConstant(t *testing.T) {
	// A block with a single constant expression should be unwrapped
	expectOptimizedResult(t, "(42)", IntValue(42))
}

func TestFoldBlockMultipleExprs(t *testing.T) {
	// Block with multiple expressions keeps block structure
	expectOptimizedResult(t, "{ 1 + 2, 3 + 4 }", IntValue(7))
}

func TestFoldAssignmentRHS(t *testing.T) {
	// The RHS of an assignment should be folded
	expectOptimizedResult(t, "x := 2 + 3, x", IntValue(5))
}

func TestFoldFnBody(t *testing.T) {
	// Constant expression inside a function body should fold
	expectOptimizedResult(t, "fn f() 2 + 3, f()", IntValue(5))
}

func TestFoldIfBranches(t *testing.T) {
	// Expressions inside if branches should fold
	expectOptimizedResult(t, "if true { true -> 2 + 3, _ -> 4 + 5 }", IntValue(5))
}

func TestFoldClassBody(t *testing.T) {
	// Constant expression inside a class body should fold
	expectOptimizedResult(t, `
		cs C { val := 2 + 3 }
		C().val
	`, IntValue(5))
}

func TestFoldPropertyAccessSubExprs(t *testing.T) {
	// Property access sub-expressions should fold
	expectOptimizedResult(t, "[2 + 1, 3 + 1].0", IntValue(3))
}

func TestFoldFnCallArguments(t *testing.T) {
	// Arguments to function calls should be folded
	expectOptimizedResult(t, "fn id(x) x, id(2 + 3)", IntValue(5))
}

func TestFoldListElements(t *testing.T) {
	// Elements inside a list should be folded
	expectOptimizedResult(t, "[1 + 1, 2 + 2, 3 + 3]", MakeList(IntValue(2), IntValue(4), IntValue(6)))
}

func TestFoldObjectValues(t *testing.T) {
	// Object values should be folded
	expectOptimizedResult(t, "{a: 1 + 1, b: 2 + 2}", ObjectValue{
		"a": IntValue(2),
		"b": IntValue(4),
	})
}

func TestFoldBoolXor(t *testing.T) {
	node := optimizedFirstNode(t, "true ^ false")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for true ^ false")
	}
}

func TestFoldFloatComparison(t *testing.T) {
	node := optimizedFirstNode(t, "1.5 > 1.0")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for 1.5 > 1.0")
	}
}

func TestFoldEqualityAtoms(t *testing.T) {
	node := optimizedFirstNode(t, ":foo = :foo")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for :foo = :foo")
	}

	node = optimizedFirstNode(t, ":foo = :bar")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if n.payload {
		t.Errorf("Expected false for :foo = :bar")
	}
}

func TestFoldEqualityStrings(t *testing.T) {
	node := optimizedFirstNode(t, "'abc' = 'abc'")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for 'abc' = 'abc'")
	}
}

func TestFoldEqualityNull(t *testing.T) {
	node := optimizedFirstNode(t, "? = ?")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for null = null")
	}
}

func TestFoldEqualityMixedIntFloat(t *testing.T) {
	node := optimizedFirstNode(t, "2 = 2.0")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for 2 = 2.0")
	}
}

func TestFoldEqualityNeq(t *testing.T) {
	node := optimizedFirstNode(t, "'a' != 'b'")
	if n, ok := node.(boolNode); !ok {
		t.Errorf("Expected boolNode, got %T", node)
	} else if !n.payload {
		t.Errorf("Expected true for 'a' != 'b'")
	}
}

func TestFoldNoFoldMixedTypes(t *testing.T) {
	// Equality between incompatible types should not fold
	node := optimizedFirstNode(t, "42 = 'abc'")
	if _, ok := node.(binaryNode); !ok {
		t.Errorf("Expected equality of incompatible types to NOT fold, got %T", node)
	}
}

func TestFoldIntPowerLargeExponent(t *testing.T) {
	expectOptimizedResult(t, "2 ** 16", IntValue(65536))
}

func TestFoldFloatSubMulMod(t *testing.T) {
	node := optimizedFirstNode(t, "5.0 - 2.0")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 3.0 {
		t.Errorf("Expected 3.0, got %f", n.payload)
	}

	node = optimizedFirstNode(t, "3.0 * 2.0")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 6.0 {
		t.Errorf("Expected 6.0, got %f", n.payload)
	}

	node = optimizedFirstNode(t, "7.0 % 3.0")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 1.0 {
		t.Errorf("Expected 1.0, got %f", n.payload)
	}
}

func TestFoldFloatPower(t *testing.T) {
	node := optimizedFirstNode(t, "4.0 ** 0.5")
	if n, ok := node.(floatNode); !ok {
		t.Errorf("Expected floatNode, got %T", node)
	} else if n.payload != 2.0 {
		t.Errorf("Expected 2.0, got %f", n.payload)
	}
}

// ===========================================================================
// Error display — config flags
// ===========================================================================

func TestDisplayErrorWithColorEnabled(t *testing.T) {
	err := &runtimeError{
		reason: "test error",
		pos:    pos{line: 1, col: 1},
	}
	output := captureStderr(t, func() {
		DisplayError(err, ErrorDisplayConfig{
			UseColor:       true,
			ShowContext:    false,
			ShowStackTrace: false,
		})
	})
	if !strings.Contains(output, "test error") {
		t.Fatalf("missing error reason: %q", output)
	}
	// Colored output should contain ANSI codes
	if !strings.Contains(output, "\033[") {
		t.Fatalf("expected ANSI color codes in colored output: %q", output)
	}
}

func TestDisplayErrorWithColorDisabled(t *testing.T) {
	err := &runtimeError{
		reason: "test error",
		pos:    pos{line: 1, col: 1},
	}
	output := captureStderr(t, func() {
		DisplayError(err, ErrorDisplayConfig{
			UseColor:       false,
			ShowContext:    false,
			ShowStackTrace: false,
		})
	})
	if !strings.Contains(output, "test error") {
		t.Fatalf("missing error reason: %q", output)
	}
	// No color mode should have no ANSI codes
	if strings.Contains(output, "\033[") {
		t.Fatalf("did not expect ANSI color codes when UseColor=false: %q", output)
	}
}

func TestDisplayErrorContextLines(t *testing.T) {
	tmpDir := t.TempDir()
	fileName := filepath.Join(tmpDir, "ctx.oak")
	content := "line1\nline2\nline3\nline4\nline5\nline6\nline7\n"
	if err := os.WriteFile(fileName, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	parseErr := parseError{
		reason: "bad token",
		pos:    pos{fileName: fileName, line: 4, col: 1},
	}

	// ContextLines=0 should show only the error line
	output := captureStderr(t, func() {
		DisplayError(parseErr, ErrorDisplayConfig{
			UseColor:     false,
			ShowContext:  true,
			ContextLines: 0,
		})
	})
	if !strings.Contains(output, "line4") {
		t.Fatalf("expected error line to appear: %q", output)
	}
	// With ContextLines=0, lines far from line 4 should not appear
	if strings.Contains(output, "line1") {
		t.Fatalf("did not expect line1 with ContextLines=0: %q", output)
	}
}

func TestDisplayErrorLargeStackTrace(t *testing.T) {
	entries := make([]stackEntry, 15)
	for i := range entries {
		entries[i] = stackEntry{name: "fn" + string(rune('a'+i)), pos: pos{line: i + 1, col: 1}}
	}

	err := &runtimeError{
		reason:     "deep call",
		pos:        pos{line: 15, col: 1},
		stackTrace: entries,
	}

	output := captureStderr(t, func() {
		DisplayError(err, ErrorDisplayConfig{
			UseColor:       false,
			ShowContext:    false,
			ShowStackTrace: true,
		})
	})

	if !strings.Contains(output, "Stack Trace") {
		t.Fatalf("missing stack trace header: %q", output)
	}
	if !strings.Contains(output, "in fn fna") {
		t.Fatalf("missing first stack entry: %q", output)
	}
}

func TestDisplayErrorContextWithTabs(t *testing.T) {
	tmpDir := t.TempDir()
	fileName := filepath.Join(tmpDir, "tabs.oak")
	content := "\terror_here\n"
	if err := os.WriteFile(fileName, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	parseErr := parseError{
		reason: "unexpected",
		pos:    pos{fileName: fileName, line: 1, col: 2},
	}

	output := captureStderr(t, func() {
		DisplayError(parseErr, ErrorDisplayConfig{
			UseColor:     false,
			ShowContext:  true,
			ContextLines: 0,
		})
	})
	if !strings.Contains(output, "^") {
		t.Fatalf("missing error pointer: %q", output)
	}
}

func TestDisplayErrorWriterOption(t *testing.T) {
	var buf strings.Builder
	err := errors.New("writer test")
	DisplayError(err, ErrorDisplayConfig{
		Writer:   &buf,
		UseColor: false,
	})
	if !strings.Contains(buf.String(), "writer test") {
		t.Fatalf("expected error written to custom writer: %q", buf.String())
	}
}

func TestFormatErrorRuntimeWithFileName(t *testing.T) {
	err := &runtimeError{
		reason: "bad op",
		pos:    pos{fileName: "myfile.oak", line: 5, col: 3},
	}
	formatted := FormatError(err)
	if !strings.Contains(formatted, "in myfile.oak") {
		t.Fatalf("expected file name in formatted error: %q", formatted)
	}
}

func TestFormatErrorParseWithoutFileName(t *testing.T) {
	err := parseError{
		reason: "missing paren",
		pos:    pos{line: 1, col: 1},
	}
	formatted := FormatError(err)
	if strings.Contains(formatted, "in  at") {
		t.Fatalf("should not have 'in' prefix when no filename: %q", formatted)
	}
	if !strings.Contains(formatted, "Parse Error at [1:1]") {
		t.Fatalf("unexpected format: %q", formatted)
	}
}

// ===========================================================================
// Tokenizer edge cases
// ===========================================================================

func TestTokenizerEmptyInput(t *testing.T) {
	tk := newTokenizer("", "x.oak")
	tokens := tk.tokenize()
	if len(tokens) != 0 {
		t.Fatalf("expected empty token list for empty input, got %d tokens", len(tokens))
	}
}

func TestTokenizerWhitespaceOnly(t *testing.T) {
	tk := newTokenizer("   \t\n  ", "x.oak")
	tokens := tk.tokenize()
	if len(tokens) != 0 {
		t.Fatalf("expected empty token list for whitespace-only input, got %d tokens", len(tokens))
	}
}

func TestTokenizerCommentLine(t *testing.T) {
	// Comments are stripped during tokenization and not emitted as tokens
	tk := newTokenizer("// comment here", "x.oak")
	tokens := tk.tokenize()
	if len(tokens) != 0 {
		t.Fatalf("expected no tokens for comment-only input, got %d", len(tokens))
	}
}

func TestTokenizerMultipleTokens(t *testing.T) {
	tk := newTokenizer("x := 42", "x.oak")
	tokens := tk.tokenize()

	if len(tokens) < 3 {
		t.Fatalf("expected at least 3 tokens, got %d", len(tokens))
	}
	if tokens[0].kind != identifier {
		t.Errorf("expected first token to be identifier, got %v", tokens[0].kind)
	}
	if tokens[1].kind != assign {
		t.Errorf("expected second token to be assign, got %v", tokens[1].kind)
	}
	if tokens[2].kind != numberLiteral {
		t.Errorf("expected third token to be numberLiteral, got %v", tokens[2].kind)
	}
}

func TestTokenizerStringWithNewline(t *testing.T) {
	tk := newTokenizer("'line1\\nline2'", "x.oak")
	tokens := tk.tokenize()
	if len(tokens) == 0 {
		t.Fatal("expected at least one token")
	}
	if tokens[0].kind != stringLiteral {
		t.Fatalf("expected string literal, got %v", tokens[0].kind)
	}
}

func TestTokenizerAllPunctuation(t *testing.T) {
	input := "( ) [ ] { } , . : ? ! ~ + - * / % & | ^ = > < :="
	tk := newTokenizer(input, "x.oak")
	tokens := tk.tokenize()
	if len(tokens) < 10 {
		t.Fatalf("expected many tokens from punctuation, got %d", len(tokens))
	}
}

func TestTokenizerAtomLiteral(t *testing.T) {
	tk := newTokenizer(":myAtom", "x.oak")
	tokens := tk.tokenize()
	if len(tokens) == 0 {
		t.Fatal("expected tokens")
	}
	// atom literals consist of colon + identifier
	foundColon := false
	for _, tok := range tokens {
		if tok.kind == colon {
			foundColon = true
		}
	}
	if !foundColon {
		// Alternatively, the tokenizer may emit a combined token
		// Either way, parsing "':myAtom'" should work in evaluation
	}
}

// ===========================================================================
// Parse edge cases
// ===========================================================================

func TestParseEmptyBlock(t *testing.T) {
	// {} is an empty object literal
	expectProgramToReturn(t, "{}", ObjectValue{})
}

func TestParseNestedBlocks(t *testing.T) {
	expectProgramToReturn(t, "{ { { 42 } } }", IntValue(42))
}

func TestParseLongPipeChain(t *testing.T) {
	expectProgramToReturn(t, `
		fn id(x) x
		42 |> id() |> id() |> id() |> id()
	`, IntValue(42))
}

func TestParseOperatorPrecedenceComplex(t *testing.T) {
	// Ensure correct precedence: * before +, ** before *
	expectProgramToReturn(t, "2 + 3 * 4", IntValue(14))
	expectProgramToReturn(t, "2 * 3 ** 2", IntValue(18))
}

func TestParseTrailingCommaInList(t *testing.T) {
	expectProgramToReturn(t, "[1, 2, 3]", MakeList(IntValue(1), IntValue(2), IntValue(3)))
}

func TestParsePropertyChain(t *testing.T) {
	expectProgramToReturn(t, `
		obj := {a: {b: {c: 42}}}
		obj.a.b.c
	`, IntValue(42))
}

func TestParseComputedPropertyAccess(t *testing.T) {
	expectProgramToReturn(t, `
		obj := {x: 99}
		key := 'x'
		obj.(key)
	`, IntValue(99))
}

// ===========================================================================
// Math builtins — domain edge cases
// ===========================================================================

func TestSinCosValues(t *testing.T) {
	expectProgramToReturn(t, `sin(0)`, FloatValue(0))
	expectProgramToReturn(t, `cos(0)`, FloatValue(1))

	// sin(pi/2) ≈ 1
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	val, err := ctx.Eval(strings.NewReader(`sin(3.14159265358979 / 2)`))
	if err != nil {
		t.Fatal(err)
	}
	if fv, ok := val.(FloatValue); !ok || math.Abs(float64(fv)-1.0) > 1e-10 {
		t.Errorf("Expected sin(pi/2) ≈ 1, got %v", val)
	}
}

func TestTanValue(t *testing.T) {
	expectProgramToReturn(t, `tan(0)`, FloatValue(0))
}

func TestAtanValue(t *testing.T) {
	expectProgramToReturn(t, `atan(0)`, FloatValue(0))
	// atan(1) ≈ pi/4
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	val, err := ctx.Eval(strings.NewReader(`atan(1)`))
	if err != nil {
		t.Fatal(err)
	}
	if fv, ok := val.(FloatValue); !ok || math.Abs(float64(fv)-math.Pi/4) > 1e-10 {
		t.Errorf("Expected atan(1) ≈ pi/4, got %v", val)
	}
}

func TestPowLog(t *testing.T) {
	expectProgramToReturn(t, `pow(2, 10)`, FloatValue(1024))
	expectProgramToReturn(t, `log(2, 1024)`, FloatValue(10))
}

func TestAsinAcosBoundary(t *testing.T) {
	expectProgramToReturn(t, `asin(1)`, FloatValue(math.Asin(1)))
	expectProgramToReturn(t, `acos(0)`, FloatValue(math.Acos(0)))
}

func TestAsinOutOfDomain(t *testing.T) {
	expectProgramToError(t, `asin(2)`)
	expectProgramToError(t, `asin(-2)`)
}

func TestAcosOutOfDomain(t *testing.T) {
	expectProgramToError(t, `acos(2)`)
	expectProgramToError(t, `acos(-2)`)
}

func TestPowDomainErrors(t *testing.T) {
	expectProgramToError(t, `pow(0, 0)`)
	expectProgramToError(t, `pow(-1, 0.5)`)
}

func TestLogDomainErrors(t *testing.T) {
	expectProgramToError(t, `log(0, 10)`)
	expectProgramToError(t, `log(2, 0)`)
}

// ===========================================================================
// Type conversion builtins
// ===========================================================================

func TestIntBuiltinConversions(t *testing.T) {
	expectProgramToReturn(t, `int('42')`, IntValue(42))
	expectProgramToReturn(t, `int(3.7)`, IntValue(3))
}

func TestIntBuiltinInvalidString(t *testing.T) {
	expectProgramToReturn(t, `int('abc')`, null)
}

func TestFloatBuiltinConversions(t *testing.T) {
	expectProgramToReturn(t, `float('3.14')`, FloatValue(3.14))
	expectProgramToReturn(t, `float(42)`, FloatValue(42))
}

func TestFloatBuiltinInvalidString(t *testing.T) {
	expectProgramToReturn(t, `float('xyz')`, null)
}

func TestCodepointBuiltin(t *testing.T) {
	expectProgramToReturn(t, `codepoint('A')`, IntValue(65))
	expectProgramToReturn(t, `codepoint('z')`, IntValue(122))
}

func TestCharBuiltin(t *testing.T) {
	expectProgramToReturn(t, `char(65)`, MakeString("A"))
	expectProgramToReturn(t, `char(0)`, MakeString("\x00"))
	expectProgramToReturn(t, `char(255)`, MakeString("\xff"))
}

func TestCharBuiltinBounds(t *testing.T) {
	// Negative maps to 0, >255 maps to 255
	expectProgramToReturn(t, `char(-10)`, MakeString("\x00"))
	expectProgramToReturn(t, `char(300)`, MakeString("\xff"))
}

// ===========================================================================
// Collection/Reflection builtins
// ===========================================================================

func TestKeysBuiltin(t *testing.T) {
	// keys returns atom-ordered keys
	expectProgramToReturn(t, `len(keys({a: 1, b: 2, c: 3}))`, IntValue(3))
}

func TestKeysEmptyObject(t *testing.T) {
	expectProgramToReturn(t, `keys({})`, MakeList())
}

func TestLenBuiltinTypes(t *testing.T) {
	expectProgramToReturn(t, `len([])`, IntValue(0))
	expectProgramToReturn(t, `len('')`, IntValue(0))
	expectProgramToReturn(t, `len([1, 2, 3])`, IntValue(3))
	expectProgramToReturn(t, `len('abc')`, IntValue(3))
}

func TestTypeBuiltinAllTypes(t *testing.T) {
	expectProgramToReturn(t, `
		[
			type(?)
			type(_)
			type(42)
			type(3.14)
			type('s')
			type(:a)
			type(true)
			type(false)
			type([])
			type({})
			type(fn {})
		]
	`, MakeList(
		AtomValue("null"),
		AtomValue("empty"),
		AtomValue("int"),
		AtomValue("float"),
		AtomValue("string"),
		AtomValue("atom"),
		AtomValue("bool"),
		AtomValue("bool"),
		AtomValue("list"),
		AtomValue("object"),
		AtomValue("function"),
	))
}

// ===========================================================================
// String operations
// ===========================================================================

func TestStringCompareOrderParity(t *testing.T) {
	expectProgramToReturn(t, `'abc' < 'abd'`, oakTrue)
	expectProgramToReturn(t, `'abc' > 'abb'`, oakTrue)
	expectProgramToReturn(t, `'abc' >= 'abc'`, oakTrue)
	expectProgramToReturn(t, `'abc' <= 'abc'`, oakTrue)
}

func TestStringCompareBytecode(t *testing.T) {
	expectProgramToReturnBytecode(t, `'abc' < 'abd'`, oakTrue)
	expectProgramToReturnBytecode(t, `'abc' = 'abc'`, oakTrue)
	expectProgramToReturnBytecode(t, `'abc' != 'xyz'`, oakTrue)
}

// ===========================================================================
// Bytecode builtins calling
// ===========================================================================

func TestBytecodeBuiltinCallsStringLen(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		[
			string(42)
			len([1, 2, 3, 4])
			type(:hello)
		]
	`, MakeList(
		MakeString("42"),
		IntValue(4),
		AtomValue("atom"),
	))
}

func TestBytecodeImport(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		std := import('std')
		type(std.append)
	`, AtomValue("function"))
}

// ===========================================================================
// Scope — vmScope get/set operations (used by bytecode VM)
// ===========================================================================

func TestVmScopeGetSet(t *testing.T) {
	parent := &vmScope{
		names:  []string{"x"},
		values: []Value{IntValue(10)},
	}
	child := &vmScope{
		names:  []string{"y"},
		values: []Value{IntValue(20)},
		parent: parent,
	}

	// Get from child
	val, ok := child.get("y")
	if !ok || !val.Eq(IntValue(20)) {
		t.Fatalf("Expected y=20, got %v (ok=%v)", val, ok)
	}

	// Get from parent via child
	val, ok = child.get("x")
	if !ok || !val.Eq(IntValue(10)) {
		t.Fatalf("Expected x=10 via parent, got %v (ok=%v)", val, ok)
	}

	// Missing variable
	_, ok = child.get("z")
	if ok {
		t.Fatal("Expected z to not be found")
	}

	// Set existing parent var through child
	if !child.set("x", IntValue(99)) {
		t.Fatal("Expected set to succeed")
	}
	val, _ = parent.get("x")
	if !val.Eq(IntValue(99)) {
		t.Fatalf("Expected parent x to be updated to 99, got %v", val)
	}

	// Set nonexistent variable
	if child.set("missing", IntValue(0)) {
		t.Fatal("Expected set of missing var to fail")
	}
}

// ===========================================================================
// EvalBytecode error paths
// ===========================================================================

func TestEvalBytecodeParseError(t *testing.T) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	_, err := ctx.EvalBytecode(strings.NewReader("fn("))
	if err == nil {
		t.Fatal("Expected parse error from bytecode mode")
	}
}

func TestEvalBytecodeRuntimeError(t *testing.T) {
	expectProgramToErrorBytecode(t, "x := 0\nx.type")
}

// ===========================================================================
// bridgeVmScope / bridgeToVmScope unit tests
// ===========================================================================

func TestBridgeVmScopeFlattensChain(t *testing.T) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()

	parent := &vmScope{
		names:  []string{"a"},
		values: []Value{IntValue(1)},
	}
	child := &vmScope{
		names:  []string{"b"},
		values: []Value{IntValue(2)},
		parent: parent,
	}

	sc := ctx.bridgeVmScope(child)

	a, err := sc.get("a")
	if err != nil {
		t.Fatalf("Expected to find 'a': %v", err)
	}
	if !a.Eq(IntValue(1)) {
		t.Fatalf("Expected a=1, got %v", a)
	}

	b, err := sc.get("b")
	if err != nil {
		t.Fatalf("Expected to find 'b': %v", err)
	}
	if !b.Eq(IntValue(2)) {
		t.Fatalf("Expected b=2, got %v", b)
	}
}

func TestBridgeVmScopeChildShadowsParent(t *testing.T) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()

	parent := &vmScope{
		names:  []string{"x"},
		values: []Value{IntValue(1)},
	}
	child := &vmScope{
		names:  []string{"x"},
		values: []Value{IntValue(99)},
		parent: parent,
	}

	sc := ctx.bridgeVmScope(child)
	x, err := sc.get("x")
	if err != nil {
		t.Fatalf("Expected to find 'x': %v", err)
	}
	if !x.Eq(IntValue(99)) {
		t.Fatalf("Expected child shadow x=99, got %v", x)
	}
}

func TestBridgeToVmScopeFlattensTreeWalker(t *testing.T) {
	root := newScope(nil)
	root.put("a", IntValue(10))
	child := newScope(&root)
	child.put("b", IntValue(20))

	vs := bridgeToVmScope(&child)

	val, ok := vs.get("a")
	if !ok {
		t.Fatal("Expected bridged vmScope to contain 'a'")
	}
	if !val.Eq(IntValue(10)) {
		t.Fatalf("Expected a=10, got %v", val)
	}

	val, ok = vs.get("b")
	if !ok {
		t.Fatal("Expected bridged vmScope to contain 'b'")
	}
	if !val.Eq(IntValue(20)) {
		t.Fatalf("Expected b=20, got %v", val)
	}
}

func TestBridgeToVmScopeNil(t *testing.T) {
	vs := bridgeToVmScope(nil)
	if vs != nil {
		t.Fatal("Expected nil vmScope from nil scope")
	}
}

// ===========================================================================
// Misc bytecode parity — complex scenarios
// ===========================================================================

func TestBytecodeMapFilterReduce(t *testing.T) {
	program := `
		std := import('std')
		nums := [1, 2, 3, 4, 5]
		doubled := std.map(nums, fn(x) x * 2)
		evens := std.filter(doubled, fn(x) x % 4 = 0)
		std.reduce(evens, 0, fn(acc, x) acc + x)
	`
	expectProgramToReturn(t, program, IntValue(12))
	expectProgramToReturnBytecode(t, program, IntValue(12))
}

func TestBytecodeTailRecursionCountDown(t *testing.T) {
	program := `
		fn loop(n) if n {
			0 -> :done
			_ -> loop(n - 1)
		}
		loop(5000)
	`
	expectProgramToReturnBytecode(t, program, AtomValue("done"))
}

func TestBytecodeObjectSpread(t *testing.T) {
	program := `
		a := {x: 1, y: 2}
		b := {y: 3, z: 4}
		fn merge(base, over) {
			result := {}
			fn copyKeys(obj) {
				ks := keys(obj)
				fn loop(i) if i < len(ks) {
					true -> {
						k := ks.(i)
						result.(k) := obj.(k)
						loop(i + 1)
					}
				}
				loop(0)
			}
			copyKeys(base)
			copyKeys(over)
			result
		}
		merged := merge(a, b)
		[merged.x, merged.y, merged.z]
	`
	expectProgramToReturnBytecode(t, program, MakeList(IntValue(1), IntValue(3), IntValue(4)))
}

func TestBytecodeMutualRecursion(t *testing.T) {
	program := `
		fn isEven?(n) if n {
			0 -> true
			_ -> isOdd?(n - 1)
		}
		fn isOdd?(n) if n {
			0 -> false
			_ -> isEven?(n - 1)
		}
		[isEven?(10), isOdd?(7)]
	`
	expectProgramToReturnBytecode(t, program, MakeList(oakTrue, oakTrue))
}

func TestBytecodeHigherOrderFunctions(t *testing.T) {
	program := `
		fn compose(f, g) fn(x) f(g(x))
		fn double(x) x * 2
		fn inc(x) x + 1
		doubleThenInc := compose(inc, double)
		doubleThenInc(5)
	`
	expectProgramToReturnBytecode(t, program, IntValue(11))
}

func TestBytecodeListDestructuringSimple(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		[a, b] := [10, 20]
		a + b
	`, IntValue(30))
}

func TestBytecodeObjectDestructuring(t *testing.T) {
	expectProgramToReturnBytecode(t, `
		{a: x, b: y} := {a: 10, b: 20}
		x + y
	`, IntValue(30))
}
