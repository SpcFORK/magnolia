package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestPackedVFSReaderAcceptsBoundedBinaryAndLegacyPack(t *testing.T) {
	payload := []byte{0x82,
		0xa6, 'a', 's', 's', 'e', 't', 's',
		0xc4, 0x04, 0x00, 0x80, 0xff, 0x41,
		0xa1, 'x', 0xa1, 'y'}
	legacy := []byte(strings.Repeat("legacy Oak bundle and footer", 3))
	packed := append(append(append([]byte{}, legacy...), payload...),
		[]byte(fmt.Sprintf("%24d%s", len(payload), packedVFSMagic))...)
	end, files, err := readPackedVFSTrailer(bytes.NewReader(packed), int64(len(packed)))
	if err != nil || end != int64(len(legacy)) ||
		!bytes.Equal(files["assets"], []byte{0, 0x80, 0xff, 'A'}) ||
		!bytes.Equal(files["x"], []byte("y")) {
		t.Fatalf("VFS bytes/boundary: end=%d files=%v err=%v", end, files, err)
	}
	readFile := packedVFSValue(files)["readFile"].(BuiltinFnValue)
	got, callErr := readFile.fn([]Value{MakeString("assets")})
	if callErr != nil || !bytes.Equal([]byte(got.(*StringValue).stringContent()), files["assets"]) {
		t.Fatalf("native Oak readFile did not preserve binary: %v", callErr)
	}
	missing, callErr := readFile.fn([]Value{MakeString("missing")})
	if callErr != nil || missing != null {
		t.Fatalf("VFS missing file did not fail closed: %v, %v", missing, callErr)
	}
	end, files, err = readPackedVFSTrailer(bytes.NewReader(legacy), int64(len(legacy)))
	if err != nil || end != int64(len(legacy)) || files != nil {
		t.Fatalf("legacy pack changed: end=%d files=%v err=%v", end, files, err)
	}
}

func TestPackedVFSReaderRejectsCorruptMapsAndTrailer(t *testing.T) {
	cases := [][]byte{
		{0x81, 0xa1, 'x', 0xda, 0, 10, 1},                  // truncated str16 value
		{0x82, 0xa1, 'x', 0xa1, '1', 0xa1, 'x', 0xa1, '2'}, // duplicate
		{0x81, 0xa1, 'x', 0xa1, '1', 0x00},                 // trailing data
		{0x81, 0xa1, 'x', 0xc4, 0xff},                      // truncated binary
		{0x91, 0xa1, 'x'},                                  // array is not file map
	}
	for _, data := range cases {
		if _, err := decodePackedVFSFiles(data); err == nil {
			t.Errorf("corrupt packed VFS was accepted: %x", data)
		}
	}
	payload := []byte{0x80}
	for _, text := range []string{"invalid", "-1", "67108865", "99999"} {
		packed := append(append([]byte("bundle"), payload...),
			[]byte(fmt.Sprintf("%24s%s", text, packedVFSMagic))...)
		_, _, err := readPackedVFSTrailer(bytes.NewReader(packed), int64(len(packed)))
		if err == nil {
			t.Errorf("bad VFS length %q accepted", text)
		}
	}
	if _, err := decodePackedVFSFiles([]byte{0xdf, 0, 0, 0, 0x01}); err == nil {
		t.Fatal("truncated map32 accepted")
	}
	if _, err := decodePackedVFSFiles([]byte{0xde, 0x10, 0x01}); err == nil ||
		!strings.Contains(err.Error(), "count") {
		t.Fatal("over-bound map16 accepted")
	}
}