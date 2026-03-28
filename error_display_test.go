package main

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func captureStderr(t *testing.T, fn func()) string {
	t.Helper()

	original := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stderr pipe: %v", err)
	}

	os.Stderr = w
	defer func() {
		os.Stderr = original
	}()

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close write pipe: %v", err)
	}
	data, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read captured stderr: %v", err)
	}
	return string(data)
}

func TestDefaultErrorConfig(t *testing.T) {
	config := DefaultErrorConfig()
	if !config.UseColor {
		t.Fatalf("expected UseColor to default to true")
	}
	if !config.ShowContext {
		t.Fatalf("expected ShowContext to default to true")
	}
	if config.ContextLines != 2 {
		t.Fatalf("expected ContextLines to be 2, got %d", config.ContextLines)
	}
	if !config.ShowStackTrace {
		t.Fatalf("expected ShowStackTrace to default to true")
	}
}

func TestFormatError(t *testing.T) {
	if got := FormatError(nil); got != "" {
		t.Fatalf("expected empty string for nil error, got %q", got)
	}

	pe := parseError{
		reason: "unexpected token",
		pos: pos{
			fileName: "sample.oak",
			line:     3,
			col:      7,
		},
	}
	if got := FormatError(pe); got != "Parse Error in sample.oak at [3:7]: unexpected token" {
		t.Fatalf("unexpected parse error format: %q", got)
	}

	re := &runtimeError{
		reason: "bad call",
		pos: pos{
			line: 9,
			col:  2,
		},
		stackTrace: []stackEntry{
			{name: "outer", pos: pos{line: 1, col: 1}},
			{name: "", pos: pos{line: 2, col: 5}},
		},
	}
	formatted := FormatError(re)
	if !strings.Contains(formatted, "Runtime Error at [9:2]: bad call") {
		t.Fatalf("unexpected runtime error header: %q", formatted)
	}
	if !strings.Contains(formatted, "in fn outer [1:1]") {
		t.Fatalf("missing named stack entry: %q", formatted)
	}
	if !strings.Contains(formatted, "in anonymous fn [2:5]") {
		t.Fatalf("missing anonymous stack entry: %q", formatted)
	}

	generic := errors.New("plain failure")
	if got := FormatError(generic); got != "plain failure" {
		t.Fatalf("unexpected generic error format: %q", got)
	}
}

func TestDisplayError_Generic(t *testing.T) {
	output := captureStderr(t, func() {
		DisplayError(errors.New("generic boom"), ErrorDisplayConfig{UseColor: false})
	})

	if !strings.Contains(output, "generic boom") {
		t.Fatalf("expected generic error output, got: %q", output)
	}
}

func TestDisplayError_Nil(t *testing.T) {
	output := captureStderr(t, func() {
		DisplayError(nil, ErrorDisplayConfig{UseColor: false})
	})
	if output != "" {
		t.Fatalf("expected no output for nil error, got: %q", output)
	}
}

func TestDisplayError_ParseWithContext(t *testing.T) {
	tmpDir := t.TempDir()
	fileName := filepath.Join(tmpDir, "program.oak")
	content := "line 1\nline two\nline 3\n"
	if err := os.WriteFile(fileName, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	parseErr := parseError{
		reason: "expected expression",
		pos: pos{
			fileName: fileName,
			line:     2,
			col:      5,
		},
	}

	output := captureStderr(t, func() {
		DisplayError(parseErr, ErrorDisplayConfig{
			UseColor:     false,
			ShowContext:  true,
			ContextLines: 1,
		})
	})

	if !strings.Contains(output, "Parse Error") {
		t.Fatalf("missing parse header: %q", output)
	}
	if !strings.Contains(output, "expected expression") {
		t.Fatalf("missing parse reason: %q", output)
	}
	if !strings.Contains(output, "Context:") {
		t.Fatalf("missing source context block: %q", output)
	}
	if !strings.Contains(output, "line two") {
		t.Fatalf("missing highlighted source line: %q", output)
	}
	if !strings.Contains(output, "^") {
		t.Fatalf("missing error pointer: %q", output)
	}
}

func TestDisplayError_RuntimeWithStackTrace(t *testing.T) {
	err := &runtimeError{
		reason: "division by zero",
		pos: pos{
			line: 4,
			col:  11,
		},
		stackTrace: []stackEntry{{name: "compute", pos: pos{line: 1, col: 1}}},
	}

	output := captureStderr(t, func() {
		DisplayError(err, ErrorDisplayConfig{
			UseColor:       false,
			ShowContext:    false,
			ShowStackTrace: true,
		})
	})

	if !strings.Contains(output, "Runtime Error") {
		t.Fatalf("missing runtime header: %q", output)
	}
	if !strings.Contains(output, "division by zero") {
		t.Fatalf("missing runtime reason: %q", output)
	}
	if !strings.Contains(output, "Stack Trace") {
		t.Fatalf("missing stack trace header: %q", output)
	}
	if !strings.Contains(output, "in fn compute [1:1]") {
		t.Fatalf("missing stack trace entry: %q", output)
	}
}

func TestDisplayError_RuntimeWithoutStackTrace(t *testing.T) {
	err := &runtimeError{
		reason: "overflow",
		pos: pos{
			line: 1,
			col:  1,
		},
		stackTrace: []stackEntry{{name: "f", pos: pos{line: 1, col: 1}}},
	}

	output := captureStderr(t, func() {
		DisplayError(err, ErrorDisplayConfig{
			UseColor:       false,
			ShowContext:    false,
			ShowStackTrace: false,
		})
	})

	if strings.Contains(output, "Stack Trace") {
		t.Fatalf("did not expect stack trace block when disabled: %q", output)
	}
}

func TestDisplayError_ParseWithoutFileContext(t *testing.T) {
	parseErr := parseError{
		reason: "missing close brace",
		pos: pos{
			line: 5,
			col:  3,
		},
	}

	output := captureStderr(t, func() {
		DisplayError(parseErr, ErrorDisplayConfig{UseColor: false, ShowContext: true})
	})

	if strings.Contains(output, "Context:") {
		t.Fatalf("did not expect context when file name is empty: %q", output)
	}
}

func TestDisplaySourceContext_FileMissing(t *testing.T) {
	var buf strings.Builder
	displaySourceContext(&buf, "definitely-missing-file.oak", 1, 1, ErrorDisplayConfig{UseColor: false, ShowContext: true, ContextLines: 1})

	if buf.Len() != 0 {
		t.Fatalf("expected no output for missing context file, got: %q", buf.String())
	}
}
