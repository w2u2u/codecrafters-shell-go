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
			execExitCommand(commands[1])
		case "echo":
			execEchoCommand(commands[1:])
		case "type":
			execTypeCommand(commands[1])
		case "cd":
			execCdCommand(commands[1])
		default:
			if output, err := tryExecuteCommand(commands[0], commands[1:]); err == nil {
				fmt.Fprintf(os.Stdout, "%s\n", output)
			} else {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", commands[0])
			}
		}
	}
}

func execExitCommand(arg string) {
	code, err := strconv.Atoi(arg)
	if err != nil {
		os.Exit(1)
	}
	os.Exit(code)
}

func execEchoCommand(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}

func execTypeCommand(command string) {
	if slices.Contains([]string{"echo", "exit", "type"}, command) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
	} else if pathCommand, err := tryGetPathCommand(command); err == nil {
		fmt.Fprintf(os.Stdout, "%s is %s\n", command, pathCommand)
	} else {
		fmt.Fprintf(os.Stdout, "%s not found\n", command)
	}
}

func execCdCommand(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", path)
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
