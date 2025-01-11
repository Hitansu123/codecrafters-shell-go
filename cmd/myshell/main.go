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
		first := strings.Split(command, " ")
		if first[0] == "echo" {
			var after string
			for _, item := range first[1:] {
				after = after + item
				after += " "
			}
			fmt.Println(after)
		} else if command == "exit 0" {
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
		}
		fmt.Fprint(os.Stdout, "$ ")
	}
}
