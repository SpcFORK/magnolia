package main

import (
	"strings"
	"testing"
)

func tok(kind tokKind, line, col int) *token {
	return &token{kind: kind, pos: pos{line: line, col: col}}
}

func tokv(kind tokKind, payload string, line, col int) token {
	return token{kind: kind, payload: payload, pos: pos{line: line, col: col}}
}

func TestAstNodeStringHelpers(t *testing.T) {
	baseTok := tok(identifier, 1, 1)

	strNode := stringNode{payload: []byte("hello"), tok: baseTok}
	intN := intNode{payload: 42, tok: baseTok}
	floatN := floatNode{payload: 3.5, tok: baseTok}
	trueN := boolNode{payload: true, tok: baseTok}
	falseN := boolNode{payload: false, tok: baseTok}
	atomN := atomNode{payload: "ok", tok: baseTok}
	identN := identifierNode{payload: "x", tok: baseTok}

	tests := []struct {
		name string
		node astNode
		want string
	}{
		{name: "empty", node: emptyNode{tok: baseTok}, want: "_"},
		{name: "null", node: nullNode{tok: baseTok}, want: "?"},
		{name: "string", node: strNode, want: "\"hello\""},
		{name: "int", node: intN, want: "42"},
		{name: "float", node: floatN, want: "3.5"},
		{name: "bool true", node: trueN, want: "true"},
		{name: "bool false", node: falseN, want: "false"},
		{name: "atom", node: atomN, want: ":ok"},
		{name: "identifier", node: identN, want: "x"},
		{name: "list", node: listNode{elems: []astNode{intN, atomN}, tok: baseTok}, want: "[42, :ok]"},
		{name: "object", node: objectNode{entries: []objectEntry{{key: identN, val: strNode}}, tok: baseTok}, want: "{ x: \"hello\" }"},
		{name: "assignment local", node: assignmentNode{isLocal: true, left: identN, right: intN, tok: baseTok}, want: "x := 42"},
		{name: "assignment nonlocal", node: assignmentNode{isLocal: false, left: identN, right: intN, tok: baseTok}, want: "x <- 42"},
		{name: "property access", node: propertyAccessNode{left: identN, right: identifierNode{payload: "y", tok: baseTok}, tok: baseTok}, want: "(x.y)"},
		{name: "unary", node: unaryNode{op: minus, right: intN, tok: baseTok}, want: "-42"},
		{name: "binary", node: binaryNode{op: plus, left: intN, right: intNode{payload: 1, tok: baseTok}, tok: baseTok}, want: "(42 + 1)"},
		{name: "fn call", node: fnCallNode{fn: identN, args: []astNode{intN}, restArg: atomN, tok: baseTok}, want: "call[x](42, :ok...)"},
		{name: "if expr", node: ifExprNode{cond: identN, branches: []ifBranch{{target: atomN, body: intN}}, tok: baseTok}, want: "if x {:ok -> 42}"},
		{name: "block", node: blockNode{exprs: []astNode{intN, identN}, tok: baseTok}, want: "{ 42, x }"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.String(); got != tt.want {
				t.Fatalf("unexpected string form: got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFnAndClassNodeString(t *testing.T) {
	baseTok := tok(identifier, 1, 1)
	body := blockNode{exprs: []astNode{intNode{payload: 1, tok: baseTok}}, tok: baseTok}

	anon := fnNode{args: []string{"a"}, restArg: "rest", body: body, tok: baseTok}
	if got := anon.String(); got != "fn(a, rest...) { 1 }" {
		t.Fatalf("unexpected anonymous fn string: %q", got)
	}

	named := fnNode{name: "sum", args: []string{"x", "y"}, body: body, tok: baseTok}
	if got := named.String(); got != "fn sum(x, y) { 1 }" {
		t.Fatalf("unexpected named fn string: %q", got)
	}

	plainClass := classNode{name: "User", args: []string{"id"}, body: body, tok: baseTok}
	if got := plainClass.String(); got != "cs User(id) { 1 }" {
		t.Fatalf("unexpected class string without parents: %q", got)
	}

	derived := classNode{
		name:        "Admin",
		args:        []string{"id"},
		restArg:     "rest",
		staticExprs: []astNode{identifierNode{payload: "x", tok: baseTok}},
		parents:     []astNode{identifierNode{payload: "Base", tok: baseTok}},
		body:        body,
		tok:         baseTok,
	}
	got := derived.String()
	if !strings.Contains(got, "cs Admin(id, rest...)") {
		t.Fatalf("missing class head in %q", got)
	}
	if !strings.Contains(got, "(Base) -> { 1 }") {
		t.Fatalf("missing inheritance body in %q", got)
	}
}

func TestParserPeekAheadAndReadUntilTokenKind(t *testing.T) {
	p := newParser([]token{{kind: identifier}, {kind: plus}, {kind: numberLiteral}})

	if got := p.peekAhead(1).kind; got != plus {
		t.Fatalf("expected plus token, got %v", got)
	}
	if got := p.peekAhead(10).kind; got != comma {
		t.Fatalf("expected comma sentinel, got %v", got)
	}

	read := p.readUntilTokenKind(numberLiteral)
	if len(read) != 2 {
		t.Fatalf("expected 2 tokens read, got %d", len(read))
	}
	if p.peek().kind != numberLiteral {
		t.Fatalf("expected parser to stop on delimiter token")
	}
}

func TestParseErrorError(t *testing.T) {
	err := parseError{reason: "bad syntax", pos: pos{line: 4, col: 2}}
	if got := err.Error(); got != "Parse error at [4:2]: bad syntax" {
		t.Fatalf("unexpected parseError.Error output: %q", got)
	}
}

func TestBlockNodePos(t *testing.T) {
	blockTok := tok(leftBrace, 8, 3)
	n := blockNode{tok: blockTok}
	p := n.pos()
	if p.line != 8 || p.col != 3 {
		t.Fatalf("unexpected block position: %+v", p)
	}
}

func TestParserExpect(t *testing.T) {
	p := newParser([]token{{kind: identifier, payload: "x", pos: pos{line: 1, col: 1}}})

	tok, err := p.expect(identifier)
	if err != nil {
		t.Fatalf("expected successful expect, got error: %v", err)
	}
	if tok.kind != identifier {
		t.Fatalf("unexpected token kind: %v", tok.kind)
	}

	p = newParser([]token{{kind: numberLiteral, payload: "1", pos: pos{line: 3, col: 4}}})
	_, err = p.expect(identifier)
	if err == nil {
		t.Fatalf("expected mismatch error")
	}
	if !strings.Contains(err.Error(), "Unexpected token number(1), expected var()") {
		t.Fatalf("unexpected mismatch message: %q", err.Error())
	}

	p = newParser([]token{})
	_, err = p.expect(identifier)
	if err == nil {
		t.Fatalf("expected EOF error")
	}
	if !strings.Contains(err.Error(), "Unexpected end of input") {
		t.Fatalf("unexpected EOF error message: %q", err.Error())
	}
}

func TestParseUnitErrors(t *testing.T) {
	p := newParser([]token{
		tokv(colon, "", 1, 1),
		tokv(comma, "", 1, 2),
	})
	_, err := p.parseUnit()
	if err == nil || !strings.Contains(err.Error(), "Expected identifier after ':'") {
		t.Fatalf("expected atom parse error, got %v", err)
	}

	p = newParser([]token{
		tokv(leftBracket, "", 1, 1),
		tokv(numberLiteral, "1", 1, 2),
		tokv(numberLiteral, "2", 1, 3),
		tokv(rightBracket, "", 1, 4),
	})
	_, err = p.parseUnit()
	if err == nil || !strings.Contains(err.Error(), "expected comma or ]") {
		t.Fatalf("expected list comma error, got %v", err)
	}

	p = newParser([]token{
		tokv(leftBrace, "", 1, 1),
		tokv(identifier, "x", 1, 2),
	})
	_, err = p.parseUnit()
	if err == nil || !strings.Contains(err.Error(), "Unexpected end of input inside block or object") {
		t.Fatalf("expected block/object EOF error, got %v", err)
	}
}

func TestParseNodePipeAndWithErrors(t *testing.T) {
	p := newParser([]token{
		tokv(identifier, "x", 1, 1),
		tokv(pipeArrow, "", 1, 2),
		tokv(identifier, "y", 1, 3),
	})
	_, err := p.parseNode()
	if err == nil || !strings.Contains(err.Error(), "Expected function call after |>") {
		t.Fatalf("expected pipe call error, got %v", err)
	}

	p = newParser([]token{
		tokv(withKeyword, "", 1, 1),
		tokv(identifier, "x", 1, 2),
		tokv(numberLiteral, "1", 1, 3),
	})
	_, err = p.parseNode()
	if err == nil || !strings.Contains(err.Error(), "with keyword should be followed by a function call") {
		t.Fatalf("expected with-call parse error, got %v", err)
	}
}

func TestParseTopLevelMissingComma(t *testing.T) {
	p := newParser([]token{
		tokv(identifier, "x", 1, 1),
		tokv(identifier, "y", 1, 2),
	})
	_, err := p.parse()
	if err == nil || !strings.Contains(err.Error(), "Unexpected token var(y), expected ,") {
		t.Fatalf("expected top-level comma error, got %v", err)
	}
}
