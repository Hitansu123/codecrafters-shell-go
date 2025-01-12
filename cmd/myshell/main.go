package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
func typecmd(path string, first []string) string {

	if first[1] == "echo" || first[1] == "exit" || first[1] == "type" {
		return fmt.Sprintf("%s is a shell builtin", first[1])
	}
	directory := strings.Split(path, ":")
	for _, item := range directory {
		fullPath := fmt.Sprintf("%s/%s", item, first[1])
		if _, err := os.Stat(fullPath); err == nil {
			return fmt.Sprintf("%s is %s", first[1], fullPath)
		}
	}

	return fmt.Sprintf("%s: not found", first[1])
}
func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	path := os.Getenv("PATH")

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
			fmt.Println(typecmd(path, first))
		} else if command == "exit 0" {
			os.Exit(0)
		} else {
			paths := strings.Split(path, ":")
			cmdName := first[0]
			cmdArgs := first[1:]
			isfound := false
			for _, item := range paths {
				fullpath := filepath.Join(item, first[1])
				if _, err := os.Stat(fullpath); err == nil {
					isfound = true

					// Execute the command
					cmd := exec.Command(cmdName, cmdArgs...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						fmt.Printf("%s: error executing command\n", cmdName)
					}
					break
				}
			}
			if !isfound {
				fmt.Printf("%s: not found\n", first[1])
			}
		}
		fmt.Fprint(os.Stdout, "$ ")
	}
}
