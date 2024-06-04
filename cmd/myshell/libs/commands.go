package libs

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ExecExitCommand(arg string) {
	code, err := strconv.Atoi(arg)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(code)
}

func ExecEchoCommand(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}

func ExecTypeCommand(command string) {
	if slices.Contains([]string{"echo", "exit", "type"}, command) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
	} else if pathCommand, err := tryGetPathCommand(command); err == nil {
		fmt.Fprintf(os.Stdout, "%s is %s\n", command, pathCommand)
	} else {
		fmt.Fprintf(os.Stdout, "%s not found\n", command)
	}
}

func ExecCdCommand(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", path)
	}
}

func ExecNonBuiltInCommand(command string, args []string) {
	if output, err := tryExecuteCommand(command, args); err == nil {
		fmt.Fprintf(os.Stdout, "%s\n", output)
	} else {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
