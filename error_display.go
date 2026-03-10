package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGray   = "\033[90m"
	colorBold   = "\033[1m"
	colorCyan   = "\033[36m"
)

// ErrorDisplayConfig controls how errors are displayed
type ErrorDisplayConfig struct {
	UseColor       bool
	ShowContext    bool
	ContextLines   int
	ShowStackTrace bool
}

// DefaultErrorConfig returns a default configuration for error display
func DefaultErrorConfig() ErrorDisplayConfig {
	return ErrorDisplayConfig{
		UseColor:       true,
		ShowContext:    true,
		ContextLines:   2,
		ShowStackTrace: true,
	}
}

// DisplayError formats and displays an error with enhanced formatting
func DisplayError(err error, config ErrorDisplayConfig) {
	if err == nil {
		return
	}

	// Helper function to apply color
	applyColor := func(color, text string) string {
		if config.UseColor {
			return color + text + colorReset
		}
		return text
	}

	// Display different error types differently
	switch e := err.(type) {
	case parseError:
		displayParseError(e, config, applyColor)
	case *runtimeError:
		displayRuntimeError(e, config, applyColor)
	default:
		// Generic error display
		fmt.Fprintf(os.Stderr, "%s\n", applyColor(colorRed, err.Error()))
	}
}

func displayParseError(e parseError, config ErrorDisplayConfig, applyColor func(string, string) string) {
	// Header
	fmt.Fprintf(os.Stderr, "\n%s\n", applyColor(colorBold+colorRed, "╭─ Parse Error ─────────────────────────────────────────────"))
	fmt.Fprintf(os.Stderr, "│\n")

	// Position information
	if e.fileName != "" {
		fmt.Fprintf(os.Stderr, "│ %s %s\n",
			applyColor(colorCyan, "File:"),
			applyColor(colorBold, e.fileName))
	}
	fmt.Fprintf(os.Stderr, "│ %s %s\n",
		applyColor(colorCyan, "Position:"),
		applyColor(colorBold, e.pos.String()))
	fmt.Fprintf(os.Stderr, "│\n")

	// Error message
	fmt.Fprintf(os.Stderr, "│ %s\n", applyColor(colorRed, e.reason))

	// Source context
	if config.ShowContext && e.fileName != "" {
		displaySourceContext(e.fileName, e.line, e.col, config, applyColor)
	}

	// Footer
	fmt.Fprintf(os.Stderr, "╰───────────────────────────────────────────────────────────\n\n")
}

func displayRuntimeError(e *runtimeError, config ErrorDisplayConfig, applyColor func(string, string) string) {
	// Header
	fmt.Fprintf(os.Stderr, "\n%s\n", applyColor(colorBold+colorRed, "╭─ Runtime Error ───────────────────────────────────────────"))
	fmt.Fprintf(os.Stderr, "│\n")

	// Position information
	if e.fileName != "" {
		fmt.Fprintf(os.Stderr, "│ %s %s\n",
			applyColor(colorCyan, "File:"),
			applyColor(colorBold, e.fileName))
	}
	fmt.Fprintf(os.Stderr, "│ %s %s\n",
		applyColor(colorCyan, "Position:"),
		applyColor(colorBold, e.pos.String()))
	fmt.Fprintf(os.Stderr, "│\n")

	// Error message
	fmt.Fprintf(os.Stderr, "│ %s\n", applyColor(colorRed, e.reason))

	// Source context
	if config.ShowContext && e.fileName != "" {
		displaySourceContext(e.fileName, e.line, e.col, config, applyColor)
	}

	// Stack trace
	if config.ShowStackTrace && len(e.stackTrace) > 0 {
		fmt.Fprintf(os.Stderr, "│\n")
		fmt.Fprintf(os.Stderr, "│ %s\n", applyColor(colorYellow, "Stack Trace:"))
		for _, entry := range e.stackTrace {
			fmt.Fprintf(os.Stderr, "│   %s\n", applyColor(colorGray, entry.String()))
		}
	}

	// Footer
	fmt.Fprintf(os.Stderr, "╰───────────────────────────────────────────────────────────\n\n")
}

func displaySourceContext(fileName string, line, col int, config ErrorDisplayConfig, applyColor func(string, string) string) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	if line < 1 || line > len(lines) {
		return
	}

	fmt.Fprintf(os.Stderr, "│\n")
	fmt.Fprintf(os.Stderr, "│ %s\n", applyColor(colorCyan, "Context:"))

	// Calculate the range of lines to show
	startLine := 1
	if line-config.ContextLines > 1 {
		startLine = line - config.ContextLines
	}
	endLine := len(lines)
	if line+config.ContextLines < len(lines) {
		endLine = line + config.ContextLines
	}

	// Display context lines
	for i := startLine; i <= endLine; i++ {
		lineNum := fmt.Sprintf("%4d", i)

		if i == line {
			// The error line - highlight it
			fmt.Fprintf(os.Stderr, "│ %s │ %s\n",
				applyColor(colorRed+colorBold, lineNum),
				applyColor(colorRed, lines[i-1]))

			// Show the error pointer
			if col > 0 {
				pointer := strings.Repeat(" ", col-1) + "^"
				fmt.Fprintf(os.Stderr, "│      │ %s\n", applyColor(colorRed+colorBold, pointer))
			}
		} else {
			// Context line
			fmt.Fprintf(os.Stderr, "│ %s │ %s\n",
				applyColor(colorGray, lineNum),
				lines[i-1])
		}
	}
}

// FormatError returns a formatted error string (without printing)
func FormatError(err error) string {
	if err == nil {
		return ""
	}

	switch e := err.(type) {
	case parseError:
		result := fmt.Sprintf("Parse Error at %s", e.pos.String())
		if e.fileName != "" {
			result = fmt.Sprintf("Parse Error in %s at %s", e.fileName, e.pos.String())
		}
		return result + ": " + e.reason
	case *runtimeError:
		result := fmt.Sprintf("Runtime Error at %s", e.pos.String())
		if e.fileName != "" {
			result = fmt.Sprintf("Runtime Error in %s at %s", e.fileName, e.pos.String())
		}
		result += ": " + e.reason

		if len(e.stackTrace) > 0 {
			trace := make([]string, len(e.stackTrace))
			for i, entry := range e.stackTrace {
				trace[i] = entry.String()
			}
			result += "\n" + strings.Join(trace, "\n")
		}
		return result
	default:
		return err.Error()
	}
}
