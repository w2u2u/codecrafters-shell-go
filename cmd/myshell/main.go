package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/libs"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		message = strings.TrimSpace(message)

		commands := strings.Split(message, " ")
		switch commands[0] {
		case "exit":
			libs.ExecExitCommand(commands[1])
		case "echo":
			libs.ExecEchoCommand(commands[1:])
		case "type":
			libs.ExecTypeCommand(commands[1])
		case "cd":
			libs.ExecCdCommand(commands[1])
		default:
			libs.ExecNonBuiltInCommand(commands[0], commands[1:])
		}
	}
}
