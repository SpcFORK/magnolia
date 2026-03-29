package main

import (
	"os"
	"strings"
)

func cutArg(index int) {
	os.Args = append(os.Args[:index], os.Args[index+1:]...)
}

func main() {
	if runPackFile() {
		return
	}

	if len(os.Args) > 1 {
		arg := os.Args[1]
		argsLen := len(os.Args)

		if argsLen > 2 {
			switch arg {
			case "--normal", "-n":
				cutArg(1)
				runFile(os.Args[1])
				return
			case "--bytecode", "-b":
				cutArg(1)
				runFileBytecode(os.Args[1])
				return
			case "--executable", "-x":
				cutArg(1)
				runFileBinary(os.Args[1])
				return
			}
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
