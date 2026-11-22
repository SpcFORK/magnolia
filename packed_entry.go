package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// A native Oak bundle's generic renderer cannot yet round-trip the large
// selected JVM adapter. An explicitly marked packed application may instead
// interpret its original Oak source and imports directly from the immutable
// appended VFS. It never extracts source/class/movie bytes to disk and never
// resolves an Oak import against the host filesystem.
const packedSourceEntry = "__oak_entry__"
const packedSourcesZip = "__oak_sources.zip"

// The source ZIP is a small, bounded part of the executable's read-only VFS.
// It is decoded in memory solely for Oak import resolution; no byte is
// extracted to the working directory or to /tmp.
func unpackPackedOakSources(files map[string][]byte) (map[string][]byte, error) {
	combined := make(map[string][]byte, len(files)+256)
	for name, data := range files {
		combined[name] = data
	}
	archive, present := files[packedSourcesZip]
	if !present {
		return combined, nil
	}
	reader, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
	if err != nil || len(reader.File) > 512 {
		return nil, fmt.Errorf("invalid or oversized packed Oak source archive")
	}
	total := uint64(0)
	for _, member := range reader.File {
		name := member.Name
		if !safePackedFilename(name) || member.FileInfo().IsDir() ||
			member.UncompressedSize64 > 1<<20 || total+member.UncompressedSize64 > 8<<20 {
			return nil, fmt.Errorf("unsafe or oversized packed Oak source %q", name)
		}
		if _, exists := combined[name]; exists {
			return nil, fmt.Errorf("duplicate packed Oak source %q", name)
		}
		opened, err := member.Open()
		if err != nil {
			return nil, fmt.Errorf("open packed Oak source %q: %w", name, err)
		}
		data, readErr := io.ReadAll(io.LimitReader(opened, (1<<20)+1))
		closeErr := opened.Close()
		if readErr != nil || closeErr != nil || uint64(len(data)) != member.UncompressedSize64 {
			return nil, fmt.Errorf("invalid packed Oak source bytes for %q", name)
		}
		combined[name] = data
		total += uint64(len(data))
	}
	return combined, nil
}

func safePackedFilename(name string) bool {
	if name == "" || strings.ContainsAny(name, "\\:\x00") || strings.HasPrefix(name, "/") {
		return false
	}
	for _, part := range strings.Split(name, "/") {
		if part == "" || part == "." || part == ".." {
			return false
		}
	}
	return true
}

func runPackedSourceEntry(ctx *Context, files map[string][]byte) (bool, error) {
	manifest, exists := files[packedSourceEntry]
	if !exists {
		return false, nil // Ordinary native bundle packs still use their bundle.
	}
	allFiles, err := unpackPackedOakSources(files)
	if err != nil {
		return true, err
	}
	for name := range allFiles {
		if name != packedSourceEntry && !safePackedFilename(name) {
			return true, fmt.Errorf("unsafe packed Oak source or resource name %q", name)
		}
	}
	entry := strings.TrimSpace(string(manifest))
	if !safePackedFilename(entry) || len(entry) > 4096 {
		return true, fmt.Errorf("invalid packed Oak entry path")
	}
	source, found := allFiles[entry]
	if !found || len(source) == 0 || len(source) > 1<<20 {
		return true, fmt.Errorf("packed Oak entry missing or oversized")
	}
	// This is only a virtual import root; no directory or source is written.
	root := filepath.Join(os.TempDir(), "magnolia-packed-oak-virtual-root")
	ctx.eng.packedRoot = root
	ctx.eng.packedFiles = allFiles
	ctx.currentFile = filepath.Join(root, filepath.FromSlash(entry))
	ctx.rootPath = filepath.Dir(ctx.currentFile)
	ctx.LoadFunc("___packed_vfs", func(args []Value) (Value, *runtimeError) {
		if len(args) != 0 {
			return nil, &runtimeError{reason: "___packed_vfs takes no arguments"}
		}
		return packedVFSValue(allFiles), nil
	})
	_, err = ctx.Eval(bytes.NewReader(source))
	return true, err
}