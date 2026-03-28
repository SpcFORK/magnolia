package main

import "os"

func main() {
	if runPackFile() {
		return
	}

	if len(os.Args) > 1 {
		arg := os.Args[1]

		// --bytecode flag: run next arg as file in bytecode VM mode
		if arg == "--bytecode" && len(os.Args) > 2 {
			runFileBytecode(os.Args[2])
			return
		}

		if isCommand := performCommandIfExists(arg); !isCommand {
			runFile(arg)
		}
		return
	}

	if isStdinReadable() {
		runStdin()
		return
	}

	runRepl()
}
