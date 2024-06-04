package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
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
			code, err := strconv.Atoi(commands[1])
			if err != nil {
				os.Exit(1)
			}
			os.Exit(code)
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(commands[1:], " "))
		case "type":
			if slices.Contains([]string{"echo", "exit", "type"}, commands[1]) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", commands[1])
			} else if pathCommand, err := tryGetPathCommand(commands[1]); err == nil {
				fmt.Fprintf(os.Stdout, "%s is %s\n", commands[1], pathCommand)
			} else {
				fmt.Fprintf(os.Stdout, "%s not found\n", commands[1])
			}
		case "cd":
			if err := os.Chdir(commands[1]); err != nil {
				fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", commands[1])
			}
		default:
			if output, err := tryExecuteCommand(commands[0], commands[1:]); err == nil {
				fmt.Fprintf(os.Stdout, "%s\n", output)
			} else {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", commands[0])
			}
		}
	}
}

func tryGetPathCommand(command string) (string, error) {
	paths := os.Getenv("PATH")

	for _, path := range strings.Split(paths, ":") {
		pathCommand := path + "/" + command
		if _, err := os.Stat(pathCommand); err == nil {
			return pathCommand, nil
		}
	}

	return "", fmt.Errorf("command not found")
}

func tryExecuteCommand(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
