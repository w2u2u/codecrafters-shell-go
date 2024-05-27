package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
