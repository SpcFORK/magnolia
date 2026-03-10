package main

import "testing"

func TestPosString(t *testing.T) {
	p := pos{line: 12, col: 34}
	if got := p.String(); got != "[12:34]" {
		t.Fatalf("unexpected pos string: %q", got)
	}
}

func TestTokenString(t *testing.T) {
	tests := []struct {
		name string
		tok  token
		want string
	}{
		{name: "unknown", tok: token{kind: unknown}, want: "(unknown token)"},
		{name: "comment", tok: token{kind: comment, payload: "hello"}, want: "//(hello)"},
		{name: "comma", tok: token{kind: comma}, want: ","},
		{name: "dot", tok: token{kind: dot}, want: "."},
		{name: "left paren", tok: token{kind: leftParen}, want: "("},
		{name: "right paren", tok: token{kind: rightParen}, want: ")"},
		{name: "left bracket", tok: token{kind: leftBracket}, want: "["},
		{name: "right bracket", tok: token{kind: rightBracket}, want: "]"},
		{name: "left brace", tok: token{kind: leftBrace}, want: "{"},
		{name: "right brace", tok: token{kind: rightBrace}, want: "}"},
		{name: "assign", tok: token{kind: assign}, want: ":="},
		{name: "nonlocal assign", tok: token{kind: nonlocalAssign}, want: "<-"},
		{name: "pipe", tok: token{kind: pipeArrow}, want: "|>"},
		{name: "branch", tok: token{kind: branchArrow}, want: "->"},
		{name: "push", tok: token{kind: pushArrow}, want: "<<"},
		{name: "colon", tok: token{kind: colon}, want: ":"},
		{name: "ellipsis", tok: token{kind: ellipsis}, want: "..."},
		{name: "question", tok: token{kind: qmark}, want: "?"},
		{name: "exclam", tok: token{kind: exclam}, want: "!"},
		{name: "tilde", tok: token{kind: tilde}, want: "~"},
		{name: "plus", tok: token{kind: plus}, want: "+"},
		{name: "minus", tok: token{kind: minus}, want: "-"},
		{name: "times", tok: token{kind: times}, want: "*"},
		{name: "divide", tok: token{kind: divide}, want: "/"},
		{name: "modulus", tok: token{kind: modulus}, want: "%"},
		{name: "power", tok: token{kind: power}, want: "**"},
		{name: "xor", tok: token{kind: xor}, want: "^"},
		{name: "and", tok: token{kind: and}, want: "&"},
		{name: "or", tok: token{kind: or}, want: "|"},
		{name: "greater", tok: token{kind: greater}, want: ">"},
		{name: "less", tok: token{kind: less}, want: "<"},
		{name: "eq", tok: token{kind: eq}, want: "="},
		{name: "deep eq", tok: token{kind: deepEq}, want: "=="},
		{name: "geq", tok: token{kind: geq}, want: ">="},
		{name: "leq", tok: token{kind: leq}, want: "<="},
		{name: "neq", tok: token{kind: neq}, want: "!="},
		{name: "rshift", tok: token{kind: rshift}, want: ">>"},
		{name: "if keyword", tok: token{kind: ifKeyword}, want: "if"},
		{name: "fn keyword", tok: token{kind: fnKeyword}, want: "fn"},
		{name: "with keyword", tok: token{kind: withKeyword}, want: "with"},
		{name: "cs keyword", tok: token{kind: csKeyword}, want: "cs"},
		{name: "underscore", tok: token{kind: underscore}, want: "_"},
		{name: "identifier", tok: token{kind: identifier, payload: "name"}, want: "var(name)"},
		{name: "true literal", tok: token{kind: trueLiteral}, want: "true"},
		{name: "string", tok: token{kind: stringLiteral, payload: "a\nb"}, want: "string(\"a\\nb\")"},
		{name: "number", tok: token{kind: numberLiteral, payload: "12.5"}, want: "number(12.5)"},
		{name: "false literal", tok: token{kind: falseLiteral}, want: "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tok.String(); got != tt.want {
				t.Fatalf("unexpected token string: got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTokenizerPeekAheadEOF(t *testing.T) {
	tk := newTokenizer("ab", "x.oak")
	if got := tk.peekAhead(5); got != ' ' {
		t.Fatalf("expected EOF sentinel whitespace, got %q", got)
	}
}
