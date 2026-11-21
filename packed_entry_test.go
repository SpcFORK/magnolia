package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestPackedSourceEntryImportsOnlyEmbeddedOakFiles(t *testing.T) {
	ctx := NewContextWithCwd()
	defer ctx.Wait()
	ctx.LoadBuiltins()
	files := map[string][]byte{
		packedSourceEntry: []byte("tools/main.oak"),
		"tools/main.oak":  []byte("helper := import('../lib/helper.oak')\nresult := helper.answer\n"),
		"lib/helper.oak":  []byte("answer := 42\n"),
	}
	handled, err := runPackedSourceEntry(&ctx, files)
	if !handled || err != nil {
		t.Fatalf("embedded source/import failed: %v", err)
	}
	result, errScope := ctx.scope.get("result")
	if errScope != nil || result != IntValue(42) {
		t.Fatalf("embedded source returned %v, %v", result, errScope)
	}

	// An existing host file must not be imported into an embedded app.
	host, err := os.CreateTemp("", "oak-packed-host-*.oak")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(host.Name())
	host.WriteString("answer := 999\n")
	host.Close()
	ctx2 := NewContextWithCwd()
	defer ctx2.Wait()
	ctx2.LoadBuiltins()
	files["tools/main.oak"] = []byte(fmt.Sprintf("import(%q)\n", host.Name()))
	if handled, err := runPackedSourceEntry(&ctx2, files); !handled || err == nil {
		t.Fatal("packed source imported an existing host file")
	}
}

func TestPackedSourceEntryReadsCompressedOakImportsInMemory(t *testing.T) {
	var buffer bytes.Buffer
	archive := zip.NewWriter(&buffer)
	for name, source := range map[string]string{
		"tools/main.oak": "helper := import('../lib/helper.oak')\nembedded := ___packed_vfs()\nresult := helper.answer + len(embedded.readFile('original/test'))\n",
		"lib/helper.oak": "answer := 71\n",
	} {
		entry, err := archive.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := entry.Write([]byte(source)); err != nil {
			t.Fatal(err)
		}
	}
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	ctx := NewContextWithCwd()
	defer ctx.Wait()
	ctx.LoadBuiltins()
	files := map[string][]byte{
		packedSourceEntry: []byte("tools/main.oak"),
		packedSourcesZip:  buffer.Bytes(),
		"original/test":   []byte{0, 0xff},
	}
	if handled, err := runPackedSourceEntry(&ctx, files); !handled || err != nil {
		t.Fatalf("compressed Oak source pack: handled=%v err=%v", handled, err)
	}
	result, scopeErr := ctx.scope.get("result")
	if scopeErr != nil || result != IntValue(73) {
		t.Fatalf("compressed VFS import returned %v, %v", result, scopeErr)
	}
	if _, present := files["lib/helper.oak"]; present {
		t.Fatal("unpack mutated original immutable VFS map")
	}

	var bad bytes.Buffer
	broken := zip.NewWriter(&bad)
	member, err := broken.Create("../host.oak")
	if err != nil {
		t.Fatal(err)
	}
	member.Write([]byte("not allowed"))
	broken.Close()
	if _, err := unpackPackedOakSources(map[string][]byte{
		packedSourcesZip: bad.Bytes(),
	}); err == nil {
		t.Fatal("ZIP path traversal was admitted")
	}
}

func TestPackedSourceEntryRejectsPathsAndMissingSource(t *testing.T) {
	for _, unsafe := range []string{"/absolute", "../outside", "inside/../outside",
		"a\\b", "C:/escape", "", "a//b", "a/./b"} {
		if safePackedFilename(unsafe) {
			t.Errorf("unsafe VFS path accepted: %q", unsafe)
		}
	}
	if !safePackedFilename("tools/fliplay-oak-cli.oak") {
		t.Fatal("valid selected entry refused")
	}
	ctx := NewContextWithCwd()
	ctx.LoadBuiltins()
	if handled, err := runPackedSourceEntry(&ctx, nil); handled || err != nil {
		t.Fatalf("ordinary packed bundle changed: handled=%v err=%v", handled, err)
	}
	for _, files := range []map[string][]byte{
		{packedSourceEntry: []byte("../outside")},
		{packedSourceEntry: []byte("tools/missing.oak")},
		{packedSourceEntry: []byte("tools/main.oak"), "../outside": []byte("wrong")},
	} {
		if handled, err := runPackedSourceEntry(&ctx, files); !handled || err == nil {
			t.Fatalf("bad source manifest accepted: %#v", files)
		}
	}
}