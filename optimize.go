package main

import "math"

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

	return n
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

	return n
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
	if n.op == plus {
		result := make([]byte, len(left.payload)+len(right.payload))
		copy(result, left.payload)
		copy(result[len(left.payload):], right.payload)
		return stringNode{payload: result, tok: n.tok}
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
	for i, expr := range n.exprs {
		n.exprs[i] = foldConstants(expr)
	}
	// Unwrap single-expression blocks when the expression is a constant,
	// so that e.g. (2 + 3) * 4 can be further folded.
	if len(n.exprs) == 1 && isConstant(n.exprs[0]) {
		return n.exprs[0]
	}
	return n
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
	return n
}

func foldPropertyAccess(n propertyAccessNode) astNode {
	n.left = foldConstants(n.left)
	n.right = foldConstants(n.right)
	return n
}
