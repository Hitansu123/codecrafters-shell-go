package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func echocmd(first []string) string {
	var after string
	for _, item := range first[1:] {
		after = after + item
		after += " "
	}
	return after
}
func typecmd(first []string) string {
	if first[1] == "echo" || first[1] == "exit" || first[1] == "type" {
		return fmt.Sprintf("%s is a shell builtin", first[1])
	}
	return fmt.Sprintf("invalid_command: not found")
}
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
		first := strings.Split(command, " ")
		if first[0] == "echo" {
			fmt.Println(echocmd(first))
		} else if first[0] == "type" {
			fmt.Println(typecmd(first))
		} else if command == "exit 0" {
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
		fmt.Fprint(os.Stdout, "$ ")
	}
}
