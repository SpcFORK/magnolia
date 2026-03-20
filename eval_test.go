package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func expectProgramToReturn(t *testing.T, program string, expected Value) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	val, err := ctx.Eval(strings.NewReader(program))
	if err != nil {
		t.Errorf("Did not expect program to exit with error: %s", err.Error())
	}
	if val == nil {
		t.Errorf("Return value of program should not be nil")
	} else if !val.Eq(expected) {
		t.Errorf("Expected and returned values don't match: %s != %s",
			strconv.Quote(expected.String()),
			strconv.Quote(val.String()))
	}
}

func expectProgramToError(t *testing.T, program string) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	_, err := ctx.Eval(strings.NewReader(program))
	if err == nil {
		t.Errorf("Expected program to exit with an error")
	}
}

func TestEvalEmptyProgram(t *testing.T) {
	expectProgramToReturn(t, "", null)
	expectProgramToReturn(t, "   \n", null)
}

func TestCommentProgram(t *testing.T) {
	expectProgramToReturn(t, "// this is a comment", null)
	expectProgramToReturn(t, "// this is a comment\n", null)
}

func TestCommentInBinaryExpr(t *testing.T) {
	expectProgramToReturn(t, "1 + // this is a comment\n2", IntValue(3))
}

func TestCommentAndNewline(t *testing.T) {
	expectProgramToReturn(t, "1 + 2 // this is a comment\n", IntValue(3))
}

func TestIdentifierAfterComment(t *testing.T) {
	expectProgramToReturn(t, "x := 10 // this is a comment\nx + x", IntValue(20))
}

func TestEmptyLiteral(t *testing.T) {
	expectProgramToReturn(t, "_", empty)
}

func TestNullLiteral(t *testing.T) {
	expectProgramToReturn(t, "?", null)
}

func TestStringLiteral(t *testing.T) {
	expectProgramToReturn(t, "'Hello, World!\\n'", MakeString("Hello, World!\n"))
}

func TestQuotedStringLiteral(t *testing.T) {
	expectProgramToReturn(t, "'a\\'b'", MakeString("a'b"))
}

func TestStringLiteralOverflow(t *testing.T) {
	expectProgramToReturn(t, "'\\", MakeString(""))
	expectProgramToReturn(t, "'\\'", MakeString("'"))
	expectProgramToReturn(t, "'\\x'", MakeString("x"))
	expectProgramToReturn(t, "'\\x1'", MakeString("x1"))
	expectProgramToReturn(t, "'\\x1g'", MakeString("x1g"))
}

func TestHexStringLiteral(t *testing.T) {
	expectProgramToReturn(t, "'a\\x!'", MakeString("ax!"))
	expectProgramToReturn(t, "'a\\x1!'", MakeString("ax1!"))
	expectProgramToReturn(t, "'a\\x0a!'", MakeString("a\n!"))
	expectProgramToReturn(t, "'a\\x0A!'", MakeString("a\n!"))
	expectProgramToReturn(t, "'a\\x1z!'", MakeString("ax1z!"))

	nonAsciiStringValue := StringValue([]byte{0x98})
	expectProgramToReturn(t, "'\\x98'", &nonAsciiStringValue) // test out of ascii range
}

func TestIntegerLiteral(t *testing.T) {
	expectProgramToReturn(t, "64710", IntValue(64710))
}

func TestFloatLiteral(t *testing.T) {
	expectProgramToReturn(t, "100.0", FloatValue(100))
	expectProgramToReturn(t, "3.141592", FloatValue(3.141592))
}

func TestAtomLiteral(t *testing.T) {
	atomNames := []string{
		"_?", "if", "fn", "with", "true", "false", "_if", "not_found_404",
	}

	for _, atomName := range atomNames {
		expectProgramToReturn(t, ":"+atomName, AtomValue(atomName))
	}
}

func TestKeywordLikeAtomLiteral(t *testing.T) {
	expectProgramToReturn(t, ":if", AtomValue("if"))
	expectProgramToReturn(t, ":fn", AtomValue("fn"))
	expectProgramToReturn(t, ":with", AtomValue("with"))
	expectProgramToReturn(t, ":cs", AtomValue("cs"))
	expectProgramToReturn(t, ":true", AtomValue("true"))
	expectProgramToReturn(t, ":false", AtomValue("false"))
}

func TestListLiteral(t *testing.T) {
	expectProgramToReturn(t, `[1, [2, 'three'], :four]`, MakeList(
		IntValue(1),
		MakeList(
			IntValue(2),
			MakeString("three"),
		),
		AtomValue("four"),
	))
}

func TestObjectLiteral(t *testing.T) {
	expectProgramToReturn(t, `{ a: 'ay', :be: 200, 100: {('d' + 'i'): :dee } }`, ObjectValue{
		"a":  MakeString("ay"),
		"be": IntValue(200),
		"100": ObjectValue{
			"di": AtomValue("dee"),
		},
	})
}

func TestListStringify(t *testing.T) {
	expectProgramToReturn(t, `
		[
			string(:atomValue)
			string(3)
			string(2.51)
		]
	`, MakeList(
		MakeString("atomValue"),
		MakeString("3"),
		MakeString("2.51"),
	))
}

func TestObjectStringify(t *testing.T) {
	expectProgramToReturn(t, `
		x := {
			first: {}
			second: :two
			_third: {
				_fourth: 'four'
			}
		}
		x |> string()
	`, MakeString("{_third: {_fourth: 'four'}, first: {}, second: :two}"))
}

func TestFunctionDefAndCall(t *testing.T) {
	expectProgramToReturn(t, `fn getThree() { x := 4, 3 }, getThree()`, IntValue(3))
}

func TestFunctionDefWithEmpty(t *testing.T) {
	expectProgramToReturn(t, `fn getThird(_, _, third) third, getThird(1, 2, 3)`, IntValue(3))
}

func TestFunctionCreatesScope(t *testing.T) {
	expectProgramToReturn(t, `
	x := 3
	fn defineX x := 10
	defineX()
	x
	`, IntValue(3))
}

func TestBlockCreatesScope(t *testing.T) {
	expectProgramToReturn(t, `
	x := 3
	{ x := 10 }
	x
	`, IntValue(3))
}

func TestEmptyFunctionBody(t *testing.T) {
	expectProgramToReturn(t, `
	fn do {
		a: :bee
	}
	do()
	`, ObjectValue{
		"a": AtomValue("bee"),
	})
}

func TestObjectLiteralFunctionBody(t *testing.T) {
	expectProgramToReturn(t, `
	fn do {}
	do()
	`, null)
}

func TestLocalAssignment(t *testing.T) {
	expectProgramToReturn(t, `x := 100, y := 200, x`, IntValue(100))
}

func TestChainedLocalAssignment(t *testing.T) {
	expectProgramToReturn(t, `
	a := b := 10
	a + b
	`, IntValue(20))
}

func TestChainedNonlocalAssignment(t *testing.T) {
	expectProgramToReturn(t, `
	a := b := 0
	{
		a <- b <- 20
	}
	a + b
	`, IntValue(40))
}

func TestDestructureList(t *testing.T) {
	expectProgramToReturn(t, `
	list := [1, 2, 3]
	[a] := list
	[_, _, b, c] := list
	[a, b, c]
	`, MakeList(
		IntValue(1),
		IntValue(3),
		null,
	))
}

func TestDestructureObject(t *testing.T) {
	expectProgramToReturn(t, `
	obj := {
		a: 'ay'
		b: 'bee'
		12: 'see'
	}
	{a: a} := obj
	{:b: b, 10 + 2: see} := {'whatever': dee} := obj
	[a, b, see, dee]
	`, MakeList(
		MakeString("ay"),
		MakeString("bee"),
		MakeString("see"),
		null,
	))
}

func TestDestrctureToReassignList(t *testing.T) {
	expectProgramToReturn(t, `
	v := [:aa, :bbb]
	[v, w] := v
	v
	`, AtomValue("aa"))
}

func TestDestrctureToReassignObject(t *testing.T) {
	expectProgramToReturn(t, `
	a := {a: :aa, b: :bbb}
	{a: a} := a
	a
	`, AtomValue("aa"))
}

func TestUnderscoreVarNames(t *testing.T) {
	expectProgramToReturn(t, `
	_a := 'A'
	b_ := 'B'
	c_d := 'CD'
	_a + b_ + c_d
	`, MakeString("ABCD"))
}

func TestNonlocalAssignment(t *testing.T) {
	expectProgramToReturn(t, `
	x := 100
	y := 200
	fn do {
		x <- x + 100
		y := y + 100
	}
	do()
	x + y
	`, IntValue(400))
}

func TestPushToString(t *testing.T) {
	expectProgramToReturn(t, `
	s := 'hi'
	[s << 'world', s]
	`, MakeList(
		MakeString("hiworld"),
		MakeString("hiworld"),
	))
}

func TestPushToList(t *testing.T) {
	expectProgramToReturn(t, `
	arr := [:a]
	[arr << :b, arr]
	`, MakeList(
		MakeList(AtomValue("a"), AtomValue("b")),
		MakeList(AtomValue("a"), AtomValue("b")),
	))
}

func TestPushArrowPrecedence(t *testing.T) {
	expectProgramToReturn(t, `
	arr := [2] << 1 + 3
	arr << 10 << 20
	arr << x := 100
	`, MakeList(
		IntValue(2),
		IntValue(4),
		IntValue(10),
		IntValue(20),
		IntValue(100),
	))
}

func TestUnaryExpr(t *testing.T) {
	expectProgramToReturn(t, `!true`, oakFalse)
	expectProgramToReturn(t, `!(false | true)`, oakFalse)

	expectProgramToReturn(t, `-546`, IntValue(-546))
	expectProgramToReturn(t, `-3.250`, FloatValue(-3.25))
}

func TestUnaryBindToProperty(t *testing.T) {
	expectProgramToReturn(t, `!!false`, oakFalse)
	expectProgramToReturn(t, `--3`, IntValue(3))
	expectProgramToReturn(t, `
	obj := {k: false, n: 10}
	[!obj.k, -obj.n]
	`, MakeList(
		oakTrue,
		IntValue(-10),
	))
}

func TestBasicBinaryExpr(t *testing.T) {
	expectProgramToReturn(t, `2 * 3 + 1`, IntValue(7))
	expectProgramToReturn(t, `1 + 2 * 3`, IntValue(7))
}

func TestFloatDivide(t *testing.T) {
	expectProgramToReturn(t, "10 / 4", FloatValue(2.5))
}

func TestOrderedBinaryExpr(t *testing.T) {
	expectProgramToReturn(t, `-1.5 + -3.5 - 5 / 5 * 2`, FloatValue(-7))
	expectProgramToReturn(t, `(-1.5 + -3.5 - 5) / 5 * 2`, FloatValue(-4))
}

func TestBinaryExprWithParens(t *testing.T) {
	expectProgramToReturn(t, `(1 + 2) / 3 - 1 + (10 + (20 / 5)) % 3`, FloatValue(2))
}

func TestLongBinaryExprWithPrecedence(t *testing.T) {
	expectProgramToReturn(t, `x := 1 + 2 * 3 + 4 / 2 + 10 % 4, x % 5 + x`, FloatValue(12))
}

func TestBinaryExprWithComplexTerms(t *testing.T) {
	expectProgramToReturn(t, `
	fn double(n) 2 * n
	fn decrement(n) n - 1
	double(10) + if decrement(10) { 9 -> 2, _ -> 1 } + 8
	`, IntValue(30))
}

func TestBinaryExprWithinComplexTermsWithinBinaryExpr(t *testing.T) {
	expectProgramToReturn(t, `
	fn inc(n) n + 1
	2 * inc(3 + 4)
	`, IntValue(16))
}

func TestStringCompare(t *testing.T) {
	expectProgramToReturn(t, `
	[
		// empty string
		'long string' > ''
		// length comparison
		'hi' < 'hiworld'
		// lexicographical
		'heels' < 'hi'
		// space
		'abc' > ' abc'
		// equality
		'abc' = 'abc'
	]
	`, MakeList(oakTrue, oakTrue, oakTrue, oakTrue, oakTrue))
}

func TestShallowVsDeepListEquality(t *testing.T) {
	expectProgramToReturn(t, `
		a := [1, 2]
		b := [1, 2]
		[
			a = b
			a == b
			a = a
		]
	`, MakeList(oakFalse, oakTrue, oakTrue))
}

func TestShallowVsDeepObjectEquality(t *testing.T) {
	expectProgramToReturn(t, `
		a := {x: 1,}
		b := {x: 1,}
		[
			a = b
			a == b
			a = a
		]
	`, MakeList(oakFalse, oakTrue, oakTrue))
}

func TestAndOperator(t *testing.T) {
	expectProgramToReturn(t, `
	[
		true & true
		true & false
		false & true
		false & false
		'abcd' & '    '
		'    ' & 'wxyz'
		'abcdef' & '   '
		'   ' & 'abcdef'
	]
	`, MakeList(
		oakTrue,
		oakFalse,
		oakFalse,
		oakFalse,
		MakeString("    "),
		MakeString("    "),
		MakeString("   \x00\x00\x00"),
		MakeString("   \x00\x00\x00"),
	))
}

func TestXorOperator(t *testing.T) {
	expectProgramToReturn(t, `
	[
		true ^ true
		true ^ false
		false ^ true
		false ^ false
		'ABCD' ^ '    '
		'    ' ^ 'WXYZ'
		'ABCDEF' ^ '   '
		'   ' ^ 'ABCDEF'
	]
	`, MakeList(
		oakFalse,
		oakTrue,
		oakTrue,
		oakFalse,
		MakeString("abcd"),
		MakeString("wxyz"),
		MakeString("abcDEF"),
		MakeString("abcDEF"),
	))
}

func TestOrOperator(t *testing.T) {
	expectProgramToReturn(t, `
	[
		true | true
		true | false
		false | true
		false | false
		'ABCD' | '    '
		'    ' | 'WXYZ'
		'ABCDEF' | '   '
		'   ' | 'ABCDEF'
	]
	`, MakeList(
		oakTrue,
		oakTrue,
		oakTrue,
		oakFalse,
		MakeString("abcd"),
		MakeString("wxyz"),
		MakeString("abcDEF"),
		MakeString("abcDEF"),
	))
}

func TestShortCircuitingAnd(t *testing.T) {
	expectProgramToReturn(t, `
	x := 3
	[
		false & 2
		false & { x <- 10, true }
		x
	]
	`, MakeList(oakFalse, oakFalse, IntValue(3)))
}

func TestShortCircuitingOr(t *testing.T) {
	expectProgramToReturn(t, `
	x := 3
	[
		true | 2
		true | { x <- 10, false }
		x
	]
	`, MakeList(oakTrue, oakTrue, IntValue(3)))
}

func TestEmptyIfExpr(t *testing.T) {
	expectProgramToReturn(t, `if 100 {}`, null)
}

func TestEmptyIfCondition(t *testing.T) {
	expectProgramToReturn(t, `if {
		false -> 10
		true -> 20
		_ -> 30
	}`, IntValue(20))
}

func TestBasicIfExpr(t *testing.T) {
	expectProgramToReturn(t, `if 2 * 2 {
		? -> 100
		{ a: 'b' } -> 200
		5 -> 'five'
		4 -> 'four'
	}`, MakeString("four"))
}

func TestIfExprWithMultiTarget(t *testing.T) {
	for _, i := range []int{11, 12, 13} {
		expectProgramToReturn(t, fmt.Sprintf(`if %d {
			10 -> :wrong
			11, 5 + 7, { 10 + 3 } -> :right
			_ -> :wrong2
		}`, i), AtomValue("right"))
	}
}

func TestNestedIfExpr(t *testing.T) {
	expectProgramToReturn(t, `if 3 {
		10, if true {
			true -> 10
			_ -> 3
		} -> 'hi'
		100, 3 -> 'hello'
	}`, MakeString("hello"))
}

func TestIfExprWithEmpty(t *testing.T) {
	expectProgramToReturn(t, `if 10 + 2 {
		12 -> 'twelve'
		_ -> 'wrong'
	}`, MakeString("twelve"))
}

func TestIfExprWithAssignmentCond(t *testing.T) {
	expectProgramToReturn(t, `if x := 2 + 4 {
		6 -> x * x
		_ -> x
	}`, IntValue(36))
}

func TestIfExprInFunction(t *testing.T) {
	expectProgramToReturn(t, `
	fn even?(n) if n % 2 {
		0 -> true
		_ -> false
	}
	even?(100)
	`, oakTrue)
}

func TestComplexIfExprTarget(t *testing.T) {
	expectProgramToReturn(t, `
	fn double(n) 2 * n
	fn xyz(n) if n {
		1 + 2 -> :abc
		2 * double(3) -> :xyz
		_ -> false
	}
	[xyz(3), xyz(12), xyz(24)]
	`, MakeList(
		AtomValue("abc"),
		AtomValue("xyz"),
		oakFalse,
	))
}

func TestBasicWithExpr(t *testing.T) {
	expectProgramToReturn(t, `fn add(a, b) { a + b }, with add(10) 40`, IntValue(50))
}

func TestWithExprWithCallback(t *testing.T) {
	expectProgramToReturn(t, `fn applyThrice(x, f) f(f(f(x))), with applyThrice(10) fn(n) n + 1`, IntValue(13))
}

func TestRecursiveFunction(t *testing.T) {
	expectProgramToReturn(t, `
	fn times(n, f) {
		fn sub(i) if i {
			n -> ?
			_ -> {
				f(i)
				sub(i + 1)
			}
		}
		sub(0)
	}

	counter := 0
	with times(10) fn(i) {
		counter <- counter + i * 10
	}
	counter
	`, IntValue(450))
}

func TestRecursiveFunctionOnList(t *testing.T) {
	expectProgramToReturn(t, `
	fn each(list, f) {
		fn sub(i) if i {
			len(list) -> ?
			_ -> {
				f(list.(i))
				sub(i + 1)
			}
		}
		sub(0)
	}

	sum := 0
	list := [1, 2, 3, 4, 5]
	with each(list) fn(it) {
		sum <- sum + it
	}
	sum
	`, IntValue(15))
}

func TestCurriedFunctionDef(t *testing.T) {
	expectProgramToReturn(t, `
	addThree := fn(a) fn(b) fn(c) {
		a + b + c
	}

	almost := addThree(15)(20)
	almost(8)
	`, IntValue(15+20+8))
}

// string ops

func TestStringAccess(t *testing.T) {
	expectProgramToReturn(t, `
	s := 'Hello, World!'
	[
		s.0 + s.2
		s.-2
		s.15
	]
	`, MakeList(MakeString("Hl"), null, null))
}

func TestStringAssign(t *testing.T) {
	expectProgramToReturn(t, `
	s := {
		payload: 'Magnolia'
	}
	t := s.payload
	[s.payload.3 := 'pie', t]
	`, MakeList(
		MakeString("Magpieia"),
		MakeString("Magpieia"),
	))
}

func TestStringAppendByPush(t *testing.T) {
	expectProgramToReturn(t, `
	s := {
		payload: 'Oak'
	}
	[s.payload << ' language', s.payload]
	`, MakeList(
		MakeString("Oak language"),
		MakeString("Oak language"),
	))
}

func TestStringAppendByAssign(t *testing.T) {
	expectProgramToReturn(t, `
	s := {
		payload: 'Oak'
	}
	t := s.payload
	[s.payload.(len(s.payload)) := ' language', s.payload]
	`, MakeList(
		MakeString("Oak language"),
		MakeString("Oak language"),
	))
}

// list ops

func TestListAccess(t *testing.T) {
	expectProgramToReturn(t, `
	s := [1, 2, 3, 4, 5]
	[
		s.0 + s.3
		s.-2
		s.15
	]
	`, MakeList(IntValue(5), null, null))
}

func TestListAssign(t *testing.T) {
	result := MakeList(
		IntValue(1),
		IntValue(2),
		MakeString("three"),
		IntValue(4),
	)

	expectProgramToReturn(t, `
	s := {
		numbers: [1, 2, 3, 4]
	}
	t := s.numbers
	[s.numbers.2 := 'three', t]
	`, MakeList(result, result))
}

func TestListAppendByPush(t *testing.T) {
	result := MakeList(
		IntValue(1),
		IntValue(2),
		IntValue(3),
		IntValue(4),
		IntValue(100),
	)

	expectProgramToReturn(t, `
	s := {
		numbers: [1, 2, 3, 4]
	}
	t := s.numbers
	[s.numbers << 100, t]
	`, MakeList(result, result))
}

func TestListAppendByAssign(t *testing.T) {
	result := MakeList(
		IntValue(1),
		IntValue(2),
		IntValue(3),
		IntValue(4),
		IntValue(100),
	)

	expectProgramToReturn(t, `
	s := {
		numbers: [1, 2, 3, 4]
	}
	[s.numbers.(len(s.numbers)) := 100, s.numbers]
	`, MakeList(result, result))
}

// object ops

func TestObjectAccess(t *testing.T) {
	expectProgramToReturn(t, `
	obj := {
		a: 'ay'
		b: 'bee'
		c: ['see', {
			d: 'd'
		}]
	}
	[
		obj.c.(1).:d
		obj.c.(1).(:d)
	]
	`, MakeList(
		MakeString("d"),
		MakeString("d"),
	))
}

func TestObjectAssign(t *testing.T) {
	expectProgramToReturn(t, `
	obj := {
		a: 'ay'
		b: 'bee'
		c: ['see', {
			d: 'd'
		}]
	}
	[
		obj.c.(1).:e := 'hello_e'
		obj.c.(1).(:f) := 'hello_f'
		obj.c
	]
	`, MakeList(
		ObjectValue{
			"d": MakeString("d"),
			"e": MakeString("hello_e"),
			"f": MakeString("hello_f"),
		},
		ObjectValue{
			"d": MakeString("d"),
			"e": MakeString("hello_e"),
			"f": MakeString("hello_f"),
		},
		MakeList(MakeString("see"), ObjectValue{
			"d": MakeString("d"),
			"e": MakeString("hello_e"),
			"f": MakeString("hello_f"),
		}),
	))
}

func TestObjectDelete(t *testing.T) {
	expectProgramToReturn(t, `
	obj := {
		a: 'ay'
		b: 'bee'
		c: {
			d: 'dee'
			e: 'ee'
		}
	}
	[
		obj.nonexistent := _
		obj.b := { 1, 2, _ }
		obj
		obj.c.d := _
		obj.c
	]
	`, MakeList(
		ObjectValue{
			"a": MakeString("ay"),
			"c": ObjectValue{
				"e": MakeString("ee"),
			},
		},
		ObjectValue{
			"a": MakeString("ay"),
			"c": ObjectValue{
				"e": MakeString("ee"),
			},
		},
		ObjectValue{
			"a": MakeString("ay"),
			"c": ObjectValue{
				"e": MakeString("ee"),
			},
		},
		ObjectValue{
			"e": MakeString("ee"),
		},
		ObjectValue{
			"e": MakeString("ee"),
		},
	))
}

func TestSinglePipe(t *testing.T) {
	expectProgramToReturn(t, `
	fn append(a, b) a + b
	'hello' |> append('world')
	`, MakeString("helloworld"))
}

func TestMultiPipe(t *testing.T) {
	expectProgramToReturn(t, `
	fn append(a, b) a + b
	'hello' |> append('world') |> append('!')
	`, MakeString("helloworld!"))
}

func TestComplexPipe(t *testing.T) {
	expectProgramToReturn(t, `
	lib := {
		add1: fn(n) n + 1
		double: fn(n) 2 * n
	}
	fn getAdder(env) { env.add1 }
	100 |> lib.add1() |> lib.double() |> getAdder(lib)()
	`, IntValue(203))
}

func TestPipeWithExpr(t *testing.T) {
	expectProgramToReturn(t, `
	fn add(a, b) a + b
	fn double(n) 2 * n
	fn apply(x, f) f(x)

	10 |> add(20) |> with apply() fn(n) n |> double() + 40
	`, IntValue(100))
}

func TestExtraArgs(t *testing.T) {
	expectProgramToReturn(t, `
	fn getExtra(a, b, c) {
		[b, c]
	}
	getExtra(1, ?)
	`, MakeList(null, null))
}

func TestRestArgs(t *testing.T) {
	expectProgramToReturn(t, `
	fn getRest(first, rest...) {
		rest
	}
	getRest(1, 2, 3, 4, 5)
	`, MakeList(
		IntValue(2),
		IntValue(3),
		IntValue(4),
		IntValue(5),
	))
}

// bitwise operations for integers

func TestBitwiseLeftShift(t *testing.T) {
	expectProgramToReturn(t, `
	[
		1 << 0
		1 << 1
		1 << 2
		1 << 3
		5 << 2
		8 << 4
	]
	`, MakeList(
		IntValue(1),
		IntValue(2),
		IntValue(4),
		IntValue(8),
		IntValue(20),
		IntValue(128),
	))
}

func TestBitwiseRightShift(t *testing.T) {
	expectProgramToReturn(t, `
	[
		8 >> 0
		8 >> 1
		8 >> 2
		8 >> 3
		20 >> 2
		128 >> 4
	]
	`, MakeList(
		IntValue(8),
		IntValue(4),
		IntValue(2),
		IntValue(1),
		IntValue(5),
		IntValue(8),
	))
}

func TestBitwiseAnd(t *testing.T) {
	expectProgramToReturn(t, `
	[
		12 & 10
		15 & 7
		255 & 15
		0 & 255
		255 & 255
	]
	`, MakeList(
		IntValue(8),
		IntValue(7),
		IntValue(15),
		IntValue(0),
		IntValue(255),
	))
}

func TestBitwiseOr(t *testing.T) {
	expectProgramToReturn(t, `
	[
		12 | 10
		8 | 4
		1 | 2
		0 | 255
		15 | 240
	]
	`, MakeList(
		IntValue(14),
		IntValue(12),
		IntValue(3),
		IntValue(255),
		IntValue(255),
	))
}

func TestBitwiseXor(t *testing.T) {
	expectProgramToReturn(t, `
	[
		12 ^ 10
		15 ^ 7
		255 ^ 0
		255 ^ 255
		1 ^ 1
		1 ^ 0
	]
	`, MakeList(
		IntValue(6),
		IntValue(8),
		IntValue(255),
		IntValue(0),
		IntValue(0),
		IntValue(1),
	))
}

func TestBitwiseNot(t *testing.T) {
	expectProgramToReturn(t, `
	[
		~0
		~1
		~(-1)
		~(~5)
	]
	`, MakeList(
		IntValue(-1),
		IntValue(-2),
		IntValue(0),
		IntValue(5),
	))
}

func TestBitwiseCombined(t *testing.T) {
	expectProgramToReturn(t, `
	[
		(1 << 4) | (1 << 2)
		(5 << 1) & 15
		~(~10)
		(12 & 10) | (8 ^ 4)
	]
	`, MakeList(
		IntValue(20),
		IntValue(10),
		IntValue(10),
		IntValue(12),
	))
}

func TestBitwiseShiftPrecedence(t *testing.T) {
	expectProgramToReturn(t, `
	[
		1 + 2 << 1
		1 << 2 + 1
		10 - 2 >> 1
	]
	`, MakeList(
		IntValue(6),
		IntValue(8),
		IntValue(4),
	))
}

// power operator tests

func TestPowerInt(t *testing.T) {
	expectProgramToReturn(t, `
	[
		2 ** 0
		2 ** 1
		2 ** 2
		2 ** 3
		2 ** 10
		5 ** 2
	]
	`, MakeList(
		IntValue(1),
		IntValue(2),
		IntValue(4),
		IntValue(8),
		IntValue(1024),
		IntValue(25),
	))
}

func TestPowerFloat(t *testing.T) {
	expectProgramToReturn(t, `
	[
		2.0 ** 0.0
		2.0 ** 1.0
		2.0 ** 0.5
		4.0 ** 0.5
		10.0 ** 2.0
	]
	`, MakeList(
		FloatValue(1.0),
		FloatValue(2.0),
		FloatValue(math.Sqrt(2)),
		FloatValue(2.0),
		FloatValue(100.0),
	))
}

func TestPowerMixed(t *testing.T) {
	expectProgramToReturn(t, `
	[
		2 ** 2.0
		2.0 ** 2
		3 ** 0.5
	]
	`, MakeList(
		FloatValue(4.0),
		FloatValue(4.0),
		FloatValue(math.Sqrt(3)),
	))
}

func TestPowerNegativeExponent(t *testing.T) {
	expectProgramToReturn(t, `
	[
		2 ** -1
		2 ** -2
		10.0 ** -1.0
	]
	`, MakeList(
		FloatValue(0.5),
		FloatValue(0.25),
		FloatValue(0.1),
	))
}

func TestPowerPrecedence(t *testing.T) {
	expectProgramToReturn(t, `
	[
		2 + 3 ** 2
		2 * 3 ** 2
		2 ** 3 * 2
		10 - 2 ** 2
	]
	`, MakeList(
		IntValue(11),
		IntValue(18),
		IntValue(16),
		IntValue(6),
	))
}

// Virtual interpreter tests

func TestVirtualInterpreterCreation(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        type(Virtual.createVM())
        `, AtomValue("object"))
}

func TestVirtualStandardVM(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        type(vm.globalScope.print)
        `, AtomValue("function"))
}

func TestVirtualContextCreation(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        type(Virtual.createVirtualContext())
        `, AtomValue("object"))
}

func TestVirtualLiteralEvaluation(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        [
                vm.run('?', {})
                vm.run('_', {})
                vm.run('42', {})
                vm.run('3.5', {})
                vm.run('\'hello\'', {})
                vm.run(':test', {})
                vm.run('true', {})
                vm.run('false', {})
        ]
        `, MakeList(
		null,
		empty,
		IntValue(42),
		FloatValue(3.5),
		MakeString("hello"),
		AtomValue("test"),
		oakTrue,
		oakFalse,
	))
}

func TestVirtualArithmeticAndComparison(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        [
                vm.run('2 + 3', {})
                vm.run('3 * 4', {})
                vm.run('17 % 5', {})
                vm.run('3.5 + 2.1', {})
                vm.run('5 = 5', {})
                vm.run('5 != 3', {})
                vm.run('4 >= 3', {})
        ]
        `, MakeList(
		IntValue(5),
		IntValue(12),
		IntValue(2),
		FloatValue(5.6),
		oakTrue,
		oakTrue,
		oakTrue,
	))
}

func TestVirtualLogicalAndUnaryOperations(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        [
                vm.run('true & false', {})
                vm.run('true | false', {})
                vm.run('!false', {})
                vm.run('-5', {})
        ]
        `, MakeList(
		oakFalse,
		oakTrue,
		oakTrue,
		IntValue(-5),
	))
}

func TestVirtualVariableOperations(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        ctx := Virtual.createVirtualContext()
        ctx.define('x', 42)
        ctx.evalExpr('x')
        `, IntValue(42))
}

func TestVirtualVariableAssignment(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        ctx := Virtual.createVirtualContext()
        ctx.define('y', 10)
        ctx.evalExpr('y <- 20')
        ctx.evalExpr('y')
        `, IntValue(20))
}

func TestVirtualFunctionDefinition(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        ctx := Virtual.createVirtualContext()
        ctx.defineFunction('double', ['n'], {
                type: :binary
                op: :times
                left: { type: :identifier, val: 'n' }
                right: { type: :int, val: 2 }
        })
        ctx.evalExpr('double(5)')
        `, IntValue(10))
}

func TestVirtualFunctionWithClosure(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        ctx := Virtual.createVirtualContext()
        ctx.define('multiplier', 3)
        ctx.defineFunction('multiply', ['x'], {
                type: :binary
                op: :times
                left: { type: :identifier, val: 'multiplier' }
                right: { type: :identifier, val: 'x' }
        })
        ctx.evalExpr('multiply(4)')
        `, IntValue(12))
}

func TestVirtualIfExpressions(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        [
                vm.run('if true { true -> 42, _ -> 24, }', {})
                vm.run('if false { false -> 24, _ -> 42, }', {})
                vm.run('if 5 > 3 { true -> \'yes\', _ -> \'no\', }', {})
        ]
        `, MakeList(
		IntValue(42),
		IntValue(24),
		MakeString("yes"),
	))
}

func TestVirtualDataStructures(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        [
                vm.run('[1, 2, 3]', {})
                vm.run('{a: 1, b: 2}', {})
        ]
        `, MakeList(
		MakeList(IntValue(1), IntValue(2), IntValue(3)),
		ObjectValue{"a": IntValue(1), "b": IntValue(2)},
	))
}

func TestVirtualStandardLibrary(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        [
                vm.run('type(42)', {})
                vm.run('len([1,2,3,4])', {})
                vm.run('len(\'hello\')', {})
                vm.run('len(keys({a:1, b:2}))', {})
                vm.run('string(123)', {})
                vm.run('int(\'456\')', {})
                vm.run('float(\'3.14\')', {})
        ]
        `, MakeList(
		AtomValue("int"),
		IntValue(4),
		IntValue(5),
		IntValue(2),
		MakeString("123"),
		IntValue(456),
		FloatValue(3.14),
	))
}

func TestVirtualUnsupportedFeaturesReturnErrors(t *testing.T) {
	expectProgramToReturn(t, `
        Virtual := import('Virtual')
        vm := Virtual.createStandardVM()
        [
                vm.run('2 ** 3', {}).type
                vm.run('if true { 42 } else { 24 }', {}).type
        ]
        `, MakeList(
		AtomValue("error"),
		AtomValue("error"),
	))
}

func TestVirtualTokenConstructors(t *testing.T) {
	expectProgramToReturn(t, `
		VirtualToken := import('VirtualToken')
		[
			VirtualToken.Comma()
			VirtualToken.Identifier('myVar', VirtualToken.at(10, 2, 7))
			VirtualToken.NumberLiteral('42')
			VirtualToken.Comment('note', [3, 1, 9])
		]
	`, MakeList(
		ObjectValue{
			"type": AtomValue("comma"),
			"val":  null,
			"pos":  MakeList(IntValue(0), IntValue(1), IntValue(1)),
		},
		ObjectValue{
			"type": AtomValue("identifier"),
			"val":  MakeString("myVar"),
			"pos":  MakeList(IntValue(10), IntValue(2), IntValue(7)),
		},
		ObjectValue{
			"type": AtomValue("numberLiteral"),
			"val":  MakeString("42"),
			"pos":  MakeList(IntValue(0), IntValue(1), IntValue(1)),
		},
		ObjectValue{
			"type": AtomValue("comment"),
			"val":  MakeString("note"),
			"pos":  MakeList(IntValue(3), IntValue(1), IntValue(9)),
		},
	))
}

func TestSyscallFunctionExists(t *testing.T) {
	expectProgramToReturn(t, `
		result := syscall(-1) // Invalid syscall number
		result.type
	`, AtomValue("error"))
}

func TestUTF16Builtin(t *testing.T) {
	expectProgramToReturn(t, `
		buf := utf16('Az')
		[
			len(buf)
			codepoint(buf.0)
			codepoint(buf.1)
			codepoint(buf.2)
			codepoint(buf.3)
			codepoint(buf.4)
			codepoint(buf.5)
		]
	`, MakeList(
		IntValue(6),
		IntValue(65),
		IntValue(0),
		IntValue(122),
		IntValue(0),
		IntValue(0),
		IntValue(0),
	))
}

func TestSysprocMissingLibraryReturnsError(t *testing.T) {
	expectProgramToReturn(t, `
		sysproc('definitely-missing-oak-syscall.dll', 'MissingProc').type
	`, AtomValue("error"))
}

func TestVirtualSyscallBuiltinsExist(t *testing.T) {
	expectProgramToReturn(t, `
		Virtual := import('Virtual')
		vm := Virtual.createStandardVM()
		[
			type(vm.globalScope.bits)
			type(vm.globalScope.addr)
			type(vm.globalScope.memread)
			type(vm.globalScope.memwrite)
			type(vm.globalScope.utf16)
			type(vm.globalScope.sysproc)
			type(vm.globalScope.syscall)
		]
	`, MakeList(
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
	))
}

func TestSyscallCanInvokeResolvedProcOnWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("sysproc-backed syscall test is Windows-specific")
	}

	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()
	val, err := ctx.Eval(strings.NewReader(`
		proc := sysproc('kernel32.dll', 'GetCurrentProcessId')
		syscall(proc)
	`))
	if err != nil {
		t.Fatalf("Did not expect program to exit with error: %s", err.Error())
	}

	result, ok := val.(ObjectValue)
	if !ok {
		t.Fatalf("Expected syscall result object, got %T", val)
	}

	if typeVal, ok := result["type"].(AtomValue); !ok || typeVal != AtomValue("success") {
		t.Fatalf("Expected syscall success result, got %v", result["type"])
	}

	pid, ok := result["r1"].(IntValue)
	if !ok {
		t.Fatalf("Expected syscall result to contain integer r1, got %T", result["r1"])
	}

	if pid != IntValue(os.Getpid()) {
		t.Fatalf("Expected PID %d, got %d", os.Getpid(), pid)
	}
}

func TestMakeChanBufferedRoundTrip(t *testing.T) {
	expectProgramToReturn(t, `
		ch := make_chan(1)
		chan_send(ch, 42)
		chan_recv(ch).data
	`, IntValue(42))
}

func TestGoBuiltinCoordinatesOverChannel(t *testing.T) {
	expectProgramToReturn(t, `
		ch := make_chan()
		go(fn {
			chan_send(ch, :ready)
		})
		chan_recv(ch).data
	`, AtomValue("ready"))
}

func TestGoRejectsBuiltinTarget(t *testing.T) {
	expectProgramToError(t, `go(wait, 0)`)
}

func TestChanRecvAsyncCallback(t *testing.T) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()

	_, err := ctx.Eval(strings.NewReader(`
		result := ?
		ch := make_chan()
		with chan_recv(ch) fn(evt) {
			result <- evt.data
		}
		go(fn {
			chan_send(ch, 99)
		})
	`))
	if err != nil {
		t.Fatalf("Did not expect program to exit with error: %s", err.Error())
	}

	ctx.Wait()

	result, scopeErr := ctx.scope.get("result")
	if scopeErr != nil {
		t.Fatalf("Could not read result from scope: %s", scopeErr.Error())
	}

	if !result.Eq(IntValue(99)) {
		t.Fatalf("Expected async channel callback to store 99, got %s", result)
	}
}

func TestGoRuntimeMetadataBuiltins(t *testing.T) {
	expectProgramToReturn(t, `
		[
			type(___runtime_go_version())
			___runtime_sys_info().os
			___runtime_sys_info().arch
			type(___runtime_sys_info().cpus)
		]
	`, MakeList(
		AtomValue("string"),
		MakeString(runtime.GOOS),
		MakeString(runtime.GOARCH),
		AtomValue("int"),
	))
}

func TestClassConstructorSugar(t *testing.T) {
	expectProgramToReturn(t, `
		cs Pair(left, right) {
			{
				left: left
				right: right
			}
		}
		Pair(1, 2).left + Pair(1, 2).right
	`, IntValue(3))
}

func TestClassConstructorWithoutArgs(t *testing.T) {
	expectProgramToReturn(t, `
		cs Empty {
			{}
		}
		type(Empty())
	`, AtomValue("object"))
}

func TestClassConstructorEmptyBodyActsLikeEmptyBlock(t *testing.T) {
	expectProgramToReturn(t, `
		cs Empty {}
		Empty()
	`, null)
}

func TestClassSupportsMultipleParentsAndStaticMembers(t *testing.T) {
	expectProgramToReturn(t, `
		cs Parent1 { a := 2, b := 1 }
		cs Parent2 { c := 3, b := 2 }

		cs Hi(make) {
			testStaticVar := 2
			fn testStaticFn {}

			(Parent1, Parent2) -> {}
		}

		[
			Hi('sedan').a
			Hi('sedan').b
			Hi('sedan').c
			Hi.testStaticVar
			type(Hi.testStaticFn)
		]
	`, MakeList(
		IntValue(2),
		IntValue(2),
		IntValue(3),
		IntValue(2),
		AtomValue("function"),
	))
}

func TestClassMatchBuiltin(t *testing.T) {
	expectProgramToReturn(t, `
		cs Alpha {}
		cs Beta {}

		[
			csof(Alpha, Alpha)
			csof(Alpha, Beta)
			csof(Alpha, :Alpha)
			csof(:Alpha, Alpha)
			csof(:Alpha, Beta)
			csof(1, :Alpha)
			csof(Alpha)
			csof(1)
		]
	`, MakeList(
		BoolValue(true),
		BoolValue(false),
		BoolValue(true),
		BoolValue(true),
		BoolValue(false),
		BoolValue(false),
		BoolValue(true),
		BoolValue(false),
	))
}

func TestBitsBuiltinRoundTrip(t *testing.T) {
	expectProgramToReturn(t, `
		bits(bits([65, 66, 67]))
	`, MakeList(
		IntValue(65),
		IntValue(66),
		IntValue(67),
	))
}

func TestMemReadWriteViaAddress(t *testing.T) {
	expectProgramToReturn(t, `
		buf := bits([65, 66, 67])
		ptr := addr(buf)
		memwrite(ptr + 1, bits([90]))
		bits(memread(ptr, 3))
	`, MakeList(
		IntValue(65),
		IntValue(90),
		IntValue(67),
	))
}

func TestMemReadWriteViaPointerBuiltin(t *testing.T) {
	expectProgramToReturn(t, `
			buf := bits([65, 66, 67])
			ptr := pointer(addr(buf))
			memwrite(ptr + 1, bits([90]))
			bits(memread(ptr, 3))
		`, MakeList(
		IntValue(65),
		IntValue(90),
		IntValue(67),
	))
}
func TestPointerBuiltinAndArithmetic(t *testing.T) {
	expectProgramToReturn(t, `[type(pointer(0)), pointer(0) == 0, pointer(5) == pointer(5), pointer(100) + 5 == pointer(105), 5 + pointer(100) == pointer(105), pointer(105) - 5 == pointer(100), pointer(20) > pointer(10), pointer(20) < pointer(10) == false]`, MakeList(
		AtomValue("pointer"),
		BoolValue(true),
		BoolValue(true),
		BoolValue(true),
		BoolValue(true),
		BoolValue(true),
		BoolValue(true),
		BoolValue(true),
	))
}

func TestIntFromPointerBuiltin(t *testing.T) {
	expectProgramToReturn(t, `int(pointer(123))`, IntValue(123))
}

func TestNameBuiltinAndAtomPointerRefs(t *testing.T) {
	expectProgramToReturn(t, `
		cs Alpha {}
		p := pointer(:Alpha)
		memwrite(:slot, :hello)
		q := pointer(:slot)
		[
			name(Alpha)
			name(:Beta)
			type(p)
			name(p)
			name(q)
			memwrite(:slot, [65, 66, 67])
			name(pointer(:slot))
		]
	`, MakeList(
		AtomValue("Alpha"),
		AtomValue("Beta"),
		AtomValue("pointer"),
		AtomValue("Alpha"),
		AtomValue("hello"),
		IntValue(3),
		AtomValue("ABC"),
	))
}

func TestBuildLibraryParseIncludesFromString(t *testing.T) {
	expectProgramToReturn(t, `
		std := import('std')
		build := import('build')
		build.parseIncludes('alpha:./one,beta:./two,third') |> std.map(:name)
	`, MakeList(
		MakeString("alpha"),
		MakeString("beta"),
		MakeString("third"),
	))
}

func TestBuildLibraryParseIncludesFromList(t *testing.T) {
	expectProgramToReturn(t, `
		build := import('build')
		str := import('str')
		parsed := build.parseIncludes(['foo:./x', { name: 'bar', path: '/tmp/bar.oak', }])
		[
			parsed.0.name,
			str.endsWith?(parsed.0.path, '.oak'),
			parsed.1.name,
			parsed.1.path,
			type(build.run)
		]
	`, MakeList(
		MakeString("foo"),
		BoolValue(true),
		MakeString("bar"),
		MakeString("/tmp/bar.oak"),
		AtomValue("function"),
	))
}

func TestBuildLibraryParseIncludeKeepsExplicitExtension(t *testing.T) {
	expectProgramToReturn(t, `
		build := import('build')
		str := import('str')
		parsed := build.parseInclude('std.test:test/std.test.oak')
		[
			parsed.name
			str.endsWith?(parsed.path, 'test/std.test.oak')
		]
	`, MakeList(
		MakeString("std.test"),
		BoolValue(true),
	))
}

func TestBuildWebBundleIncludesStaticStdImport(t *testing.T) {
	expectProgramToReturn(t, `
		buildImports := import('build-imports')
		moduleNodes := {}
		collected := []
		
		buildImports.addImportsFromSource(
			'/entry.oak'
			''
			moduleNodes
			{}
			false
			fn(_) false
			fn(_) ?
			fn(name, _) '/deps/' + name
			fn(_) '/root'
			fn(path) collected << path
			fn(_, _) [{
				type: :assignment
				local?: true
				left: { type: :identifier, val: 'std' }
				right: {
					type: :fnCall
					function: { type: :identifier, val: 'import' }
					args: [{ type: :string, val: 'std' }]
					restArg: ?
				}
			}]
		)
		collected
	`, MakeList(
		MakeString("/deps/std.oak"),
	))
}

func TestPackLibraryBuildArgs(t *testing.T) {
	expectProgramToReturn(t, `
		pack := import('pack')
		[
			pack.buildArgs('main.oak', 'bundle.oak', ?)
			pack.buildArgs('main.oak', 'bundle.oak', 'dyn:./dyn')
			type(pack.run)
		]
	`, MakeList(
		MakeList(
			MakeString("build"),
			MakeString("--entry"),
			MakeString("main.oak"),
			MakeString("--output"),
			MakeString("bundle.oak"),
		),
		MakeList(
			MakeString("build"),
			MakeString("--entry"),
			MakeString("main.oak"),
			MakeString("--output"),
			MakeString("bundle.oak"),
			MakeString("--include"),
			MakeString("dyn:./dyn"),
		),
		AtomValue("function"),
	))
}

func TestImportSupportsAlternativeModuleExtensions(t *testing.T) {
	tmpDir := t.TempDir()

	if err := os.WriteFile(filepath.Join(tmpDir, "mod_ok.ok"), []byte("value := 11"), 0o644); err != nil {
		t.Fatalf("Could not write .ok module: %s", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "mod_mag.mag"), []byte("value := 12"), 0o644); err != nil {
		t.Fatalf("Could not write .mag module: %s", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "mod_mg.mg"), []byte("value := 13"), 0o644); err != nil {
		t.Fatalf("Could not write .mg module: %s", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "explicit.ok"), []byte("value := 14"), 0o644); err != nil {
		t.Fatalf("Could not write explicit .ok module: %s", err)
	}

	ctx := NewContext(tmpDir)
	ctx.LoadBuiltins()

	val, err := ctx.Eval(strings.NewReader(`
		[
			import('mod_ok').value
			import('mod_mag').value
			import('mod_mg').value
			import('explicit.ok').value
		]
	`))
	if err != nil {
		t.Fatalf("Did not expect program to exit with error: %s", err.Error())
	}

	expected := MakeList(
		IntValue(11),
		IntValue(12),
		IntValue(13),
		IntValue(14),
	)
	if val == nil || !val.Eq(expected) {
		t.Fatalf("Expected %s, got %v", expected, val)
	}
}

func TestTrigAndPowLogBuiltins(t *testing.T) {
	expectProgramToReturn(t, `
		[
			sin(0)
			cos(0)
			tan(0)
			atan(0)
			pow(2, 3)
			log(2, 8)
		]
	`, MakeList(
		FloatValue(0),
		FloatValue(1),
		FloatValue(0),
		FloatValue(0),
		FloatValue(8),
		FloatValue(3),
	))
}

func TestAsinAndAcosBuiltins(t *testing.T) {
	expectProgramToReturn(t, `
		[
			asin(0)
			acos(1)
		]
	`, MakeList(
		FloatValue(0),
		FloatValue(0),
	))
}

func TestMathBuiltinDomainErrors(t *testing.T) {
	expectProgramToError(t, `asin(2)`)
	expectProgramToError(t, `acos(-2)`)
	expectProgramToError(t, `pow(0, 0)`)
	expectProgramToError(t, `pow(-1, 0.5)`)
	expectProgramToError(t, `log(0, 10)`)
	expectProgramToError(t, `log(2, 0)`)
}

func TestProcessTimeAndRandomBuiltins(t *testing.T) {
	expectProgramToReturn(t, `
		[
			type(args())
			len(args()) > 0
			type(env())
			type(time())
			type(nanotime())
			type(rand())
			len(srand(8))
			rand() >= 0 & rand() < 1
		]
	`, MakeList(
		AtomValue("list"),
		oakTrue,
		AtomValue("object"),
		AtomValue("float"),
		AtomValue("int"),
		AtomValue("float"),
		IntValue(8),
		oakTrue,
	))
}

func TestPrintBuiltinWritesAndReturnsCount(t *testing.T) {
	expectProgramToReturn(t, `print('oak')`, IntValue(3))
}

func TestWaitBuiltinAsyncCallback(t *testing.T) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()

	_, err := ctx.Eval(strings.NewReader(`
		done := false
		with wait(0) fn(_) {
			done <- true
		}
	`))
	if err != nil {
		t.Fatalf("Did not expect program to exit with error: %s", err.Error())
	}

	ctx.Wait()

	done, scopeErr := ctx.scope.get("done")
	if scopeErr != nil {
		t.Fatalf("Could not read done from scope: %s", scopeErr.Error())
	}

	if !done.Eq(oakTrue) {
		t.Fatalf("Expected async wait callback to set done to true, got %s", done)
	}
}

func TestAsyncEventBusSyncFlow(t *testing.T) {
	expectProgramToReturn(t, `
		buslib := import('async/event-bus')
		bus := buslib.create()
		seen := []
		persistId := bus.on('tick', fn(payload) {
			seen << payload
		})
		bus.once('tick', fn(payload) {
			seen << payload + 100
		})

		first := bus.emit('tick', 1)
		second := bus.emitSync('tick', 2)
		removed := bus.off('tick', persistId)
		third := bus.emit('tick', 3)

		[
			seen
			first
			second
			removed
			third
			bus.listenerCount('tick')
		]
	`, MakeList(
		MakeList(
			IntValue(1),
			IntValue(101),
			IntValue(2),
		),
		IntValue(2),
		IntValue(1),
		IntValue(1),
		IntValue(0),
		IntValue(0),
	))
}

func TestAsyncEventBusAsyncEmit(t *testing.T) {
	ctx := NewContext("/tmp")
	ctx.LoadBuiltins()

	val, err := ctx.Eval(strings.NewReader(`
		buslib := import('async/event-bus')
		bus := buslib.create()
		calls := 0
		doneCount := -1
		bus.on(:evt, fn(payload) {
			calls <- calls + payload
		})
		scheduled := bus.emitAsync(:evt, 2, fn(dispatched) {
			doneCount <- dispatched
		})
		[scheduled, calls, doneCount]
	`))
	if err != nil {
		t.Fatalf("Did not expect program to exit with error: %s", err.Error())
	}

	if val == nil || !val.Eq(MakeList(IntValue(1), IntValue(0), IntValue(-1))) {
		t.Fatalf("Expected immediate async emit state [1, 0, -1], got %v", val)
	}

	ctx.Wait()

	calls, scopeErr := ctx.scope.get("calls")
	if scopeErr != nil {
		t.Fatalf("Could not read calls from scope: %s", scopeErr.Error())
	}
	if !calls.Eq(IntValue(2)) {
		t.Fatalf("Expected async event handler to run once with payload 2, got %s", calls)
	}

	doneCount, scopeErr := ctx.scope.get("doneCount")
	if scopeErr != nil {
		t.Fatalf("Could not read doneCount from scope: %s", scopeErr.Error())
	}
	if !doneCount.Eq(IntValue(1)) {
		t.Fatalf("Expected async emit done callback to receive dispatched count 1, got %s", doneCount)
	}
}

func TestAtomAndCharBuiltins(t *testing.T) {
	expectProgramToReturn(t, `
		[
			atom('xyz')
			atom(10)
			char(65)
			char(-10)
			char(300)
		]
	`, MakeList(
		AtomValue("xyz"),
		AtomValue("10"),
		MakeString("A"),
		MakeString("\x00"),
		MakeString("\xff"),
	))
}

func TestRuntimeStdlibAndIntrospectionBuiltins(t *testing.T) {
	expectProgramToReturn(t, `
		[
			type(___stdlibs())
			type(___stdlibs().std)
			type(___stdlibs().sys)
			type(___stdlibs().windows)
			type(___stdlibs().gpus)
			type(___stdlibs().GUI)
			type(___stdlibs().websocket)
			type(___runtime_lib('std'))
			type(___runtime_lib('sys'))
			type(___runtime_lib('windows'))
			type(___runtime_lib('gpus'))
			type(___runtime_lib('GUI'))
			type(___runtime_lib('websocket'))
			___runtime_lib?('std')
			___runtime_lib?('sys')
			___runtime_lib?('windows')
			___runtime_lib?('gpus')
			___runtime_lib?('GUI')
			___runtime_lib?('websocket')
			type(___runtime_lib('definitely_missing_lib'))
			type(___runtime_gc())
			type(___runtime_mem().heap)
			type(___runtime_mem().allocs)
			type(___runtime_proc().pid)
			(type(___runtime_proc().exe) = :string) | (type(___runtime_proc().exe) = :null)
		]
	`, MakeList(
		AtomValue("object"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		AtomValue("string"),
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		AtomValue("null"),
		AtomValue("null"),
		AtomValue("int"),
		AtomValue("int"),
		AtomValue("int"),
		oakTrue,
	))
}

func TestWindowsStdlibSafeSurface(t *testing.T) {
	expectProgramToReturn(t, `
		win := import('windows')
		os := ___runtime_sys_info().os
		api := win.kernel32('GetCurrentProcessId')

		[
			type(win.Kernel32)
			win.Kernel32 = 'kernel32.dll'
			type(win.PROCESS_VM_READ)
			type(win.MEM_COMMIT)
			type(win.PAGE_READWRITE)
			type(win.wstr('oak'))
			type(win.cstr('oak'))
			len(win.cstr('oak')) = 4
			if os {
				'windows' -> {
					win.isWindows?() &
					((api.type = :ok) | (api.type = :success)) &
					(win.currentProcessId() > 0) &
					(win.imageBase() > 0)
				}
				_ -> {
					(!win.isWindows?()) &
					(api.type = :error) &
					(win.currentProcessId() = -1) &
					(win.imageBase() = 0)
				}
			}
		]
	`, MakeList(
		AtomValue("string"),
		oakTrue,
		AtomValue("int"),
		AtomValue("int"),
		AtomValue("int"),
		AtomValue("string"),
		AtomValue("string"),
		oakTrue,
		oakTrue,
	))
}

func TestLinuxStdlibSafeSurface(t *testing.T) {
	expectProgramToReturn(t, `
		lin := import('linux')
		os := ___runtime_sys_info().os
		uid := lin.getuid()
		gid := lin.getgid()
		fileAccess := lin.access('./README.md', lin.F_OK)

		[
			type(lin.LibC)
			lin.SEEK_SET = 0
			lin.F_OK = 0
			type(lin.cstr('oak'))
			len(lin.cstr('oak')) = 4
			if os {
				'linux' -> {
					lin.isLinux?() &
					((uid.type = :ok) | (uid.type = :success)) &
					((gid.type = :ok) | (gid.type = :success)) &
					((fileAccess.type = :ok) | (fileAccess.type = :success))
				}
				_ -> {
					(!lin.isLinux?()) &
					(uid.type = :error) &
					(gid.type = :error) &
					(fileAccess.type = :error)
				}
			}
		]
	`, MakeList(
		AtomValue("list"),
		oakTrue,
		oakTrue,
		AtomValue("string"),
		oakTrue,
		oakTrue,
	))
}

func TestGUIStdlibSafeSurface(t *testing.T) {
	expectProgramToReturn(t, `
		gui := import('GUI')
		os := ___runtime_sys_info().os
		window := gui.createWindow('GUI Smoke', 320, 220)
		
		[
			type(gui.backend())
			type(gui.rgb(1, 2, 3))
			type(gui.Vec2)
			type(gui.Rect2)
			type(gui.vec2Add)
			type(gui.vec2Sub)
			type(gui.vec2Scale)
			type(gui.vec2Dot)
			type(gui.vec2Len)
			type(gui.vec2Normalize)
			type(gui.rectTranslate)
			type(gui.rectContains)
			type(gui.rectIntersects)
			type(gui.Transform2D)
			type(gui.applyTransform2D)
			type(gui.Camera2D)
			type(gui.worldToScreen2D)
			type(gui.screenToWorld2D)
			type(gui.drawRect2D)
			type(gui.drawCircle2D)
			type(gui.drawPolyline2D)
			type(gui.drawPolygon2D)
			type(gui.drawGrid2D)
			type(gui.createCanvas)
			type(gui.initWebGL)
			type(gui.webglCreateShader)
			type(gui.webglCreateProgram)
			type(gui.webglUseProgram)
			type(gui.webglClearColor)
			type(gui.webglViewport)
			type(gui.webglClear)
			type(gui.webglDrawArrays)
			type(gui.webglFlush)
			type(gui.Vec3)
			type(gui.CubeMesh)
			type(gui.Mesh)
			type(gui.GridMesh)
			type(gui.AxesMesh)
			type(gui.VoxelMesh)
			type(gui.VoxelGrid)
			type(gui.degToRad)
			type(gui.drawLine)
			type(gui.drawTriangleFilled)
			type(gui.drawMeshSolid)
			type(gui.drawMeshWireframe)
			type(gui.Renderer3D)
			gui.GL_COLOR_BUFFER_BIT = 16384
			gui.GL_TRIANGLES = 4
			type(window)
			window.type = :ok
			type(window.backend)
			if os {
				'windows' -> {
					(gui.backend() = :windows) &
					(window.backend = :windows)
				}
				'linux' -> {
					(gui.backend() = :linux) &
					(window.backend = :linux)
				}
				_ -> {
					(gui.backend() = :web) | (gui.backend() = :unknown)
				}
			}
			type(gui.close(window))
		]
	`, MakeList(
		AtomValue("atom"),
		AtomValue("int"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		oakTrue,
		oakTrue,
		AtomValue("object"),
		oakTrue,
		AtomValue("atom"),
		oakTrue,
		AtomValue("int"),
	))
}

func TestGUIOpacityHelpers(t *testing.T) {
	expectProgramToReturn(t, `
		gui := import('GUI')
		bg := gui.rgb(18, 20, 26)
		overlay := gui.opacity(gui.rgb(248, 232, 242), 0.2, bg)
		overlayRgba := gui.rgba(248, 232, 242, 0.2, bg)

		[
			overlay
			overlay = overlayRgba
			gui.colorR(overlay)
			gui.colorG(overlay)
			gui.colorB(overlay)
		]
	`, MakeList(
		IntValue(4537920),
		oakTrue,
		IntValue(64),
		IntValue(62),
		IntValue(69),
	))
}

func TestGUIDrawOpsPreserveColorOnWeb(t *testing.T) {
	expectProgramToReturn(t, `
		gui := import('GUI')
		bg := gui.rgb(18, 20, 26)
		textColor := gui.opacity(gui.rgb(248, 232, 242), 0.2, bg)
		rectColor := gui.opacity(gui.rgb(80, 120, 160), 0.5, bg)
		win := {
			type: :ok
			backend: :web
			messages: []
		}

		gui.drawText(win, 10, 20, 'alpha', textColor)
		gui.fillRect(win, 1, 2, 30, 40, rectColor)
		gui.drawLine(win, 0, 0, 5, 5, textColor)

		[
			len(win.messages)
			win.messages.(0).type = :text
			win.messages.(0).color = textColor
			win.messages.(1).type = :rect
			win.messages.(1).color = rectColor
			win.messages.(2).type = :line
			win.messages.(2).color = textColor
		]
	`, MakeList(
		IntValue(3),
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
	))
}

func TestGUIEventBusHelpers(t *testing.T) {
	expectProgramToReturn(t, `
		gui := import('GUI')
		win := {type: :ok}
		seen := []

		persist := gui.on(win, :dispatch, fn(step, name) {
			seen << (string(name) + ':' + string(step.type))
		})
		gui.once(win, :dispatch, fn(_, name) {
			seen << (string(name) + ':once')
		})

		first := gui.emit(win, :dispatch, {type: :dispatch})
		second := gui.emit(win, :dispatch, {type: :dispatch})
		removed := gui.off(win, :dispatch, persist)
		third := gui.emit(win, :dispatch, {type: :dispatch})

		[
			seen
			first
			second
			removed
			third
			gui.listenerCount(win, :dispatch)
			type(gui.eventBus(win).emitAsync)
		]
	`, MakeList(
		MakeList(
			MakeString("dispatch:dispatch"),
			MakeString("dispatch:once"),
			MakeString("dispatch:dispatch"),
		),
		IntValue(2),
		IntValue(1),
		IntValue(1),
		IntValue(0),
		IntValue(0),
		AtomValue("function"),
	))
}

func TestGUIGraphingHelpers(t *testing.T) {
	expectProgramToReturn(t, `
		gui := import('GUI')
		win := {
			type: :ok
			backend: :web
			messages: []
			width: 320
			height: 200
		}

		values := [2, 5, 9, 4, 7]
		range := gui.graphRange(values, {padding: 1})
		x2 := gui.graphMapX(2, len(values), 10, 100)
		yTop := gui.graphMapY(range.max, 20, 70, range)

		gui.drawGraphAxes(win, 10, 10, 180, 90, {
			showGrid: true
			xTicks: 3
			yTicks: 2
		})
		gui.drawLineGraph(win, 10, 120, 180, 60, values, {
			showPoints: true
			showLabels: true
		})
		gui.drawBarGraph(win, 200, 120, 100, 60, values, {
			showLabels: true
			barGap: 2
		})
		gui.drawSparkline(win, 200, 20, 100, 40, values, {
			rangePadding: 1
		})

		fn countType(i, target, acc) if i < len(win.messages) {
			true -> {
				msg := win.messages.(i)
				next := if msg.type = target {
					true -> acc + 1
					_ -> acc
				}
				countType(i + 1, target, next)
			}
			_ -> acc
		}

		lineCount := countType(0, :line, 0)
		rectCount := countType(0, :rect, 0)
		textCount := countType(0, :text, 0)

		[
			type(gui.graphRange)
			type(gui.graphMapX)
			type(gui.graphMapY)
			type(gui.drawGraphAxes)
			type(gui.drawLineGraph)
			type(gui.drawBarGraph)
			type(gui.drawSparkline)
			range.min = 1
			range.max = 10
			range.span = 9
			x2 = 60
			yTop = 20
			len(win.messages) > 0
			lineCount > 0
			rectCount > 0
			textCount > 0
		]
	`, MakeList(
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
		oakTrue,
	))
}

func TestBuiltinPresenceForRemainingIoAndProcessFeatures(t *testing.T) {
	expectProgramToReturn(t, `
		[
			type(input)
			type(exit)
			type(exec)
			type(ls)
			type(rm)
			type(mkdir)
			type(stat)
			type(open)
			type(close)
			type(read)
			type(write)
			type(listen)
			type(req)
			type(ws_dial)
			type(ws_send)
			type(ws_recv)
			type(ws_close)
			type(ws_listen)
		]
	`, MakeList(
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
		AtomValue("function"),
	))
}
