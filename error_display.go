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

const errorBoxWidth = 60

// ErrorDisplayConfig controls how errors are displayed
type ErrorDisplayConfig struct {
	Writer         io.Writer
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

func (c ErrorDisplayConfig) writer() io.Writer {
	if c.Writer != nil {
		return c.Writer
	}
	return os.Stderr
}

func (c ErrorDisplayConfig) colorize(color, text string) string {
	if c.UseColor {
		return color + text + colorReset
	}
	return text
}

// DisplayError formats and displays an error with enhanced formatting
func DisplayError(err error, config ErrorDisplayConfig) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case parseError:
		displayErrorBox(config, "Parse Error", e.fileName, e.pos, e.reason, nil)
	case *runtimeError:
		displayErrorBox(config, "Runtime Error", e.fileName, e.pos, e.reason, e.stackTrace)
	default:
		fmt.Fprintf(config.writer(), "%s\n", config.colorize(colorRed, err.Error()))
	}
}

func displayErrorBox(config ErrorDisplayConfig, kind, fileName string, p pos, reason string, stackTrace []stackEntry) {
	w := config.writer()

	// Header with dynamic width
	dashCount := max(3, errorBoxWidth-4-len(kind))
	header := "╭─ " + kind + " " + strings.Repeat("─", dashCount)
	fmt.Fprintf(w, "\n%s\n", config.colorize(colorBold+colorRed, header))
	fmt.Fprintf(w, "│\n")

	// Position information
	if fileName != "" {
		fmt.Fprintf(w, "│ %s %s\n",
			config.colorize(colorCyan, "File:"),
			config.colorize(colorBold, fileName))
	}
	fmt.Fprintf(w, "│ %s %s\n",
		config.colorize(colorCyan, "Position:"),
		config.colorize(colorBold, p.String()))
	fmt.Fprintf(w, "│\n")

	// Error message
	fmt.Fprintf(w, "│ %s\n", config.colorize(colorRed, reason))

	// Source context
	if config.ShowContext && fileName != "" {
		displaySourceContext(w, fileName, p.line, p.col, config)
	}

	// Stack trace
	if config.ShowStackTrace && len(stackTrace) > 0 {
		fmt.Fprintf(w, "│\n")
		fmt.Fprintf(w, "│ %s\n", config.colorize(colorYellow, "Stack Trace:"))
		for _, entry := range stackTrace {
			fmt.Fprintf(w, "│   %s\n", config.colorize(colorGray, entry.String()))
		}
	}

	// Footer
	fmt.Fprintf(w, "╰%s\n\n", strings.Repeat("─", errorBoxWidth-1))
}

func displaySourceContext(w io.Writer, fileName string, line, col int, config ErrorDisplayConfig) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")
	if line < 1 || line > len(lines) {
		return
	}

	fmt.Fprintf(w, "│\n")
	fmt.Fprintf(w, "│ %s\n", config.colorize(colorCyan, "Context:"))

	startLine := max(1, line-config.ContextLines)
	endLine := min(len(lines), line+config.ContextLines)

	for i := startLine; i <= endLine; i++ {
		lineNum := fmt.Sprintf("%4d", i)

		if i == line {
			fmt.Fprintf(w, "│ %s │ %s\n",
				config.colorize(colorRed+colorBold, lineNum),
				config.colorize(colorRed, lines[i-1]))

			// Show error pointer, preserving tab alignment
			if col > 0 {
				src := lines[i-1]
				prefixLen := min(col-1, len(src))
				pointer := strings.Map(func(r rune) rune {
					if r == '\t' {
						return '\t'
					}
					return ' '
				}, src[:prefixLen]) + "^"
				fmt.Fprintf(w, "│      │ %s\n", config.colorize(colorRed+colorBold, pointer))
			}
		} else {
			fmt.Fprintf(w, "│ %s │ %s\n",
				config.colorize(colorGray, lineNum),
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
