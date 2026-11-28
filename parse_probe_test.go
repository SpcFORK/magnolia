package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParseProbe(t *testing.T) {
	srcBytes, err := os.ReadFile(os.Args[0])
	_ = srcBytes
	_ = err
	path := "lib/codegen-c.oak"
	if len(os.Args) > 0 {
		if p := os.Getenv("PROBE_FILE"); p != "" {
			path = p
		}
	}
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	tk := newTokenizer(string(src), path)
	toks := tk.tokenize()
	p := newParser(toks)
	nodes, perr := p.parse()
	if perr != nil {
		fmt.Println("PARSE ERROR:", perr)
		os.Exit(1)
	}
	fmt.Println("OK, statements:", len(nodes))
}
