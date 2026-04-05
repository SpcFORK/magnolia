package main

import (
	"math"
	"strconv"
)

// optimizeAST applies constant folding to a list of parsed AST nodes.
// This is called between parsing and evaluation to simplify expressions
// that can be computed at compile time.
func optimizeAST(nodes []astNode) []astNode {
	optimized := make([]astNode, len(nodes))
	for i, node := range nodes {
		optimized[i] = foldConstants(node)
	}
	return optimized
}

// foldConstants recursively walks an AST node and folds constant expressions
// into their resulting literal values.
func foldConstants(node astNode) astNode {
	switch n := node.(type) {
	case unaryNode:
		return foldUnary(n)
	case binaryNode:
		return foldBinary(n)
	case blockNode:
		return foldBlock(n)
	case listNode:
		return foldList(n)
	case objectNode:
		return foldObject(n)
	case fnNode:
		return foldFn(n)
	case classNode:
		return foldClass(n)
	case ifExprNode:
		return foldIf(n)
	case assignmentNode:
		return foldAssignment(n)
	case fnCallNode:
		return foldFnCall(n)
	case propertyAccessNode:
		return foldPropertyAccess(n)
	default:
		return node
	}
}

func isConstant(node astNode) bool {
	switch node.(type) {
	case intNode, floatNode, boolNode, stringNode, atomNode, nullNode:
		return true
	}
	return false
}

func foldUnary(n unaryNode) astNode {
	n.right = foldConstants(n.right)

	switch right := n.right.(type) {
	case intNode:
		switch n.op {
		case plus:
			return right
		case minus:
			return intNode{payload: -right.payload, tok: n.tok}
		case tilde:
			return intNode{payload: ^right.payload, tok: n.tok}
		}
	case floatNode:
		switch n.op {
		case plus:
			return right
		case minus:
			return floatNode{payload: -right.payload, tok: n.tok}
		}
	case boolNode:
		switch n.op {
		case exclam:
			return boolNode{payload: !right.payload, tok: n.tok}
		}
	}

	return simplifyUnary(n)
}

func foldBinary(n binaryNode) astNode {
	n.left = foldConstants(n.left)
	n.right = foldConstants(n.right)

	if n.op == eq || n.op == deepEq || n.op == neq {
		return foldEquality(n)
	}

	switch left := n.left.(type) {
	case intNode:
		switch right := n.right.(type) {
		case intNode:
			return foldIntInt(n, left, right)
		case floatNode:
			return foldFloatFloat(n,
				floatNode{payload: float64(left.payload), tok: left.tok},
				right)
		}
	case floatNode:
		switch right := n.right.(type) {
		case floatNode:
			return foldFloatFloat(n, left, right)
		case intNode:
			return foldFloatFloat(n, left,
				floatNode{payload: float64(right.payload), tok: right.tok})
		}
	case stringNode:
		if right, ok := n.right.(stringNode); ok {
			return foldStringString(n, left, right)
		}
	case boolNode:
		if right, ok := n.right.(boolNode); ok {
			return foldBoolBool(n, left, right)
		}
	}

	return simplifyBinary(n)
}

func foldIntInt(n binaryNode, left, right intNode) astNode {
	switch n.op {
	case plus:
		return intNode{payload: left.payload + right.payload, tok: n.tok}
	case minus:
		return intNode{payload: left.payload - right.payload, tok: n.tok}
	case times:
		return intNode{payload: left.payload * right.payload, tok: n.tok}
	case divide:
		if right.payload == 0 {
			return n
		}
		if left.payload%right.payload == 0 {
			return intNode{payload: left.payload / right.payload, tok: n.tok}
		}
		return floatNode{payload: float64(left.payload) / float64(right.payload), tok: n.tok}
	case modulus:
		if right.payload == 0 {
			return n
		}
		return intNode{payload: left.payload % right.payload, tok: n.tok}
	case power:
		if right.payload < 0 {
			return floatNode{payload: math.Pow(float64(left.payload), float64(right.payload)), tok: n.tok}
		}
		result := math.Pow(float64(left.payload), float64(right.payload))
		if result > math.MaxInt64 || result < math.MinInt64 {
			return floatNode{payload: result, tok: n.tok}
		}
		return intNode{payload: int64(result), tok: n.tok}
	case xor:
		return intNode{payload: left.payload ^ right.payload, tok: n.tok}
	case and:
		return intNode{payload: left.payload & right.payload, tok: n.tok}
	case or:
		return intNode{payload: left.payload | right.payload, tok: n.tok}
	case pushArrow:
		if right.payload < 0 {
			return n
		}
		return intNode{payload: left.payload << uint(right.payload), tok: n.tok}
	case rshift:
		if right.payload < 0 {
			return n
		}
		return intNode{payload: left.payload >> uint(right.payload), tok: n.tok}
	case greater:
		return boolNode{payload: left.payload > right.payload, tok: n.tok}
	case less:
		return boolNode{payload: left.payload < right.payload, tok: n.tok}
	case geq:
		return boolNode{payload: left.payload >= right.payload, tok: n.tok}
	case leq:
		return boolNode{payload: left.payload <= right.payload, tok: n.tok}
	}
	return n
}

func foldFloatFloat(n binaryNode, left, right floatNode) astNode {
	switch n.op {
	case plus:
		return floatNode{payload: left.payload + right.payload, tok: n.tok}
	case minus:
		return floatNode{payload: left.payload - right.payload, tok: n.tok}
	case times:
		return floatNode{payload: left.payload * right.payload, tok: n.tok}
	case divide:
		if right.payload == 0 {
			return n
		}
		return floatNode{payload: left.payload / right.payload, tok: n.tok}
	case modulus:
		if right.payload == 0 {
			return n
		}
		return floatNode{payload: math.Mod(left.payload, right.payload), tok: n.tok}
	case power:
		return floatNode{payload: math.Pow(left.payload, right.payload), tok: n.tok}
	case greater:
		return boolNode{payload: left.payload > right.payload, tok: n.tok}
	case less:
		return boolNode{payload: left.payload < right.payload, tok: n.tok}
	case geq:
		return boolNode{payload: left.payload >= right.payload, tok: n.tok}
	case leq:
		return boolNode{payload: left.payload <= right.payload, tok: n.tok}
	}
	return n
}

func foldStringString(n binaryNode, left, right stringNode) astNode {
	switch n.op {
	case plus:
		result := make([]byte, len(left.payload)+len(right.payload))
		copy(result, left.payload)
		copy(result[len(left.payload):], right.payload)
		return stringNode{payload: result, tok: n.tok}
	case greater:
		return boolNode{payload: string(left.payload) > string(right.payload), tok: n.tok}
	case less:
		return boolNode{payload: string(left.payload) < string(right.payload), tok: n.tok}
	case geq:
		return boolNode{payload: string(left.payload) >= string(right.payload), tok: n.tok}
	case leq:
		return boolNode{payload: string(left.payload) <= string(right.payload), tok: n.tok}
	}
	return n
}

func foldBoolBool(n binaryNode, left, right boolNode) astNode {
	switch n.op {
	case plus, or:
		return boolNode{payload: left.payload || right.payload, tok: n.tok}
	case times, and:
		return boolNode{payload: left.payload && right.payload, tok: n.tok}
	case xor:
		return boolNode{payload: left.payload != right.payload, tok: n.tok}
	}
	return n
}

func foldEquality(n binaryNode) astNode {
	// Self-comparison: x = x → true, x != x → false (when x is a pure identifier)
	if isSimpleExpr(n.left) && isPure(n.left) {
		if leftId, ok := n.left.(identifierNode); ok {
			if rightId, ok := n.right.(identifierNode); ok {
				if leftId.payload == rightId.payload {
					if n.op == neq {
						return boolNode{payload: false, tok: n.tok}
					}
					return boolNode{payload: true, tok: n.tok}
				}
			}
		}
	}

	if !isConstant(n.left) || !isConstant(n.right) {
		return n
	}

	var equal bool
	switch left := n.left.(type) {
	case intNode:
		if right, ok := n.right.(intNode); ok {
			equal = left.payload == right.payload
		} else if right, ok := n.right.(floatNode); ok {
			equal = float64(left.payload) == right.payload
		} else {
			return n
		}
	case floatNode:
		if right, ok := n.right.(floatNode); ok {
			equal = left.payload == right.payload
		} else if right, ok := n.right.(intNode); ok {
			equal = left.payload == float64(right.payload)
		} else {
			return n
		}
	case boolNode:
		if right, ok := n.right.(boolNode); ok {
			equal = left.payload == right.payload
		} else {
			return n
		}
	case stringNode:
		if right, ok := n.right.(stringNode); ok {
			equal = string(left.payload) == string(right.payload)
		} else {
			return n
		}
	case atomNode:
		if right, ok := n.right.(atomNode); ok {
			equal = left.payload == right.payload
		} else {
			return n
		}
	case nullNode:
		_, ok := n.right.(nullNode)
		equal = ok
	default:
		return n
	}

	if n.op == neq {
		equal = !equal
	}
	return boolNode{payload: equal, tok: n.tok}
}

func foldBlock(n blockNode) astNode {
	return foldBlockAdvanced(n)
}

func foldList(n listNode) astNode {
	for i, elem := range n.elems {
		n.elems[i] = foldConstants(elem)
	}
	return n
}

func foldObject(n objectNode) astNode {
	for i, entry := range n.entries {
		n.entries[i].key = foldConstants(entry.key)
		n.entries[i].val = foldConstants(entry.val)
	}
	return n
}

func foldFn(n fnNode) astNode {
	n.body = foldConstants(n.body)
	return n
}

func foldClass(n classNode) astNode {
	n.body = foldConstants(n.body)
	for i, expr := range n.staticExprs {
		n.staticExprs[i] = foldConstants(expr)
	}
	return n
}

func foldIf(n ifExprNode) astNode {
	n.cond = foldConstants(n.cond)
	for i, branch := range n.branches {
		n.branches[i].target = foldConstants(branch.target)
		n.branches[i].body = foldConstants(branch.body)
	}

	// Dead branch elimination: if the condition is a constant, we can
	// statically determine which branch matches and eliminate the rest.
	// We can only fold away the entire if when ALL prior targets are constants
	// (so we know at compile time that none of them match).
	if isConstant(n.cond) {
		allPriorConstant := true
		for _, branch := range n.branches {
			if isConstant(branch.target) && constantsEqual(n.cond, branch.target) {
				return branch.body
			}
			// _ (emptyNode) is a wildcard that always matches
			if _, ok := branch.target.(emptyNode); ok {
				if allPriorConstant {
					return branch.body
				}
				break
			}
			if !isConstant(branch.target) {
				allPriorConstant = false
			}
		}
	}

	// If a trailing branch has a _ (wildcard) target and all prior targets
	// are constants, we can prune unreachable branches after the wildcard.
	// (The wildcard always matches, so branches after it are dead.)
	for i, branch := range n.branches {
		if _, ok := branch.target.(emptyNode); ok {
			n.branches = n.branches[:i+1]
			break
		}
	}

	// Single-branch simplification: if only one branch remains and it
	// is a wildcard, the entire if-expression reduces to just evaluating
	// the condition (for side effects) and then the body. However, if
	// the condition is pure, we can skip it entirely.
	if len(n.branches) == 1 {
		if _, ok := n.branches[0].target.(emptyNode); ok {
			if isPure(n.cond) {
				return n.branches[0].body
			}
		}
	}

	return n
}

func foldAssignment(n assignmentNode) astNode {
	n.right = foldConstants(n.right)
	return n
}

func foldFnCall(n fnCallNode) astNode {
	n.fn = foldConstants(n.fn)
	for i, arg := range n.args {
		n.args[i] = foldConstants(arg)
	}
	if n.restArg != nil {
		n.restArg = foldConstants(n.restArg)
	}

	// Builtin constant folding: if calling a known pure builtin with
	// constant arguments, evaluate at compile time.
	if ident, ok := n.fn.(identifierNode); ok && n.restArg == nil {
		if len(n.args) == 1 {
			arg := n.args[0]
			switch ident.payload {
			case "len":
				return foldBuiltinLen(n, arg)
			case "type":
				return foldBuiltinType(n, arg)
			case "int":
				return foldBuiltinInt(n, arg)
			case "float":
				return foldBuiltinFloat(n, arg)
			case "string":
				return foldBuiltinString(n, arg)
			case "codepoint":
				return foldBuiltinCodepoint(n, arg)
			case "char":
				return foldBuiltinChar(n, arg)
			}
		}
		if len(n.args) == 2 {
			switch ident.payload {
			case "pow":
				return foldBuiltinPow(n, n.args[0], n.args[1])
			}
		}
	}

	return n
}

func foldPropertyAccess(n propertyAccessNode) astNode {
	n.left = foldConstants(n.left)
	n.right = foldConstants(n.right)

	// Constant list index: [1, 2, 3].0 → 1
	if list, ok := n.left.(listNode); ok {
		if idx, ok := n.right.(intNode); ok {
			if idx.payload >= 0 && int(idx.payload) < len(list.elems) {
				elem := list.elems[idx.payload]
				if isConstant(elem) {
					return elem
				}
			}
		}
	}

	// Constant string index: 'hello'.0 → single-char string
	if str, ok := n.left.(stringNode); ok {
		if idx, ok := n.right.(intNode); ok {
			if idx.payload >= 0 && int(idx.payload) < len(str.payload) {
				return stringNode{payload: []byte{str.payload[idx.payload]}, tok: n.tok}
			}
		}
	}

	// Constant object property access: {a: 1, b: 2}.a → 1
	if obj, ok := n.left.(objectNode); ok {
		if keyIdent, ok := n.right.(identifierNode); ok {
			for _, entry := range obj.entries {
				var keyStr string
				switch k := entry.key.(type) {
				case identifierNode:
					keyStr = k.payload
				case stringNode:
					keyStr = string(k.payload)
				default:
					continue
				}
				if keyStr == keyIdent.payload && isConstant(entry.val) {
					return entry.val
				}
			}
		}
	}

	return n
}

// constantsEqual checks whether two constant AST nodes represent the same value.
// Both nodes must pass isConstant() before calling this.
func constantsEqual(a, b astNode) bool {
	switch av := a.(type) {
	case intNode:
		if bv, ok := b.(intNode); ok {
			return av.payload == bv.payload
		}
		if bv, ok := b.(floatNode); ok {
			return float64(av.payload) == bv.payload
		}
	case floatNode:
		if bv, ok := b.(floatNode); ok {
			return av.payload == bv.payload
		}
		if bv, ok := b.(intNode); ok {
			return av.payload == float64(bv.payload)
		}
	case boolNode:
		if bv, ok := b.(boolNode); ok {
			return av.payload == bv.payload
		}
	case stringNode:
		if bv, ok := b.(stringNode); ok {
			return string(av.payload) == string(bv.payload)
		}
	case atomNode:
		if bv, ok := b.(atomNode); ok {
			return av.payload == bv.payload
		}
	case nullNode:
		_, ok := b.(nullNode)
		return ok
	}
	return false
}

// isZero checks if a constant node is a numeric zero.
func isZero(node astNode) bool {
	switch n := node.(type) {
	case intNode:
		return n.payload == 0
	case floatNode:
		return n.payload == 0.0
	}
	return false
}

// isOne checks if a constant node is a numeric one.
func isOne(node astNode) bool {
	switch n := node.(type) {
	case intNode:
		return n.payload == 1
	case floatNode:
		return n.payload == 1.0
	}
	return false
}

// isIntTwo checks if a node is the integer literal 2.
func isIntTwo(node astNode) bool {
	n, ok := node.(intNode)
	return ok && n.payload == 2
}

// isIntLiteral checks if a node is any integer literal.
func isIntLiteral(node astNode) bool {
	_, ok := node.(intNode)
	return ok
}

// simplifyBinary applies algebraic identity, absorption, strength-reduction,
// and boolean simplification rules after constant folding has been performed
// on the operands. It is called from foldBinary when full constant folding
// did not reduce the node.
func simplifyBinary(n binaryNode) astNode {
	// --- Identity / Absorption rules ---
	switch n.op {
	case plus:
		// x + 0 → x, 0 + x → x
		if isZero(n.right) {
			return n.left
		}
		if isZero(n.left) {
			return n.right
		}
	case minus:
		// x - 0 → x
		if isZero(n.right) {
			return n.left
		}
	case times:
		// x * 1 → x, 1 * x → x
		if isOne(n.right) {
			return n.left
		}
		if isOne(n.left) {
			return n.right
		}
		// x * 0 → 0, 0 * x → 0 (when x is side-effect free)
		if isZero(n.right) && isPure(n.left) {
			return n.right
		}
		if isZero(n.left) && isPure(n.right) {
			return n.left
		}
	case divide:
		// x / 1 → x
		if isOne(n.right) {
			return n.left
		}
	case power:
		// x ** 1 → x
		if isOne(n.right) {
			return n.left
		}
		// x ** 0 → 1 (when x is side-effect free)
		if isZero(n.right) && isPure(n.left) {
			return intNode{payload: 1, tok: n.tok}
		}
		// x ** 2 → x * x (strength reduction) — only when x is a simple identifier
		if isIntTwo(n.right) && isSimpleExpr(n.left) {
			return binaryNode{op: times, left: n.left, right: n.left, tok: n.tok}
		}
	case or:
		// x | 0 → x, 0 | x → x
		if isZero(n.right) {
			return n.left
		}
		if isZero(n.left) {
			return n.right
		}
	case and:
		// x & 0 → 0 (when x is side-effect free)
		if isZero(n.right) && isPure(n.left) {
			return n.right
		}
		if isZero(n.left) && isPure(n.right) {
			return n.left
		}
	case xor:
		// x ^ 0 → x, 0 ^ x → x
		if isZero(n.right) {
			return n.left
		}
		if isZero(n.left) {
			return n.right
		}
	case pushArrow:
		// x << 0 → x (only for integer left-shift, NOT list push)
		// Since pushArrow is overloaded (bitwise shift for ints, push for lists),
		// we can only safely fold when the left side is a known integer.
		if isZero(n.right) && isIntLiteral(n.left) {
			return n.left
		}
	case rshift:
		// x >> 0 → x
		if isZero(n.right) {
			return n.left
		}
	}

	// --- Boolean simplification ---
	switch n.op {
	case plus, or:
		// true | x → true (short-circuit), x | true → true
		if isBoolTrue(n.left) && isPure(n.right) {
			return n.left
		}
		if isBoolTrue(n.right) && isPure(n.left) {
			return n.right
		}
		// false | x → x, x | false → x
		if isBoolFalse(n.left) {
			return n.right
		}
		if isBoolFalse(n.right) {
			return n.left
		}
	case times, and:
		// false & x → false (short-circuit), x & false → false
		if isBoolFalse(n.left) && isPure(n.right) {
			return n.left
		}
		if isBoolFalse(n.right) && isPure(n.left) {
			return n.right
		}
		// true & x → x, x & true → x
		if isBoolTrue(n.left) {
			return n.right
		}
		if isBoolTrue(n.right) {
			return n.left
		}
	}

	return n
}

// isPure returns true if an expression has no side effects and can safely
// be eliminated (e.g., for x*0 → 0 when x is pure).
func isPure(node astNode) bool {
	switch n := node.(type) {
	case intNode, floatNode, boolNode, stringNode, atomNode, nullNode, emptyNode:
		return true
	case identifierNode:
		return true
	case unaryNode:
		return isPure(n.right)
	case binaryNode:
		return isPure(n.left) && isPure(n.right)
	case blockNode:
		for _, expr := range n.exprs {
			if !isPure(expr) {
				return false
			}
		}
		return true
	case listNode:
		for _, elem := range n.elems {
			if !isPure(elem) {
				return false
			}
		}
		return true
	case propertyAccessNode:
		return isPure(n.left) && isPure(n.right)
	}
	return false
}

// isSimpleExpr returns true if an expression is a single identifier or literal,
// safe to duplicate without concern about evaluation cost or side effects.
func isSimpleExpr(node astNode) bool {
	switch node.(type) {
	case identifierNode, intNode, floatNode, boolNode, stringNode, atomNode, nullNode:
		return true
	}
	return false
}

func isBoolTrue(node astNode) bool {
	n, ok := node.(boolNode)
	return ok && n.payload
}

func isBoolFalse(node astNode) bool {
	n, ok := node.(boolNode)
	return ok && !n.payload
}

// simplifyUnary applies simplification rules for unary expressions
// after the operand has been folded.
func simplifyUnary(n unaryNode) astNode {
	// Double negation elimination: !!x → x
	if n.op == exclam {
		if inner, ok := n.right.(unaryNode); ok && inner.op == exclam {
			return inner.right
		}
		// Negation of comparison: !(a > b) → a <= b, etc.
		if inner, ok := n.right.(binaryNode); ok {
			var invertedOp tokKind
			switch inner.op {
			case greater:
				invertedOp = leq
			case less:
				invertedOp = geq
			case geq:
				invertedOp = less
			case leq:
				invertedOp = greater
			case eq:
				invertedOp = neq
			case neq:
				invertedOp = eq
			default:
				return n
			}
			return binaryNode{op: invertedOp, left: inner.left, right: inner.right, tok: inner.tok}
		}
	}
	// Double numeric negation: --x → x
	if n.op == minus {
		if inner, ok := n.right.(unaryNode); ok && inner.op == minus {
			return inner.right
		}
	}
	// Double bitwise NOT: ~~x → x
	if n.op == tilde {
		if inner, ok := n.right.(unaryNode); ok && inner.op == tilde {
			return inner.right
		}
	}
	return n
}

// --- Builtin constant folding ---

func foldBuiltinLen(n fnCallNode, arg astNode) astNode {
	switch a := arg.(type) {
	case stringNode:
		return intNode{payload: int64(len(a.payload)), tok: n.tok}
	case listNode:
		// Only fold if all elements are constants (no spread/rest)
		return intNode{payload: int64(len(a.elems)), tok: n.tok}
	}
	return n
}

func foldBuiltinType(n fnCallNode, arg astNode) astNode {
	var typeName string
	switch arg.(type) {
	case nullNode:
		typeName = "null"
	case intNode:
		typeName = "int"
	case floatNode:
		typeName = "float"
	case boolNode:
		typeName = "bool"
	case atomNode:
		typeName = "atom"
	case stringNode:
		typeName = "string"
	default:
		return n
	}
	return atomNode{payload: typeName, tok: n.tok}
}

func foldBuiltinInt(n fnCallNode, arg astNode) astNode {
	switch a := arg.(type) {
	case intNode:
		return a
	case floatNode:
		return intNode{payload: int64(math.Floor(a.payload)), tok: n.tok}
	case stringNode:
		v, err := strconv.ParseInt(string(a.payload), 10, 64)
		if err != nil {
			// int('not_a_number') → ? at runtime; don't fold
			return n
		}
		return intNode{payload: v, tok: n.tok}
	}
	return n
}

func foldBuiltinFloat(n fnCallNode, arg astNode) astNode {
	switch a := arg.(type) {
	case intNode:
		return floatNode{payload: float64(a.payload), tok: n.tok}
	case floatNode:
		return a
	case stringNode:
		v, err := strconv.ParseFloat(string(a.payload), 64)
		if err != nil {
			return n
		}
		return floatNode{payload: v, tok: n.tok}
	}
	return n
}

func foldBuiltinString(n fnCallNode, arg astNode) astNode {
	switch a := arg.(type) {
	case stringNode:
		return a
	case intNode:
		return stringNode{payload: []byte(strconv.FormatInt(a.payload, 10)), tok: n.tok}
	case floatNode:
		return stringNode{payload: []byte(strconv.FormatFloat(a.payload, 'f', -1, 64)), tok: n.tok}
	case boolNode:
		if a.payload {
			return stringNode{payload: []byte("true"), tok: n.tok}
		}
		return stringNode{payload: []byte("false"), tok: n.tok}
	case atomNode:
		return stringNode{payload: []byte(a.payload), tok: n.tok}
	}
	return n
}

func foldBuiltinCodepoint(n fnCallNode, arg astNode) astNode {
	// codepoint('A') → 65
	if s, ok := arg.(stringNode); ok {
		if len(s.payload) == 1 {
			return intNode{payload: int64(s.payload[0]), tok: n.tok}
		}
	}
	return n
}

func foldBuiltinChar(n fnCallNode, arg astNode) astNode {
	// char(65) → 'A'
	if i, ok := arg.(intNode); ok {
		cp := i.payload
		if cp < 0 {
			cp = 0
		}
		if cp > 255 {
			cp = 255
		}
		return stringNode{payload: []byte{byte(cp)}, tok: n.tok}
	}
	return n
}

func foldBuiltinPow(n fnCallNode, base, exp astNode) astNode {
	// Get numeric values for base and exponent
	var bv, ev float64
	var bOk, eOk bool

	switch b := base.(type) {
	case intNode:
		bv, bOk = float64(b.payload), true
	case floatNode:
		bv, bOk = b.payload, true
	}
	switch e := exp.(type) {
	case intNode:
		ev, eOk = float64(e.payload), true
	case floatNode:
		ev, eOk = e.payload, true
	}
	if !bOk || !eOk {
		return n
	}

	// Guard domain errors: pow(0,0), negative base with fractional exponent
	if bv == 0 && ev == 0 {
		return n // runtime error: 0^0 undefined
	}
	result := math.Pow(bv, ev)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		return n // runtime error for domain violations
	}

	// Return int if base/exp were both ints with non-negative exponent
	if _, bIsInt := base.(intNode); bIsInt {
		if e, eIsInt := exp.(intNode); eIsInt && e.payload >= 0 {
			if result <= math.MaxInt64 && result >= math.MinInt64 {
				return intNode{payload: int64(result), tok: n.tok}
			}
		}
	}
	return floatNode{payload: result, tok: n.tok}
}

// --- Block-level optimizations ---

// foldBlock applies constant folding to block expressions, then performs
// additional simplification passes on the block structure itself.
func foldBlockAdvanced(n blockNode) astNode {
	// First fold all sub-expressions
	for i, expr := range n.exprs {
		n.exprs[i] = foldConstants(expr)
	}

	// Flatten nested single-expression blocks: { { expr } } → { expr }
	// and multi-expression inner blocks that contain no local bindings
	// get their expressions inlined into the parent.
	flattened := make([]astNode, 0, len(n.exprs))
	for _, expr := range n.exprs {
		if inner, ok := expr.(blockNode); ok && len(inner.exprs) == 1 && !inner.hasLocal {
			flattened = append(flattened, inner.exprs[0])
		} else {
			flattened = append(flattened, expr)
		}
	}
	n.exprs = flattened

	// Unwrap single-expression blocks when the expression is a constant,
	// so that e.g. (2 + 3) * 4 can be further folded.
	if len(n.exprs) == 1 && isConstant(n.exprs[0]) {
		return n.exprs[0]
	}

	// NOTE: We do NOT unwrap single-expression non-constant blocks here.
	// In Oak/Magnolia, parenthesized expressions like obj.(key) parse as
	// propertyAccessNode{right: blockNode{[identifierNode]}}, and the block
	// serves as a marker for computed property access. Unwrapping it would
	// change obj.(key) to obj.key, breaking computed access semantics.

	// Dead assignment elimination: if a local assignment's target variable
	// is immediately reassigned (next expression) without being read, the
	// first assignment is dead and can be removed — but only when the RHS
	// is pure (no side effects to preserve).
	if len(n.exprs) >= 2 {
		cleaned := make([]astNode, 0, len(n.exprs))
		for i := 0; i < len(n.exprs); i++ {
			if i < len(n.exprs)-1 {
				if isDeadAssignment(n.exprs[i], n.exprs[i+1]) {
					continue
				}
			}
			cleaned = append(cleaned, n.exprs[i])
		}
		if len(cleaned) < len(n.exprs) {
			n.exprs = cleaned
		}
	}

	// After dead assignment elimination, re-check single-expression unwrap
	if len(n.exprs) == 1 && isConstant(n.exprs[0]) {
		return n.exprs[0]
	}

	// Compute hasLocal: mark whether this block has any local (:=) assignments.
	// This allows eval to skip scope creation for blocks without bindings.
	n.hasLocal = blockHasLocalAssignment(n.exprs)

	return n
}

// blockHasLocalAssignment checks if any expression in the list is a local assignment.
func blockHasLocalAssignment(exprs []astNode) bool {
	for _, expr := range exprs {
		if assign, ok := expr.(assignmentNode); ok && assign.isLocal {
			return true
		}
	}
	return false
}

// isDeadAssignment checks if expr is a local assignment whose variable is
// immediately reassigned in nextExpr without being used. Only pure RHS
// values can be eliminated.
func isDeadAssignment(expr, nextExpr astNode) bool {
	assign, ok := expr.(assignmentNode)
	if !ok || !assign.isLocal {
		return false
	}
	ident, ok := assign.left.(identifierNode)
	if !ok {
		return false
	}
	if !isPure(assign.right) {
		return false
	}
	// Check that nextExpr is a local assignment to the same variable
	nextAssign, ok := nextExpr.(assignmentNode)
	if !ok || !nextAssign.isLocal {
		return false
	}
	nextIdent, ok := nextAssign.left.(identifierNode)
	if !ok {
		return false
	}
	return ident.payload == nextIdent.payload
}
