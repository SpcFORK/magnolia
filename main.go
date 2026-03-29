package main

import (
	"os"
	"strings"
)

func main() {
	if runPackFile() {
		return
	}

	if len(os.Args) > 1 {
		arg := os.Args[1]

		if arg == "--bytecode" && len(os.Args) > 2 {
			os.Args = append(os.Args[:1], os.Args[2:]...)
			runFileBytecode(os.Args[1])
			return
		}

		if arg == "--binary" && len(os.Args) > 2 {
			os.Args = append(os.Args[:1], os.Args[2:]...)
			runFileBinary(os.Args[1])
			return
		}

		if isCommand := performCommandIfExists(arg); !isCommand {
			switch {
			case strings.HasSuffix(arg, ".mb"), strings.HasSuffix(arg, ".mgb"), strings.HasSuffix(arg, ".magb"):
				runFileBinary(arg)
			default:
				runFile(arg)
			}
		}
		return
	}

	if isStdinReadable() {
		runStdin()
		return
	}

	runRepl()
}
