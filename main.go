package main

import (
	"fmt"
	"os"
)

func debugStartup(msg string) {
	if os.Getenv("MAGNOLIA_DEBUG_STARTUP") == "1" {
		fmt.Fprintln(os.Stderr, "[startup] "+msg)
	}
}

func main() {
	debugStartup("begin")
	if runPackFile() {
		debugStartup("runPackFile=true")
		return
	}
	debugStartup("runPackFile=false")

	if len(os.Args) > 1 {
		debugStartup("args>1")
		arg := os.Args[1]
		if isCommand := performCommandIfExists(arg); !isCommand {
			debugStartup("running file arg")
			runFile(arg)
		}
		return
	}

	if isStdinReadable() {
		// No explicit command: avoid interpreting redirected shell/control input
		// as Oak source. Use `eval` or `pipe` for stdin-driven execution.
		debugStartup("stdin redirected without command; exiting")
		return
	}

	debugStartup("running repl")
	runRepl()
}
