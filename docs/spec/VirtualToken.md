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

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

- `__std_builtin_int` — constant
- `__std_builtin_float` — constant
- `__std_builtin_string` — constant
- `__std_builtin_type` — constant
- `__std_builtin_len` — constant
- `__std_builtin_char` — constant
- `__std_builtin_codepoint` — constant
- `__std_builtin_print` — constant
- `__std_builtin_keys` — constant
- `__std_builtin_atom` — constant
- `__std_builtin_name` — constant
- `__std_builtin_rand` — constant
- `__std_builtin_srand` — constant
- `__std_builtin_input` — constant
- `__std_builtin_exit` — constant
- `__std_builtin_args` — constant
- `__std_builtin_addr` — constant
- `__std_builtin_bits` — constant
- `__std_builtin_memread` — constant
- `__std_builtin_memwrite` — constant
- `__std_builtin_utf16` — constant
- `__std_builtin_go` — constant
- `__std_builtin_wait` — constant
- `__std_builtin_exec` — constant
- `__std_builtin_env` — constant
- `__std_builtin_write` — constant
- `__std_builtin_read` — constant
- `__std_builtin_interpreter` — constant
- `__std_builtin_pointer` — constant
- `__std_builtin_log` — constant
- `__std_builtin_time` — constant
- `__std_builtin_nanotime` — constant
- `int` — constant
- `float` — constant
- `string` — constant
- `type` — constant
- `len` — constant
- `print` — constant
- `time` — constant
- `nanotime` — constant
- `char` — constant
- `codepoint` — constant
- `keys` — constant
- `atom` — constant
- `name` — constant
- `rand` — constant
- `srand` — constant
- `input` — constant
- `args` — constant
- `addr` — constant
- `bits` — constant
- `memread` — constant
- `memwrite` — constant
- `utf16` — constant
- `exec` — constant
- `env` — constant
- `write` — constant
- `read` — constant
- `interpreter` — constant
- `pointer` — constant
- `log` — constant
- `wait` — constant
- `__std_builtin_sin` — constant
- `__std_builtin_cos` — constant
- `__std_builtin_tan` — constant
- `__std_builtin_asin` — constant
- `__std_builtin_acos` — constant
- `__std_builtin_atan` — constant
- `__std_builtin_pow` — constant
- `sin` — constant
- `cos` — constant
- `tan` — constant
- `asin` — constant
- `acos` — constant
- `atan` — constant
- `pow` — constant
