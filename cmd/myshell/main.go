package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	// Wait for user input
	userInput := bufio.NewReader(os.Stdin)

	for {
		command, err := userInput.ReadString('\n')

		if err != nil {
			fmt.Println("invalid_command: command not found")
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		if command == "exit 0" {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		fmt.Fprint(os.Stdout, "$ ")
	}
}
