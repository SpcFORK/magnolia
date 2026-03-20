# Virtual Token Constructors (VirtualToken)

## Overview

`VirtualToken` provides factory functions for programmatically constructing Oak AST nodes and token objects without running the tokenizer or parser. It is used by code-generation tools, AST transformers, and the macro system to synthesize well-formed AST trees.

## Import

```oak
vt := import('VirtualToken')
// or destructure specific constructors
{ IdentifierNode: IdentifierNode, BinaryAdd: BinaryAdd, FnCallNode: FnCallNode } := import('VirtualToken')
```

## Position Helpers

### `at(index, line, col)`

Creates a position triple `[index, line, col]`.

```oak
at(0, 1, 1)  // => [0, 1, 1]
```

### Token and Node Primitives

#### `token(type, val, pos?)`

Creates a raw token object. `pos` defaults to `[0, 1, 1]`.

```oak
token(:identifier, 'foo', at(0, 1, 1))
```

#### `node(type, tok, fields?)`

Creates a bare AST node merged with `fields`.

```oak
node(:string, ?, { val: 'hello' })
```

#### `objectEntry(key, val)`

Creates an object entry pair for use inside `ObjectNode`.

```oak
objectEntry(StringNode('x'), IntNode(42))
```

## Literal Node Constructors

All constructors accept an optional `tok` that defaults to `?`.

| Constructor             | Node type    | Extra fields          |
|-------------------------|--------------|-----------------------|
| `NullNode(tok?)`        | `:null`      | —                     |
| `EmptyNode(tok?)`       | `:empty`     | —                     |
| `IntNode(val, tok?)`    | `:int`       | `val`                 |
| `FloatNode(val, tok?)`  | `:float`     | `val`                 |
| `StringNode(val, tok?)` | `:string`    | `val`                 |
| `AtomNode(val, tok?)`   | `:atom`      | `val`                 |
| `BoolNode(val, tok?)`   | `:bool`      | `val`                 |
| `IdentifierNode(name, tok?)` | `:identifier` | `val: name`    |

```oak
IdentifierNode('x')           // { type: :identifier, val: 'x', tok: ? }
StringNode('hello')           // { type: :string, val: 'hello', tok: ? }
IntNode(42)                   // { type: :int, val: 42, tok: ? }
```

## Composite Node Constructors

| Constructor                   | Node type   | Extra fields                        |
|-------------------------------|-------------|-------------------------------------|
| `ListNode(elems?, tok?)`      | `:list`     | `elems` (default `[]`)              |
| `ObjectNode(entries?, tok?)`  | `:object`   | `entries` (default `[]`)            |

```oak
ListNode([IntNode(1), IntNode(2)])
ObjectNode([objectEntry(StringNode('a'), IntNode(1))])
```

## Unary Node Constructors

| Constructor                  | Op       | Default token |
|------------------------------|----------|---------------|
| `UnaryNode(op, right, tok?)` | any      | —             |
| `UnaryNegate(right, tok?)`   | `:minus` | `Minus()`     |
| `UnaryNot(right, tok?)`      | `:exclam`| `Exclam()`    |
| `UnaryBitNot(right, tok?)`   | `:tilde` | `Tilde()`     |

## Binary Node Constructors

| Constructor                         | Op         |
|-------------------------------------|------------|
| `BinaryNode(op, left, right, tok?)` | any        |
| `BinaryAdd(left, right, tok?)`      | `:plus`    |
| `BinarySub(left, right, tok?)`      | `:minus`   |
| `BinaryMul(left, right, tok?)`      | `:times`   |
| `BinaryDiv(left, right, tok?)`      | `:divide`  |
| `BinaryMod(left, right, tok?)`      | `:modulus` |
| `BinaryPow(left, right, tok?)`      | `:power`   |
| `BinaryAnd(left, right, tok?)`      | `:and`     |
| `BinaryOr(left, right, tok?)`       | `:or`      |
| `BinaryXor(left, right, tok?)`      | `:xor`     |
| `BinaryEq(left, right, tok?)`       | `:eq`      |
| `BinaryNeq(left, right, tok?)`      | `:neq`     |
| `BinaryGreater(left, right, tok?)`  | `:greater` |
| `BinaryLess(left, right, tok?)`     | `:less`    |
| `BinaryGeq(left, right, tok?)`      | `:geq`     |
| `BinaryLeq(left, right, tok?)`      | `:leq`     |

```oak
BinaryAdd(IdentifierNode('x'), IntNode(1))
// => { type: :binary, op: :plus, left: {type: :identifier, val: 'x'}, right: {type: :int, val: 1} }
```

## Structural Node Constructors

| Constructor                                        | Node type         |
|----------------------------------------------------|-------------------|
| `AssignmentNode(left, right, tok?)`                | `:assignment`     |
| `PropertyAccessNode(left, right, tok?)`            | `:propertyAccess` |
| `FnCallNode(function, args?, restArg?, tok?)`      | `:fnCall`         |
| `FunctionNode(args?, body, restArg?, name?, tok?)` | `:function`       |
| `IfBranchNode(target, body, tok?)`                 | `:ifBranch`       |
| `IfExprNode(cond, branches?, tok?)`                | `:ifExpr`         |
| `BlockNode(exprs?, tok?)`                          | `:block`          |

```oak
FnCallNode(IdentifierNode('print'), [StringNode('hello')])
```

## Token Constructors

Punctuation and keyword tokens for use when a real `tok` is needed (e.g. as the `tok` field for a synthesized node).

```oak
Comma(), Dot(), LeftParen(), RightParen()
LeftBracket(), RightBracket(), LeftBrace(), RightBrace()
Assign(), NonlocalAssign(), PipeArrow(), BranchArrow(), PushArrow()
Colon(), Ellipsis(), Qmark(), Exclam(), Tilde()
Plus(), Minus(), Times(), Divide(), Modulus(), Power()
Xor(), And(), Or(), Greater(), Less(), Eq(), DeepEq(), Geq(), Leq(), Neq(), Rshift()
IfKeyword(), FnKeyword(), WithKeyword(), CsKeyword(), Underscore()
TrueLiteral(), FalseLiteral()
// Value-carrying tokens:
Comment(payload?, pos?)
Identifier(name, pos?)
StringLiteral(value, pos?)
NumberLiteral(value, pos?)
```

All accept an optional `pos` argument (a position triple from `at()`).
