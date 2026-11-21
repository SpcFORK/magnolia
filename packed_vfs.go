package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Oak's packBundleWithVFS writes [executable][Oak bundle][24-byte bundle
// length][Oak magic][MessagePack map][24-byte map length][VFS magic]. The
// previous native launcher only recognized Oak magic at EOF, so every
// VFS-bearing standalone executable was mistaken for an ordinary interpreter.
// This reader admits only flat string/binary-file maps, exactly the format
// pack.collectVFSFiles produces, with hard length/count limits.
const packedVFSMagic = "vfs \x01\x00\x00\x00"
const maxPackedVFSBytes = 64 << 20
const maxPackedVFSFiles = 4096

func readPackedVFSTrailer(file io.ReaderAt, size int64) (int64, map[string][]byte, error) {
	if size < 32 {
		return size, nil, nil
	}
	var tail [32]byte
	if _, err := file.ReadAt(tail[:], size-32); err != nil {
		return 0, nil, err
	}
	if !bytes.Equal(tail[24:], []byte(packedVFSMagic)) {
		return size, nil, nil // The established no-VFS Oak pack remains unchanged.
	}
	text := string(tail[:24])
	for _, ch := range text {
		if ch != ' ' && (ch < '0' || ch > '9') {
			return 0, nil, fmt.Errorf("invalid packed VFS length")
		}
	}
	vfsLen, err := strconv.ParseInt(strings.TrimSpace(text), 10, 64)
	if err != nil || vfsLen < 1 || vfsLen > maxPackedVFSBytes || vfsLen > size-64 {
		return 0, nil, fmt.Errorf("packed VFS exceeds finite size or has invalid length")
	}
	bundleEnd := size - 32 - vfsLen
	data := make([]byte, int(vfsLen))
	if _, err := file.ReadAt(data, bundleEnd); err != nil {
		return 0, nil, err
	}
	files, err := decodePackedVFSFiles(data)
	if err != nil {
		return 0, nil, err
	}
	return bundleEnd, files, nil
}

type packedVFSReader struct {
	data []byte
	off  int
}

func (r *packedVFSReader) read(n uint64) ([]byte, error) {
	if n > uint64(len(r.data)-r.off) {
		return nil, fmt.Errorf("truncated packed VFS map")
	}
	part := r.data[r.off : r.off+int(n)]
	r.off += int(n)
	return part, nil
}

func (r *packedVFSReader) length(n uint64) (uint64, error) {
	part, err := r.read(n)
	if err != nil {
		return 0, err
	}
	switch n {
	case 1:
		return uint64(part[0]), nil
	case 2:
		return uint64(binary.BigEndian.Uint16(part)), nil
	case 4:
		return uint64(binary.BigEndian.Uint32(part)), nil
	default:
		return 0, fmt.Errorf("unsupported packed VFS length width")
	}
}

func (r *packedVFSReader) fileBytes() ([]byte, error) {
	prefix, err := r.read(1)
	if err != nil {
		return nil, err
	}
	var n uint64
	switch tag := prefix[0]; {
	case tag >= 0xa0 && tag <= 0xbf:
		n = uint64(tag & 31)
	case tag == 0xd9 || tag == 0xc4: // str8 / bin8
		n, err = r.length(1)
	case tag == 0xda || tag == 0xc5: // str16 / bin16
		n, err = r.length(2)
	case tag == 0xdb || tag == 0xc6: // str32 / bin32
		n, err = r.length(4)
	default:
		return nil, fmt.Errorf("packed VFS requires flat string/binary entries")
	}
	if err != nil || n > maxPackedVFSBytes {
		return nil, fmt.Errorf("packed VFS entry exceeds finite size")
	}
	return r.read(n)
}

func decodePackedVFSFiles(data []byte) (map[string][]byte, error) {
	r := packedVFSReader{data: data}
	prefix, err := r.read(1)
	if err != nil {
		return nil, err
	}
	var count uint64
	switch tag := prefix[0]; {
	case tag >= 0x80 && tag <= 0x8f:
		count = uint64(tag & 15)
	case tag == 0xde:
		count, err = r.length(2)
	case tag == 0xdf:
		count, err = r.length(4)
	default:
		return nil, fmt.Errorf("packed VFS must be a file map")
	}
	if err != nil || count > maxPackedVFSFiles {
		return nil, fmt.Errorf("packed VFS file count exceeds bound")
	}
	files := make(map[string][]byte, int(count))
	for i := uint64(0); i < count; i++ {
		key, err := r.fileBytes()
		if err != nil || len(key) == 0 || len(key) > 4096 || bytes.IndexByte(key, 0) >= 0 {
			return nil, fmt.Errorf("invalid packed VFS filename")
		}
		name := string(key)
		if _, exists := files[name]; exists {
			return nil, fmt.Errorf("duplicate packed VFS filename")
		}
		value, err := r.fileBytes()
		if err != nil {
			return nil, err
		}
		files[name] = value
	}
	if r.off != len(data) {
		return nil, fmt.Errorf("packed VFS has trailing bytes")
	}
	return files, nil
}

// The selected native CLI consumes only readFile. Keep this VFS read-only:
// the original classes/movie cannot be replaced by a write from Oak code.
func packedVFSValue(files map[string][]byte) ObjectValue {
	return ObjectValue{
		"readFile": BuiltinFnValue{name: "packedVFS.readFile", fn: func(args []Value) (Value, *runtimeError) {
			if len(args) != 1 {
				return nil, &runtimeError{reason: "packed VFS readFile expects one path"}
			}
			path, ok := args[0].(*StringValue)
			if !ok {
				return nil, &runtimeError{reason: "packed VFS path must be a string"}
			}
			if data, found := files[path.stringContent()]; found {
				return MakeString(string(data)), nil
			}
			return null, nil
		}},
	}
}