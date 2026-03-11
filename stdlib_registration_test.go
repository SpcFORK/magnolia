package main

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestStdlibRegistryCoversLibDirectory(t *testing.T) {
	entries, err := os.ReadDir("lib")
	if err != nil {
		t.Fatalf("failed to read lib directory: %v", err)
	}

	libFiles := make(map[string]struct{})
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if filepath.Ext(name) != ".oak" {
			continue
		}
		base := strings.TrimSuffix(name, ".oak")
		libFiles[base] = struct{}{}
	}

	contentBytes, err := os.ReadFile("lib.go")
	if err != nil {
		t.Fatalf("failed to read lib.go: %v", err)
	}
	content := string(contentBytes)

	embedByVar := map[string]string{}
	embedPairs := regexp.MustCompile(`(?m)^//go:embed\s+lib/([^\r\n]+)\.oak\s*\r?\nvar\s+(lib[a-zA-Z0-9_]+)\s+string\s*$`).FindAllStringSubmatch(content, -1)
	for _, match := range embedPairs {
		embedByVar[match[2]] = match[1]
	}

	mapEntries := regexp.MustCompile(`(?m)^\s*"([^"]+)"\s*:\s*(lib[a-zA-Z0-9_]+),\s*$`).FindAllStringSubmatch(content, -1)
	mapKeys := map[string]string{}
	for _, match := range mapEntries {
		mapKeys[match[1]] = match[2]
	}

	embeddedFiles := map[string]struct{}{}
	for _, file := range embedByVar {
		embeddedFiles[file] = struct{}{}
	}

	mappedFiles := map[string]struct{}{}
	for key, varName := range mapKeys {
		file, ok := embedByVar[varName]
		if !ok {
			t.Fatalf("stdlib key %q references non-embedded variable %q", key, varName)
		}
		mappedFiles[file] = struct{}{}
	}

	missingEmbed := missingKeys(libFiles, embeddedFiles)
	if len(missingEmbed) > 0 {
		t.Fatalf("lib files missing //go:embed in lib.go: %s", strings.Join(missingEmbed, ", "))
	}

	missingMap := missingKeys(embeddedFiles, mappedFiles)
	if len(missingMap) > 0 {
		t.Fatalf("embedded lib files missing stdlibs map entries: %s", strings.Join(missingMap, ", "))
	}
}

func missingKeys(expected map[string]struct{}, actual map[string]struct{}) []string {
	missing := make([]string, 0)
	for key := range expected {
		if _, ok := actual[key]; !ok {
			missing = append(missing, key)
		}
	}
	sort.Strings(missing)
	return missing
}
