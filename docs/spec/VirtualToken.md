# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\VirtualToken.oak`

- `std` · `import(...)`
- `default` — constant
- `merge` — constant
### `at(index, line, col)`

> returns `:list`

### `defaultPos(pos)`

### `token(type, val, pos)`

> returns `:object`

### `node(type, tok, fields)`

### `objectEntry(key, val)`

> returns `:object`

### `cs NullNode(tok)`

### `cs EmptyNode(tok)`

### `cs IntNode(val, tok)`

### `cs FloatNode(val, tok)`

### `cs StringNode(val, tok)`

### `cs AtomNode(val, tok)`

### `cs BoolNode(val, tok)`

### `cs IdentifierNode(name, tok)`

### `cs ListNode(elems, tok)`

### `cs ObjectNode(entries, tok)`

### `cs UnaryNode(op, right, tok)`

### `cs UnaryNegate(right, tok)`

### `cs UnaryNot(right, tok)`

### `cs UnaryBitNot(right, tok)`

### `cs BinaryNode(op, left, right, tok)`

### `cs BinaryAdd(left, right, tok)`

### `cs BinarySub(left, right, tok)`

### `cs BinaryMul(left, right, tok)`

### `cs BinaryDiv(left, right, tok)`

### `cs BinaryMod(left, right, tok)`

### `cs BinaryPow(left, right, tok)`

### `cs BinaryAnd(left, right, tok)`

### `cs BinaryOr(left, right, tok)`

### `cs BinaryXor(left, right, tok)`

### `cs BinaryEq(left, right, tok)`

### `cs BinaryNeq(left, right, tok)`

### `cs BinaryGreater(left, right, tok)`

### `cs BinaryLess(left, right, tok)`

### `cs BinaryGeq(left, right, tok)`

### `cs BinaryLeq(left, right, tok)`

### `cs AssignmentNode(left, right, tok)`

### `cs PropertyAccessNode(left, right, tok)`

### `cs FnCallNode(function, args, restArg, tok)`

### `cs FunctionNode(args, body, restArg, name, tok)`

### `cs IfBranchNode(target, body, tok)`

### `cs IfExprNode(cond, branches, tok)`

### `cs BlockNode(exprs, tok)`

### `cs Comment(payload, pos)`

### `cs Comma(pos)`

### `cs Dot(pos)`

### `cs LeftParen(pos)`

### `cs RightParen(pos)`

### `cs LeftBracket(pos)`

### `cs RightBracket(pos)`

### `cs LeftBrace(pos)`

### `cs RightBrace(pos)`

### `cs Assign(pos)`

### `cs NonlocalAssign(pos)`

### `cs PipeArrow(pos)`

### `cs BranchArrow(pos)`

### `cs PushArrow(pos)`

### `cs Colon(pos)`

### `cs Ellipsis(pos)`

### `cs Qmark(pos)`

### `cs Exclam(pos)`

### `cs Tilde(pos)`

### `cs Plus(pos)`

### `cs Minus(pos)`

### `cs Times(pos)`

### `cs Divide(pos)`

### `cs Modulus(pos)`

### `cs Power(pos)`

### `cs Xor(pos)`

### `cs And(pos)`

### `cs Or(pos)`

### `cs Greater(pos)`

### `cs Less(pos)`

### `cs Eq(pos)`

### `cs DeepEq(pos)`

### `cs Geq(pos)`

### `cs Leq(pos)`

### `cs Neq(pos)`

### `cs Rshift(pos)`

### `cs IfKeyword(pos)`

### `cs FnKeyword(pos)`

### `cs WithKeyword(pos)`

### `cs CsKeyword(pos)`

### `cs Underscore(pos)`

### `cs Identifier(name, pos)`

### `cs TrueLiteral(pos)`

### `cs FalseLiteral(pos)`

### `cs StringLiteral(value, pos)`

### `cs NumberLiteral(value, pos)`

