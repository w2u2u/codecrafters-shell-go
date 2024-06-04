package libs

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
