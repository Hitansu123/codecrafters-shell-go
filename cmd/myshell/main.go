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
	userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("invalid_command: command not found")
	}
	userInput = strings.TrimSpace(userInput)
	output := fmt.Sprintf("%s : command not found", userInput)
	fmt.Fprint(os.Stdout, output)
}
